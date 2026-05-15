package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/vhs/command"
	"github.com/spf13/cobra"
)

const (
	// Version is the current version of vhs.
	Version = "0.1.0"

	// DefaultPort is the default port for the vhs server.
	// Changed from 1976 to 8080 to avoid conflicts with other local services.
	DefaultPort = 8080
)

var rootCmd = &cobra.Command{
	Use:     "vhs <file>",
	Short:   "Run a VHS tape (a series of commands) to create a GIF or video.",
	Long:    `VHS is a tool for recording terminal GIFs and videos using a simple scripting language.`,
	Version: Version,
	Args:    cobra.MaximumNArgs(1),
	RunE:    command.Run,
}

func init() {
	rootCmd.AddCommand(command.NewServeCmd(DefaultPort))
	rootCmd.AddCommand(command.NewRecordCmd())
	rootCmd.AddCommand(command.NewNewCmd())

	rootCmd.Flags().StringP("output", "o", "", "output file path (e.g. out.gif, out.mp4, out.webm)")
	rootCmd.Flags().StringP("publish", "p", "", "publish output to charm cloud")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode (no output)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
