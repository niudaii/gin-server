package request

import uuid "github.com/satori/go.uuid"

type PageInfo struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}

type UUID struct {
	UUID uuid.UUID `json:"uuid" binding:"required"`
}
