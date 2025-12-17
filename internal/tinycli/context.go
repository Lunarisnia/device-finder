package tinycli

import (
	"os"
	"strings"
)

type Context interface {
	Argument(arg string) string
	Arguments() map[string]string
}

type ctx struct {
	formattedArguments map[string]string
}

func newContext() Context {
	args := os.Args
	formattedArguments := make(map[string]string)
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			value := ""
			valueIndex := i + 1
			if valueIndex < len(args) {
				value = args[valueIndex]
			}
			formattedArguments[arg] = value
		}
	}

	return &ctx{
		formattedArguments: formattedArguments,
	}
}

func (a *ctx) Argument(arg string) string {
	return a.formattedArguments[arg]
}

func (a *ctx) Arguments() map[string]string {
	return a.formattedArguments
}
