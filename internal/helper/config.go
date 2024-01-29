package helper

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Project      GCPInfo      `mapstructure:"gcp"`
	FunctionInfo FunctionInfo `mapstructure:"function"`
}

type GCPInfo struct {
	GCPProjectId string `mapstructure:"projectId"`
	Location     string `mapstructure:"location"`
}

type FunctionInfo struct {
	FunctionName        string `mapstructure:"functionName"`
	Runtime             string `mapstructure:"runtime"`
	Entrypoint          string `mapstructure:"entrypoint"`
	Bucket              string `mapstructure:"bucketName"`
	AvailableCpu        string `mapstructure:"availableCpu"`
	AvailableMemory     string `mapstructure:"availableMemory"`
	MinInstanceCount    string `mapstructure:"minInstanceCount"`
	MaxInstanceCount    string `mapstructure:"maxInstanceCount"`
	ServiceAccountEmail string `mapstructure:"serviceAccountEmail"`
}

var FunctionConfig Config

func SetupConfig() {
	log.Println("Converting config into struct...")
	err := viper.Unmarshal(&FunctionConfig)
	if err != nil {
		panic(err)
	}
}
