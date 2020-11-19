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
	eventJson    map[string]interface{}
)

func main() {
	lambda.Start(Handler)
}

func Handler(event events.CloudWatchEvent) {

	log.Printf("Processing Lambda request %s\n", event.ID)

	err := json.Unmarshal(event.Detail, &eventJson)
	if err != nil {
		log.Fatal("Could not unmarshal scheduled event: ", err)
	}

	detail := eventJson["additionalEventData"].(map[string]interface{})

	secretArn := detail["SecretId"].(string)

	log.Printf("About to get secret value for %s\n", secretArn)
	response, _ := getNewSecret(secretArn)

	log.Printf("Retrived secret value for %s\n", secretArn)
	log.Printf("Updating secret value in %s\n", targetRegion)
	secretValue := *response.SecretString
	secretVersionID := *response.VersionId
	resp, err := updateSecretValue(secretValue, secretVersionID)
	if err != nil {
		log.Fatalf("There was an issue updating the secret value.")
	}

	log.Printf("SecretValue updated for %s\n", resp.ARN)
}

func getNewSecret(secretArn string) (*secretsmanager.GetSecretValueOutput, error) {

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(sourceRegion)},
	)
	svc := secretsmanager.New(awsSession)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretArn),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Fatalf("Could not GetSecretValue for: %s\n,%s\n", secretArn, err)
	}

	return result, nil
}

func updateSecretValue(secret string, versionID string) (*secretsmanager.PutSecretValueOutput, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(targetRegion)},
	)
	svc := secretsmanager.New(awsSession)
	input := &secretsmanager.PutSecretValueInput{
		SecretString: aws.String(secret),
	}

	result, err := svc.PutSecretValue(input)
	if err != nil {
		log.Fatalf("Could not PutSecretValue: %s\n", err)
	}

	return result, nil
}
