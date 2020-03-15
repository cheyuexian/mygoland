package models

type T1 struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Name   string `xorm:"not null index VARCHAR(255)"`
	Num    int    `xorm:"not null INT(11)"`
	Status int    `xorm:"not null index INT(11)"`
}
