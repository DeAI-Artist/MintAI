package bridges_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmoiron/sqlx"

	"github.com/DeAI-Artist/MintAI/core/auth"
	"github.com/DeAI-Artist/MintAI/core/bridges"
	"github.com/DeAI-Artist/MintAI/core/internal/cltest"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils/configtest"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils/pgtest"
	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services/pg"
	"github.com/DeAI-Artist/MintAI/core/services/pipeline"
	"github.com/DeAI-Artist/MintAI/core/store/models"
)

func setupORM(t *testing.T) (*sqlx.DB, bridges.ORM) {
	t.Helper()

	cfg := configtest.NewGeneralConfig(t, nil)
	db := pgtest.NewSqlxDB(t)
	orm := bridges.NewORM(db, logger.TestLogger(t), cfg.Database())

	return db, orm
}

func TestORM_FindBridges(t *testing.T) {
	t.Parallel()
	_, orm := setupORM(t)

	bt := bridges.BridgeType{
		Name: "bridge1",
		URL:  cltest.WebURL(t, "https://bridge1.com"),
	}
	assert.NoError(t, orm.CreateBridgeType(&bt))
	bt2 := bridges.BridgeType{
		Name: "bridge2",
		URL:  cltest.WebURL(t, "https://bridge2.com"),
	}
	assert.NoError(t, orm.CreateBridgeType(&bt2))
	bts, err := orm.FindBridges([]bridges.BridgeName{"bridge2", "bridge1"})
	require.NoError(t, err)
	require.Equal(t, 2, len(bts))

	bts, err = orm.FindBridges([]bridges.BridgeName{"bridge1"})
	require.NoError(t, err)
	require.Equal(t, 1, len(bts))
	require.Equal(t, "bridge1", bts[0].Name.String())

	// One invalid bridge errors
	bts, err = orm.FindBridges([]bridges.BridgeName{"bridge1", "bridgeX"})
	require.Error(t, err, bts)

	// All invalid bridges error
	bts, err = orm.FindBridges([]bridges.BridgeName{"bridgeY", "bridgeX"})
	require.Error(t, err, bts)

	// Requires at least one bridge
	bts, err = orm.FindBridges([]bridges.BridgeName{})
	require.Error(t, err, bts)
}

func TestORM_FindBridge(t *testing.T) {
	t.Parallel()

	_, orm := setupORM(t)

	bt := bridges.BridgeType{}
	bt.Name = bridges.MustParseBridgeName("solargridreporting")
	bt.URL = cltest.WebURL(t, "https://denergy.eth")
	assert.NoError(t, orm.CreateBridgeType(&bt))

	cases := []struct {
		description string
		name        bridges.BridgeName
		want        bridges.BridgeType
		errored     bool
	}{
		{"actual external adapter", bt.Name, bt, false},
		{"core adapter", "ethtx", bridges.BridgeType{}, true},
		{"non-existent adapter", "nonExistent", bridges.BridgeType{}, true},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			tt, err := orm.FindBridge(test.name)
			tt.CreatedAt = test.want.CreatedAt
			tt.UpdatedAt = test.want.UpdatedAt
			if test.errored {
				require.Error(t, err)
			} else {
				// we can't make any assumptions about the return type if scanning failed
				require.Equal(t, test.want, tt)
			}
		})
	}
}
func TestORM_UpdateBridgeType(t *testing.T) {
	_, orm := setupORM(t)

	firstBridge := &bridges.BridgeType{
		Name: "UniqueName",
		URL:  cltest.WebURL(t, "http:/oneurl.com"),
	}

	require.NoError(t, orm.CreateBridgeType(firstBridge))

	updateBridge := &bridges.BridgeTypeRequest{
		URL: cltest.WebURL(t, "http:/updatedurl.com"),
	}

	require.NoError(t, orm.UpdateBridgeType(firstBridge, updateBridge))

	foundbridge, err := orm.FindBridge("UniqueName")
	require.NoError(t, err)
	require.Equal(t, updateBridge.URL, foundbridge.URL)

	bs, count, err := orm.BridgeTypes(0, 10)
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.Len(t, bs, 1)

	require.NoError(t, orm.DeleteBridgeType(&foundbridge))

	bs, count, err = orm.BridgeTypes(0, 10)
	require.NoError(t, err)
	require.Equal(t, 0, count)
	require.Len(t, bs, 0)
}

func TestORM_TestCachedResponse(t *testing.T) {
	cfg := configtest.NewGeneralConfig(t, nil)
	db := pgtest.NewSqlxDB(t)
	orm := bridges.NewORM(db, logger.TestLogger(t), cfg.Database())

	trORM := pipeline.NewORM(db, logger.TestLogger(t), cfg.Database(), cfg.JobPipeline().MaxSuccessfulRuns())
	specID, err := trORM.CreateSpec(pipeline.Pipeline{}, *models.NewInterval(5 * time.Minute), pg.WithParentCtx(testutils.Context(t)))
	require.NoError(t, err)

	_, err = orm.GetCachedResponse("dot", specID, 1*time.Second)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no rows in result set")

	err = orm.UpsertBridgeResponse("dot", specID, []byte{111, 222, 2})
	require.NoError(t, err)

	val, err := orm.GetCachedResponse("dot", specID, 1*time.Second)
	require.NoError(t, err)
	require.Equal(t, []byte{111, 222, 2}, val)
}

func TestORM_CreateExternalInitiator(t *testing.T) {
	_, orm := setupORM(t)

	token := auth.NewToken()
	name := uuid.New().String()
	req := bridges.ExternalInitiatorRequest{
		Name: name,
	}
	exi, err := bridges.NewExternalInitiator(token, &req)
	require.NoError(t, err)
	require.NoError(t, orm.CreateExternalInitiator(exi))

	exi2, err := bridges.NewExternalInitiator(token, &req)
	require.NoError(t, err)
	require.Contains(t, orm.CreateExternalInitiator(exi2).Error(), `ERROR: duplicate key value violates unique constraint "external_initiators_name_key" (SQLSTATE 23505)`)
}

func TestORM_DeleteExternalInitiator(t *testing.T) {
	_, orm := setupORM(t)

	token := auth.NewToken()
	name := uuid.New().String()
	req := bridges.ExternalInitiatorRequest{
		Name: name,
	}
	exi, err := bridges.NewExternalInitiator(token, &req)
	require.NoError(t, err)
	require.NoError(t, orm.CreateExternalInitiator(exi))

	_, err = orm.FindExternalInitiator(token)
	require.NoError(t, err)
	_, err = orm.FindExternalInitiatorByName(exi.Name)
	require.NoError(t, err)

	err = orm.DeleteExternalInitiator(exi.Name)
	require.NoError(t, err)

	_, err = orm.FindExternalInitiator(token)
	require.Error(t, err)
	_, err = orm.FindExternalInitiatorByName(exi.Name)
	require.Error(t, err)

	require.NoError(t, orm.CreateExternalInitiator(exi))
}
