package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-south-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("ap-south-1"))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("keywords"),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		panic(err)
	}
  
	value := *param.Parameter.Value
	fmt.Println("Value of Parameter Store: ", value)
  
  //JSON Example-> Value of Parameter Store: {"primary": ["anime", "games", "music"], "secondary": ["study", "eat", "drink"]}
  
	keywordmap := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(value), &keywordmap)

	if err1 != nil {
		panic(err1)
	}

	// for key, value := range keywordmap {
	// 	fmt.Println("index : ", key, " value : ", value)
	// }

	// fmt.Println(keywordmap["primary"])
	keywordString := fmt.Sprintf("%v", keywordmap["primary"])

	// fmt.Println(keywordmap["secondary"])
	keywordString2 := fmt.Sprintf("%v", keywordmap["secondary"])

	keywordString = strings.TrimSuffix(keywordString, "]")
	keywordString = strings.TrimPrefix(keywordString, "[")

	keywordString2 = strings.TrimSuffix(keywordString2, "]")
	keywordString2 = strings.TrimPrefix(keywordString2, "[")

	keywordlist := strings.Split(keywordString, " ")
	keywordlist2 := strings.Split(keywordString2, " ")

	keywords := make([][]string, 2)
	keywords[0] = keywordlist
	keywords[1] = keywordlist2

	fmt.Println(keywords)

	fmt.Println("Am I going to do this Task? Answer: ", PriorityCategorize("i am going to listen to some cool music", keywords))

}

func PriorityCategorize(str string, keywords [][]string) bool {

	for i := range keywords[0] {
		if strings.Contains(str, keywords[0][i]) {
			fmt.Println(keywords[0][i])
			return true
		}
	}

	matches := 0

	for i := range keywords[1] {
		if strings.Contains(str, keywords[1][i]) {
			matches += 1
			if matches >= 2 {
				return true
			}
		}
	}

	return matches > 1
}
