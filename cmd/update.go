/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"
	"github.com/ladiesman2127/netconfig/cmd/client"
	"github.com/spf13/cobra"
)

var updateHostname = &cobra.Command{
	Use:   "update",
	Short: "Update a host name of a device",
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
		hostname := args[0]

		client.UpdateHostname(ctx, &pb.UpdateHostnameRequest{
			NewHostname: hostname,
		})
	},
}

func init() {
	rootCmd.AddCommand(updateHostname)
}
