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
package delete

import (
	"github.com/opencord/voltadm/internal/pkg/cmds/delete/cluster"
	"github.com/opencord/voltadm/internal/pkg/cmds/delete/env"
	"github.com/opencord/voltadm/internal/pkg/cmds/delete/stack"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command

func NewCommand() *cobra.Command {
	deleteCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "delete",
		Short: "Deletes a VOLTHA environment or stack",
		Long:  "Deletes a VOLTHA environment or stack",
	}

	deleteCmd.AddCommand(env.NewCommand())
	deleteCmd.AddCommand(stack.NewCommand())
	deleteCmd.AddCommand(cluster.NewCommand())

	return deleteCmd
}
