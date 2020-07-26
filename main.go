package main

import (
	"BudgetApi/src/config"
	"BudgetApi/src/config/route"
)

func main() {
	config.LoadEnv()
	var s route.Routes
	s.StartGin()
}
