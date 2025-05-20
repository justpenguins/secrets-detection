package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)


// Rotate rotates an AWS access key by creating a new one and deleting the old one
func Rotate(oldKey string) {

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}

	client := iam.NewFromConfig(cfg)

	// Retrieve the username associated with the old key
	userName, err := getUserNameForKey(client, oldKey)
	if err != nil {
		fmt.Println("Error retrieving username for key:", err)
		return
	}

	// Create a new access key
	newKey, err := client.CreateAccessKey(context.TODO(), &iam.CreateAccessKeyInput{
		UserName: &userName,
	})
	if err != nil {
		fmt.Println("Error creating new access key:", err)
		return
	}

	fmt.Printf("New access key created: %s\n", *newKey.AccessKey.AccessKeyId)

	// Delete the old access key
	_, err = client.DeleteAccessKey(context.TODO(), &iam.DeleteAccessKeyInput{
		UserName:    &userName,
		AccessKeyId: &oldKey,
	})
	if err != nil {
		fmt.Println("Error deleting old access key:", err)
		return
	}

	fmt.Println("Old access key deleted successfully.")
}

// getUserNameForKey retrieves the username associated with the given access key
func getUserNameForKey(client *iam.Client, accessKey string) (string, error) {
	output, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{})
	if err != nil {
		return "", err
	}

	for _, key := range output.AccessKeyMetadata {
		if *key.AccessKeyId == accessKey {
			return *key.UserName, nil
		}
	}

	return "", fmt.Errorf("User not found")
}