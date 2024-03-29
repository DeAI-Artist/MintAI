package pg

import (
	"testing"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/store/dialects"
)

func Test_disallowReplica(t *testing.T) {

	testutils.SkipShortDB(t)
	db, err := sqlx.Open(string(dialects.TransactionWrappedPostgres), uuid.New().String())
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, db.Close()) })

	_, err = db.Exec("SET session_replication_role= 'origin'")
	require.NoError(t, err)
	err = disallowReplica(db)
	require.NoError(t, err)

	_, err = db.Exec("SET session_replication_role= 'replica'")
	require.NoError(t, err)
	err = disallowReplica(db)
	require.Error(t, err, "replica role should be disallowed")

	_, err = db.Exec("SET session_replication_role= 'not_valid_role'")
	require.Error(t, err)

}
