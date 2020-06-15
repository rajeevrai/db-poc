package bootstrap

import (
	"github.com/razorpay/db-poc/constants"
	"github.com/razorpay/db-poc/data_generator"
	"github.com/razorpay/db-poc/dynamodb"
)

func Initialize(dbType string) {
	data_generator.GenerateMerchantIds()

	switch dbType {
	case constants.DynamoDbFlag:
		dynamodb.Benchmark(1)
	}
}
