package main

import (
	"github.com/ericfmonteiro/travel-network/app/db"
	"github.com/ericfmonteiro/travel-network/app/routes"
)

func main() {
	db.NewDatabase()
	routes.InitGin()
}
