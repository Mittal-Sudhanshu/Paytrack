package model

import "time"

type ClockInRequest struct {
	ClockInTime time.Time `json:"clockInTime" binding:"required"`
	Latitude    string    `json:"latitude" binding:"required"`
	Longitude   string    `json:"longitude" binding:"required"`
}

type ClockOutRequest struct {
	ClockOutTime time.Time `json:"clockOutTime" binding:"required"`
	Latitude     string `json:"latitude" binding:"required"`
	Longitude    string `json:"longitude" binding:"required"`
}
