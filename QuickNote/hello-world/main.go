package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	db = dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:aws.String("ap-south-1"),
	})))
	tableName = "NotesTable"
)

type Note struct{
	NoteID string `json:"noteId"`
	Content string `json:"content"`
}

func handler(ctx context.Context,request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	switch request.HTTPMethod{
	case "GET":
		return getNoteHandler(request)
	default:
		return events.APIGatewayProxyResponse{
		Body:       "Method not allowed",
		StatusCode: 200,
	}, nil
	}
}

func  getNoteHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	noteID := request.QueryStringParameters["noteId"] 
	result,err := db.GetItem(&dynamodb.GetItemInput{
		TableName:aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"NoteID": {
				S: aws.String(noteID),
			},
		},
	})
	
	
}

func main() {
	lambda.Start(handler)
}
