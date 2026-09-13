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
	fileToChunk := flag.String("chunk", "", "File to split into chunks")
	filesToAssemble := flag.String("assemble", "", "Comma separated list of chunks to assemble")
	outputFile := flag.String("out", "restored_file", "Output file for assembled chunks")
	flag.Parse()

	// 1. Sharding Operations
	if *fileToChunk != "" {
		fmt.Printf("Chunking file: %s\n", *fileToChunk)
		outDir := filepath.Dir(*fileToChunk)
		chunkPaths, err := sharding.ChunkFile(*fileToChunk, 1024*1024, outDir) // 1 MB chunks
		if err != nil {
			log.Fatalf("Error chunking file: %v", err)
		}
		fmt.Println("File split into chunks:")
		for _, p := range chunkPaths {
			fmt.Printf(" - %s\n", p)
		}
		return
	}

	// 2. Assembler Operations
	if *filesToAssemble != "" {
		chunks := strings.Split(*filesToAssemble, ",")
		fmt.Printf("Assembling %d chunks into %s\n", len(chunks), *outputFile)
		err := sharding.AssembleChunks(chunks, *outputFile)
		if err != nil {
			log.Fatalf("Error assembling chunks: %v", err)
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
