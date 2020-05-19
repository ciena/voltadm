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
package create

import (
	"github.com/opencord/voltadm/internal/pkg/cmds/create/env"
	"github.com/opencord/voltadm/internal/pkg/cmds/create/stack"
	"github.com/spf13/cobra"
)

// createCmd represents the create command

func NewCommand() *cobra.Command {
	createCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create",
		Short: "Creates a VOLTHA environment or stack",
		Long:  "Creates a VOLTHA environment or stack",
	}

	createCmd.AddCommand(env.NewCommand())
	createCmd.AddCommand(stack.NewCommand())

	return createCmd
}