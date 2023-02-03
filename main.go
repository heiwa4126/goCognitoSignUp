package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"github.com/joho/godotenv"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// AWS SDK for Go v2の設定
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("REGION")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// cognitoIdp Clientの作成
	cognitoIdp := cognitoidentityprovider.NewFromConfig(cfg)

	_, err = cognitoIdp.SignUp(
		context.TODO(),
		&cognitoidentityprovider.SignUpInput{
			ClientId: aws.String(os.Getenv("CLIENT_ID")),
			Password: aws.String(password),
			UserAttributes: []types.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(username),
				},
				{
					Name:  aws.String("given_name"),
					Value: aws.String("g"),
				},
				{
					Name:  aws.String("family_name"),
					Value: aws.String("f"),
				},
			},
			Username: aws.String(username),
		})
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = cognitoIdp.AdminUpdateUserAttributes(
		context.TODO(),
		&cognitoidentityprovider.AdminUpdateUserAttributesInput{
			UserAttributes: []types.AttributeType{
				{
					Name:  aws.String("email_verified"),
					Value: aws.String("true"),
				},
			},
			UserPoolId: aws.String(os.Getenv("USER_POOL_ID")),
			Username:   aws.String(username),
		})
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = cognitoIdp.AdminConfirmSignUp(
		context.TODO(),
		&cognitoidentityprovider.AdminConfirmSignUpInput{
			UserPoolId: aws.String(os.Getenv("USER_POOL_ID")),
			Username:   aws.String(username),
		})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("SignUp, UpdateUserAttributes and ConfirmSignUp success.")
}
