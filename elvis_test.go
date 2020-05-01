// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmThree(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algThree(in, find, repl, &output)
	}
}

func TestCheckAlgoThreeHappyCase(t *testing.T) {
	var output bytes.Buffer
	in := []byte("elvis")
	find := []byte("elvis")
	repl := []byte("Elvis")
	expected := []byte("Elvis")

	algThree(in, find, repl, &output)
	if bytes.Compare(output.Bytes(), expected) != 0 {
		t.Errorf("got %s, expected %s", output.Bytes(), expected)
	}
}

func TestCheckAlgoThreeHappyCasePreBytes(t *testing.T) {
	var output bytes.Buffer
	in := []byte("aeelvis")
	find := []byte("elvis")
	repl := []byte("Elvis")
	expected := []byte("aeElvis")

	algThree(in, find, repl, &output)
	if bytes.Compare(output.Bytes(), expected) != 0 {
		t.Errorf("got '%s', expected '%s'", output.Bytes(), expected)
	}
}

func TestCheckAlgoThreeHappyCasePostBytes(t *testing.T) {
	var output bytes.Buffer
	in := []byte("elvisae")
	find := []byte("elvis")
	repl := []byte("Elvis")
	expected := []byte("Elvisae")

	algThree(in, find, repl, &output)
	if bytes.Compare(output.Bytes(), expected) != 0 {
		t.Errorf("got '%s', expected '%s'", output.Bytes(), expected)
	}
}

func TestCheckAlgoThreeHappyCaseNotFound(t *testing.T) {
	var output bytes.Buffer
	in := []byte("elxvisae")
	find := []byte("elvis")
	repl := []byte("Elvis")
	expected := []byte("elxvisae")

	algThree(in, find, repl, &output)
	if bytes.Compare(output.Bytes(), expected) != 0 {
		t.Errorf("got '%s', expected '%s'", output.Bytes(), expected)
	}
}

func TestCheckAlgoThreeHappyCaseWholeReplUsed(t *testing.T) {
	var output bytes.Buffer
	in := []byte("elvisae")
	find := []byte("elvis")
	repl := []byte("ElviX")
	expected := []byte("ElviXae")

	algThree(in, find, repl, &output)
	if bytes.Compare(output.Bytes(), expected) != 0 {
		t.Errorf("got '%s', expected '%s'", output.Bytes(), expected)
	}
}
