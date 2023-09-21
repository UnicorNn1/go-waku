//go:build include_postgres_tests
// +build include_postgres_tests

package postgres

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/waku-org/go-waku/waku/persistence"
)

func TestQueries(t *testing.T) {
	db := persistence.NewMockPgDB()

	queries, err := NewQueries("test_queries", db)
	require.NoError(t, err)

	query := queries.Delete()
	require.NotEmpty(t, query)

	query = queries.Exists()
	require.NotEmpty(t, query)

	query = queries.Get()
	require.NotEmpty(t, query)

	query = queries.Put()
	require.NotEmpty(t, query)

	query = queries.Query()
	require.NotEmpty(t, query)

	query = queries.Prefix()
	require.NotEmpty(t, query)

	query = queries.Limit()
	require.NotEmpty(t, query)

	query = queries.Offset()
	require.NotEmpty(t, query)

	query = queries.GetSize()
	require.NotEmpty(t, query)
}

func TestCreateTable(t *testing.T) {
	db := persistence.NewMockPgDB()

	err := CreateTable(db, "test_create_table")
	require.NoError(t, err)
}
