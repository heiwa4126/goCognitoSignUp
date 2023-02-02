package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/joho/godotenv"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}

	cognitoIdp := cognitoidentityprovider.New(sess)

	signupInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("CLIENT_ID")),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
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
	}

	signupOutput, err := cognitoIdp.SignUp(signupInput)
	if err != nil {
		fmt.Println("Error signing up:", err)
		os.Exit(1)
	}

	updateInput := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
		UserPoolId: aws.String(os.Getenv("USER_POOL_ID")),
		Username:   aws.String(username),
	}

	updateOutput, err := cognitoIdp.AdminUpdateUserAttributes(updateInput)
	if err != nil {
		fmt.Println("Error updating user attributes:", err)
		os.Exit(1)
	}

	fmt.Println("Signup Output:", signupOutput)
	fmt.Println("Update Output:", updateOutput)
}
