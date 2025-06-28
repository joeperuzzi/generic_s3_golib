package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	s3client "github.com/josephperuzzi/generic_s3_golib"
)

func main() {
	cfg := s3client.Config{
		Endpoint:       "https://[your-id].r2.cloudflarestorage.com",
		Region:         "auto",
		AccessKey:      "",
		SecretKey:      "",
		Bucket:         "bardshare",
		ForcePathStyle: false,
		DisableSSL:     false,
	}

	client, err := s3client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	content := strings.NewReader("Hello from Cloudflare R2!")
	err = client.Upload(ctx, "r2-test-file.txt", content, "text/plain")
	if err != nil {
		log.Printf("Upload failed: %v", err)
	} else {
		fmt.Println("File uploaded to Cloudflare R2 successfully!")
	}

	exists, err := client.Exists(ctx, "r2-test-file.txt")
	if err != nil {
		log.Printf("Exists check failed: %v", err)
	} else {
		fmt.Printf("File exists in R2: %v\n", exists)
	}

	body, err := client.Download(ctx, "r2-test-file.txt")
	if err != nil {
		log.Printf("Download failed: %v", err)
	} else {
		defer body.Close()
		fmt.Println("File downloaded from R2 successfully!")
	}

	files, err := client.List(ctx, "")
	if err != nil {
		log.Printf("List failed: %v", err)
	} else {
		fmt.Printf("Found %d files in R2 bucket:\n", len(files))
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	err = client.Delete(ctx, "r2-test-file.txt")
	if err != nil {
		log.Printf("Delete failed: %v", err)
	} else {
		fmt.Println("File deleted from R2 successfully!")
	}
}
