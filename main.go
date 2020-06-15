package main

import (
	"flag"

	"github.com/razorpay/db-poc/bootstrap"
	"github.com/razorpay/db-poc/constants"
)

func main() {
	dbType := flag.String(constants.DbTypeFlag, constants.DynamoDbFlag, "db type")
	flag.Parse()

	bootstrap.Initialize(*dbType)
}
