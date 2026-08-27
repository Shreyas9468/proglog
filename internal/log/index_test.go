package log

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	// 1. Create a temporary file for testing
	f, err := os.CreateTemp("", "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	c := Config{}
	c.Segment.MaxIndexBytes = 1024

	idx, err := newIndex(f, c)
	require.NoError(t, err)

	// 2. Reading an empty index should error
	_, _, err = idx.Read(-1)
	require.Error(t, err)
	require.Equal(t, f.Name(), idx.Name())

	// 3. Define sample entries to write
	entries := []struct {
		Off uint32
		Pos uint64
	}{
		{Off: 0, Pos: 0},
		{Off: 1, Pos: 10},
	}

	// 4. Write entries and verify immediate read-back
	for _, want := range entries {
		err = idx.Write(want.Off, want.Pos)
		require.NoError(t, err)

		_, pos, err := idx.Read(int64(want.Off))
		require.NoError(t, err)
		require.Equal(t, want.Pos, pos)
	}

	// 5. Reading past existing entries should return io.EOF
	_, _, err = idx.Read(int64(len(entries)))
	require.Equal(t, io.EOF, err)

	// 6. Close the index (Syncs memory map to disk & truncates empty space)
	_ = idx.Close()

	// 7. Test state recovery after simulated restart
	f, _ = os.OpenFile(f.Name(), os.O_RDWR, 0600)
	idx, err = newIndex(f, c)
	require.NoError(t, err)

	off, pos, err := idx.Read(-1)
	require.NoError(t, err)
	require.Equal(t, uint32(1), off)
	require.Equal(t, entries[1].Pos, pos)
}
