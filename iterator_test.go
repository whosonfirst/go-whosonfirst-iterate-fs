package fs

import (
	"context"
	"io"
	"sync/atomic"
	"testing"

	"github.com/whosonfirst/go-whosonfirst-iterate-fs/v2/fixtures"
)

func TestFSIterator(t *testing.T) {

	ctx := context.Background()

	i := int32(0)

	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {
		atomic.AddInt32(&i, 1)
		return nil
	}

	iter, err := NewIteratorWithFS(ctx, "fs://", iter_cb, fixtures.FS)

	if err != nil {
		t.Fatalf("Failed to create new iterator, %v", err)
	}

	err = iter.IterateURIs(ctx, ".")

	if err != nil {
		t.Fatalf("Failed to iterate FS, %v", err)
	}

	count := atomic.LoadInt32(&i)
	expected := int32(37)

	if count != expected {
		t.Fatalf("Expected %d records, but counted %d", expected, count)
	}

}
