package usecases

import (
	"context"
	"github.com/Oppodelldog/jsonrpc-remote-cli/usecases/add"
	"github.com/Oppodelldog/jsonrpc-remote-cli/usecases/split"
	"github.com/swaggest/usecase"
)

func Interactors() []usecase.Interactor {
	var u1 = usecase.NewIOI(new(split.Input), new(split.Output), func(ctx context.Context, input, output interface{}) error {
		var in = input.(*split.Input)
		var out = output.(*split.Output)

		return split.Do(ctx, *in, out)
	})
	u1.SetName("split")
	u1.SetDescription("splits a string with the given delimiter")

	var u2 = usecase.NewIOI(new(add.Input), new(add.Output), func(ctx context.Context, input, output interface{}) error {
		var in = input.(*add.Input)
		var out = output.(*add.Output)

		return add.Do(ctx, *in, out)
	})
	u2.SetName("add")
	u2.SetDescription("adds two float values")

	return []usecase.Interactor{u1, u2}
}
