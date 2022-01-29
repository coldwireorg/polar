package models

import (
	"polar/database"

	"gorm.io/gorm/clause"
)

type Node struct {
	PublicKey string `gorm:"not null;primaryKey"`
	Address   string `gorm:"not null;unique"` // ip:port
	Network   string `gorm:"not null"`        // Network the node is part of: Mainnet/Testnet
	Status    int    `gorm:"not null"`        // 0 = dead, 1 = alive, 2 = timeout
	Storage   uint64 `gorm:"not null"`        // Free storage left
}

func (n Node) Add() {
	database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&n)
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
