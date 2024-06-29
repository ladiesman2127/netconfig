package cmd

import (
	"context"
	"fmt"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"
	"github.com/ladiesman2127/netconfig/cmd/client"
	common "github.com/ladiesman2127/netconfig/internal/pkg"
	"github.com/spf13/cobra"
)

var deleteDns = &cobra.Command{
	Use:   "delete",
	Short: "Delete dns server from a config file. You have to pass serverName(optionally), serverAddress and sudo password",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		client, connection := client.New()
		defer connection.Close()
		if len(args) != 1 {
			fmt.Println("invalid args count")
			return
		}
		dns, err := common.ProcessAddress(&args[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		client.DeleteDns(ctx, &pb.DeleteDNSRequst{
			DnsToRemove: dns,
		})
	},
}

func init() {
	rootCmd.AddCommand(deleteDns)
}
