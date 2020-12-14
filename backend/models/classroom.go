package models

import "backend/utils"

type Classroom struct {
	ID      uint   `gorm:"primarykey"`
	Major   string `gorm:"size:20"` // IPA-MIA/IPS_IIS/BHS
	Level   string `gorm:"type:ENUM('X', 'XI', 'XII')"`
	Seq     uint8
	Members []UserClassroom `gorm:"foreignKey:ClassroomId"`
	utils.Created
	utils.Updated
}

func (Classroom) CreateHistory() interface{} {
	return &ClassroomHistory{}
}

type ClassroomHistory struct {
	utils.History
	ID    uint   `gorm:"primarykey"`
	Major string `gorm:"size:20"` // IPA-MIA/IPS_IIS/BHS
	Level string `gorm:"type:ENUM('X', 'XI', 'XII')"`
	Seq   uint8
	utils.Updated
}
