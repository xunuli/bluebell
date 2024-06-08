package config

import (
	"fmt"
	"testing"
)

func TestGetGlobalConf(t *testing.T) {
	res := GetGlobalConf()

	fmt.Println(res)
}
