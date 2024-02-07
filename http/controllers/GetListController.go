package controllers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ryugenxd/todolist-api/http/resources"
)

func GetListController(e *echo.Echo, db *sql.DB) {

	e.GET("/api/list", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM list")

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []resources.ListResource

		for rows.Next() {
			var id int
			var title string
			var description string
			var done int

			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var list resources.ListResource

			list.Id = id
			list.Title = title
			list.Description = description

			if done == 1 {
				list.Done = true
			}

			res = append(res, list)
		}

		return ctx.JSON(http.StatusOK, res)
	})
}
