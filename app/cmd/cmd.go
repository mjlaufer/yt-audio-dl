package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mjlaufer/yt-audio-dl/app/yt"
	"github.com/spf13/cobra"
)

// Verbose flag - allows the CLI to provide more information to the user.
var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use: "yt",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a video url argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yt-audio-dl CLI")

		options := yt.Options{
			Verbose: verbose,
		}

		url := args[0]
		if err := yt.Download(url, &options); err != nil {
			fmt.Println("err:", err)
			return
		}
	},
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
