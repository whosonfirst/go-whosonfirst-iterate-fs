package fs

import (
	"context"
	"io"
	"sync/atomic"
	"testing"

	"github.com/whosonfirst/go-whosonfirst-iterate-fs/v2/fixtures"
)

func TestFSEmitter(t *testing.T) {

	ctx := context.Background()

	e, err := NewFSEmitter(ctx, "fs://", fixtures.FS)

	if err != nil {
		t.Fatalf("Failed to create FS emitter, %v", err)
	}

	i := int32(0)

	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {
		atomic.AddInt32(&i, 1)
		return nil
	}

	err = e.WalkURI(ctx, iter_cb, ".")

	if err != nil {
		t.Fatalf("Failed to walk FS, %v", err)
	}

	count := atomic.LoadInt32(&i)
	expected := int32(37)

	if count != expected {
		t.Fatalf("Expected %d records, but counted %d", expected, count)
	}

}
