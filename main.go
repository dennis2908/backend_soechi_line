package main

import (
	_ "backend_soechi_line/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "backend_soechi_line/routers"
	"backend_soechi_line/models"
	"github.com/astaxie/beego/plugins/cors"
	_ "os"

)


func init(){ // init instead of int
	 
	

	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	
	beego.InsertFilter("/Driver/*", beego.BeforeRouter, models.FilterUser)
	
	// Setting postgres
	set_user := beego.AppConfig.String("set_user")
	set_port := beego.AppConfig.String("set_port")
	set_dbname := beego.AppConfig.String("set_dbname")
	set_sslmode := beego.AppConfig.String("set_sslmode")
	set_password := beego.AppConfig.String("set_password")
	set_host := beego.AppConfig.String("set_host")
	//
	
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", 
        "postgres",
        "user=" + set_user + " password="+ set_password +" host="+ set_host +" port=" + set_port + " dbname="+ set_dbname +" sslmode="+ set_sslmode);
	orm.RegisterModel(new(models.Customers),new(models.Drivers),new(models.Items),new(models.Shipments))	
    orm.RunSyncdb("default", false, true)
	orm.RunCommand()
}
 func main() {
 		beego.Run()
 }
 