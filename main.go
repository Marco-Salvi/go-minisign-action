package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Get the minisign location from the env variable
	minisignLocation := os.Getenv("MINISIGN_PATH")
	// Get the file to sign from the env variable
	fileToSign := os.Getenv("FILE_TO_SIGN_PATH")
	// Get the INPUT_PASSWORD environment variable
	inputPassword := os.Getenv("INPUT_PASSWORD")
	// Get the INPUT_MINISIGN_KEY environment variable
	inputMinisignKey := os.Getenv("INPUT_MINISIGN_KEY")

	// If any of the env variables are not set, print an error and exit
	if minisignLocation == "" {
		fmt.Println("Error: MINISIGN_PATH is not set")
		os.Exit(1)
	}
	if fileToSign == "" {
		fmt.Println("Error: FILE_TO_SIGN_PATH is not set")
		os.Exit(1)
	}
	if inputPassword == "" {
		fmt.Println("Error: INPUT_PASSWORD is not set")
		os.Exit(1)
	}
	if inputMinisignKey == "" {
		fmt.Println("Error: INPUT_MINISIGN_KEY is not set")
		os.Exit(1)
	}

	// Env variables are set
	fmt.Println("Env variables loaded successfully")

	// Set the minisignKeyPath to the path of the minisign.key file in the .minisign directory in the user's home directory
	minisignKeyPath := filepath.Join(os.Getenv("HOME"), ".minisign", "minisign.key")

	// Get the directory part of minisignKeyPath
	dir := filepath.Dir(minisignKeyPath)
	// Create the directory, including any necessary parents, with mode 0755
	err := os.MkdirAll(dir, 0755)
	// If there's an error, print it and exit
	if err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}
	fmt.Println("Directory created successfully")

	// Write the inputMinisignKey to the minisignKeyPath file with mode 0644
	err = os.WriteFile(minisignKeyPath, []byte(inputMinisignKey), 0644)
	// If there's an error, print it and exit
	if err != nil {
		fmt.Println("Error writing key to file:", err)
		os.Exit(1)
	}
	fmt.Println("Key written to file successfully")

	// Create a new reader with the inputPassword followed by a newline
	r := strings.NewReader(inputPassword + "\n")

	// Create a new command to execute /minisign with the arguments passed to this program
	cmd := exec.Command(minisignLocation, "-Sm", fileToSign)
	// Set the command's stdin to the reader
	cmd.Stdin = io.MultiReader(r)

	// Run the command
	err = cmd.Run()
	// If there's an error, print it and exit
	if err != nil {
		fmt.Println("Error running minisign:", err)
		os.Exit(1)
	}
	fmt.Println("Minisign command executed successfully")
}
