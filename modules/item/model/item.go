package model

import (
	"errors"
	"http_api/common"
)

var (
	ErrPathEmpty = errors.New("path is empty")
	ErrDataTypeEmpty = errors.New("data type is empty")
	ErrInvalidID = errors.New("invalid id")
	ErrDataNotFound = errors.New("data not found")
	ErrDeleteFailed = errors.New("delete failed")
)

type Parameter struct {
	common.SqlModel
	Path        string     `json:"path" gorm:"column:pathh;type:varchar(255);not null"`
	DataType    string     `json:"data_type" gorm:"column:data_type;type:varchar(255);not null"`
	Description string     `json:"description" gorm:"column:description;type:varchar(255);omitempty"`
	Status 		ItemStatus `json:"status" gorm:"column:status;type:public.status;not null"`
	UpdatedBy   string     `json:"updated_by" gorm:"column:updated_by;type:varchar(255);not null"`
}

func (Parameter) TableName() string { return "parameters" }

type ParameterCreateRequest struct {
	Path        string `json:"path" binding:"required"`
	DataType    string `json:"data_type" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

func (ParameterCreateRequest) TableName() string { return Parameter{}.TableName() }

type ParameterUpdateRequest struct {
	Path        *string `gorm:"column:path;type:varchar;not null" json:"path"`
	DataType    *string `gorm:"column:data_type;type:varchar(32);not null" json:"data_type"`
	Description *string `gorm:"column:description;type:string;default:null" json:"description,omitempty"`
	Status      *ItemStatus `gorm:"column:status;type:public.status;default:'ENABLE':::public.status" json:"status"`
	UpdatedBy   *string `gorm:"column:updated_by" json:"updated_by"`
}

func (ParameterUpdateRequest) TableName() string { return Parameter{}.TableName() }