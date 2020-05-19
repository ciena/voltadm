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
package stacks

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command

func NewCommand() *cobra.Command {
	listStacksCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "stacks",
		Short: "List the known VOLTHA stacks",
		Long:  "List the known VOLTHA stacks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list stacks called")
		},
	}
	return listStacksCmd
}
