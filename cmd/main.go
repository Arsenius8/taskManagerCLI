package main

import (
	"cli/taskmanager/internal/database"
	"cli/taskmanager/internal/model"
	"cli/taskmanager/internal/service"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	err := database.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer database.DB.Close()

	if err := database.CreateTable(); err != nil {
		fmt.Println(err)
		return
	}

	flagsValue := model.InitFlags()

	err = service.TaskManager(flagsValue)

	if err != nil {
		fmt.Println(err)
	}
}
