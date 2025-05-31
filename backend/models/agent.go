package models

import "gorm.io/gorm"

type SC_AGENT struct {
	gorm.Model
	Name    string `gorm:name`
	AgentId string `gorm:agentId "primaryKey"`
}
