package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Parameters []Parameter

func (p Parameters) read() {
	url := "https://stparampiper.blob.core.windows.net/"
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Error getting credentials: %v", err)
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Assuming you have a container and blob name
	containerName := "sqlite"
	blobName := "parampiper.json"

	containerClient := client.ServiceClient().NewContainerClient(containerName)
	blobClient := containerClient.NewBlobClient(blobName)

	props, err := blobClient.GetProperties(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting properties: %v", err)
	}

	blobSize := props.ContentLength
	jsonData := make([]byte, int(*blobSize))

	_, err = blobClient.DownloadBuffer(ctx, jsonData, nil)
	if err != nil {
		log.Fatalf("Error downloading blob: %v", err)
	}

	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Print(p)
}
func (p Parameters) save() {
	url := "https://stparampiper.blob.core.windows.net/"
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Println(err)
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		log.Println(err)
	}

	// data := []byte("\nHello, world! This is a blob.\n")
	// blobName := "sample-blob"
	jsonData, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Println(err)
	}

	// Upload to data to blob storage
	// fmt.Printf("Uploading a blob named %s\n", blobName)
	_, err = client.UploadBuffer(ctx, "sqlite", "parampiper.json", jsonData, &azblob.UploadBufferOptions{})
	if err != nil {
		log.Println(err)
	}
}

func (p Parameters) String() string {
	s := ""
	for _, parameter := range p {
		s += parameter.String() + "\n"
	}
	return s
}
