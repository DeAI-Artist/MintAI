package keystore

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jmoiron/sqlx"

	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services/keystore/keys/ethkey"
	"github.com/DeAI-Artist/MintAI/core/services/pg"
	"github.com/DeAI-Artist/MintAI/core/utils"
)

func mustNewEthKey(t *testing.T) *ethkey.KeyV2 {
	key, err := ethkey.NewV2()
	require.NoError(t, err)
	return &key
}

func ExposedNewMaster(t *testing.T, db *sqlx.DB, cfg pg.QConfig) *master {
	return newMaster(db, utils.FastScryptParams, logger.TestLogger(t), cfg)
}

func (m *master) ExportedSave() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.save()
}

func (m *master) ResetXXXTestOnly() {
	m.keyRing = newKeyRing()
	m.keyStates = newKeyStates()
	m.password = ""
}

func (m *master) SetPassword(pw string) {
	m.password = pw
}
