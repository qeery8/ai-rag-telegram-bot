package repository

import (
	"al-rag-tgbot/internal/domain"
	"al-rag-tgbot/internal/repository/postgres/user"
	"context"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Suite struct {
	db *sqlx.DB
}

func setupDB(t *testing.T) *Suite {
	t.Helper()

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dsn = "host=localhost port=5432 user=admin dbname=rag_test sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Skipf("no database avaible, skipping: %v", err)
	}

	return &Suite{db: db}
}

func TestUserRepo_Integration_SaveUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	suite := setupDB(t)

	repo := user.NewUserRepoSQL(suite.db)
	ctx := context.Background()

	t.Run("save and return id", func(t *testing.T) {
		u := &domain.User{Username: "integrationuser"}

		err := repo.SaveUser(ctx, u)

		require.NoError(t, err)
		assert.NotZero(t, u.ID)
		assert.NotZero(t, u.CreatedAt)
	})

	t.Run("find save user", func(t *testing.T) {
		u := &domain.User{Username: "findme"}

		err := repo.SaveUser(ctx, u)
		require.NoError(t, err)

		found, err := repo.FindByID(ctx, u.ID)

		require.NoError(t, err)
		assert.Equal(t, u.ID, found.ID)
		assert.Equal(t, "findme", found.Username)
	})
}
