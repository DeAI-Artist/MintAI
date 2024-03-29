package localauth_test

import (
	"testing"
	"time"

	"github.com/onsi/gomega"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/DeAI-Artist/MintAI/core/internal/cltest"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils/pgtest"
	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/logger/audit"
	"github.com/DeAI-Artist/MintAI/core/sessions"
	"github.com/DeAI-Artist/MintAI/core/sessions/localauth"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
)

type sessionReaperConfig struct{}

func (c sessionReaperConfig) SessionTimeout() commonconfig.Duration {
	return *commonconfig.MustNewDuration(42 * time.Second)
}

func (c sessionReaperConfig) SessionReaperExpiration() commonconfig.Duration {
	return *commonconfig.MustNewDuration(142 * time.Second)
}

func TestSessionReaper_ReapSessions(t *testing.T) {
	t.Parallel()

	db := pgtest.NewSqlxDB(t)
	config := sessionReaperConfig{}
	lggr := logger.TestLogger(t)
	orm := localauth.NewORM(db, config.SessionTimeout().Duration(), lggr, pgtest.NewQConfig(true), audit.NoopLogger)

	r := localauth.NewSessionReaper(db.DB, config, lggr)
	t.Cleanup(func() {
		assert.NoError(t, r.Stop())
	})

	tests := []struct {
		name     string
		lastUsed time.Time
		wantReap bool
	}{
		{"current", time.Now(), false},
		{"expired", time.Now().Add(-config.SessionTimeout().Duration()), false},
		{"almost stale", time.Now().Add(-config.SessionReaperExpiration().Duration()), false},
		{"stale", time.Now().Add(-config.SessionReaperExpiration().Duration()).
			Add(-config.SessionTimeout().Duration()), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				_, err2 := db.Exec("DELETE FROM sessions where email = $1", cltest.APIEmailAdmin)
				require.NoError(t, err2)
			})

			_, err := db.Exec("INSERT INTO sessions (last_used, email, id, created_at) VALUES ($1, $2, $3, now())", test.lastUsed, cltest.APIEmailAdmin, test.name)
			require.NoError(t, err)

			r.WakeUp()

			if test.wantReap {
				gomega.NewWithT(t).Eventually(func() []sessions.Session {
					sessions, err := orm.Sessions(0, 10)
					assert.NoError(t, err)
					return sessions
				}).Should(gomega.HaveLen(0))
			} else {
				gomega.NewWithT(t).Consistently(func() []sessions.Session {
					sessions, err := orm.Sessions(0, 10)
					assert.NoError(t, err)
					return sessions
				}).Should(gomega.HaveLen(1))
			}
		})
	}
}
