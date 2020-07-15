package main

import (
	"strings"

	"github.com/neovim/go-client/nvim/plugin"
)

// original code: https://github.com/neovim/go-client/tree/acbae1ac0e05229056569fdc54285da8aa88f1b4

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "Hello"}, hello)
		return nil
	})
}

func hello(args []string) (string, error) {
	return "Hello" + strings.Join(args, ""), nil
}
