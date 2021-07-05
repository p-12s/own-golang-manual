package internal

import "testing"

func TestNewServer(t *testing.T) {
	s, err := NewServer("8090")

	if err != nil {
		t.Fatalf("server creating err: %v", err)
	}

	if s.Addr != ":8090" {
		t.Fatal("invalid server addr")
	}
}
