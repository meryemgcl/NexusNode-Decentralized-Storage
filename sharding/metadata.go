package sharding

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ShardMetadata holds information needed to reassemble the original file.
type ShardMetadata struct {
	OriginalFileName string `json:"original_file_name"`
	OriginalSize     int64  `json:"original_size"`
	DataShards       int    `json:"data_shards"`
	ParityShards     int    `json:"parity_shards"`
	TotalShards      int    `json:"total_shards"`
}

// MetadataFileName returns the conventional metadata filename for a given file.
func MetadataFileName(baseName string) string {
	return fmt.Sprintf("%s.meta.json", baseName)
}

// WriteMetadata saves shard metadata as a JSON file alongside the shards.
func WriteMetadata(outputDir string, meta ShardMetadata) (string, error) {
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return "", fmt.Errorf("metadata: marshal: %w", err)
	}

	path := filepath.Join(outputDir, MetadataFileName(meta.OriginalFileName))
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", fmt.Errorf("metadata: write %s: %w", path, err)
	}

	return path, nil
}

// ReadMetadata loads shard metadata from a JSON file.
func ReadMetadata(metaPath string) (ShardMetadata, error) {
	data, err := os.ReadFile(metaPath)
	if err != nil {
		return ShardMetadata{}, fmt.Errorf("metadata: read %s: %w", metaPath, err)
	}

	var meta ShardMetadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return ShardMetadata{}, fmt.Errorf("metadata: unmarshal: %w", err)
	}

	return meta, nil
}
