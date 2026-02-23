package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	ctx := context.Background()
	client := openai.NewClient()

	response, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("今日のニュースを教えて?"),
		},
		Model: openai.ChatModelGPT5_2,
		// Web Search Tool を使用するためのツールパラメータを指定
		Tools: []responses.ToolUnionParam{
			{
				OfWebSearch: &responses.WebSearchToolParam{
					Type:              responses.WebSearchToolTypeWebSearch,
					SearchContextSize: responses.WebSearchToolSearchContextSizeMedium,
					UserLocation: responses.WebSearchToolUserLocationParam{
						Type:     "approximate",
						Country:  openai.String("JP"),
						Region:   openai.String("Shizuoka"),
						City:     openai.String("Kannami"),
						Timezone: openai.String("Asia/Tokyo"),
					},
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response.OutputText())
}
