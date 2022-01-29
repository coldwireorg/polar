package models

import (
	"errors"
	"polar/database"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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

func (n Node) FindWithPbKey() Node {
	var node Node
	err := database.DB.Raw("SELECT * FROM nodes WHERE public_key = ? LIMIT 1", n.PublicKey).Scan(&node).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Node{}
		} else {
			log.Err(err)
		}
	}

	return node
}

func (n Node) List() []Node {
	var nodes []Node
	database.DB.Find(&nodes)
	return nodes
}

func (n Node) Number() int {
	var node []Node
	database.DB.Find(&node)
	return len(node)
}
