/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/amasukakarot/google-cloud-builder/internal/cloudfunction"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// DeployCmd represents the deploy command
var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your commandss",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deploy called...")
		ctx := context.Background()
		functionExists := cloudfunction.IfFunctionExists(ctx, "projects/groovy-iris-412518/locations/europe-west2/functions/my-gcb-function-970601")
		if functionExists {
			log.Println("Updating cloud function...")
			cloudfunction.UpdateCloudFunction()
		} else {
			log.Println("Creating cloud function...")
			start := time.Now()
			cloudfunction.CreateCloudFunction(ctx)
			finished := time.Since(start)
			log.Printf("Function took %s", finished)
		}
	},
}

func init() {
	//add flags here

}
