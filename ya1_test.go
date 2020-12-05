package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stroyeka/fuzzysearch"
)

func TestFuzzySearch(t *testing.T) {
	require.Equal(t, 1, 1)
	require.Equal(t, main.Fuzzysearch("car", "cartwheel"), true)
	require.Equal(t, main.Fuzzysearch("cwhl", "cartwheel"), true)
	require.Equal(t, main.Fuzzysearch("cwheel", "cartwheel"), true)
	require.Equal(t, main.Fuzzysearch("cartwheel", "cartwheel"), true)
	require.Equal(t, main.Fuzzysearch("cwheeel", "cartwheel"), false)
	require.Equal(t, main.Fuzzysearch("lw", "cartwheel"), false)
	require.Equal(t, main.Fuzzysearch("wl", "cartwheel"), true)
	require.Equal(t, main.Fuzzysearch("cewe", "caertwheel"), true)
	require.Equal(t, main.Fuzzysearch("ceew", "caertwheel"), false)
}
