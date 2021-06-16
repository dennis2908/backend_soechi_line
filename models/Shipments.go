package models

type Shipments struct {
	Id int `orm:"pk;auto"`
	Customer_id int `orm`
	Driver_id int `orm`
	Item_id int `orm`
	Qty int `orm`
}

func (a *Shipments) TableName() string {
	return "shipments"
}