package main

import (
	"fmt"
)

func main() {
	// Example AWS credentials (these are not real keys!)
	const awsAccessKeyID = "AKIAIOSFODNN7EXAMPLE" // Detected by GitHub
	const awsSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" // Detected by GitHub

	fmt.Println("AWS Access Key ID:", awsAccessKeyID)
	fmt.Println("AWS Secret Access Key:", awsSecretAccessKey)
}
