package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/PullRequestInc/go-gpt3"
)


func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == ""{
		log.Fatalln("Missing API KEY")
	}


	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	resp ,err :=client.Completion(ctx, gpt3.CompletionRequest{
		Prompt: []string{"One piece the anime is an "},
		MaxTokens: gpt3.IntPtr(30),
		Stop:[]string{"."},
		Echo:true,

	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}