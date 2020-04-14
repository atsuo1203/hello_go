package sub

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func GetServer() {
	e := echo.New()
	e.GET("/users/:name", getUserName)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}
