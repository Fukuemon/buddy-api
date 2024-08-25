package cognito

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var Actions *CognitoClient

type CognitoClient struct {
	Client   *cognitoidentityprovider.Client
	ClientId string
}

