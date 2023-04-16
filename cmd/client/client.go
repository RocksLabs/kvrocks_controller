package main

import (
	"github.com/c-bata/go-prompt"
)

func main() {
	// TODO: Parse endpoint from command line arguments
	promptCtx := NewPromptContext()
	request := NewRequest("")
	completer := NewCompleter(promptCtx, request)
	executor := NewExecutor(promptCtx, request)

	for {
		input := prompt.Input(promptCtx.Prefix(), completer.Complete)
		if quit := executor.Execute(input); quit {
			break
		}
	}
}
