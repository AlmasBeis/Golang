package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"thirdassignment/controller"
	"thirdassignment/model"
	"thirdassignment/repository"
	"thirdassignment/service"
)

func main() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	//dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = db.AutoMigrate(&model.Book{}); err != nil {
		fmt.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	c := controller.NewController(serv)

	router := c.Routes()

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
