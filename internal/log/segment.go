package log

import (
	"fmt"
	"os"
	"path"

	"google.golang.org/protobuf/proto"

	api "github.com/Shreyas9468/proglog/internal/api/v1"
)

// segment represents a segment of log data.
type segment struct {
	store                  *store
	index                  *index
	baseOffset, nextOffset uint64
	config                 Config
}

// newSegment creates a new segment of log data.
// It takes a directory path, base offset, and configuration as parameters.
// It returns a pointer to the segment and an error.
func newSegment(dir string, baseOffset uint64, c Config) (*segment, error) {
	// Create a new segment.
	s := &segment{
		baseOffset: baseOffset,
		config:     c,
	}

	var err error

	// Open the store file for reading and writing.
	storefile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d.store", baseOffset)),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return nil, err
	}

	if s.store, err = newStore(storefile); err != nil {
		return nil, err
	}

	indexfile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d.index", baseOffset)),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return nil, err
	}

	if s.index, err = newIndex(indexfile, c); err != nil {
		return nil, err
	}

	// Read the offset from the index.
	if off, _, err := s.index.Read(-1); err != nil {
		// If there's an error reading the offset, set the next offset to the base offset.
		s.nextOffset = baseOffset
	} else {
		// Otherwise, set the next offset to the base offset plus the offset plus one.
		s.nextOffset = baseOffset + uint64(off) + 1
	}

	// Return the segment and no error.
	return s, nil
}

// Append appends a record to the segment.
// It takes a record pointer as a parameter.
// It returns the offset and an error.
func (s *segment) Append(record *api.Record) (offset uint64, err error) {
	cur := s.nextOffset
	record.Offset = cur
	p, err := proto.Marshal(record)
	if err != nil {
		return 0, err
	}
	_, pos, err := s.store.Append(p)
	if err != nil {
		return 0, err
	}
	if err = s.index.Write(
		// index offsets are relative to base offset
		uint32(s.nextOffset-uint64(s.baseOffset)),
		pos,
	); err != nil {
		return 0, err
	}
	s.nextOffset++
	return cur, nil
}
func (s *segment) Read(offset uint64) (*api.Record, error) {

	_, pos, err := s.index.Read((int64(offset - s.baseOffset)))
	if err != nil {
		return nil, err
	}
	b, err := s.store.Read(pos)
	if err != nil {
		return nil, err
	}
	record := &api.Record{}
	err = proto.Unmarshal(b, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *segment) IsMaxed() bool {
	return s.store.size >= s.config.Segment.MaxStoreBytes ||
		s.index.size >= s.config.Segment.MaxIndexBytes
}

func (s *segment) Remove() error {
	if err := s.Close(); err != nil {
		return err
	}

	if err := os.Remove(s.store.Name()); err != nil {
		return err
	}

	if err := os.Remove(s.index.Name()); err != nil {
		return err
	}
	return nil
}

func (s *segment) Close() error {
	if err := s.index.Close(); err != nil {
		return err
	}
	if err := s.store.Close(); err != nil {
		return err
	}
	return nil
}

func nearestMultiple(j, k uint64) uint64 {
	if j >= 0 {
		return (j / k) * k
	}
	return ((j - k + 1) / k) * k
}
