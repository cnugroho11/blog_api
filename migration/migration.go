package main

import (
	"fmt"
	"log"

	"github.com/cnugroho11/blog_api/initializer"
	"github.com/cnugroho11/blog_api/model"
)

func init() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env")
	}

	initializer.ConnectDatabase(&config)
}

func main() {
	initializer.DB.AutoMigrate(&model.Blog{})

	fmt.Println("Migration complete")
}
