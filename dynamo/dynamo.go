package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Repo type
type Repo struct {
	session *session.Session
}

// NewDynamoRepo returns a new dynamo repo
func NewDynamoRepo(session *session.Session) *Repo {
	return &Repo{session}
}

// Create adds an item to dynamodb
func (r Repo) Create(item Item) error {
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-1")},
	// )
	// sess := session.Must(session.NewSession())
	svc := dynamodb.New(r.session)
	data, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String("articles"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("[ERROR] error calling PutItem:")
		fmt.Println(err.Error())
	}
	// r.svc.Put
	return err
}
