package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ryugenxd/todolist-api/http/requests"
)

func UpdateStatusController(e *echo.Echo, db *sql.DB) {
	e.PUT("/api/list/status/:id", func(ctx echo.Context) error {

		id := ctx.Param("id")
		var request requests.StatusRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)
		var done int
		if request.Done {
			done = 1
		}
		_, err := db.Exec(
			"UPDATE list SET done = ? WHERE id = ?",
			done,
			id,
		)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")

	})
}
