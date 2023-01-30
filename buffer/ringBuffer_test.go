package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingBuffer_Write(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedLength int
		expectWritePt  int
	}{
		{"default", []byte{1, 2, 3}, 3, 3},
		{"write full", []byte{1, 2, 3, 4}, 4, 0},
		{"empty", []byte{}, 0, 0},
	}

	for _, tt := range tests {
		b := NewRingBuffer[byte](4)
		b.Write(tt.input)
		assert.Equal(t, tt.expectedLength, b.Readable(), tt.name)
		assert.Equal(t, tt.input, b.buf[:len(tt.input)], tt.name)
		assert.Equal(t, tt.expectWritePt, b.writePt, tt.name)
	}
}

func TestRingBuffer_Read(t *testing.T) {
	tests := []struct {
		name       string
		input      []byte
		readCnt    int
		expectRead []byte
		expectCnt  int
	}{
		{"default", []byte{1, 2, 3, 4}, 2, []byte{1, 2}, 2},
		{"read all", []byte{1, 2, 3, 4}, 4, []byte{1, 2, 3, 4}, 4},
		{"read over", []byte{1, 2, 3, 4}, 5, []byte{1, 2, 3, 4}, 4},
	}

	for _, tt := range tests {
		b := NewRingBuffer[byte](4)
		b.Write(tt.input)
		data := b.Read(tt.readCnt)
		assert.Equal(t, tt.expectCnt, len(data), tt.name)
		assert.Equal(t, tt.expectRead, data, tt.name)
	}
}

func TestRingBuffer_ReadAndWrite(t *testing.T) {
	b := NewRingBuffer[byte](4)
	b.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, b.Readable())

	data := b.Read(4)
	assert.Equal(t, 4, len(data))
	assert.Equal(t, []byte{1, 2, 3, 4}, data)

	b.Write([]byte{5, 6, 7})
	data = b.Read(4)
	assert.Equal(t, 3, len(data))
	assert.Equal(t, []byte{5, 6, 7}, data)
}
