package net

import (
	"encoding/json"
	"fmt"

	"github.com/rgb-24bit/kit/pkg/net"
	"github.com/spf13/cobra"
)

var cidr = &cobra.Command{
	Use:   "cidr [cidr]",
	Short: "get cidr info, like cidr 192.168.0.1/24",
	Args:  cobra.ExactArgs(1),
	Run:   cidrRun,
}

func cidrRun(cmd *cobra.Command, args []string) {
	network, _ := net.ParseIPNetwork(args[0])
	bytes, _ := json.MarshalIndent(network, "  ", "  ")
	fmt.Println(string(bytes))
}

func init() {
	_net.AddCommand(cidr)
}
