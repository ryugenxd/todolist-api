package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ryugenxd/todolist-api/http/requests"
)

func UpdateListController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/api/list/update/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		var request requests.UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)
		_, err := db.Exec(
			"UPDATE list SET title = ? , description = ? WHERE id = ?",
			request.Title,
			request.Description,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
