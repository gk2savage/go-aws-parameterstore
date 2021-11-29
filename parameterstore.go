package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	ssmsession, _ := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-south-1")},
		SharedConfigState: session.SharedConfigEnable,
	})

	ssmsvc := ssm.New(ssmsession, aws.NewConfig().WithRegion("ap-south-1"))
	param, _ := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/dev/keywords"),
		WithDecryption: aws.Bool(true),
	})

	keywords := *param.Parameter.Value
	fmt.Println(keywords)
}
