package helper

import "fmt"

// ShowClusterOutput ...
func ShowClusterOutput(clusters []Cluster) {
	for _, clust := range clusters {
		fmt.Printf("%10s          %s            %s\n", clust.Name, clust.Region, clust.Project)
	}
}

// ShowEnvironsOutput ...
func ShowEnvironsOutput(list []string) {
	for _, environ := range list {
		fmt.Println(environ)
	}
}

// ShowResourceOutput ...
func ShowResourceOutput(data map[string]interface{}) {
	for key := range data {
		fmt.Println(key)
	}
}

// ShowBackupBucketOutput ...
func ShowBackupBucketOutput(stateFileLocation StatefileLocation, clusters []Cluster) {
	for _, clust := range clusters {
		fmt.Printf("%s          %s        %10s          %s            %s\n", stateFileLocation.BackupInfo.IsTrue, stateFileLocation.BackupInfo.Bucket, clust.Name, clust.Region, clust.Project)
	}
}

// ShowVaultClusterOutput ...
func ShowVaultClusterOutput(clusters []Cluster) {
	for _, clust := range clusters {
		fmt.Printf("%10s          %s            %s\n", clust.Name, clust.Region, clust.Project)
	}
}

// ShowRedisOutput ...
func ShowRedisOutput(redisIP string) {
	fmt.Printf("%s\n", redisIP)
}
