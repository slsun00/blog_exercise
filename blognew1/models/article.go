package models

import "time"

//Categor 类别
type Category struct {
	ID              int64
	Title           string
	Created         time.Time `gorm:"index"`
	Views           int64     `gorm:"index"`
	TopicTime       time.Time `gorm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

//Topic文章
type Topic struct {
	ID               int64
	Uid              int64
	Title            string
	content          string `gorm:"size(5000)"`
	Attachment       string
	Created          time.Time `gorm:"index"`
	Updated          time.Time `gorm:"index"`
	Views            int64     `gorm:"index"`
	Author           string
	Replaytime       time.Time `gorm:"index"`
	ReplayCount      int64
	ReplayLastUserId int64
}

