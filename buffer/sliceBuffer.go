package buffer

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (b *SliceBuffer[T]) Write(data []T) {
	b.buf = append(b.buf, data...)
}

func (b *SliceBuffer[T]) Readable() int {
	return len(b.buf)
}

func (b *SliceBuffer[T]) Read(size int) []T {
	size = min(size, b.Readable())
	data := make([]T, size)
	copy(data, b.buf[:size])
	b.buf = b.buf[size:]
	return data
}
