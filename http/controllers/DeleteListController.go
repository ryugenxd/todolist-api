package controllers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteListController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/api/list/delete/:id", func(ctx echo.Context) error {

		id := ctx.Param("id")
		_, err := db.Exec(
			"DELETE FROM list WHERE id = ?",
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
