package cmd

import (
	"github.com/aakashshankar/claude-cli/api"
	"github.com/spf13/cobra"
)

var stream bool
var tokens int
var model string
var system string
var clr bool
var askCmd = &cobra.Command{
	Use:   "ask <prompt>",
	Short: "Prompt claude",
	Long:  `This command prompts the Claude API with the given prompt and returns the response`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api.PromptClaude(args[0], stream, tokens, model, system, clr)
	},
}

func init() {
	askCmd.Flags().StringVarP(&model, "model", "m", "claude-3-sonnet-20240229", "Model that responds to the prompt")
	askCmd.Flags().BoolVarP(&stream, "stream", "s", false, "Stream the response")
	askCmd.Flags().IntVarP(&tokens, "tokens", "t", 1024, "Maximum number of tokens preferred in the output")
	askCmd.Flags().StringVarP(&system, "system", "S", "You are a helpful assistant", "Set the system prompt")
	askCmd.Flags().BoolVarP(&clr, "clear", "c", false, "Clear the current session and start a new one")
	rootCmd.AddCommand(askCmd)
}
