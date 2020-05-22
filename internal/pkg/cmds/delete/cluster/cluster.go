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
	"github.com/opencord/voltadm/internal/pkg/cmds/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type flagpole struct {
	Name string
}

// deleteCmd represents the delete command
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	deleteClusterCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "Deletes a KinD cluster",
		Long:  "Deletes a KinD cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ExecArrayE(
				"kind",
				"delete",
				"cluster",
				"--name",
				flags.Name,
			)

		},
	}

	//viper.AutomaticEnv()

	viper.BindEnv("name", "CLUSTER_NAME")
	viper.SetDefault("name", "default")

	deleteClusterCmd.Flags().StringVarP(&flags.Name, "name", "n", viper.GetString("name"),
		"name of the KinD cluster to delete")

	return deleteClusterCmd
}
