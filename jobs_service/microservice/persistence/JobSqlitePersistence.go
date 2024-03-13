package persistence

import (
	"context"
	"strings"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cerr "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	csqlite "github.com/pip-services3-gox/pip-services3-sqlite-gox/persistence"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type JobSqlitePersistence struct {
	*csqlite.IdentifiableSqlitePersistence[data1.JobV1, string]
}

func NewJobSqlitePersistence() *JobSqlitePersistence {
	c := &JobSqlitePersistence{}
	c.IdentifiableSqlitePersistence = csqlite.InheritIdentifiableSqlitePersistence[data1.JobV1, string](c, "jobs")
	return c
}

func (c *JobSqlitePersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiableSqlitePersistence.DefineSchema()
	c.EnsureSchema("CREATE TABLE " + c.QuotedTableName() + " (\"id\" TEXT PRIMARY KEY, \"owner\" TEXT, \"status\" TEXT)")

}

func (c *JobSqlitePersistence) composeFilter(filter cdata.FilterParams) string {
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

func (c *JobSqlitePersistence) GetNotStartedJob(ctx context.Context, correlationId string) (data1.JobV1, error) {
	query := "SELECT * FROM " + c.QuotedTableName() + " WHERE \"status\"= \"not_started\" ORDER BY \"id\" LIMIT 1"

	qResult, err := c.Client.QueryContext(ctx, query)
	if err != nil {
		return data1.JobV1{}, err
	}
	defer qResult.Close()

	if !qResult.Next() {
		return data1.JobV1{}, qResult.Err()
	}

	result, err := c.Overrides.ConvertToPublic(qResult)

	if err == nil {
		c.Logger.Trace(ctx, correlationId, "Retrieved from %s with owner = %s", c.TableName)
		return result, err
	}
	c.Logger.Trace(ctx, correlationId, "Nothing found from %s with owner = %s", c.TableName)
	return data1.JobV1{}, err
}

func (c *JobSqlitePersistence) GetPageByFilter(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[data1.JobV1], err error) {
	return c.IdentifiableSqlitePersistence.GetPageByFilter(ctx, correlationId,
		c.composeFilter(filter), paging,
		"", "",
	)
}

func (c *JobSqlitePersistence) GetPageByStatus(ctx context.Context, correlationId string, status string) (page cdata.DataPage[data1.JobV1], err error) {
	query := "SELECT * FROM " + c.QuotedTableName() + " WHERE \"status\"=$1"

	rows, err := c.Client.QueryContext(ctx, query, status)
	if err != nil {
		return *cdata.NewEmptyDataPage[data1.JobV1](), err
	}
	defer rows.Close()

	// if !rows.Next() {
	// 	return *cdata.NewEmptyDataPage[data1.JobV1](), rows.Err()
	// }

	items := make([]data1.JobV1, 0)
	for rows.Next() {
		if c.IsTerminated() {
			rows.Close()
			return *cdata.NewEmptyDataPage[data1.JobV1](), cerr.
				NewError("query terminated").
				WithCorrelationId(correlationId)
		}
		item, convErr := c.Overrides.ConvertToPublic(rows)
		if convErr != nil {
			return page, convErr
		}
		items = append(items, item)
	}

	if items != nil {
		c.Logger.Trace(ctx, correlationId, "Retrieved %d from %s", len(items), c.TableName)
	}

	return *cdata.NewDataPage(items, cdata.EmptyTotalValue), rows.Err()
}
