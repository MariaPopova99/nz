package main

import (
	nz "api-character"
	"fmt"
	"time"
)

func main() {
	characterClient, err := nz.NewClient(time.Second * 10)
	if err != nil {
		return err
	}
	for _, asset := range characterClient {
		fmt.Println(asset.GetInfo())
	}
}
