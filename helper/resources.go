package helper

// Cluster ...
type Cluster struct {
	Name, Region, Project string
}

// StatefileLocation object gets the state file information from status file
type StatefileLocation struct {
	Bucket, Path, Project string
	BackupInfo            BackupInfo
}

// Status object reads the status file
type Status struct {
	ApplicationCluster map[string]string `json:"application_cluster"`
	AdminCluster       map[string]string `json:"admin_cluster"`
	BackupCluster      map[string]string `json:"backup_cluster"`
	VaultConnect       map[string]string `json:"vault_connect"`
	ConsulConnect      map[string]string `json:"consul_connect"`
	RedisSpin          map[string]string `json:"redis_spin"`
}

// BackupInfo backup info from the status string
type BackupInfo struct {
	Bucket, IsTrue string
}
