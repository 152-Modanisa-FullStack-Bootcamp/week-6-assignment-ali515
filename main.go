package main

import (
	"bootcamp/config"
	"fmt"
)

func main() {
	fmt.Println(config.Env.IsLocal())
}
