package model

import (
	"time"
)

type Employee struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	LeaveType  string    `json:"leavetype"`
	Start_Date time.Time `json:"start_date"`
	End_Date   time.Time `json:"end_date"`
	TeamName   string    `json:"teamname"`
	File       string    `json:"file"`
	Reporter   string    `json:"reporter"`
}
