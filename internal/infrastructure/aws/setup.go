package aws

import (
	"api-buddy/infrastructure/aws/cognito"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type AWSClients struct{}

func (c *AWSClients) SetupCognitoClient(region, clientId string) {
	cfg, err := LoadAWSConfig(region)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	cognito.Actions = &cognito.CognitoClient{
		Client:   cognitoidentityprovider.NewFromConfig(cfg),
		ClientId: clientId,
	}
}

func LoadAWSConfig(region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS SDK config, %v", err)
	}
	return cfg, nil
}
