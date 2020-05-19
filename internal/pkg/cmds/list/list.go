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
package list

import (
	"github.com/opencord/voltadm/internal/pkg/cmds/list/envs"
	"github.com/opencord/voltadm/internal/pkg/cmds/list/stacks"
	"github.com/spf13/cobra"
)

// listCmd represents the list command

func NewCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "list",
		Short: "Lists the existing VOLTHA environments of stacks",
		Long:  "Lists the existing VOLTHA environments of stacks",
	}

	listCmd.AddCommand(envs.NewCommand())
	listCmd.AddCommand(stacks.NewCommand())

	return listCmd
}
