package main

import (
	"fmt"
	"os"
	"path/filepath"

	"aead.dev/minisign"
)

func main() {
	// // Get the file to sign from the env variable
	// fileToSign := os.Getenv("FILE_TO_SIGN_PATH")
	// Get the INPUT_PASSWORD environment variable
	inputPassword := os.Getenv("INPUT_PASSWORD")
	// Get the INPUT_RAW_PRIVATE_KEY environment variable
	inputRawPrivateKey := os.Getenv("INPUT_RAW_PRIVATE_KEY")

	// If any of the env variables are not set, print an error and exit
	// if fileToSign == "" {
	// 	fmt.Println("Error: FILE_TO_SIGN_PATH is not set")
	// 	os.Exit(1)
	// }
	if inputPassword == "" {
		fmt.Println("Error: INPUT_PASSWORD is not set")
		os.Exit(1)
	}
	if inputRawPrivateKey == "" {
		fmt.Println("Error: INPUT_RAW_PRIVATE_KEY is not set")
		os.Exit(1)
	}

	// Env variables are set
	fmt.Println("Env variables loaded successfully")

	// Decrypt the raw private key with the password
	privateKey, err := minisign.DecryptKey(inputPassword, []byte(inputRawPrivateKey))
	if err != nil {
		fmt.Println("Failed to decrypt the private key: ", err)
		os.Exit(1)
	}

	// Private key decrypted successfully
	fmt.Println("Private key decrypted successfully")

	// Get the build directory
	buildDir := filepath.Join("..", "build", "bin")

	fmt.Println("Looking for files in: ", buildDir)

	// Find all the files in the build/bin directory
	files, err := filepath.Glob(filepath.Join(buildDir, "*"))
	if err != nil {
		fmt.Println("Failed to find files: ", err)
		os.Exit(1)
	}

	// Sign the first file found
	for _, file := range files {
		fmt.Println("Signing file: ", file)

		// Read the file
		fileBytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Failed to read the file: ", err)
			os.Exit(1)
		}

		// Generate the signature
		signature := minisign.Sign(privateKey, fileBytes)

		fmt.Println("Signature generated successfully")

		// Write the signature to a file
		err = os.WriteFile("signature.minisig", signature, 0644)
		if err != nil {
			fmt.Println("Failed to write the signature to a file: ", err)
			os.Exit(1)
		}

		fmt.Println("Signature written to signature.minisig")

		return
	}
}
