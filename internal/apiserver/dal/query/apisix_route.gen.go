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

func newApisixRoute(db *gorm.DB, opts ...gen.DOOption) apisixRoute {
	_apisixRoute := apisixRoute{}

	_apisixRoute.apisixRouteDo.UseDB(db, opts...)
	_apisixRoute.apisixRouteDo.UseModel(&model.ApisixRoute{})

	tableName := _apisixRoute.apisixRouteDo.TableName()
	_apisixRoute.ALL = field.NewAsterisk(tableName)
	_apisixRoute.ID = field.NewInt32(tableName, "id")
	_apisixRoute.RouteID = field.NewString(tableName, "route_id")
	_apisixRoute.Content = field.NewString(tableName, "content")
	_apisixRoute.ContentYaml = field.NewString(tableName, "content_yaml")
	_apisixRoute.Type = field.NewInt16(tableName, "type")
	_apisixRoute.Status = field.NewInt16(tableName, "status")
	_apisixRoute.CreateAt = field.NewTime(tableName, "create_at")
	_apisixRoute.UpdateAt = field.NewTime(tableName, "update_at")

	_apisixRoute.fillFieldMap()

	return _apisixRoute
}

type apisixRoute struct {
	apisixRouteDo apisixRouteDo

	ALL         field.Asterisk
	ID          field.Int32
	RouteID     field.String
	Content     field.String
	ContentYaml field.String
	Type        field.Int16
	Status      field.Int16
	CreateAt    field.Time
	UpdateAt    field.Time

	fieldMap map[string]field.Expr
}

func (a apisixRoute) Table(newTableName string) *apisixRoute {
	a.apisixRouteDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a apisixRoute) As(alias string) *apisixRoute {
	a.apisixRouteDo.DO = *(a.apisixRouteDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *apisixRoute) updateTableName(table string) *apisixRoute {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.RouteID = field.NewString(table, "route_id")
	a.Content = field.NewString(table, "content")
	a.ContentYaml = field.NewString(table, "content_yaml")
	a.Type = field.NewInt16(table, "type")
	a.Status = field.NewInt16(table, "status")
	a.CreateAt = field.NewTime(table, "create_at")
	a.UpdateAt = field.NewTime(table, "update_at")

	a.fillFieldMap()

	return a
}

func (a *apisixRoute) WithContext(ctx context.Context) *apisixRouteDo {
	return a.apisixRouteDo.WithContext(ctx)
}

func (a apisixRoute) TableName() string { return a.apisixRouteDo.TableName() }

func (a apisixRoute) Alias() string { return a.apisixRouteDo.Alias() }

func (a *apisixRoute) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *apisixRoute) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["route_id"] = a.RouteID
	a.fieldMap["content"] = a.Content
	a.fieldMap["content_yaml"] = a.ContentYaml
	a.fieldMap["type"] = a.Type
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
}

func (a apisixRoute) clone(db *gorm.DB) apisixRoute {
	a.apisixRouteDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a apisixRoute) replaceDB(db *gorm.DB) apisixRoute {
	a.apisixRouteDo.ReplaceDB(db)
	return a
}

type apisixRouteDo struct{ gen.DO }

func (a apisixRouteDo) Debug() *apisixRouteDo {
	return a.withDO(a.DO.Debug())
}

func (a apisixRouteDo) WithContext(ctx context.Context) *apisixRouteDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a apisixRouteDo) ReadDB() *apisixRouteDo {
	return a.Clauses(dbresolver.Read)
}

func (a apisixRouteDo) WriteDB() *apisixRouteDo {
	return a.Clauses(dbresolver.Write)
}

func (a apisixRouteDo) Session(config *gorm.Session) *apisixRouteDo {
	return a.withDO(a.DO.Session(config))
}

func (a apisixRouteDo) Clauses(conds ...clause.Expression) *apisixRouteDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a apisixRouteDo) Returning(value interface{}, columns ...string) *apisixRouteDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a apisixRouteDo) Not(conds ...gen.Condition) *apisixRouteDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a apisixRouteDo) Or(conds ...gen.Condition) *apisixRouteDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a apisixRouteDo) Select(conds ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a apisixRouteDo) Where(conds ...gen.Condition) *apisixRouteDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a apisixRouteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *apisixRouteDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a apisixRouteDo) Order(conds ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a apisixRouteDo) Distinct(cols ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a apisixRouteDo) Omit(cols ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a apisixRouteDo) Join(table schema.Tabler, on ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a apisixRouteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a apisixRouteDo) RightJoin(table schema.Tabler, on ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a apisixRouteDo) Group(cols ...field.Expr) *apisixRouteDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a apisixRouteDo) Having(conds ...gen.Condition) *apisixRouteDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a apisixRouteDo) Limit(limit int) *apisixRouteDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a apisixRouteDo) Offset(offset int) *apisixRouteDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a apisixRouteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *apisixRouteDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a apisixRouteDo) Unscoped() *apisixRouteDo {
	return a.withDO(a.DO.Unscoped())
}

func (a apisixRouteDo) Create(values ...*model.ApisixRoute) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a apisixRouteDo) CreateInBatches(values []*model.ApisixRoute, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a apisixRouteDo) Save(values ...*model.ApisixRoute) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a apisixRouteDo) First() (*model.ApisixRoute, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApisixRoute), nil
	}
}

func (a apisixRouteDo) Take() (*model.ApisixRoute, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApisixRoute), nil
	}
}

func (a apisixRouteDo) Last() (*model.ApisixRoute, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApisixRoute), nil
	}
}

func (a apisixRouteDo) Find() ([]*model.ApisixRoute, error) {
	result, err := a.DO.Find()
	return result.([]*model.ApisixRoute), err
}

func (a apisixRouteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ApisixRoute, err error) {
	buf := make([]*model.ApisixRoute, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a apisixRouteDo) FindInBatches(result *[]*model.ApisixRoute, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a apisixRouteDo) Attrs(attrs ...field.AssignExpr) *apisixRouteDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a apisixRouteDo) Assign(attrs ...field.AssignExpr) *apisixRouteDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a apisixRouteDo) Joins(fields ...field.RelationField) *apisixRouteDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a apisixRouteDo) Preload(fields ...field.RelationField) *apisixRouteDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a apisixRouteDo) FirstOrInit() (*model.ApisixRoute, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApisixRoute), nil
	}
}

func (a apisixRouteDo) FirstOrCreate() (*model.ApisixRoute, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApisixRoute), nil
	}
}

func (a apisixRouteDo) FindByPage(offset int, limit int) (result []*model.ApisixRoute, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a apisixRouteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a apisixRouteDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a apisixRouteDo) Delete(models ...*model.ApisixRoute) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *apisixRouteDo) withDO(do gen.Dao) *apisixRouteDo {
	a.DO = *do.(*gen.DO)
	return a
}
