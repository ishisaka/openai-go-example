package main

import (
	"context"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	// コンテキストを作成
	ctx := context.Background()

	// OpenAI クライアントを初期化
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	// AIに与える質問
	question := "Write me a haiku about computers"

	// OpenAI APIにリクエストを送信
	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: openai.ChatModelGPT5_2,
	})

	// エラー処理
	if err != nil {
		panic(err)
	}

	// レスポンスの出力テキストを表示
	println(resp.OutputText())
}
