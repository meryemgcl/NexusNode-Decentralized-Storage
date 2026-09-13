package main

import (
	"flag"
	"fmt"
	"log"
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
	fileToChunk := flag.String("chunk", "", "File to split into shards")
	filesToAssemble := flag.String("assemble", "", "Comma separated list of shards to assemble (use 'missing' for lost shards)")
	outputFile := flag.String("out", "restored_file", "Output file for assembled shards")
	
	dataShards := flag.Int("data", 10, "Number of data shards")
	parityShards := flag.Int("parity", 4, "Number of parity shards")
	encryptionKeyStr := flag.String("key", "0123456789abcdef0123456789abcdef", "AES-256 Encryption key (must be exactly 32 bytes)")
	
	flag.Parse()

	if len(*encryptionKeyStr) != 32 {
		log.Fatalf("Encryption key must be exactly 32 bytes long.")
	}
	encKey := []byte(*encryptionKeyStr)

	// 1. Sharding Operations
	if *fileToChunk != "" {
		fmt.Printf("Splitting file: %s (Data Shards: %d, Parity Shards: %d)\n", *fileToChunk, *dataShards, *parityShards)
		outDir := filepath.Dir(*fileToChunk)
		
		chunkPaths, err := sharding.SplitFile(*fileToChunk, *dataShards, *parityShards, encKey, outDir)
		if err != nil {
			log.Fatalf("Error chunking file: %v", err)
		}
		
		fmt.Println("File split and encrypted into shards:")
		for _, p := range chunkPaths {
			fmt.Printf(" - %s\n", p)
		}
		return
	}

	// 2. Assembler Operations
	if *filesToAssemble != "" {
		chunks := strings.Split(*filesToAssemble, ",")
		
		// If user explicitly writes 'missing', we treat it as an empty string (missing shard)
		for i, c := range chunks {
			if strings.TrimSpace(c) == "missing" {
				chunks[i] = ""
			}
		}

		fmt.Printf("Assembling %d shards into %s\n", len(chunks), *outputFile)
		err := sharding.AssembleShards(chunks, *dataShards, *parityShards, encKey, *outputFile)
		if err != nil {
			log.Fatalf("Error assembling shards: %v", err)
		}
		fmt.Println("File assembled successfully.")
		return
	}

	// 3. Network Node Operations
	log.Println("Starting NexusNode...")
	node, err := network.NewNode(*port)
	if err != nil {
		log.Fatalf("Failed to start node: %v", err)
	}

	log.Printf("Node started. Listening on:")
	for _, addr := range node.Host.Addrs() {
		log.Printf(" - %s/p2p/%s\n", addr, node.Host.ID())
	}

	// Start mDNS discovery
	rendezvous := "nexusnode-discovery-v1"
	log.Printf("Setting up mDNS discovery with rendezvous string: %s\n", rendezvous)
	if err := node.SetupDiscovery(rendezvous); err != nil {
		log.Fatalf("Failed to set up discovery: %v", err)
	}

	// Wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Println("Shutting down node...")

	if err := node.Host.Close(); err != nil {
		panic(err)
	}
}
