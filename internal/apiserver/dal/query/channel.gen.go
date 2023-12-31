// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gitlab.intsig.net/textin-gateway/internal/apiserver/dal/model"
)

func newChannel(db *gorm.DB, opts ...gen.DOOption) channel {
	_channel := channel{}

	_channel.channelDo.UseDB(db, opts...)
	_channel.channelDo.UseModel(&model.Channel{})

	tableName := _channel.channelDo.TableName()
	_channel.ALL = field.NewAsterisk(tableName)
	_channel.ID = field.NewInt32(tableName, "id")
	_channel.Name = field.NewString(tableName, "name")
	_channel.CreatorID = field.NewInt32(tableName, "creator_id")
	_channel.LastEditorID = field.NewInt32(tableName, "last_editor_id")
	_channel.Abandoned = field.NewBool(tableName, "abandoned")
	_channel.DelUniqueKey = field.NewInt32(tableName, "del_unique_key")
	_channel.Ctime = field.NewTime(tableName, "ctime")
	_channel.LastUpdateTime = field.NewTime(tableName, "last_update_time")

	_channel.fillFieldMap()

	return _channel
}

type channel struct {
	channelDo channelDo

	ALL            field.Asterisk
	ID             field.Int32
	Name           field.String
	CreatorID      field.Int32
	LastEditorID   field.Int32
	Abandoned      field.Bool
	DelUniqueKey   field.Int32
	Ctime          field.Time
	LastUpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (c channel) Table(newTableName string) *channel {
	c.channelDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c channel) As(alias string) *channel {
	c.channelDo.DO = *(c.channelDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *channel) updateTableName(table string) *channel {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Name = field.NewString(table, "name")
	c.CreatorID = field.NewInt32(table, "creator_id")
	c.LastEditorID = field.NewInt32(table, "last_editor_id")
	c.Abandoned = field.NewBool(table, "abandoned")
	c.DelUniqueKey = field.NewInt32(table, "del_unique_key")
	c.Ctime = field.NewTime(table, "ctime")
	c.LastUpdateTime = field.NewTime(table, "last_update_time")

	c.fillFieldMap()

	return c
}

func (c *channel) WithContext(ctx context.Context) *channelDo { return c.channelDo.WithContext(ctx) }

func (c channel) TableName() string { return c.channelDo.TableName() }

func (c channel) Alias() string { return c.channelDo.Alias() }

func (c *channel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *channel) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["creator_id"] = c.CreatorID
	c.fieldMap["last_editor_id"] = c.LastEditorID
	c.fieldMap["abandoned"] = c.Abandoned
	c.fieldMap["del_unique_key"] = c.DelUniqueKey
	c.fieldMap["ctime"] = c.Ctime
	c.fieldMap["last_update_time"] = c.LastUpdateTime
}

func (c channel) clone(db *gorm.DB) channel {
	c.channelDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c channel) replaceDB(db *gorm.DB) channel {
	c.channelDo.ReplaceDB(db)
	return c
}

type channelDo struct{ gen.DO }

func (c channelDo) Debug() *channelDo {
	return c.withDO(c.DO.Debug())
}

func (c channelDo) WithContext(ctx context.Context) *channelDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c channelDo) ReadDB() *channelDo {
	return c.Clauses(dbresolver.Read)
}

func (c channelDo) WriteDB() *channelDo {
	return c.Clauses(dbresolver.Write)
}

func (c channelDo) Session(config *gorm.Session) *channelDo {
	return c.withDO(c.DO.Session(config))
}

func (c channelDo) Clauses(conds ...clause.Expression) *channelDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c channelDo) Returning(value interface{}, columns ...string) *channelDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c channelDo) Not(conds ...gen.Condition) *channelDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c channelDo) Or(conds ...gen.Condition) *channelDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c channelDo) Select(conds ...field.Expr) *channelDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c channelDo) Where(conds ...gen.Condition) *channelDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c channelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *channelDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c channelDo) Order(conds ...field.Expr) *channelDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c channelDo) Distinct(cols ...field.Expr) *channelDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c channelDo) Omit(cols ...field.Expr) *channelDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c channelDo) Join(table schema.Tabler, on ...field.Expr) *channelDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c channelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *channelDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c channelDo) RightJoin(table schema.Tabler, on ...field.Expr) *channelDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c channelDo) Group(cols ...field.Expr) *channelDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c channelDo) Having(conds ...gen.Condition) *channelDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c channelDo) Limit(limit int) *channelDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c channelDo) Offset(offset int) *channelDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c channelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *channelDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c channelDo) Unscoped() *channelDo {
	return c.withDO(c.DO.Unscoped())
}

func (c channelDo) Create(values ...*model.Channel) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c channelDo) CreateInBatches(values []*model.Channel, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c channelDo) Save(values ...*model.Channel) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c channelDo) First() (*model.Channel, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Channel), nil
	}
}

func (c channelDo) Take() (*model.Channel, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Channel), nil
	}
}

func (c channelDo) Last() (*model.Channel, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Channel), nil
	}
}

func (c channelDo) Find() ([]*model.Channel, error) {
	result, err := c.DO.Find()
	return result.([]*model.Channel), err
}

func (c channelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Channel, err error) {
	buf := make([]*model.Channel, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c channelDo) FindInBatches(result *[]*model.Channel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c channelDo) Attrs(attrs ...field.AssignExpr) *channelDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c channelDo) Assign(attrs ...field.AssignExpr) *channelDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c channelDo) Joins(fields ...field.RelationField) *channelDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c channelDo) Preload(fields ...field.RelationField) *channelDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c channelDo) FirstOrInit() (*model.Channel, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Channel), nil
	}
}

func (c channelDo) FirstOrCreate() (*model.Channel, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Channel), nil
	}
}

func (c channelDo) FindByPage(offset int, limit int) (result []*model.Channel, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c channelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c channelDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c channelDo) Delete(models ...*model.Channel) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *channelDo) withDO(do gen.Dao) *channelDo {
	c.DO = *do.(*gen.DO)
	return c
}
