package hw11

import "unsafe"

const EmptyPointer = uintptr(0x00)

func Trace(stacks [][]uintptr) []uintptr {
	visited := make(map[uintptr]struct{})
	result := make([]uintptr, 0)

	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks[i]); j++ {
			ptr := stacks[i][j]
			if ptr != EmptyPointer {
				result = traverse(ptr, visited, result)
			}
		}
	}

	return result
}

func traverse(ptr uintptr, visited map[uintptr]struct{}, results []uintptr) (result []uintptr) {
	if _, exists := visited[ptr]; exists {
		return results
	}

	defer func() {
		if r := recover(); r != nil {
			result = results
		}
	}()

	visited[ptr] = struct{}{}
	next := (*uintptr)(unsafe.Pointer(ptr))

	if next == nil {
		return results
	}

	return traverse(*next, visited, append(results, ptr))
}
