package main

import (
	"fmt"
	"github.com/douyu/jupiter"
	"jupiter-demo/internal/app/engine"
	"jupiter-demo/internal/app/model"
	"jupiter-demo/internal/app/service"
	"log"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
		fmt.Println("exit jupiter app ...")
		return nil
	})

	model.Init()
	service.Init()
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
