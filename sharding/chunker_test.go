package sharding_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/meryemgcl/NexusNode-Decentralized-Storage/sharding"
)

var testKey = []byte("12345678901234567890123456789012") // 32 bytes

func TestSplitFileCreatesCorrectNumberOfShards(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("NexusNode test data for sharding. Hello World!"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	dataShards := 4
	parityShards := 2

	shardPaths, metaPath, err := sharding.SplitFile(testFile, dataShards, parityShards, testKey, tmpDir)
	if err != nil {
		t.Fatalf("SplitFile failed: %v", err)
	}

	expectedShards := dataShards + parityShards
	if len(shardPaths) != expectedShards {
		t.Fatalf("expected %d shards, got %d", expectedShards, len(shardPaths))
	}

	// Metadata file should exist
	if metaPath == "" {
		t.Fatal("expected non-empty metaPath")
	}
	if _, err := os.Stat(metaPath); os.IsNotExist(err) {
		t.Fatalf("metadata file not found at %s", metaPath)
	}

	// Each shard file should exist and be non-empty
	for _, p := range shardPaths {
		info, err := os.Stat(p)
		if err != nil {
			t.Fatalf("shard file not found: %s", p)
		}
		if info.Size() == 0 {
			t.Fatalf("shard file is empty: %s", p)
		}
	}
}

func TestSplitFileWritesCorrectMetadata(t *testing.T) {
	tmpDir := t.TempDir()
	content := []byte("metadata accuracy check content 1234")
	testFile := filepath.Join(tmpDir, "meta_test.bin")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	_, metaPath, err := sharding.SplitFile(testFile, 4, 2, testKey, tmpDir)
	if err != nil {
		t.Fatalf("SplitFile failed: %v", err)
	}

	meta, err := sharding.ReadMetadata(metaPath)
	if err != nil {
		t.Fatalf("ReadMetadata failed: %v", err)
	}

	if meta.OriginalSize != int64(len(content)) {
		t.Fatalf("metadata OriginalSize mismatch: got %d, want %d", meta.OriginalSize, len(content))
	}
	if meta.DataShards != 4 || meta.ParityShards != 2 {
		t.Fatalf("unexpected shard counts in metadata")
	}
}
