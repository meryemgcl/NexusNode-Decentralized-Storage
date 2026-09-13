package sharding

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/klauspost/reedsolomon"
	"github.com/meryemgcl/NexusNode-Decentralized-Storage/crypto"
)

// AssembleShards reads encrypted shards from disk, decrypts them, reconstructs any
// missing shards using Reed-Solomon parity, then writes the original file to outputPath.
//
// shardPaths must have exactly (dataShards + parityShards) entries.
// Pass an empty string ("") for any shard that is unavailable (simulates offline node).
// metaPath is the path to the .meta.json file produced by SplitFile.
// encryptionKey must be the same 32-byte key used during SplitFile.
func AssembleShards(shardPaths []string, metaPath string, encryptionKey []byte, outputPath string) error {
	meta, err := ReadMetadata(metaPath)
	if err != nil {
		return err
	}

	totalShards := meta.DataShards + meta.ParityShards
	if len(shardPaths) != totalShards {
		return fmt.Errorf("AssembleShards: expected %d shard paths, got %d", totalShards, len(shardPaths))
	}

	enc, err := reedsolomon.New(meta.DataShards, meta.ParityShards)
	if err != nil {
		return fmt.Errorf("AssembleShards: create encoder: %w", err)
	}

	shards := make([][]byte, totalShards)

	for i, path := range shardPaths {
		if path == "" {
			slog.Warn("shard is missing, will attempt reconstruction", "index", i)
			continue
		}

		encryptedData, err := os.ReadFile(path)
		if err != nil {
			slog.Warn("failed to read shard, treating as missing", "path", path, "error", err)
			continue
		}

		decryptedData, err := crypto.Decrypt(encryptionKey, encryptedData)
		if err != nil {
			slog.Warn("failed to decrypt shard, treating as missing", "path", path, "error", err)
			continue
		}

		shards[i] = decryptedData
	}

	// Reconstruct missing shards if needed.
	ok, _ := enc.Verify(shards)
	if !ok {
		slog.Info("reconstructing missing/corrupted shards via Reed-Solomon parity")
		if err := enc.Reconstruct(shards); err != nil {
			return fmt.Errorf("AssembleShards: reconstruct failed: %w", err)
		}
		slog.Info("reconstruction successful")
	} else {
		slog.Info("all shards intact, no reconstruction needed")
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("AssembleShards: create output file: %w", err)
	}
	defer outFile.Close()

	// Use the exact original file size stored in metadata to avoid zero-byte padding.
	if err := enc.Join(outFile, shards, int(meta.OriginalSize)); err != nil {
		return fmt.Errorf("AssembleShards: join shards: %w", err)
	}

	slog.Info("file assembled successfully", "output", outputPath, "size", meta.OriginalSize)
	return nil
}
