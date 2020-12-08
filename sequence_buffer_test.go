package rmnp

import "testing"

func TestSequenceBuffer(t *testing.T) {
	b := newSequenceBuffer(10)

	if b.get(0) != false || b.get(10) != false {
		t.Error("Expected buffer to be empty")
	}

	b.set(10, true)

	if b.get(0) != false || b.get(10) != true {
		t.Error("Expected buffer to contain value only at index 10")
	}

	b.reset()

	if b.get(0) != false || b.get(10) != false {
		t.Error("Expected buffer to be empty after reset")
	}
}
