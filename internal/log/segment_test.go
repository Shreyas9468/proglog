package log

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/Shreyas9468/proglog/internal/api/v1"
)

// TestSegment tests the Segment struct and its methods.
func TestSegment(t *testing.T) {
	// Create a temporary directory to store the segment files.
	dir, _ := os.MkdirTemp("", "segment-test")
	defer os.RemoveAll(dir)

	// Create a sample record to use in the test.
	want := &api.Record{Value: []byte("hello world")}

	// Set up the configuration for the segment.
	c := Config{
		Segment: struct {
			MaxStoreBytes uint64
			MaxIndexBytes uint64
			InitialOffset uint64
		}{
			MaxStoreBytes: 1024,
			MaxIndexBytes: entWidth * 3,
			InitialOffset: 16,
		},
	}

	// Create a new segment with the specified configuration and initial offset.
	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)

	// Verify that the segment was created with the correct initial offset.
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)

	// Verify that the segment is not yet maxed.
	require.False(t, s.IsMaxed())

	// Append the sample record to the segment multiple times.
	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)

		// Verify that the record was appended with the correct offset.
		require.Equal(t, 16+i, off)

		// Read the record back from the segment.
		got, err := s.Read(off)
		require.NoError(t, err)

		// Verify that the record was read back correctly.
		require.Equal(t, want.Value, got.Value)
	}

	// Attempt to append the sample record again.
	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)

	// Verify that the segment is now maxed.
	require.True(t, s.IsMaxed())

	// Increase the maximum store bytes and maximum index bytes in the configuration.
	c.Segment.MaxStoreBytes = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexBytes = 1024

	// Create a new segment with the updated configuration.
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	// Verify that the segment is now maxed due to the increased store bytes.
	require.True(t, s.IsMaxed())

	// Remove the segment.
	err = s.Remove()
	require.NoError(t, err)

	// Create a new segment with the original configuration.
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	// Verify that the segment is not yet maxed.
	require.False(t, s.IsMaxed())
}
