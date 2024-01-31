package helper

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Project      GCPInfo        `mapstructure:"gcp"`
	FunctionInfo []FunctionInfo `mapstructure:"function"`
}

type GCPInfo struct {
	GCPProjectId string `mapstructure:"projectId"`
	Location     string `mapstructure:"location"`
}

type FunctionInfo struct {
	FunctionName        string `mapstructure:"functionName"`
	FunctionDescription string `mapstructure:"functionDescription"`
	Runtime             string `mapstructure:"runtime"`
	Entrypoint          string `mapstructure:"entrypoint"`
	Bucket              string `mapstructure:"bucketName"`
	SourceFile          string `mapstructure:"sourceFile"`
	AvailableCpu        string `mapstructure:"availableCpu"`
	AvailableMemory     string `mapstructure:"availableMemory"`
	MinInstanceCount    int32  `mapstructure:"minInstanceCount"`
	MaxInstanceCount    int32  `mapstructure:"maxInstanceCount"`
	ServiceAccountEmail string `mapstructure:"serviceAccountEmail"`
}

var FunctionConfig Config

func SetupConfig() {
	log.Println("Converting config into struct...")
	err := viper.Unmarshal(&FunctionConfig)
	fmt.Println(FunctionConfig)
	if err != nil {
		panic(err)
	}
}
