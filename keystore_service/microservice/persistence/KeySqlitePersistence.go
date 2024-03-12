package persistence

import (
	"context"
	"strings"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	csqlite "github.com/pip-services3-gox/pip-services3-sqlite-gox/persistence"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
)

type KeySqlitePersistence struct {
	*csqlite.IdentifiableSqlitePersistence[data1.KeyV1, string]
}

func NewKeysSqlitePersistence() *KeySqlitePersistence {
	c := &KeySqlitePersistence{}
	c.IdentifiableSqlitePersistence = csqlite.InheritIdentifiableSqlitePersistence[data1.KeyV1, string](c, "keys")
	return c
}

func (c *KeySqlitePersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiableSqlitePersistence.DefineSchema()
	c.EnsureSchema("CREATE TABLE " + c.QuotedTableName() + " (\"id\" TEXT PRIMARY KEY, \"key\" TEXT, \"owner\" TEXT)")
	c.EnsureIndex(c.TableName+"_key", map[string]string{"key": "1"}, map[string]string{"unique": "true"})
}

func (c *KeySqlitePersistence) composeFilter(filter cdata.FilterParams) string {
	criteria := make([]string, 0)

	if key, ok := filter.GetAsNullableString("key"); ok && key != "" {
		criteria = append(criteria, "key='"+key+"'")
	}

	if id, ok := filter.GetAsNullableString("id"); ok && id != "" {
		criteria = append(criteria, "id='"+id+"'")
	}

	if tempIds, ok := filter.GetAsNullableString("ids"); ok && tempIds != "" {
		ids := strings.Split(tempIds, ",")
		criteria = append(criteria, "id IN ('"+strings.Join(ids, "','")+"')")
	}

	if len(criteria) > 0 {
		return strings.Join(criteria, " AND ")
	} else {
		return ""
	}
}

func (c *KeySqlitePersistence) GetOneByOwner(ctx context.Context, correlationId string, owner string) (data1.KeyV1, error) {
	query := "SELECT * FROM " + c.QuotedTableName() + " WHERE \"owner\"=$1"

	qResult, err := c.Client.QueryContext(ctx, query, owner)
	if err != nil {
		return data1.KeyV1{}, err
	}
	defer qResult.Close()

	if !qResult.Next() {
		return data1.KeyV1{}, qResult.Err()
	}

	result, err := c.Overrides.ConvertToPublic(qResult)

	if err == nil {
		c.Logger.Trace(ctx, correlationId, "Retrieved from %s with owner = %s", c.TableName, owner)
		return result, err
	}
	c.Logger.Trace(ctx, correlationId, "Nothing found from %s with owner = %s", c.TableName, owner)
	return data1.KeyV1{}, err
}

func (c *KeySqlitePersistence) GetPageByFilter(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[data1.KeyV1], err error) {
	return c.IdentifiableSqlitePersistence.GetPageByFilter(ctx, correlationId,
		c.composeFilter(filter), paging,
		"", "",
	)
}
