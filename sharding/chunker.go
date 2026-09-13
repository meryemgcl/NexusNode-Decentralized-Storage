package sharding

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ChunkFile splits a file into multiple chunks of the given size.
// Returns a list of the chunk file paths.
func ChunkFile(filePath string, chunkSize int64, outputDir string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Calculate number of chunks
	fileSize := fileInfo.Size()
	numChunks := (fileSize + chunkSize - 1) / chunkSize
	var chunkPaths []string

	buffer := make([]byte, chunkSize)

	for i := int64(0); i < numChunks; i++ {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}

		chunkName := fmt.Sprintf("%s.chunk.%d", filepath.Base(filePath), i)
		chunkPath := filepath.Join(outputDir, chunkName)

		err = os.WriteFile(chunkPath, buffer[:bytesRead], 0644)
		if err != nil {
			return nil, err
		}

		chunkPaths = append(chunkPaths, chunkPath)
	}

	return chunkPaths, nil
}
