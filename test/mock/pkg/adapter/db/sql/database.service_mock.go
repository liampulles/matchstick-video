package sql

import (
	goSql "database/sql"

	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
)

// DatabaseServiceMock is for mocking
type DatabaseServiceMock struct {
	mock.Mock
}

var _ sql.DatabaseService = &DatabaseServiceMock{}

// Get is for mocking
func (m *DatabaseServiceMock) Get() *goSql.DB {
	args := m.Called()
	return safeArgsGetSQLDb(args, 0)
}

func safeArgsGetSQLDb(args mock.Arguments, idx int) *goSql.DB {
	if val, ok := args.Get(idx).(*goSql.DB); ok {
		return val
	}
	return nil
}
