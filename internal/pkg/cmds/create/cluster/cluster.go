/*
Copyright © 2020 Open Networking Foundation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cluster

import (
	"fmt"
	"os"
	"strings"

	"github.com/opencord/voltadm/internal/pkg/cmds/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
)

type flagpole struct {
	Name                  string
	ClusterConfig         string
	NumWorkers            int
	NumControllers        int
	ScheduleOnControllers bool
}

// createCmd represents the create command
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	createClusterCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "Creates a KinD cluster",
		Long:  "Creates a KinD cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("create kind cluster called")
			fmt.Printf("%#+v\n", flags)

			c2 := &v1alpha4.Cluster{
				TypeMeta: v1alpha4.TypeMeta{
					Kind:       "Cluster",
					APIVersion: "kind.sigs.k8s.io/v1aplha4",
				},
			}

			if flags.ClusterConfig != "" {
				f, err := os.Open(flags.ClusterConfig)
				if err != nil {
					return err
				}
				defer f.Close()
				decoder := yaml.NewDecoder(f)
				if err = decoder.Decode(&c2); err != nil {
					return err
				}
			} else {
				for i := 0; i < flags.NumControllers; i++ {
					c2.Nodes = append(c2.Nodes, v1alpha4.Node{Role: "control-plane"})
				}
				for i := 0; i < flags.NumWorkers; i++ {
					c2.Nodes = append(c2.Nodes, v1alpha4.Node{Role: "worker"})
				}
			}

			y, _ := yaml.Marshal(c2)
			return utils.ExecIOEArrayE(
				strings.NewReader(string(y)),
				nil,
				nil,
				"kind",
				"create",
				"cluster",
				"--name",
				flags.Name,
				"--wait",
				"30m",
			)
		},
	}

	//viper.AutomaticEnv()

	viper.BindEnv("name", "CLUSTER_NAME")
	viper.BindEnv("config", "CLUSTER_CONFIG")
	viper.BindEnv("workers", "NUM_WORKERS")
	viper.BindEnv("controllers", "NUM_CONTROLLERS")
	viper.BindEnv("schedule", "SCHEDULE_ON_CONTROLLERS")
	viper.SetDefault("name", "default")
	viper.SetDefault("config", "")
	viper.SetDefault("workers", 2)
	viper.SetDefault("controllers", 1)
	viper.SetDefault("schedule", true)

	createClusterCmd.Flags().StringVarP(&flags.Name, "name", "n", viper.GetString("name"),
		"name of the KinD cluster to create")
	createClusterCmd.Flags().StringVarP(&flags.ClusterConfig, "cluster", "f", viper.GetString("config"),
		"name of the KinD cluster to create")
	createClusterCmd.Flags().IntVarP(&flags.NumWorkers, "workers", "w", viper.GetInt("workers"),
		"number of k8s workers in cluster")
	createClusterCmd.Flags().IntVarP(&flags.NumControllers, "controllers", "c", viper.GetInt("controllers"),
		"number of k8s controllers in cluster")
	createClusterCmd.Flags().BoolVarP(&flags.ScheduleOnControllers, "schedule", "s", viper.GetBool("schedule"),
		"remove the taint from the k8s controllers to allow scheduling")

	return createClusterCmd
}
