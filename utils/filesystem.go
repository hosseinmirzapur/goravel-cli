package utils

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) error {

	// Open the source file for reading
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return err
	}
	defer sourceFile.Close()

	// Create the destination file for writing
	destinationFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return err
	}
	defer destinationFile.Close()

	// Create a buffer with a fixed size (e.g., 4096 bytes)
	buffer := make([]byte, 4096)

	// Copy the content from the source file to the destination file using the buffer
	for {
		// Read data from the source file into the buffer
		bytesRead, err := sourceFile.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from source file:", err)
			return err
		}

		// If no more data is available, break the loop
		if bytesRead == 0 {
			break
		}

		// Write the data from the buffer to the destination file
		_, err = destinationFile.Write(buffer[:bytesRead])
		if err != nil {
			fmt.Println("Error writing to destination file:", err)
			return err
		}
	}

	return nil
}

func WriteFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
