package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/meryemgcl/NexusNode-Decentralized-Storage/network"
	"github.com/meryemgcl/NexusNode-Decentralized-Storage/sharding"
)

func main() {
	port := flag.Int("port", 0, "Port to listen on (0 for random)")
	fileToChunk := flag.String("chunk", "", "File to split into encrypted shards")
	filesToAssemble := flag.String("assemble", "", "Comma-separated list of shard paths (use 'missing' for unavailable shards)")
	metaFile := flag.String("meta", "", "Path to the .meta.json file (required for -assemble)")
	outputFile := flag.String("out", "restored_file", "Output path for the assembled file")
	dataShards := flag.Int("data", 10, "Number of data shards")
	parityShards := flag.Int("parity", 4, "Number of parity shards")

	flag.Parse()

	// Configure structured logging.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})))

	// --- Sharding or Assembly: both require the encryption key ---
	if *fileToChunk != "" || *filesToAssemble != "" {
		encKey := readKeyFromStdin()

		if *fileToChunk != "" {
			slog.Info("splitting file", "path", *fileToChunk, "data_shards", *dataShards, "parity_shards", *parityShards)
			outDir := filepath.Dir(*fileToChunk)
			shardPaths, metaPath, err := sharding.SplitFile(*fileToChunk, *dataShards, *parityShards, encKey, outDir)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("File split into %d shards:\n", len(shardPaths))
			for _, p := range shardPaths {
				fmt.Printf("  %s\n", p)
			}
			fmt.Printf("Metadata: %s\n", metaPath)
			return
		}

		if *filesToAssemble != "" {
			if *metaFile == "" {
				log.Fatal("error: -meta flag is required when using -assemble")
			}
			chunks := strings.Split(*filesToAssemble, ",")
			for i, c := range chunks {
				if strings.TrimSpace(c) == "missing" {
					chunks[i] = ""
				}
			}
			slog.Info("assembling shards", "count", len(chunks), "output", *outputFile)
			if err := sharding.AssembleShards(chunks, *metaFile, encKey, *outputFile); err != nil {
				log.Fatalf("error: %v", err)
			}
			return
		}
	}

	// --- P2P Node operation (no key required) ---
	slog.Info("starting NexusNode...")
	node, err := network.NewNode(*port)
	if err != nil {
		log.Fatalf("failed to start node: %v", err)
	}

	for _, addr := range node.Host.Addrs() {
		slog.Info("listening", "addr", fmt.Sprintf("%s/p2p/%s", addr, node.Host.ID()))
	}

	rendezvous := "nexusnode-discovery-v1"
	slog.Info("starting mDNS discovery", "rendezvous", rendezvous)
	if err := node.SetupDiscovery(rendezvous); err != nil {
		log.Fatalf("failed to set up discovery: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	slog.Info("shutting down node...")
	if err := node.Host.Close(); err != nil {
		log.Fatalf("error closing host: %v", err)
	}
}

// readKeyFromStdin reads the AES-256 encryption key from the first line of stdin.
// This avoids exposing the key in process argument lists (ps aux, Task Manager, etc.).
func readKeyFromStdin() []byte {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error: failed to read encryption key from stdin: %v", err)
	}
	key := strings.TrimRight(line, "\r\n")
	if len(key) != 32 {
		log.Fatalf("error: encryption key must be exactly 32 characters, got %d", len(key))
	}
	return []byte(key)
}
