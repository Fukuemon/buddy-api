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

// SignUpRequest struct to encapsulate the request parameters
type CognitoSignUpRequest struct {
	Username    string
	Password    string
	Email       *string
	PhoneNumber *string
	// Add other fields as necessary
}

func (c *CognitoClient) SignUp(req *CognitoSignUpRequest) (*string, error) {
	attributes := []types.AttributeType{}

	if req.Email != nil {
		attributes = append(attributes, types.AttributeType{
			Name:  aws.String("email"),
			Value: aws.String(*req.Email),
		})
	}

	if req.PhoneNumber != nil {
		attributes = append(attributes, types.AttributeType{
			Name:  aws.String("phone_number"),
			Value: aws.String(*req.PhoneNumber),
		})
	}

	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId:       aws.String(c.ClientId),
		Username:       aws.String(req.Username),
		Password:       aws.String(req.Password),
		UserAttributes: attributes,
	}

	output, err := c.Client.SignUp(context.TODO(), signUpInput)
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't sign up user: %v", err)
		}
		return nil, err
	}

	return output.UserSub, nil
}

// SignIn function for logging in a user
func (c *CognitoClient) SignIn(username, password string) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(c.ClientId),
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
	}

	authOutput, err := c.Client.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		log.Printf("Failed to authenticate user: %v", err)
		return nil, err
	}

	return authOutput, nil
}

// ListUsers function to get a list of users in the user pool
func (c *CognitoClient) ListUsers(userPoolId string) ([]types.UserType, error) {
	listUsersInput := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(userPoolId),
	}

	output, err := c.Client.ListUsers(context.TODO(), listUsersInput)
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		return nil, err
	}

	return output.Users, nil
}
