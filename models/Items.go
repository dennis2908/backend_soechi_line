package models

type Items struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"size(100);unique"`
}

func (a *Items) TableName() string {
	return "items"
}