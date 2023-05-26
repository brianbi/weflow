// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/wegoteam/weflow/pkg/model"
)

func newModelVersion(db *gorm.DB, opts ...gen.DOOption) modelVersion {
	_modelVersion := modelVersion{}

	_modelVersion.modelVersionDo.UseDB(db, opts...)
	_modelVersion.modelVersionDo.UseModel(&model.ModelVersion{})

	tableName := _modelVersion.modelVersionDo.TableName()
	_modelVersion.ALL = field.NewAsterisk(tableName)
	_modelVersion.ID = field.NewInt64(tableName, "id")
	_modelVersion.ModelID = field.NewString(tableName, "model_id")
	_modelVersion.ModelTitle = field.NewString(tableName, "model_title")
	_modelVersion.VersionID = field.NewString(tableName, "version_id")
	_modelVersion.ProcessDefID = field.NewString(tableName, "process_def_id")
	_modelVersion.FormDefID = field.NewString(tableName, "form_def_id")
	_modelVersion.UseStatus = field.NewInt32(tableName, "use_status")
	_modelVersion.Remark = field.NewString(tableName, "remark")
	_modelVersion.CreateTime = field.NewTime(tableName, "create_time")
	_modelVersion.CreateUser = field.NewString(tableName, "create_user")
	_modelVersion.UpdateTime = field.NewTime(tableName, "update_time")
	_modelVersion.UpdateUser = field.NewString(tableName, "update_user")
	_modelVersion.NoticeURL = field.NewString(tableName, "notice_url")
	_modelVersion.TitleProps = field.NewString(tableName, "title_props")

	_modelVersion.fillFieldMap()

	return _modelVersion
}

type modelVersion struct {
	modelVersionDo modelVersionDo

	ALL          field.Asterisk
	ID           field.Int64  // 唯一id
	ModelID      field.String // 模板id
	ModelTitle   field.String // 模板版本标题
	VersionID    field.String // 版本id
	ProcessDefID field.String // 流程定义id
	FormDefID    field.String // 表单定义id
	UseStatus    field.Int32  // 使用状态【1：使用；2：未使用】
	Remark       field.String // 描述
	CreateTime   field.Time   // 创建时间
	CreateUser   field.String // 创建人
	UpdateTime   field.Time   // 更新时间
	UpdateUser   field.String // 更新人
	NoticeURL    field.String // 回调通知推送url
	TitleProps   field.String // 标题配置

	fieldMap map[string]field.Expr
}

func (m modelVersion) Table(newTableName string) *modelVersion {
	m.modelVersionDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m modelVersion) As(alias string) *modelVersion {
	m.modelVersionDo.DO = *(m.modelVersionDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *modelVersion) updateTableName(table string) *modelVersion {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.ModelID = field.NewString(table, "model_id")
	m.ModelTitle = field.NewString(table, "model_title")
	m.VersionID = field.NewString(table, "version_id")
	m.ProcessDefID = field.NewString(table, "process_def_id")
	m.FormDefID = field.NewString(table, "form_def_id")
	m.UseStatus = field.NewInt32(table, "use_status")
	m.Remark = field.NewString(table, "remark")
	m.CreateTime = field.NewTime(table, "create_time")
	m.CreateUser = field.NewString(table, "create_user")
	m.UpdateTime = field.NewTime(table, "update_time")
	m.UpdateUser = field.NewString(table, "update_user")
	m.NoticeURL = field.NewString(table, "notice_url")
	m.TitleProps = field.NewString(table, "title_props")

	m.fillFieldMap()

	return m
}

func (m *modelVersion) WithContext(ctx context.Context) *modelVersionDo {
	return m.modelVersionDo.WithContext(ctx)
}

func (m modelVersion) TableName() string { return m.modelVersionDo.TableName() }

func (m modelVersion) Alias() string { return m.modelVersionDo.Alias() }

func (m *modelVersion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *modelVersion) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 14)
	m.fieldMap["id"] = m.ID
	m.fieldMap["model_id"] = m.ModelID
	m.fieldMap["model_title"] = m.ModelTitle
	m.fieldMap["version_id"] = m.VersionID
	m.fieldMap["process_def_id"] = m.ProcessDefID
	m.fieldMap["form_def_id"] = m.FormDefID
	m.fieldMap["use_status"] = m.UseStatus
	m.fieldMap["remark"] = m.Remark
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["create_user"] = m.CreateUser
	m.fieldMap["update_time"] = m.UpdateTime
	m.fieldMap["update_user"] = m.UpdateUser
	m.fieldMap["notice_url"] = m.NoticeURL
	m.fieldMap["title_props"] = m.TitleProps
}

func (m modelVersion) clone(db *gorm.DB) modelVersion {
	m.modelVersionDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m modelVersion) replaceDB(db *gorm.DB) modelVersion {
	m.modelVersionDo.ReplaceDB(db)
	return m
}

type modelVersionDo struct{ gen.DO }

func (m modelVersionDo) Debug() *modelVersionDo {
	return m.withDO(m.DO.Debug())
}

func (m modelVersionDo) WithContext(ctx context.Context) *modelVersionDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m modelVersionDo) ReadDB() *modelVersionDo {
	return m.Clauses(dbresolver.Read)
}

func (m modelVersionDo) WriteDB() *modelVersionDo {
	return m.Clauses(dbresolver.Write)
}

func (m modelVersionDo) Session(config *gorm.Session) *modelVersionDo {
	return m.withDO(m.DO.Session(config))
}

func (m modelVersionDo) Clauses(conds ...clause.Expression) *modelVersionDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m modelVersionDo) Returning(value interface{}, columns ...string) *modelVersionDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m modelVersionDo) Not(conds ...gen.Condition) *modelVersionDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m modelVersionDo) Or(conds ...gen.Condition) *modelVersionDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m modelVersionDo) Select(conds ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m modelVersionDo) Where(conds ...gen.Condition) *modelVersionDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m modelVersionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *modelVersionDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m modelVersionDo) Order(conds ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m modelVersionDo) Distinct(cols ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m modelVersionDo) Omit(cols ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m modelVersionDo) Join(table schema.Tabler, on ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m modelVersionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m modelVersionDo) RightJoin(table schema.Tabler, on ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m modelVersionDo) Group(cols ...field.Expr) *modelVersionDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m modelVersionDo) Having(conds ...gen.Condition) *modelVersionDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m modelVersionDo) Limit(limit int) *modelVersionDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m modelVersionDo) Offset(offset int) *modelVersionDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m modelVersionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *modelVersionDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m modelVersionDo) Unscoped() *modelVersionDo {
	return m.withDO(m.DO.Unscoped())
}

func (m modelVersionDo) Create(values ...*model.ModelVersion) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m modelVersionDo) CreateInBatches(values []*model.ModelVersion, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m modelVersionDo) Save(values ...*model.ModelVersion) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m modelVersionDo) First() (*model.ModelVersion, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelVersion), nil
	}
}

func (m modelVersionDo) Take() (*model.ModelVersion, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelVersion), nil
	}
}

func (m modelVersionDo) Last() (*model.ModelVersion, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelVersion), nil
	}
}

func (m modelVersionDo) Find() ([]*model.ModelVersion, error) {
	result, err := m.DO.Find()
	return result.([]*model.ModelVersion), err
}

func (m modelVersionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ModelVersion, err error) {
	buf := make([]*model.ModelVersion, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m modelVersionDo) FindInBatches(result *[]*model.ModelVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m modelVersionDo) Attrs(attrs ...field.AssignExpr) *modelVersionDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m modelVersionDo) Assign(attrs ...field.AssignExpr) *modelVersionDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m modelVersionDo) Joins(fields ...field.RelationField) *modelVersionDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m modelVersionDo) Preload(fields ...field.RelationField) *modelVersionDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m modelVersionDo) FirstOrInit() (*model.ModelVersion, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelVersion), nil
	}
}

func (m modelVersionDo) FirstOrCreate() (*model.ModelVersion, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelVersion), nil
	}
}

func (m modelVersionDo) FindByPage(offset int, limit int) (result []*model.ModelVersion, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m modelVersionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m modelVersionDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m modelVersionDo) Delete(models ...*model.ModelVersion) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *modelVersionDo) withDO(do gen.Dao) *modelVersionDo {
	m.DO = *do.(*gen.DO)
	return m
}
