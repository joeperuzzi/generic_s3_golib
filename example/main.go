package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/joeperuzzi/generic_s3_golib"
)

func main() {
	cfg := s3client.Config{
		Endpoint:       "https://s3.amazonaws.com",
		Region:         "us-east-1",
		AccessKey:      "your-access-key",
		SecretKey:      "your-secret-key",
		Bucket:         "your-bucket-name",
		ForcePathStyle: false,
		DisableSSL:     false,
	}

	client, err := s3client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	content := strings.NewReader("Hello, S3!")
	err = client.Upload(ctx, "test-file.txt", content, "text/plain")
	if err != nil {
		log.Printf("Upload failed: %v", err)
	} else {
		fmt.Println("File uploaded successfully!")
	}

	exists, err := client.Exists(ctx, "test-file.txt")
	if err != nil {
		log.Printf("Exists check failed: %v", err)
	} else {
		fmt.Printf("File exists: %v\n", exists)
	}

	body, err := client.Download(ctx, "test-file.txt")
	if err != nil {
		log.Printf("Download failed: %v", err)
	} else {
		defer body.Close()
		fmt.Println("File downloaded successfully!")
	}

	files, err := client.List(ctx, "")
	if err != nil {
		log.Printf("List failed: %v", err)
	} else {
		fmt.Printf("Found %d files:\n", len(files))
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	err = client.Delete(ctx, "test-file.txt")
	if err != nil {
		log.Printf("Delete failed: %v", err)
	} else {
		fmt.Println("File deleted successfully!")
	}
}