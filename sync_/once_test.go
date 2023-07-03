package sync_

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyOnce(t *testing.T) {
	i := 0
	once := sync.Once{}
	twice := once
	once.Do(func() {
		i++
	})
	twice.Do(func() {
		i++
	})
	require.Equal(t, i, 2)
}
