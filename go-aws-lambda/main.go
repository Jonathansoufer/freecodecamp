package main

/**
* This simple example works with AWS Lambda, and for achieving this, we need to:
* 1. Build the binary for the Lambda function.
* 2. Create a Lambda function.
* 3. Upload the binary to the Lambda function (compressed).
* 4. Test the Lambda function.
**/

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main(){
	lambda.Start(HandleLambdaEvent)
}

