package entities

import core "HyperFlow/core/entities"

type Office struct {
	core.BaseEntity[int]
	Name string `gorm:"column:adi"`
}

func (Office) TableName() string {
	return "remote_controls.ofisler"
}
