package data

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
)

type AzureStorageAccount struct {
	StorageAccountName string
	ContainerName      string
	BlobName           string
}

func (p AzureStorageAccount) Read() (Parameters, error) {
	url := "https://" + p.StorageAccountName + ".blob.core.windows.net/"
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Error getting credentials: %v", err)
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	containerClient := client.ServiceClient().NewContainerClient(p.ContainerName)
	blobClient := containerClient.NewBlobClient(p.BlobName)

	log.Println("[DEBUG]", "SA Name:", p.StorageAccountName, "\nContainer Name:", p.ContainerName, "\nFileName:", p.BlobName)
	props, err := blobClient.GetProperties(ctx, nil)
	if bloberror.HasCode(err, bloberror.BlobNotFound) {
		log.Println("[DEBUG] Blob not found:", err)
		_, err = client.UploadBuffer(ctx, p.ContainerName, p.BlobName, []byte("[]"), &azblob.UploadBufferOptions{})
		if err != nil {
			log.Fatalf("Error on creating empty blob: %v", err)
		}

		props, err = blobClient.GetProperties(ctx, nil)
		if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Fatalf("Error getting blob properties: %v", err)

	}

	blobSize := props.ContentLength
	jsonData := make([]byte, int(*blobSize))
	log.Println("[DEBUG] Blob size:", blobSize)

	_, err = blobClient.DownloadBuffer(ctx, jsonData, nil)
	if err != nil {
		log.Fatalf("Error downloading blob: %v", err)
	}

	var parameters Parameters
	err = json.Unmarshal(jsonData, &parameters)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	log.Println("[DEBUG]", parameters)
	return parameters, nil
}

func (p AzureStorageAccount) Save(params Parameters) error {
	url := "https://" + p.StorageAccountName + ".blob.core.windows.net/"
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Println(err)
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		log.Println(err)
	}

	jsonData, err := json.MarshalIndent(params, "", "    ")
	if err != nil {
		log.Println(err)
	}

	_, err = client.UploadBuffer(ctx, p.ContainerName, p.BlobName, jsonData, &azblob.UploadBufferOptions{})
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (p AzureStorageAccount) String() string {
	s := ""
	parameters, err := p.Read()
	if err != nil {
		log.Println(err)
	}

	for _, parameter := range parameters {
		s += parameter.String() + "\n"
	}
	return s
}
