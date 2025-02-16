package main

import (
	"www.interview.com/routers"
	"www.interview.com/utils"
)

func main() {
	err := utils.InitDB()
	if err != nil {
		panic(err)
	}

	r := routers.SetupRouter()
	r.Run(":8080")
}
