// Code generated by command: go run gen.go -out encodeblock_amd64.s -stubs encodeblock_amd64.go. DO NOT EDIT.

// +build !appengine
// +build !noasm
// +build gc

package s2

// encodeBlock encodes a non-empty src to a guaranteed-large-enough dst.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
// 
//go:noescape
func genEncodeBlockAsm(dst []byte, src []byte) int
