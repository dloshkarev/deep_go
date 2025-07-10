package hw3

import (
	"unsafe"
)

// go test ./hw3

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	refs := 1
	buf := COWBuffer{
		data: data,
		refs: &refs,
	}

	return buf
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++

	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	*b.refs--
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs > 1 {
		b.Close()
		refs := 1
		b.refs = &refs
		copied := make([]byte, len(b.data))
		copy(copied, b.data)
		b.data = copied
	}

	b.data[index] = value

	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
