package data

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type AzureStorageAccount struct {
	StorageAccountName string
	ContainerName      string
	FileName           string
	Parameters         Parameters
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
	blobClient := containerClient.NewBlobClient(p.FileName)

	props, err := blobClient.GetProperties(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting properties: %v", err)
	}

	blobSize := props.ContentLength
	jsonData := make([]byte, int(*blobSize))
	log.Println("[DEBUG] Blob size:", blobSize)

	_, err = blobClient.DownloadBuffer(ctx, jsonData, nil)
	if err != nil {
		log.Fatalf("Error downloading blob: %v", err)
	}

	err = json.Unmarshal(jsonData, &p.Parameters)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	log.Println("[DEBUG]", p.Parameters)
	return p.Parameters, nil
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

	_, err = client.UploadBuffer(ctx, p.ContainerName, p.FileName, jsonData, &azblob.UploadBufferOptions{})
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (p AzureStorageAccount) String() string {
	s := ""
	for _, parameter := range p.Parameters {
		s += parameter.String() + "\n"
	}
	return s
}
