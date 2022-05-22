package main;
import (
	"net/http"
	"github.com/labstack/echo/v4"
)
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello");
}
	
func main() {
	app := echo.New();
	app.GET("/", hello);
	app.Start(":1222");
}
