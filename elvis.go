package main

import (
	"bytes"
	"fmt"
	"io"
)

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// assembleInputStream combines all the input into a
// single stream for processing.
func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

// assembleOutputStream combines all the output into a
// single stream for comparing.
func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func main() {
	var output bytes.Buffer
	in := assembleInputStream()
	out := assembleOutputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	fmt.Println("=======================================\nRunning Algorithm One")
	output.Reset()
	algOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

	fmt.Println("=======================================\nRunning Algorithm Two")
	output.Reset()
	algTwo(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

	fmt.Println("=======================================\nRunning Algorithm Three")
	output.Reset()
	algThree(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())
}

// algOne is one way to solve the problem.
func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {

		// Read in one byte from the input stream.
		if _, err := io.ReadFull(input, buf[end:]); err != nil {

			// Flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// If we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// Read a new initial number of bytes.
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// Write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// Slice that front byte out.
		copy(buf, buf[1:])
	}
}

// algTwo is a second way to solve the problem.
// Provided by Tyler Bunnell https://twitter.com/TylerJBunnell
func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use the bytes Reader to provide a stream to process.
	input := bytes.NewReader(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Create an index variable to match bytes.
	idx := 0

	for {

		// Read a single byte from our input.
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		// Does this byte match the byte at this offset?
		if b == find[idx] {

			// It matches so increment the index position.
			idx++

			// If every byte has been matched, write
			// out the replacement.
			if idx == size {
				output.Write(repl)
				idx = 0
			}

			continue
		}

		// Did we have any sort of match on any given byte?
		if idx != 0 {

			// Write what we've matched up to this point.
			output.Write(find[:idx])

			// Unread the unmatched byte so it can be processed again.
			input.UnreadByte()

			// Reset the offset to start matching from the beginning.
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset.
		output.WriteByte(b)
		idx = 0
	}
}

// algThree is another way to solve the problem, using Knuth
func algThree(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	var sz int = len(data)
	var find_sz = len(find)
	var matrix [256]byte

	// Make sure we find elvis at the end
	data = append(data, find...)

	// setup jump matrix
	for i, _ := range matrix {
		matrix[i] = 255
	}
	for i, x := range find {
		// Set up the table so that all letters in find contain the distance
		// to the first letter of the find.
		// It must contain the smallest jump, not applicable
		// for Elvis but would be for Presley
		if matrix[x] > byte(i) {
			matrix[x] = byte(i)
		}
	}
	fmt.Printf("matrix[e]: %d\n", matrix['e'])

main_loop:
	for ix := 0; true; ix += find_sz {
		p := int(matrix[data[ix]])
		if p != 255 && p <= ix {
			tmp := ix - p
			fmt.Printf("ix: %d, p: %d tmp: %d\n", ix, p, tmp)

			if bytes.Compare(data[tmp:tmp+find_sz], find) == 0 {
				fmt.Printf("match!\n")
				for i, d := range repl {
					data[tmp+i] = d
				}
				if ix >= sz {
					// done, we found the last 'find' we put there, make sure to only
					// copy the bytes excluding that
					output.Write(data[0:sz])
					break main_loop
				}
			}
		}
	}
}
