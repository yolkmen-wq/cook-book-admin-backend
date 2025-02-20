package models

import "time"

type SystemLog struct {
	ID          int       `json:"id" gorm:"column:id;primary_key;"`
	Module      string    `json:"module" gorm:"column:module"`
	Url         string    `json:"url" gorm:"column:url"`
	Method      string    `json:"method" gorm:"column:method"`
	Ip          string    `json:"ip" gorm:"column:ip"`
	Address     string    `json:"address" gorm:"column:address"`
	System      string    `json:"system" gorm:"column:system"`
	Browser     string    `json:"browser" gorm:"column:browser"`
	TakesTime   int64     `json:"takesTime" gorm:"column:takes_time"`
	RequestTime time.Time `json:"requestTime" gorm:"column:request_time"`
}
