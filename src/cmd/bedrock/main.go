package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const (
	format  = "\n\nHuman:%s\n\nAssistant:"
	modelID = "anthropic.claude-v2"
	prompt  = "Using the article located at https://medium.com/@spei/ai-without-machine-learning-47e90e5ae7c5, create a summary with maximum 150 words and add 3 hashtags."
)

type Request struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"top_p,omitempty"`
	TopK              int      `json:"top_k,omitempty"`
	StopSequences     []string `json:"stop_sequences,omitempty"`
}

type Response struct {
	Completion string `json:"completion"`
}

func main() {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bedrock := bedrockruntime.NewFromConfig(cfg)

	payload := &Request{
		Prompt:            fmt.Sprintf(format, prompt),
		MaxTokensToSample: 500,
		Temperature:       0.1,
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	out, err := bedrock.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		Body:        bytes,
		ModelId:     aws.String(modelID),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	if err := json.Unmarshal(out.Body, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response from Claude:\n", response.Completion)
}
