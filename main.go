package main

import (
	"github.com/harry/jazzcorp/config"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
}
