package models

type Drivers struct {
	Id int `orm:"pk;auto"`
	Username string `orm:"size(100);unique"`
	Password string `orm:"size(100)"`
	Name string
}

func (a *Drivers) TableName() string {
	return "drivers"
}