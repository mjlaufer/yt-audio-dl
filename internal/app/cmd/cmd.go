package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mjlaufer/yt-to-mp3/internal/app/yt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "yt",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a video url argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yt-to-mp3 CLI")

		url := args[0]
		if err := yt.Download(url); err != nil {
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
