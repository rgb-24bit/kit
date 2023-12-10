package main

import (
	"github.com/rgb-24bit/kit/cmd/kit/net"
	"github.com/spf13/cobra"
)

var kit = &cobra.Command{}

func main() {
	net.Setup(kit)

	kit.Execute()
}
