package main

import (
	"fmt"

	"github.com/blackironj/gin-test/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	fmt.Println(setting.ServerSetting.RunMode)
}
