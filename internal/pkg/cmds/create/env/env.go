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
package env

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type BaseSpec struct {
	Enabled  bool `yaml:"enabled,omitempty"`
	Replicas int  `yaml:"replicas"`
	Image    struct {
		Repository string `yaml:"repository,omitempty"`
		Tag        string `yaml:"tag,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty"`
	} `yaml:"image,omitempty"`
}

type ZookeeperSpec struct {
	BaseSpec `yaml:",inline"`
}

type KafkaSpec struct {
	BaseSpec  `yaml:",inline"`
	Zookeeper ZookeeperSpec `yaml:"zookeeper"`
}

type EtcdSpec struct {
	BaseSpec `yaml:",inline"`
}

type AtomixSpec struct {
	BaseSpec `yaml:",inline"`

	Persistence struct {
		Enabled bool `yaml:"enabled"`
	}
}

type OnosSpec struct {
	BaseSpec `yaml:",inline"`

	Atomix AtomixSpec `yaml:"atomix"`
	DHCP   bool       `yaml:"dhcp"`
	EAPOL  bool       `yaml:"eapol"`
	IGMP   bool       `yaml:"igmp"`
}

type RadiusSpec struct {
	BaseSpec `yaml:",inline"`
}

type BBSIMSpec struct {
	OltId int  `yaml:"olt_id"`
	Auth  bool `yaml:"auth"`
	DHCP  bool `yaml:"dhcp"`
	IGMP  bool `yaml:"igmp"`
	Pon   int  `yaml:"pon"`
	Onu   int  `yaml:"onu"`
}

type BasicEnvironment struct {
	Kafka  KafkaSpec  `yaml:"kafka"`
	Etcd   EtcdSpec   `yaml:etcd"`
	Onos   OnosSpec   `yaml:"onos"`
	Radius RadiusSpec `yaml:"radius"`
	BBSIM0 BBSIMSpec  `yaml:"bbsim0"`
	BBSIM1 BBSIMSpec  `yaml:"bbsim1"`
	BBSIM2 BBSIMSpec  `yaml:"bbsim2"`
	BBSIM3 BBSIMSpec  `yaml:"bbsim3"`
	BBSIM4 BBSIMSpec  `yaml:"bbsim4"`
}

type flagpole struct {
	Name string
}

// createCmd represents the create command
func NewCommand() *cobra.Command {
	createEnvCmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "env",
		Short: "Creates a VOLTHA environment and deploys required infrastructure",
		Long:  "Creates a VOLTHA environment and deploys required infrastructure",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create env called")
			b := &BasicEnvironment{}
			y, _ := yaml.Marshal(b)
			fmt.Println(string(y))
		},
	}
	return createEnvCmd
}
