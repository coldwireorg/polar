package models

import "polar/database"

type Object struct {
	Uid  string `mapstructure:"uid"`
	Data []byte `mapstructure:"data"`
}

// Append object to db
func (o Object) Add() {
	err := database.DB.Create(o).Error
	if err != nil {
		panic(err)
	}
}

// Delete object from db
func (o Object) Del() {}

// Get object from db
func (o Object) Get() Object {
	var res Object

	err := database.DB.Where("uid = ?", o.Uid).Find(&res).Error
	if err != nil {
		panic(err)
	}

	return res
}
