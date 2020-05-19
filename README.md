
# ![](./assets/voltadm64x64.png) VOLTHA Stack and Infrastructure Deployment Tool
`voltadm` is a utility that can be used to deploy a VOLTHA stack and/or any
required VOLTHA infrastructure into an environment. The simplest definition of
`voltadm` is a utility that wraps underlying commands, such as `helm install`,
with common configuration options to simplify a VOLTHA stack deployment. Any
capability provided by `voltadm` can also be executed by an operator using
the underlying commands.

## VOLTHA Environments
A VOLTHA environment represents a Kubernetes cluster along with any VOLTHA
required dependencies, such as ETCd and Kafka. `voltadm` can deploy the
infrastructure into an existing Kubernetes cluster or it can create a `KinD`
Kubernetes cluster by leveraging the `kind` command.

## VOLTHA Stack
A VOLTHA stack represents those components that are unique to VOLTHA and
manage a single OLT and its associated ONTs/ONUs. These components are the
VOLTHA Core, OpenFlow Agent, and required OLT/ONU adapters.

## Commands
`voltadm` provides commands to `create`, `delete`, `list`, and `get` VOLTHA
environments and stacks.

## Examples

### Create VOLTHA Environment Including `KinD` Cluster
```
$ voltadm create env --name="sample" --with-k8s --num-workers=2 --num-control=1 \
    --num-bbsim=10 --with-kafka --with-etcd
```

### Create VOLTHA Stack
```
$ voltadm create stack --env=sample --name"stack1"
```
