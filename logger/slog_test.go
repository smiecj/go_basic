package logger

import (
	"fmt"
	"log/slog"
	"maps"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlog(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("[test] hello slog text handler")

	slog.Info("[test] hello slog default")

	x := 1
	y := 2
	slog.Info(fmt.Sprintf("max of int: %d", max(x, y)))
	slog.Info(fmt.Sprintf("max of float: %.2f", max(1.2, 2.1)))

	testSlice := []int{3, 2, 1}
	slices.Sort(testSlice)
	logger.Info(fmt.Sprintf("after order slice: %v", testSlice))

	map1 := map[int]string{1: "1", 2: "2", 3: "3"}
	map2 := map[int]string{1: "1", 3: "3", 2: "2"}
	require.True(t, maps.Equal(map1, map2))

	// maps.Clear(map1)
}
