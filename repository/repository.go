package repository

import (
	"context"
	"strings"

	sqlclient "backend_template_personal/internal/sqlclient"

	"backend_template_personal/model"
	"gorm.io/gorm"

	"github.com/uptrace/bun"

	"github.com/uptrace/bun/schema"
)

var FusionSqlClient sqlclient.ISqlClientConn
var GormSqlClient sqlclient.IGormSqlClientConn

func CreateTableCollate(client sqlclient.ISqlClientConn, ctx context.Context, table interface{}) error {
	query := client.GetDB().NewCreateTable().Model(table).IfNotExists()
	value, _ := query.AppendQuery(schema.NewFormatter(query.Dialect()), nil)
	queryStr := string(value) + " COLLATE utf8mb4_general_ci"
	_, err := client.GetDB().ExecContext(ctx, queryStr)
	return err
}

func CreateTable(client sqlclient.ISqlClientConn, ctx context.Context, table interface{}) error {
	query := client.GetDB().NewCreateTable().Model(table).IfNotExists()
	value, _ := query.AppendQuery(schema.NewFormatter(query.Dialect()), nil)
	queryStr := string(value)
	if client.GetDriver() == sqlclient.POSTGRESQL {
		queryStr = strings.ReplaceAll(queryStr, " char(36)", " uuid")
		queryStr = strings.ReplaceAll(queryStr, " timestamp", " timestamptz")
		queryStr = strings.ReplaceAll(queryStr, " timestamptz_only", " timestamp")
	}
	_, err := client.GetDB().ExecContext(ctx, queryStr)
	return err
}

func AddColumn(client sqlclient.ISqlClientConn, ctx context.Context, table interface{}, column string) error {
	query := client.GetDB().NewAddColumn().Model(table).IfNotExists().ColumnExpr(column)
	value, _ := query.AppendQuery(schema.NewFormatter(query.Dialect()), nil)
	queryStr := string(value)
	if client.GetDriver() == sqlclient.POSTGRESQL {
		queryStr = strings.ReplaceAll(queryStr, " char(36)", " uuid")
		queryStr = strings.ReplaceAll(queryStr, " timestamp", " timestamptz")
	}
	_, err := client.GetDB().ExecContext(ctx, queryStr)
	return err
}

func AddColumns(client sqlclient.ISqlClientConn, ctx context.Context, table interface{}, query string, columns ...bun.Ident) error {
	args := make([]interface{}, len(columns))
	for i, column := range columns {
		args[i] = column
	}
	addColumnQuery := client.GetDB().NewAddColumn().Model(table).ColumnExpr(query, args...)
	value, _ := addColumnQuery.AppendQuery(schema.NewFormatter(addColumnQuery.Dialect()), nil)
	queryStr := string(value)
	if client.GetDriver() == sqlclient.POSTGRESQL {
		queryStr = strings.ReplaceAll(queryStr, " char(36)", " uuid")
		queryStr = strings.ReplaceAll(queryStr, " timestamp", " timestamptz")
	}
	_, err := client.GetDB().ExecContext(ctx, queryStr)
	return err
}

type ProjectRepository interface {
	Create(*model.ProjectInfo) error
}

type projectRepository struct {
	db *gorm.DB
}

func (r *projectRepository) migrate() error {
	return r.db.AutoMigrate(&model.ProjectInfo{})
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	rep := &projectRepository{db: db}
	_ = rep.migrate()
	return rep
}

func (r *projectRepository) Create(info *model.ProjectInfo) error {
	return r.db.Create(info).Error
}
