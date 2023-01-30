package buffer

type RingBuffer[T any] struct {
	buf             []T
	isFull          bool
	size            int
	readPt, writePt int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf:  make([]T, size),
		size: size,
	}
}

func (b *RingBuffer[T]) Write(data []T) {
	b.write(data)
}

func (b *RingBuffer[T]) write(data []T) int {
	if len(data) == 0 || b.writeable() == 0 {
		return 0
	}

	writeSize := 0
	if b.writePt >= b.readPt {
		writeSize = min(b.size-b.writePt, len(data))
	} else {
		writeSize = min(b.writeable(), len(data))
	}

	copy(b.buf[b.writePt:], data[:writeSize])
	b.writePt += writeSize
	b.writePt %= b.size

	if writeSize > 0 && b.writePt == b.readPt {
		b.isFull = true
	}

	remain := len(data) - writeSize
	if remain > 0 && b.writeable() > 0 {
		writeSize += b.write(data[writeSize:])
	}

	return writeSize
}

func (b *RingBuffer[T]) Read(size int) []T {
	if b.Readable() == 0 || size <= 0 {
		return []T{}
	}

	readCnt := min(b.Readable(), size)
	data := make([]T, readCnt)

	b.read(data, readCnt, 0)
	b.isFull = false
	return data
}

func (b *RingBuffer[T]) read(data []T, size int, start int) {
	if size == 0 {
		return
	}

	toEnd := b.size - b.readPt
	readCnt := min(size, toEnd)
	copy(data[start:], b.buf[b.readPt:b.readPt+readCnt])
	b.readPt += readCnt
	b.readPt %= b.size

	remain := size - readCnt
	b.read(data, remain, start+readCnt)
	return
}

func (b *RingBuffer[T]) Readable() int {
	if b.isFull {
		return b.size
	}
	if b.readPt > b.writePt {
		return b.size - b.readPt + b.writePt
	}
	return b.writePt - b.readPt
}

func (b *RingBuffer[T]) writeable() int {
	return b.size - b.Readable()
}
