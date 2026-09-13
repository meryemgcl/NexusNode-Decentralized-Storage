package sharding

import (
	"os"
)

// AssembleChunks reads the given chunk files and combines them into the target output file.
func AssembleChunks(chunkPaths []string, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, chunkPath := range chunkPaths {
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			return err
		}

		_, err = outFile.Write(chunkData)
		if err != nil {
			return err
		}
	}

	return nil
}
