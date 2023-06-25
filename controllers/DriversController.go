package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"backend_soechi_line/models"
	"backend_soechi_line/structs"
	_ "github.com/beego/beego/v2/core/validation"
	 _ "fmt"
	"encoding/json"
)

type DriversController struct {
	beego.Controller
}

func (api *DriversController) DriverLogin() {
    var Drivers models.Drivers
	json.Unmarshal(api.Ctx.Input.RequestBody, &Drivers)
	username := Drivers.Username;
	password := Drivers.Password;
    o := orm.NewOrm()
	o.Using("default")
	var DriversS []*models.Drivers
	sql := "select * from drivers where username = '"+username+"' limit 1"
	num, err := o.Raw(sql).QueryRows(&DriversS)
	if err != orm.ErrNoRows && num > 0 {
	    key := []byte(beego.AppConfig.String("secretkey"))
		text := models.Decrypt(key, DriversS[0].Password)
		if(password == text){
		    tokenString:=models.CreateToken(username)
			result := structs.RetArr {Token:tokenString,Id:DriversS[0].Id}
			theresult, _ := json.Marshal(result)
			api.Ctx.WriteString(string(theresult))
			return
			return
				}else{
					api.Data["json"] = "Username or password is wrong"
				}
	}else{
	    api.Data["json"] = "Username or password is wrong"
	}
	api.ServeJSON()
}

func (api *DriversController) CustomerLogin() {
    var Customers models.Customers
	json.Unmarshal(api.Ctx.Input.RequestBody, &Customers)
	username := Customers.Username;
	password := Customers.Password;
    o := orm.NewOrm()
	o.Using("default")
	var CustomersS []*models.Customers
	sql := "select * from customers where username = '"+username+"' limit 1"
	num, err := o.Raw(sql).QueryRows(&CustomersS)
	if err != orm.ErrNoRows && num > 0 {
	    key := []byte(beego.AppConfig.String("secretkey"))
		text := models.Decrypt(key, CustomersS[0].Password)
		if(password == text){
		    tokenString:=models.CreateToken(username)
			result := structs.RetArr {Token:tokenString,Id:CustomersS[0].Id}
			theresult, _ := json.Marshal(result)
			api.Ctx.WriteString(string(theresult))
			return
			return
				}else{
					api.Data["json"] = "Username or password is wrong"
				}
	}else{
	    api.Data["json"] = "Username or password is wrong"
	}
	api.ServeJSON()
}

func (api *DriversController) GetCustomer() {
	
    o := orm.NewOrm()
	o.Using("default")
	var Customers []*models.Customers
	sql := "select * from customers"
	num, err := o.Raw(sql).QueryRows(&Customers)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Customers
	}else{
	    api.Data["json"] = structs.ArryEmp{}
	}
	api.ServeJSON()
}

func (api *DriversController) GetItem() {
	
    o := orm.NewOrm()
	o.Using("default")
	var Items []*models.Items
	sql := "select * from items"
	num, err := o.Raw(sql).QueryRows(&Items)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Items
	}else{
	    api.Data["json"] = structs.ArryEmp{}
	}
	api.ServeJSON()
}


func (api *DriversController) AddShipment() {
    var TheShipments structs.PutShipments
	id := api.Ctx.Input.Param(":id")
	json.Unmarshal(api.Ctx.Input.RequestBody, &TheShipments)
	var Shipments []*models.Shipments
    o := orm.NewOrm()
	o.Using("default")
	sql := "INSERT INTO shipments (customer_id, item_id,qty) VALUES ("+id+", "+TheShipments.Item_id+","+TheShipments.Qty+")"
	o.Raw(sql).QueryRows(&Shipments)
	api.Data["json"] = structs.ArryEmp{}
	api.ServeJSON()
}

func (api *DriversController) UpdateShipment() {
    var TheShipments structs.PutShipments
	var Shipments []*models.Shipments
	id := api.Ctx.Input.Param(":id")
	id_driver := api.Ctx.Input.Param(":id_driver")
	json.Unmarshal(api.Ctx.Input.RequestBody, &TheShipments)
	o := orm.NewOrm()
	o.Using("default")
	sql := "UPDATE shipments set driver_id = "+id_driver+" where id = "+id
	o.Raw(sql).QueryRows(&Shipments)
	api.Data["json"] =structs.ArryEmp{}
	api.ServeJSON()
}

func (api *DriversController) DeleteShipment() {
    id := api.Ctx.Input.Param(":id")
    o := orm.NewOrm()
	o.Using("default")
	var Shipments []*models.Shipments
	sql := "delete from shipments where id ="+id 
	o.Raw(sql).QueryRows(&Shipments)
	api.Data["json"] = structs.ArryEmp{}
	api.ServeJSON()
}


func (api *DriversController) GetShipmentWithLimitOffset() {
	
    offset := api.Ctx.Input.Param(":offset")
	limit := api.Ctx.Input.Param(":limit")
	
	sql_offset := ""
	sql_limit := ""
	if(offset!=""){
	   sql_offset = " OFFSET "+offset
	}
	if(limit!=""){
	   sql_limit = " LIMIT "+limit
	}
    o := orm.NewOrm()
	o.Using("default")
	getShipments :=   [] structs.AmbilDataShipments {}
	sql := "select shipments.id as shipment_id,customers.name as customer_name,customers.address as customer_address,customer_id,items.name as item_name,"
	sql += "drivers.id as driver_id,drivers.name as driver_name,shipments.id as shipment_id,customers.name as customer_name,customers.address as customer_address,"
	sql += "customer_id,items.name as item_name,item_id,qty from shipments join customers on customers.id = shipments.customer_id"
	sql += " left join drivers on drivers.id = shipments.driver_id join items on items.id = shipments.item_id order by shipments.id DESC"
	sql += sql_limit+sql_offset
	num, err := o.Raw(sql).QueryRows(&getShipments)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = getShipments
	}else{
	    api.Data["json"] = structs.ArryEmp{}
	}
	api.ServeJSON()
}

func (api *DriversController) GetShipmentCustomerWithLimitOffset() {
	
    offset := api.Ctx.Input.Param(":offset")
	limit := api.Ctx.Input.Param(":limit")
	id := api.Ctx.Input.Param(":id")
	
	sql_offset := ""
	sql_limit := ""
	if(offset!=""){
	   sql_offset = " OFFSET "+offset
	}
	if(limit!=""){
	   sql_limit = " LIMIT "+limit
	}
    o := orm.NewOrm()
	o.Using("default")
	getShipments :=   [] structs.AmbilDataShipments {}
	sql := "select shipments.id as shipment_id,customers.name as customer_name,customers.address as customer_address,customer_id,items.name as item_name,"
	sql += "drivers.id as driver_id,drivers.name as driver_name,shipments.id as shipment_id,customers.name as customer_name,customers.address as customer_address,"
	sql += "customer_id,items.name as item_name,item_id,qty from shipments join customers on customers.id = shipments.customer_id"
	sql += " left join drivers on drivers.id = shipments.driver_id join items on items.id = shipments.item_id where customers.id = "+id+" order by shipments.id DESC"
	sql += sql_limit+sql_offset
	num, err := o.Raw(sql).QueryRows(&getShipments)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = getShipments
	}else{
	    api.Data["json"] = structs.ArryEmp{}
	}
	api.ServeJSON()
}