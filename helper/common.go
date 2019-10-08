package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	// ApplicationCluster ...
	ApplicationCluster = "application_cluster"
	// AdminCluster ...
	AdminCluster = "admin_cluster"
	// BackupCluster ...
	BackupCluster = "backup_cluster"
	// VaultConnect ...
	VaultConnect = "vault_connect"
	// ConsulConnect ...
	ConsulConnect = "consul_connect"
	// RedisSpin ...
	RedisSpin = "redis_spin"
	// CarbonStatusLocalFile ...
	CarbonStatusLocalFile = "carbon-status-local.json"
	// CarbonStateLocalFile ...
	CarbonStateLocalFile = "carbon-state-local.tfstate"
	// ClusterNameSuffix ...
	ClusterNameSuffix = "_cluster_name"
	// ClusterRegionSuffix ...
	ClusterRegionSuffix = "_cluster_region"
	// ProjectSuffix ...
	ProjectSuffix = "_project_name"

	// AdminClusterNameKey ...
	AdminClusterNameKey = "cluster_name"
	// AdminClusterRegionKey ...
	AdminClusterRegionKey = "cluster_region"
	// AdminClusterProjectKey ...
	AdminClusterProjectKey = "project"
)

// HomeDir get the home directory path
func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// Return true if the given file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// WriteToFile takes input and write file
func WriteToFile(filepath string, fileContent []byte) error {
	err := ioutil.WriteFile(filepath, fileContent, 0644)
	if err != nil {
		return errors.Wrap(err, "cannot write to file")
	}
	return nil
}

// MatchPattern is used to match the state file location in status file
func MatchPattern(resourceType, pathString string) bool {
	var regExp string
	if resourceType == ApplicationCluster || resourceType == AdminCluster || resourceType == RedisSpin {
		regExp = "^[a-z-]+[:][a-z-]+[:][\a-z-]+(.tfstate)$"
	} else if resourceType == BackupCluster {
		regExp = "^[true|false][a-z-]+[:][a-z-]+[:][a-z-]+[:][\a-z-]+(.tfstate)$"
	} else if resourceType == VaultConnect {
		regExp = "^[a-z-]+[:][a-z-]+[:][\a-z-]+(.tfstate)$"
	} else if resourceType == ConsulConnect {
		regExp = "^[a-z-]+[:][a-z-]+[:][\a-z-]+(.tfstate)$"
	}
	regCompile := regexp.MustCompile(regExp)
	isMatch := regCompile.MatchString(pathString)
	return isMatch
}

// ReadStatusFilePath reads the status file and creates a StatefileLocation object
func ReadStatusFilePath(resourceType, location string) StatefileLocation {
	// isMatch := MatchPattern(location)
	statefileLocation := new(StatefileLocation)
	if resourceType == ApplicationCluster || resourceType == AdminCluster || resourceType == RedisSpin {
		// if isMatch {
		strSplit := strings.Split(location, ":")
		statefileLocation.Project = strSplit[0]
		statefileLocation.Bucket = strSplit[1]
		statefileLocation.Path = strSplit[2]
	} else if resourceType == BackupCluster {
		// if isMatch {
		strSplit := strings.Split(location, ":")
		statefileLocation.BackupInfo.IsTrue = strSplit[0]
		statefileLocation.BackupInfo.Bucket = strSplit[1]
		statefileLocation.Project = strSplit[2]
		statefileLocation.Bucket = strSplit[3]
		statefileLocation.Path = strSplit[4]
	} else if resourceType == VaultConnect {
		// if isMatch {
		strSplit := strings.Split(location, ":")
		statefileLocation.Project = strSplit[0]
		statefileLocation.Bucket = strSplit[1]
		statefileLocation.Path = strSplit[2]
	} else if resourceType == ConsulConnect {
		// if isMatch {
		strSplit := strings.Split(location, ":")
		statefileLocation.Project = strSplit[0]
		statefileLocation.Bucket = strSplit[1]
		statefileLocation.Path = strSplit[2]
	}
	return *statefileLocation
}

// ReadStatusFileValues reads the status file
func ReadStatusFileValues(resourceType, environ string, statusFileObj Status) (StatefileLocation, error) {
	var statefileLocationObj StatefileLocation

	switch resourceType {
	case ApplicationCluster:
		if val, ok := statusFileObj.ApplicationCluster[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(ApplicationCluster, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	case AdminCluster:
		if val, ok := statusFileObj.AdminCluster[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(AdminCluster, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	case BackupCluster:
		if val, ok := statusFileObj.BackupCluster[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(BackupCluster, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	case VaultConnect:
		if val, ok := statusFileObj.VaultConnect[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(VaultConnect, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	case ConsulConnect:
		if val, ok := statusFileObj.ConsulConnect[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(ConsulConnect, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	case RedisSpin:
		if val, ok := statusFileObj.RedisSpin[environ]; ok {
			if MatchPattern(resourceType, val) {
				statefileLocationObj = ReadStatusFilePath(RedisSpin, val)
			} else {
				return statefileLocationObj, errors.New("statusfile: values does not follow the correct pattern, please validate")
			}
		}
	}

	return statefileLocationObj, nil
}

// ReadStatusFile parses the status file
func ReadStatusFile(metaFilelocation string) Status {
	var status Status
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadFile(metaFilelocation)
	err := json.Unmarshal(byteValue, &status)
	if err != nil {
		log.Fatalf("error parsing JSON: %s\n", err.Error())
	}
	return status
}

// GetStatusMap parses the status json file into a map
func GetStatusMap(metaFilelocation string) map[string]interface{} {
	var statusData map[string]interface{}

	byteValue, _ := ioutil.ReadFile(metaFilelocation)
	err := json.Unmarshal(byteValue, &statusData)
	if err != nil {
		log.Fatalf("error parsing JSON: %s\n", err.Error())
	}
	return statusData
}

// ReadStatusResourceEnvirons reads the status file for environs
func ReadStatusResourceEnvirons(resourceType, metaFilelocation string) []string {
	stateRecd := ReadStatusFile(metaFilelocation)
	var environsList []string

	switch resourceType {
	case ApplicationCluster:
		for key := range stateRecd.ApplicationCluster {
			environsList = append(environsList, key)
		}
	case AdminCluster:
		for key := range stateRecd.AdminCluster {
			environsList = append(environsList, key)
		}
	case BackupCluster:
		for key := range stateRecd.BackupCluster {
			environsList = append(environsList, key)
		}
	case VaultConnect:
		for key := range stateRecd.VaultConnect {
			environsList = append(environsList, key)
		}
	case ConsulConnect:
		for key := range stateRecd.ConsulConnect {
			environsList = append(environsList, key)
		}
	case RedisSpin:
		for key := range stateRecd.RedisSpin {
			environsList = append(environsList, key)
		}
	}
	return environsList
}

// ReadStateValues reads the values for the keys from the state file
func ReadStateValues(statefilelocationStr, key string) string {
	statefileObj, err := ParseTerraformStateFile(statefilelocationStr)
	if err != nil {
		log.Fatal(err)
	}
	for module := range statefileObj.Modules {
		if err != nil {
			log.Fatal(err)
		}
		if val, ok := statefileObj.Modules[module].Outputs[key]; ok {
			recdMap := val.(map[string]interface{})
			generic := recdMap["value"]
			return generic.(string)
		}
	}
	return ""
}

// GetClusterKeys is to get the cluster keys based on the number of clusters created in the environment
func GetClusterKeys(statefilelocationStr string) ([]string, []string, []string) {
	statefileObj, err := ParseTerraformStateFile(statefilelocationStr)
	nameList := []string{}
	regionList := []string{}
	projectList := []string{}

	if err != nil {
		log.Fatal(err)
	}
	for module := range statefileObj.Modules {
		if err != nil {
			log.Fatal(err)
		}
		outputsMap := statefileObj.Modules[module].Outputs
		for key := range outputsMap {
			if strings.Contains(key, ClusterNameSuffix) {
				nameList = append(nameList, key)
				clusterPrefix := strings.Split(key, ClusterNameSuffix)
				regionKey := clusterPrefix[0] + ClusterRegionSuffix
				regionList = append(regionList, regionKey)
				projectKey := clusterPrefix[0] + ProjectSuffix
				projectList = append(projectList, projectKey)
			}
		}
	}
	return nameList, regionList, projectList
}
