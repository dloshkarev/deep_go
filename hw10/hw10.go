package hw10

import (
	"unsafe"
)

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	var idx, ptr = 0, uintptr(unsafe.Pointer(&memory[0]))

	for i := 0; i < len(pointers); i++ {
		diff := uintptr(pointers[i]) - ptr
		memory[diff], memory[idx] = 0x00, memory[diff]
		pointers[i] = unsafe.Pointer(&memory[idx])
		idx++
	}
}
