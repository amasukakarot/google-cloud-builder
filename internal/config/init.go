package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Values struct {
	Project  GCPInfo    `mapstructure:"gcp"`
	Function []Function `mapstructure:"function"`
}

type GCPInfo struct {
	GCPProjectId string `mapstructure:"projectId"`
	Location     string `mapstructure:"location"`
}

type Function struct {
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

var FunctionData Values

func InitConfig() {
	log.Println("Converting config into struct...")
	err := viper.Unmarshal(&FunctionData)
	fmt.Println(FunctionData)
	if err != nil {
		panic(err)
	}
}
