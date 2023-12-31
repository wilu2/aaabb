// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Account        string     `gorm:"column:account;not null" json:"account"`
	Password       string     `gorm:"column:password;not null" json:"password"`
	Alias_         string     `gorm:"column:alias;not null" json:"alias"`
	Role           string     `gorm:"column:role;not null" json:"role"`
	Status         bool       `gorm:"column:status;not null;default:true" json:"status"`
	Channels       JSON       `gorm:"column:channels;not null" json:"channels"`
	Description    string     `gorm:"column:description;not null;default:''::character varying" json:"description"`
	Abandoned      bool       `gorm:"column:abandoned;not null" json:"abandoned"`
	DelUniqueKey   int32      `gorm:"column:del_unique_key;not null" json:"del_unique_key"`
	Ctime          time.Time  `gorm:"column:ctime;not null;default:CURRENT_TIMESTAMP" json:"ctime"`
	LastUpdateTime time.Time  `gorm:"column:last_update_time;not null;default:CURRENT_TIMESTAMP" json:"last_update_time"`
	LastLoginTime  *time.Time `gorm:"column:last_login_time" json:"last_login_time"`
	PwExpiraTime   *time.Time `gorm:"column:pw_expira_time" json:"pw_expira_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
