package cmds

import (
	"github.com/opencord/voltadm/internal/pkg/cmds/completion"
	"github.com/opencord/voltadm/internal/pkg/cmds/create"
	"github.com/opencord/voltadm/internal/pkg/cmds/delete"
	"github.com/opencord/voltadm/internal/pkg/cmds/get"
	"github.com/opencord/voltadm/internal/pkg/cmds/list"
	"github.com/spf13/cobra"
)

func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "voltadm",
		Short: "Creates and manages VOLTHA environments and stacks",
		Long: `voltadm is used to create and manage VOLTHA environments and stacks.

When being used to manage VOLTHA environments, voltadm can deploy the required VOLTHA infrastructure
against an existing Kubernetes cluster or can deploy a KinD Kubernetes cluster.

When being used to manage VOLTHA stacks, voltadm can deploy a VOLTHA stack into any
Kubernetes cluster.`,
	}
	cmd.AddCommand(create.NewCommand())
	cmd.AddCommand(delete.NewCommand())
	cmd.AddCommand(list.NewCommand())
	cmd.AddCommand(get.NewCommand())
	cmd.AddCommand(completion.NewCommand())
	return cmd
}
