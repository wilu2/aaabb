// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDocument = "document"

// Document mapped from table <document>
type Document struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Content string `gorm:"column:content;not null" json:"content"`
}

// TableName Document's table name
func (*Document) TableName() string {
	return TableNameDocument
}
