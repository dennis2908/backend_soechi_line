package models

type Customers struct {
	Id int `orm:"pk;auto"`
	Username string `orm:"size(100);unique"`
	Address string `orm:"size(150)"`
	Password string `orm:"size(50)"`
	Name string
}

func (a *Customers) TableName() string {
	return "customers"
}