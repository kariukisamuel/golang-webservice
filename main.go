package main

import (
	"fmt"

	"webservice/models"
)

func main() {
	u := models.User{
		Name:     "sam",
		Id:       20,
		Location: "ruiru",
	}
	fmt.Println(u)

}
