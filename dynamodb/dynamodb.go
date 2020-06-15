package dynamodb

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/razorpay/db-poc/aws"
	"github.com/razorpay/db-poc/data_generator"
	"github.com/razorpay/db-poc/logger"
)

const (
	NumberOfRecordsPerScript = 500000

	TestTableName = "poc_payments"

	LogPath = "/go/src/github.com/razorpay/db-poc/logs"
)

func getConnection() (*dynamodb.DynamoDB, bool) {
	awsSession, sessionError := aws.GetSession()

	if sessionError != nil {
		logger.Get().Error(fmt.Sprintf("Could not connect to AWS."))

		return nil, false
	}

	return dynamodb.New(awsSession), true
}

func pushData(awsService *dynamodb.DynamoDB, rowData map[string]interface{}) (string, bool) {
	paymentId := rowData["id"].(string)
	info, err := dynamodbattribute.MarshalMap(rowData)

	if err != nil {
		logger.Get().Error(fmt.Sprintf("Error: %s", err.Error()))
		return "", false
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(TestTableName),
	}

	_, err = awsService.PutItem(input)

	if err != nil {
		logger.Get().Error(fmt.Sprintf("Error: %s", err.Error()))
		return "", false
	}

	return paymentId, true
}

func getData(awsService *dynamodb.DynamoDB, paymentId string, consistent bool) bool {
	_, err := awsService.GetItem(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(consistent),
		TableName: aws.String(TestTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(paymentId),
			},
		},
	})

	if err != nil {
		return false
	}

	return true
}

func Benchmark(routineId int) bool {
	awsService, isConnected := getConnection()
	if !isConnected {
		return false
	}

	var logData string
	var tWriteStart, tWriteEnd, tWriteTotal int64
	var tcReadStart, tcReadEnd, tcReadTotal int64
	var tncReadStart, tncReadEnd, tncReadTotal int64

	for i := 0; i < NumberOfRecordsPerScript; i++ {
		rowData := data_generator.GetRowData()

		tWriteStart = getTimestamp()
		paymentId, ifPushed := pushData(awsService, rowData)
		tWriteEnd = getTimestamp()

		logData = fmt.Sprintf("write %d %s\n", tWriteEnd-tWriteStart, strconv.FormatBool(ifPushed))
		fileWrite(logData, routineId)

		tcReadStart = getTimestamp()
		ifRead := getData(awsService, paymentId, true)
		tcReadEnd = getTimestamp()

		logData = fmt.Sprintf("tcread %d %s\n", tcReadEnd-tcReadStart, strconv.FormatBool(ifRead))
		fileWrite(logData, routineId)

		tncReadStart = getTimestamp()
		ifRead = getData(awsService, paymentId, false)
		tncReadEnd = getTimestamp()

		logData = fmt.Sprintf("tncread %d %s\n", tncReadEnd-tncReadStart, strconv.FormatBool(ifRead))
		fileWrite(logData, routineId)

		tWriteTotal += tWriteEnd - tWriteStart
		tcReadTotal += tcReadEnd - tcReadStart
		tncReadTotal += tncReadEnd - tncReadStart
	}

	logData = fmt.Sprintf("write_total %d\n", tWriteTotal)
	fileWrite(logData, routineId)

	logData = fmt.Sprintf("tcread_total %d\n", tcReadTotal)
	fileWrite(logData, routineId)

	logData = fmt.Sprintf("tncread_total %d\n", tncReadTotal)
	fileWrite(logData, routineId)

	return true
}

func fileWrite(data string, routineId int) {
	file := fmt.Sprintf("%s/data-%d.txt", LogPath, routineId)

	_, err := os.Stat(file)
	if err != nil {
		fpCreate, errCreate := os.Create(file)

		if errCreate != nil {
			return
		}

		fpCreate.Close()
	}

	fpOpen, errOpen := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	if errOpen != nil {
		return
	}

	defer fpOpen.Close()
	_, _ = fpOpen.WriteString(data)
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
