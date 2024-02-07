package main

import (
	"github.com/labstack/echo"
	"github.com/ryugenxd/todolist-api/database"
	"github.com/ryugenxd/todolist-api/http/controllers"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// End point Get
	controllers.GetListController(e, db)
	// End point post
	controllers.PostListController(e, db)
	// End point update
	controllers.UpdateListController(e, db)
	// End point delete
	controllers.DeleteListController(e, db)
	// End point udate status
	controllers.UpdateStatusController(e, db)

	e.Start(":8080")
}
