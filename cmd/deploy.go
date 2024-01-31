/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/amasukakarot/google-cloud-builder/internal/cloudfunction"
	"github.com/spf13/cobra"
)

// DeployCmd represents the deploy command
var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your commandss",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cloudfunction.StartDeployment(ctx)
	},
}

func init() {
	//add flags here

}
