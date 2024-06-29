/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"
	"github.com/ladiesman2127/netconfig/cmd/client"
	common "github.com/ladiesman2127/netconfig/internal/pkg"
	"github.com/spf13/cobra"
)

var addDns = &cobra.Command{
	Use:   "add",
	Short: "Add dns server to a config file",
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
		client.AddDns(ctx, &pb.AddDNSRequest{
			NewDns: dns,
		})
	},
}

func init() {
	rootCmd.AddCommand(addDns)
}
