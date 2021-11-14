package add

import (
	"context"
)

type Input struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
}

type Output struct {
	Value float32 `json:"value"`
}

func Do(ctx context.Context, input Input, output *Output) error {
	output.Value = input.A + input.B

	return nil
}
