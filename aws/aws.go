package aws

import (
	"errors"
	"github.com/razorpay/db-poc/utils"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/razorpay/db-poc/constants"
)

func GetSession() (*session.Session, error) {
	var (
		awsSession   *session.Session
		sessionError error
	)

	awsKeyID := os.Getenv(constants.AwsAccessKeyID)
	awsAccessKey := os.Getenv(constants.AwsSecretAccessKey)

	if utils.IsEmpty(awsKeyID) || utils.IsEmpty(awsAccessKey) {
		return nil, errors.New("No AWS credentials.")
	}

	creds := credentials.NewStaticCredentials(awsKeyID, awsAccessKey, "")

	awsSession, sessionError = session.NewSession(
		&aws.Config{Region: aws.String(constants.AwsRegion), Credentials: creds})

	return awsSession, sessionError
}

func String(val string) *string {
	return aws.String(val)
}

func Bool(val bool) *bool {
	return aws.Bool(val)
}
