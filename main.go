package main

import (
	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/router"
)

func main() {
	config.Load()
	database.Connect()

	router.Create()
	router.Run()
}
