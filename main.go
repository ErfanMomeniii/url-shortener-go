package main;
import (
	"github.com/ErfanMomeniii/urlshortener-go/db"
	"github.com/ErfanMomeniii/urlshortener-go/router"
)
	
func main() {
	d := db.ConnectToDb();
	db.HandleMigration(d);

	r:=router.New();
	router.RegisterRoutes(r,d);

}
