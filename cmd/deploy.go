/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/amasukakarot/google-cloud-builder/internal/cloudfunction"
	"github.com/spf13/cobra"
)

// DeployCmd represents the deploy command
var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your commandss",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		cloudfunction.CreateCloudFunction()
	},
}

func init() {
	//add flags here

}
