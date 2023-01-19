package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	conf := viper.New()
	conf.SetConfigFile("folder/testjson.conf")
	conf.SetConfigType("json")
	if err := conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var dst []map[string]interface{}
	if err := conf.UnmarshalKey("config", &dst); err != nil {
		panic(fmt.Errorf("UnmarshalKey \"config\" error: %w", err))
	}
	if dst[0]["p3"].(bool) {
		fmt.Println(int(dst[0]["p4"].(float64)) + 1)
	}
	if !dst[1]["p7"].(bool) {
		fmt.Println(int(dst[1]["p8"].(float64)) + 10)
	}
	for i := 0; i < len(dst); i++ {
		fmt.Println()
		for k, v := range dst[i] {
			fmt.Println(k, v)
		}
	}
}
