package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ryugenxd/todolist-api/http/requests"
)

func PostListController(e *echo.Echo, db *sql.DB) {
	e.POST("/api/list", func(ctx echo.Context) error {

		var request requests.CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO list (title,description,done) VALUES (?,?,0)",
			request.Title, request.Description,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

}
