package context_

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// context with value
// 子 context 可以获取到 父 context 的数据，反过来不行
func TestContextWithValue(t *testing.T) {
	var (
		parentKey   = "k_parent"
		parentValue = "v_parent"
		childKey    = "k_child"
		childValue  = "v_child"
	)
	parentCtx := context.WithValue(context.Background(), parentKey, parentValue)
	childCtx := context.WithValue(parentCtx, childKey, childValue)
	require.Equal(t, parentCtx.Value(parentKey), parentValue)
	require.Equal(t, childCtx.Value(childKey), childValue)

	require.Equal(t, childCtx.Value(parentKey), parentValue)
	require.Empty(t, parentCtx.Value(childCtx))
}
