package router;

import (
	"github.com/ErfanMomeniii/urlshortener-go/db"
	"net/http"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/ErfanMomeniii/urlshortener-go/model"
)


func New() *echo.Echo {

	app := echo.New();
	app.Start(":1222");
	return app;
}

func RegisterRoutes(app *echo.Echo,d *gorm.DB){

	app.GET("url/:shortLink",func(c echo.Context)error{
		return getUrlByShortLink(c,d);
	});

	app.POST("url",func (c echo.Context)error{
		return storeUrl(c,d)
	});

}

func getUrlByShortLink(c echo.Context,d *gorm.DB)error{
	result:=d.Where("shortLink = ?",c.QueryParam("shortLink")).First(&model.URL{});
	if result.Error!=nil{
		return result.Error;
	}
	response:=c.JSON(http.StatusOK,result);
	return response;
}

func storeUrl(c echo.Context,d *gorm.DB)error{
	url := &model.URL{ShortLink:c.QueryParam("shortLink"),Address:c.QueryParam("Address")}
	if err:=url.Validate();err!=nil{
		return err;
	}
	
	if err:=db.Create(url,d);err!=nil{
		return err;
	}
	
	response:=c.JSON(http.StatusOK,url.ShortLink);
	return response;
}
