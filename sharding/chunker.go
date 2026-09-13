package sharding

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/klauspost/reedsolomon"
	"github.com/meryemgcl/NexusNode-Decentralized-Storage/crypto"
)

// SplitFile splits a file into Reed-Solomon data + parity shards, encrypts each shard
// with AES-256 GCM, and writes them plus a metadata file to outputDir.
// encryptionKey must be exactly 32 bytes.
// Returns the list of shard paths and the metadata file path.
func SplitFile(filePath string, dataShards, parityShards int, encryptionKey []byte, outputDir string) (shardPaths []string, metaPath string, err error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("SplitFile: read %s: %w", filePath, err)
	}
	originalSize := int64(len(b))

	enc, err := reedsolomon.New(dataShards, parityShards)
	if err != nil {
		return nil, "", fmt.Errorf("SplitFile: create encoder: %w", err)
	}

	shards, err := enc.Split(b)
	if err != nil {
		return nil, "", fmt.Errorf("SplitFile: split: %w", err)
	}

	if err := enc.Encode(shards); err != nil {
		return nil, "", fmt.Errorf("SplitFile: encode parity: %w", err)
	}

	baseName := filepath.Base(filePath)

	for i, shard := range shards {
		encryptedShard, err := crypto.Encrypt(encryptionKey, shard)
		if err != nil {
			return nil, "", fmt.Errorf("SplitFile: encrypt shard %d: %w", i, err)
		}

		chunkPath := filepath.Join(outputDir, fmt.Sprintf("%s.shard.%d", baseName, i))
		if err := os.WriteFile(chunkPath, encryptedShard, 0644); err != nil {
			return nil, "", fmt.Errorf("SplitFile: write shard %d: %w", i, err)
		}

		shardPaths = append(shardPaths, chunkPath)
	}

	// Save metadata so AssembleShards knows the exact original file size.
	meta := ShardMetadata{
		OriginalFileName: baseName,
		OriginalSize:     originalSize,
		DataShards:       dataShards,
		ParityShards:     parityShards,
		TotalShards:      dataShards + parityShards,
	}
	metaPath, err = WriteMetadata(outputDir, meta)
	if err != nil {
		return nil, "", err
	}

	return shardPaths, metaPath, nil
}
