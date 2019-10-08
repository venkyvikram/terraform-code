package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var project, bucket, path string

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "synergy",
	Short: "An example of synergy",
	Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "project name where the bucket of status file resides (default is none)")
	rootCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "status file bucket name (default is none)")
	rootCmd.PersistentFlags().StringVar(&path, "path", "", "object path for the status file (default is none)")
}
