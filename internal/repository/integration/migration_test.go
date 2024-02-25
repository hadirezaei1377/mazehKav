package integration

import (
	"mazehkav/internal/entity"
	"mazehkav/pkg/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrations(t *testing.T) {
	t.Parallel()
	db := testutil.GetConnection(t)
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Restaurant{},
	)
	assert.NoError(t, err)
	// check indexes
	indexInfo, err := testutil.GetIndexInfo(t, db, "restaurants", "idx_location")
	assert.Equal(t, "SPATIAL", indexInfo.IndexType)
}
