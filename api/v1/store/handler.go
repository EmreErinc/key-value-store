package store

import "github.com/labstack/echo/v4"

func Handler(group *echo.Group) {
	group.POST("", set)
	group.GET("", get)
	group.PUT("/flush", flush)
}
