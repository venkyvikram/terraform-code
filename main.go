package main

import (
	"synergy/cmd"
)

func main() {
	cmd.Execute()
}

// const statefilelocation string = "state.tfstate"

// type cluster struct {
// 	name, region, project string
// }

// func main() {
// 	clusters := [2]string{"cluster1", "cluster2"}
// 	for _, clusterprefix := range clusters {
// 		clustername := clusterprefix + "_" + "cluster_name"
// 		fmt.Println("name from state", readValues(clustername))
// 		clusterregion := clusterprefix + "_" + "cluster_region"
// 		fmt.Println("region from state", readValues(clusterregion))
// 	}
// 	fmt.Println("project from state", readValues("project"))
// }

// func readValues(key string) string {
// 	statefileObj, err := helper.ParseTerraformStateFile(statefilelocation)
// 	for module := range statefileObj.Modules {
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if val, ok := statefileObj.Modules[module].Outputs[key]; ok {
// 			clusterMap := val.(map[string]interface{})
// 			clustergeneric := clusterMap["value"]
// 			return clustergeneric.(string)
// 		}
// 	}
// 	return ""
// }
