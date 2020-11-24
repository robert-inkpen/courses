/*
package db
import (
	"fmt"
	"strings"
)
//SELECT ..... FROM table
//WHERE ......
//OFFSET .. LIMIT ..
//queryBuilder.Select(...).From(...).Where(...).Where(...).GetQuery()
type queryBuilder struct {
	selectClause string
	whereClause  string
	table        string
	offset       string
	limit        string
}
func NewQueryBuilder() queryBuilder {
	return queryBuilder{}
}
func (q queryBuilder) From(tablename string) queryBuilder {
	q.table = tablename
	return q
}
func (q queryBuilder) Select(fields ...string) queryBuilder {
	q.selectClause = strings.Join(fields, ", ")
	return q
}
func (q queryBuilder) ToSql() string {
	sql := "SELECT %s FROM %s"
	sql = fmt.Sprintf(sql, q.selectClause, q.table)
	return sql
}
*/