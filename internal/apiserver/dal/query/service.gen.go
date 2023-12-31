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

func newService(db *gorm.DB, opts ...gen.DOOption) service {
	_service := service{}

	_service.serviceDo.UseDB(db, opts...)
	_service.serviceDo.UseModel(&model.Service{})

	tableName := _service.serviceDo.TableName()
	_service.ALL = field.NewAsterisk(tableName)
	_service.ID = field.NewInt32(tableName, "id")
	_service.Name = field.NewString(tableName, "name")
	_service.UpstreamID = field.NewString(tableName, "upstream_id")
	_service.APISet = field.NewField(tableName, "api_set")
	_service.CreatorID = field.NewInt32(tableName, "creator_id")
	_service.LastEditorID = field.NewInt32(tableName, "last_editor_id")
	_service.Abandoned = field.NewBool(tableName, "abandoned")
	_service.DocumentID = field.NewInt32(tableName, "document_id")
	_service.ServiceType = field.NewString(tableName, "service_type")
	_service.DelUniqueKey = field.NewInt32(tableName, "del_unique_key")
	_service.Ctime = field.NewTime(tableName, "ctime")
	_service.LastUpdateTime = field.NewTime(tableName, "last_update_time")

	_service.fillFieldMap()

	return _service
}

type service struct {
	serviceDo serviceDo

	ALL            field.Asterisk
	ID             field.Int32
	Name           field.String
	UpstreamID     field.String
	APISet         field.Field
	CreatorID      field.Int32
	LastEditorID   field.Int32
	Abandoned      field.Bool
	DocumentID     field.Int32
	ServiceType    field.String
	DelUniqueKey   field.Int32
	Ctime          field.Time
	LastUpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (s service) Table(newTableName string) *service {
	s.serviceDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s service) As(alias string) *service {
	s.serviceDo.DO = *(s.serviceDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *service) updateTableName(table string) *service {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Name = field.NewString(table, "name")
	s.UpstreamID = field.NewString(table, "upstream_id")
	s.APISet = field.NewField(table, "api_set")
	s.CreatorID = field.NewInt32(table, "creator_id")
	s.LastEditorID = field.NewInt32(table, "last_editor_id")
	s.Abandoned = field.NewBool(table, "abandoned")
	s.DocumentID = field.NewInt32(table, "document_id")
	s.ServiceType = field.NewString(table, "service_type")
	s.DelUniqueKey = field.NewInt32(table, "del_unique_key")
	s.Ctime = field.NewTime(table, "ctime")
	s.LastUpdateTime = field.NewTime(table, "last_update_time")

	s.fillFieldMap()

	return s
}

func (s *service) WithContext(ctx context.Context) *serviceDo { return s.serviceDo.WithContext(ctx) }

func (s service) TableName() string { return s.serviceDo.TableName() }

func (s service) Alias() string { return s.serviceDo.Alias() }

func (s *service) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *service) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["upstream_id"] = s.UpstreamID
	s.fieldMap["api_set"] = s.APISet
	s.fieldMap["creator_id"] = s.CreatorID
	s.fieldMap["last_editor_id"] = s.LastEditorID
	s.fieldMap["abandoned"] = s.Abandoned
	s.fieldMap["document_id"] = s.DocumentID
	s.fieldMap["service_type"] = s.ServiceType
	s.fieldMap["del_unique_key"] = s.DelUniqueKey
	s.fieldMap["ctime"] = s.Ctime
	s.fieldMap["last_update_time"] = s.LastUpdateTime
}

func (s service) clone(db *gorm.DB) service {
	s.serviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s service) replaceDB(db *gorm.DB) service {
	s.serviceDo.ReplaceDB(db)
	return s
}

type serviceDo struct{ gen.DO }

func (s serviceDo) Debug() *serviceDo {
	return s.withDO(s.DO.Debug())
}

func (s serviceDo) WithContext(ctx context.Context) *serviceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serviceDo) ReadDB() *serviceDo {
	return s.Clauses(dbresolver.Read)
}

func (s serviceDo) WriteDB() *serviceDo {
	return s.Clauses(dbresolver.Write)
}

func (s serviceDo) Session(config *gorm.Session) *serviceDo {
	return s.withDO(s.DO.Session(config))
}

func (s serviceDo) Clauses(conds ...clause.Expression) *serviceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serviceDo) Returning(value interface{}, columns ...string) *serviceDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serviceDo) Not(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serviceDo) Or(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serviceDo) Select(conds ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serviceDo) Where(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serviceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *serviceDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serviceDo) Order(conds ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serviceDo) Distinct(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serviceDo) Omit(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serviceDo) Join(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serviceDo) Group(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serviceDo) Having(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serviceDo) Limit(limit int) *serviceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serviceDo) Offset(offset int) *serviceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *serviceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serviceDo) Unscoped() *serviceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serviceDo) Create(values ...*model.Service) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serviceDo) CreateInBatches(values []*model.Service, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serviceDo) Save(values ...*model.Service) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serviceDo) First() (*model.Service, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Take() (*model.Service, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Last() (*model.Service, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Find() ([]*model.Service, error) {
	result, err := s.DO.Find()
	return result.([]*model.Service), err
}

func (s serviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Service, err error) {
	buf := make([]*model.Service, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serviceDo) FindInBatches(result *[]*model.Service, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serviceDo) Attrs(attrs ...field.AssignExpr) *serviceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serviceDo) Assign(attrs ...field.AssignExpr) *serviceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serviceDo) Joins(fields ...field.RelationField) *serviceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serviceDo) Preload(fields ...field.RelationField) *serviceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serviceDo) FirstOrInit() (*model.Service, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) FirstOrCreate() (*model.Service, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) FindByPage(offset int, limit int) (result []*model.Service, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s serviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serviceDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serviceDo) Delete(models ...*model.Service) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serviceDo) withDO(do gen.Dao) *serviceDo {
	s.DO = *do.(*gen.DO)
	return s
}
