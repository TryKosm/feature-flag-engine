package engine

import "testing"

func TestBucketDeterministic(t *testing.T) {
	a := Bucket("u-1", 50)
	b := Bucket("u-1", 50)
	if a != b {
		t.Fatal("bucket assignment must be deterministic")
	}
}
