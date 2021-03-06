package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sourceRegion = os.Getenv("SOURCE_REGION")
	targetRegion = os.Getenv("TARGET_REGION")
	eventJSON    map[string]interface{}
)

func main() {
	log.Print("Calling the handler...")
	lambda.Start(handler)
}

func handler(event events.CloudWatchEvent) {
	log.Printf("Processing Lambda request %s\n", event.ID)

	err := json.Unmarshal(event.Detail, &eventJSON)
	if err != nil {
		log.Fatal("Could not unmarshal scheduled event: ", err)
	}

	log.Print("Getting the event detail...")
	detail := eventJSON["requestParameters"].(map[string]interface{})

	log.Print("Getting the SecretId...")
	secretID := detail["secretId"].(string)

	log.Printf("About to get secret value for %s\n", secretID)
	response, _ := getNewSecret(secretID)
	secretValue := *response.SecretString
	log.Printf("Retrived the secret value for %s\n", secretID)

	log.Printf("Updating the secret value in %s\n", targetRegion)
	resp, err := updateSecretValue(secretValue, secretID)
	if err != nil {
		log.Fatalf("There was an issue updating the secret value.")
	}

	log.Printf("SecretValue updated for %s", *resp.Name)
}

func getNewSecret(secretID string) (*secretsmanager.GetSecretValueOutput, error) {

	log.Print("Creating session for getNewSecret...")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(sourceRegion)},
	)
	svc := secretsmanager.New(awsSession)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Fatalf("Could not GetSecretValue for: %s\n,%s\n", secretID, err)
	}

	return result, nil
}

func updateSecretValue(secret string, secretID string) (*secretsmanager.PutSecretValueOutput, error) {

	log.Print("Creating session for updateSecretValue...")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(targetRegion)},
	)
	svc := secretsmanager.New(awsSession)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(secretID),
		SecretString: aws.String(secret),
	}

	result, err := svc.PutSecretValue(input)
	if err != nil {
		log.Fatalf("Could not PutSecretValue: %s\n", err)
	}

	return result, nil
}
