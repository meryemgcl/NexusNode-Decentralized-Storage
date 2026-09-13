package sharding

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/klauspost/reedsolomon"
	"github.com/meryemgcl/NexusNode-Decentralized-Storage/crypto"
)

// SplitFile splits a file into data shards and parity shards using Reed-Solomon,
// encrypts each shard using AES-256, and writes them to the output directory.
// The encryption key must be 32 bytes.
func SplitFile(filePath string, dataShards, parityShards int, encryptionKey []byte, outputDir string) ([]string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	enc, err := reedsolomon.New(dataShards, parityShards)
	if err != nil {
		return nil, err
	}

	// Split the file into shards
	shards, err := enc.Split(b)
	if err != nil {
		return nil, err
	}

	// Encode parity
	err = enc.Encode(shards)
	if err != nil {
		return nil, err
	}

	var chunkPaths []string
	baseName := filepath.Base(filePath)

	// Encrypt and write shards
	for i, shard := range shards {
		encryptedShard, err := crypto.Encrypt(encryptionKey, shard)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt shard %d: %v", i, err)
		}

		chunkName := fmt.Sprintf("%s.shard.%d", baseName, i)
		chunkPath := filepath.Join(outputDir, chunkName)

		err = os.WriteFile(chunkPath, encryptedShard, 0644)
		if err != nil {
			return nil, err
		}

		chunkPaths = append(chunkPaths, chunkPath)
	}

	return chunkPaths, nil
}
