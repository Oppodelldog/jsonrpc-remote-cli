package split

import (
	"context"
	"strings"
)

type Input struct {
	Data      string `json:"data"`
	Delimiter string `json:"delim"`
}

type Output struct {
	Value []string `json:"value"`
}

func Do(ctx context.Context, input Input, output *Output) error {
	output.Value = strings.Split(input.Data, input.Delimiter)

	return nil
}
