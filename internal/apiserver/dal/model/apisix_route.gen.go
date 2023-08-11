// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameApisixRoute = "apisix_route"

// ApisixRoute mapped from table <apisix_route>
type ApisixRoute struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RouteID     string    `gorm:"column:route_id;not null" json:"route_id"`
	Content     string    `gorm:"column:content;not null" json:"content"`
	ContentYaml string    `gorm:"column:content_yaml;not null" json:"content_yaml"`
	Type        int16     `gorm:"column:type;not null" json:"type"`
	Status      int16     `gorm:"column:status;not null" json:"status"`
	CreateAt    time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt    time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName ApisixRoute's table name
func (*ApisixRoute) TableName() string {
	return TableNameApisixRoute
}