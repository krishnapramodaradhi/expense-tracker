package main

import (
	"github.com/krishnapramodaradhi/expense-tracker/internal/config"
)

func main() {
	s := config.NewServer(":8443")
	s.Start()
}
