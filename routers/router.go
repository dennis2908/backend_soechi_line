package routers

import (
	"github.com/astaxie/beego"
	"backend_soechi_line/controllers"
)

func init() {
    beego.Router("/login", &controllers.DriversController{}, "post:DriverLogin")
	beego.Router("/CustomerLogin", &controllers.DriversController{}, "post:CustomerLogin")
	beego.Router("/Driver/Customer", &controllers.DriversController{}, "get:GetCustomer")
	beego.Router("/Driver/Item", &controllers.DriversController{}, "get:GetItem")
	beego.Router("/Driver/AddShipment/:id:int", &controllers.DriversController{}, "post:AddShipment")
	beego.Router("/Driver/GetShipmentWithLimitOffset/:limit:int/:offset:int", &controllers.DriversController{}, "get:GetShipmentWithLimitOffset")
	beego.Router("/Driver/GetShipmentCustomerWithLimitOffset/:id:int/:limit:int/:offset:int", &controllers.DriversController{}, "get:GetShipmentCustomerWithLimitOffset")
	beego.Router("/Driver/UpdateShipment/:id:int/:id_driver:int", &controllers.DriversController{}, "put:UpdateShipment")
	beego.Router("/Driver/DeleteShipment/:id:int", &controllers.DriversController{}, "delete:DeleteShipment")
}