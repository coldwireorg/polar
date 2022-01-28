package models

import (
	"polar/database"
)

type Node struct {
	Address   string `gorm:"not null;unique;primaryKey"` // ip:port
	PublicKey []byte `gorm:"not null;unique"`
	Status    int    `gorm:"not null"` // 0 = dead, 1 = alive, 2 = timeout
	Storage   int64  `gorm:"not null"` // Free storage left
}

func (n Node) Add() {
	database.DB.Create(&n)
}

func (n Node) Del() {
	database.DB.Delete(&n)
}

func (n Node) List() []Node {
	var nodes []Node
	database.DB.Find(&Node{}).Scan(&nodes)
	return nodes
}

func (n Node) Number() int64 {
	nodes := database.DB.Find(&Node{})
	return nodes.RowsAffected
}
