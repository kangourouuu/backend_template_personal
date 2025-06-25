package repository

import "context"

type (
	IElasticsearch interface {
		CheckAliasExist(ctx context.Context, index, alias string) (bool, error)
		CreateAlias(ctx context.Context, index, alias string) error
		// CreateDocRabbitMQ(ctx context.Context, index, tenant, routing, uuid string, esDoc map[string]any) (bool, error)
		// CreateAliasRabbitMQ(ctx context.Context, index, alias string) (bool, error)
		InsertLog(ctx context.Context, tenantId, index, docId string, esDoc map[string]any) error
		UpdateDocById(ctx context.Context, index, docId string, esDoc map[string]any) error
		BulkUpdateDoc(ctx context.Context, index string, esDoc map[string]any) error
		DeleteById(ctx context.Context, index, docId string) error
	}
)

var ESRepo IElasticsearch
