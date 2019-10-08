package cmd

import (
	"errors"
	"log"
	"synergy/helper"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusConfigCmd())
}

func statusConfigCmd() *cobra.Command {
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "status will information from the carbon status file",
		// Args: func(cmd *cobra.Command, args []string) error {
		// 	if IsValidPath(filepath) {
		// 		return nil
		// 	}
		// 	return fmt.Errorf("invalid filepath")
		// },
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here

		},
	}
	statusCmd.AddCommand(listValuesCmd())
	statusCmd.AddCommand(listResourcesCmd())

	return statusCmd
}

func listValuesCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list-environs",
		Short: "list will fetch the current environs related to a resource from the carbon status file",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(project) < 1 {
				return errors.New("project is mandatory")
			}
			if len(bucket) < 1 {
				return errors.New("bucket is mandatory")
			}
			if len(path) < 1 {
				return errors.New("path is mandatory")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			file := helper.CarbonStatusLocalFile
			err := helper.FetchFromRemote(project, bucket, path, file)
			if err != nil {
				log.Fatal(err)
			}
			environsList := helper.ReadStatusResourceEnvirons(resource, file)
			helper.ShowEnvironsOutput(environsList)
		},
	}
	listCmd.Flags().StringVar(&resource, "resource", "", "name of the resource")
	return listCmd
}

func listResourcesCmd() *cobra.Command {
	var listResCmd = &cobra.Command{
		Use:   "list-resources",
		Short: "list will fetch the resources that are present in the status json file",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(project) < 1 {
				return errors.New("project is mandatory")
			}
			if len(bucket) < 1 {
				return errors.New("bucket is mandatory")
			}
			if len(path) < 1 {
				return errors.New("path is mandatory")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// fetching status json file from gcs
			err := helper.FetchFromRemote(project, bucket, path, helper.CarbonStatusLocalFile)
			if err != nil {
				log.Fatal(err)
			}
			// fetching status file data
			statusFileData := helper.GetStatusMap(helper.CarbonStatusLocalFile)
			helper.ShowResourceOutput(statusFileData)
		},
	}
	return listResCmd
}
