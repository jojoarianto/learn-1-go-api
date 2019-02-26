package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	repositories "github.com/jojoarianto/learn-1-go-api/src/repositories/postgres"
	"github.com/jojoarianto/learn-1-go-api/src/usecases"
)

var (
	// PGConnection get connection from env
	PGConnection = os.Getenv("PG_CONNECTION")
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	userID, _ := strconv.Atoi(request.PathParameters["userid"])

	db, err := gorm.Open("postgres", PGConnection)
	if err != nil {
		return Response{StatusCode: 400}, err
	}
	defer db.Close()

	repo := usecases.NewGetUserById(repositories.NewUserRepository(db))
	data, _ := repo.GetUserById(userID)

	body, err := json.Marshal(map[string]interface{}{
		"data": data,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Content-Type":                     "application/json",
			"X-MyCompany-Func-Reply":           "get-all-user-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
