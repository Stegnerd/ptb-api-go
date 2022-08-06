package api

import "github.com/gin-gonic/gin"

func Run() {
	// db set up

	// db migrations

	// set up dependencies

	// set up routers from each directory

	r := gin.Default()

	err := r.Run()
	if err != nil {
		return
	}
}
