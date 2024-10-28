package main

import (
	"api-buddy/config"
	"api-buddy/infrastructure/aws"
	"api-buddy/infrastructure/aws/cognito"
	"fmt"
	"log"
)

func main() {
	// 設定をロード (例: 設定ファイルや環境変数から)
	conf := config.GetConfig()

	// Cognitoクライアントのセットアップ
	awsClients := &aws.AWSClients{}
	awsClients.SetupCognitoClient(conf.AWSConfig.Region, conf.CognitoConfig.ClientId)

	// ユーザー一覧の取得
	users, err := cognito.Actions.ListUsers(conf.CognitoConfig.UserPoolId)
	if err != nil {
		log.Fatalf("Failed to retrieve users: %v", err)
	}

	// ユーザー一覧を表示
	fmt.Println("Cognito Users:")
	for _, user := range users {
		fmt.Printf("Username: %s, Status: %s\n", *user.Username, user.UserStatus)

		// ユーザーの属性を表示
		fmt.Println("Attributes:")
		for _, attr := range user.Attributes {
			fmt.Printf("  %s: %s\n", *attr.Name, *attr.Value)
		}
		fmt.Println()
	}
}
