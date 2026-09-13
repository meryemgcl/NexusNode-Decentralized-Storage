package sharding_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/meryemgcl/NexusNode-Decentralized-Storage/sharding"
)

func TestAssembleShardsFullRoundtrip(t *testing.T) {
	tmpDir := t.TempDir()
	original := []byte("NexusNode full roundtrip test. This content must survive split and reassembly perfectly.")

	testFile := filepath.Join(tmpDir, "roundtrip.txt")
	if err := os.WriteFile(testFile, original, 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	dataShards, parityShards := 4, 2
	shardPaths, metaPath, err := sharding.SplitFile(testFile, dataShards, parityShards, testKey, tmpDir)
	if err != nil {
		t.Fatalf("SplitFile failed: %v", err)
	}

	outPath := filepath.Join(tmpDir, "restored.txt")
	if err := sharding.AssembleShards(shardPaths, metaPath, testKey, outPath); err != nil {
		t.Fatalf("AssembleShards failed: %v", err)
	}

	restored, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("failed to read restored file: %v", err)
	}

	if !bytes.Equal(restored, original) {
		t.Fatalf("restored content mismatch:\n got: %q\nwant: %q", restored, original)
	}
}

func TestAssembleShardsWithMissingShards(t *testing.T) {
	tmpDir := t.TempDir()
	original := []byte("NexusNode fault tolerance test. Some shards will be lost and must be recovered via parity!")

	testFile := filepath.Join(tmpDir, "fault_tolerance.txt")
	if err := os.WriteFile(testFile, original, 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	dataShards, parityShards := 4, 2
	shardPaths, metaPath, err := sharding.SplitFile(testFile, dataShards, parityShards, testKey, tmpDir)
	if err != nil {
		t.Fatalf("SplitFile failed: %v", err)
	}

	// Simulate 2 nodes going offline (parity count allows recovery of up to 2 missing)
	shardPaths[0] = ""
	shardPaths[2] = ""

	outPath := filepath.Join(tmpDir, "recovered.txt")
	if err := sharding.AssembleShards(shardPaths, metaPath, testKey, outPath); err != nil {
		t.Fatalf("AssembleShards with missing shards failed: %v", err)
	}

	recovered, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("failed to read recovered file: %v", err)
	}

	if !bytes.Equal(recovered, original) {
		t.Fatalf("recovered content mismatch:\n got: %q\nwant: %q", recovered, original)
	}
}

func TestAssembleShardsFailsWhenTooManyMissing(t *testing.T) {
	tmpDir := t.TempDir()
	original := []byte("This will fail because too many shards are lost.")

	testFile := filepath.Join(tmpDir, "too_many_lost.txt")
	if err := os.WriteFile(testFile, original, 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	dataShards, parityShards := 4, 2
	shardPaths, metaPath, err := sharding.SplitFile(testFile, dataShards, parityShards, testKey, tmpDir)
	if err != nil {
		t.Fatalf("SplitFile failed: %v", err)
	}

	// Remove 3 shards — exceeds parity capacity (max 2 recoverable)
	shardPaths[0] = ""
	shardPaths[1] = ""
	shardPaths[2] = ""

	outPath := filepath.Join(tmpDir, "should_fail.txt")
	err = sharding.AssembleShards(shardPaths, metaPath, testKey, outPath)
	if err == nil {
		t.Fatal("expected error when too many shards are missing, got nil")
	}
}
