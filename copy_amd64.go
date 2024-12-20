package ed25519

import "unsafe"

//go:noescape
func memcopy_avx2_32(src unsafe.Pointer, src2 unsafe.Pointer) int

//go:noescape
func memcopy_avx2_64(src unsafe.Pointer, src2 unsafe.Pointer) int
