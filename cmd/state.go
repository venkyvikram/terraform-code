package cmd

import (
	"errors"
	"log"
	"synergy/helper"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stateConfigCmd())
}

var resource, environ string

// const statefilelocationStr string = "state.tfstate"

func stateConfigCmd() *cobra.Command {
	var stateCmd = &cobra.Command{
		Use:   "state",
		Short: "state will fetch the current active and inactive clusters",
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
	stateCmd.AddCommand(readValuesCmd())

	return stateCmd
}

func readValuesCmd() *cobra.Command {
	var readCmd = &cobra.Command{
		Use:   "read",
		Short: "read will fetch the values from state",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(resource) < 1 {
				// TODO check on resource name validation
				return errors.New("resource is mandatory")
			}
			if environ != "dev" &&
				environ != "stage" &&
				environ != "prod" &&
				environ != "adm" &&
				environ != "predev" &&
				environ != "sandbox" {
				return errors.New("environ must be adm/dev/stage/prod/predev/sandbox")
			}
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
			statefileLocationObj := helper.ReadStatusFile(helper.CarbonStatusLocalFile)
			carbonstatefileLocation, err := helper.ReadStatusFileValues(resource, environ, statefileLocationObj)
			if err != nil {
				log.Fatal(err)
			}
			if resource == helper.ApplicationCluster {
				clusters := readApplicationCluster(carbonstatefileLocation)
				helper.ShowClusterOutput(clusters)

			} else if resource == helper.AdminCluster {
				error := helper.FetchFromRemote(carbonstatefileLocation.Project, carbonstatefileLocation.Bucket, carbonstatefileLocation.Path, helper.CarbonStateLocalFile)
				if error != nil {
					log.Fatal(error)
				}
				// creating a cluster object
				clusterObj := helper.Cluster{}

				// fetching values form the admin cluster state file
				readAdminClusterName := helper.ReadStateValues(helper.CarbonStateLocalFile, helper.AdminClusterNameKey)
				readAdminClusterRegion := helper.ReadStateValues(helper.CarbonStateLocalFile, helper.AdminClusterRegionKey)
				readAdminClusterProject := helper.ReadStateValues(helper.CarbonStateLocalFile, helper.AdminClusterProjectKey)
				clusterObj.Name = readAdminClusterName
				clusterObj.Region = readAdminClusterRegion
				clusterObj.Project = readAdminClusterProject

				// displaying the output
				var clusters []helper.Cluster
				clusters = append(clusters, clusterObj)
				helper.ShowClusterOutput(clusters)
			} else if resource == helper.BackupCluster {
				error := helper.FetchFromRemote(carbonstatefileLocation.Project, carbonstatefileLocation.Bucket, carbonstatefileLocation.Path, helper.CarbonStateLocalFile)
				if error != nil {
					log.Fatal(error)
				}
				clusterObj := readApplicationCluster(carbonstatefileLocation)
				helper.ShowBackupBucketOutput(carbonstatefileLocation, clusterObj)
			} else if resource == helper.VaultConnect {
				clusterObj := readApplicationCluster(carbonstatefileLocation)
				helper.ShowVaultClusterOutput(clusterObj)
			} else if resource == helper.ConsulConnect {
				clusterObj := readApplicationCluster(carbonstatefileLocation)
				helper.ShowVaultClusterOutput(clusterObj)
			} else if resource == helper.RedisSpin {
				error := helper.FetchFromRemote(carbonstatefileLocation.Project, carbonstatefileLocation.Bucket, carbonstatefileLocation.Path, helper.CarbonStateLocalFile)
				if error != nil {
					log.Fatal(error)
				}
				// fetching admin cluster project from the state file
				redisIP := helper.ReadStateValues(helper.CarbonStateLocalFile, "carbon_redis_internal_ip")
				helper.ShowRedisOutput(redisIP)
			}
		},
	}
	readCmd.Flags().StringVar(&resource, "resource", "", "name of the resource")
	readCmd.Flags().StringVar(&environ, "environ", "", "environ of the resource")
	return readCmd
}

func readApplicationCluster(carbonstatefileLocation helper.StatefileLocation) []helper.Cluster {
	// fetching terraform state file
	error := helper.FetchFromRemote(carbonstatefileLocation.Project, carbonstatefileLocation.Bucket, carbonstatefileLocation.Path, helper.CarbonStateLocalFile)
	if error != nil {
		log.Fatal(error)
	}
	var clusters []helper.Cluster
	nameStrings, regionStrings, projectStrings := helper.GetClusterKeys(helper.CarbonStateLocalFile)
	// below array needs to be dynamically picked up from the statefile
	// creating an object of type Cluster

	// reading cluster name and region values from terrafrom state
	for index := range nameStrings {
		//We only want to have objects listed if their count matches
		clusterObj := helper.Cluster{}
		clustername := nameStrings[index]
		clusterregion := regionStrings[index]
		clusterproject := projectStrings[index]
		readClusterName := helper.ReadStateValues(helper.CarbonStateLocalFile, clustername)
		readClusterRegion := helper.ReadStateValues(helper.CarbonStateLocalFile, clusterregion)
		readClusterProject := helper.ReadStateValues(helper.CarbonStateLocalFile, clusterproject)
		if len(readClusterName) > 0 &&
			len(readClusterRegion) > 0 &&
			len(readClusterProject) > 0 {
			clusterObj.Name = readClusterName
			clusterObj.Region = readClusterRegion
			clusterObj.Project = readClusterProject
			clusters = append(clusters, clusterObj)
		}
	}
	return clusters
}
