package main

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestReadTimeout(t *testing.T) {
	ch := make(chan int)
	start := time.Now()
	_, err := Read(ch)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	t.Logf("got: %v, took: %s", err, time.Since(start))
}

func TestReadTimeoutWithSynctest(t *testing.T) {
	synctest.Run(func() {
		ch := make(chan int)
		start := time.Now()
		_, err := Read(ch)
		if err == nil {
			t.Fatal("expected timeout error, got nil")
		}
		t.Logf("got: %v, took: %s", err, time.Since(start))
	})
}
