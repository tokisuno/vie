package cmd

import (
	// "fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
    cfgFile string
    userLicense string
)
var rootCmd = &cobra.Command{
	Use:   "vie",
	Short: "Journal",
	Long: "Journalling cli tool",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.PersistentFlags().Bool("viper", true, "user Viper for configuration")
}

