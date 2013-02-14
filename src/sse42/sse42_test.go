package sse42

import (
	"bytes"
	"testing"
)

const size1K = 1 << 10
const size32K = 1 << 15
const size1M = 1 << 20
const size1G = 1 << 30

type equalFunc func([]byte, []byte) bool

func benchmarkEqual(b *testing.B, equal equalFunc, size int) {
	b1 := make([]byte, size)
	b2 := make([]byte, size)
	for i := 0; i < len(b1); i++ {
		b1[i] = 'a'
		b2[i] = 'a'
	}
	b.SetBytes(int64(len(b1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !equal(b1, b2) {
			panic("failed")
		}
	}
}

func BenchmarkMemcmp1(b *testing.B)   { benchmarkEqual(b, Memcmp, 1) }
func BenchmarkMemcmp256(b *testing.B) { benchmarkEqual(b, Memcmp, 256) }
func BenchmarkMemcmp512(b *testing.B) { benchmarkEqual(b, Memcmp, 512) }
func BenchmarkMemcmp1K(b *testing.B)  { benchmarkEqual(b, Memcmp, size1K) }
func BenchmarkMemcmp32K(b *testing.B) { benchmarkEqual(b, Memcmp, size32K) }
func BenchmarkMemcmp1M(b *testing.B)  { benchmarkEqual(b, Memcmp, size1M) }
func BenchmarkMemcmp1G(b *testing.B)  { benchmarkEqual(b, Memcmp, size1G) }

func BenchmarkStrncmp1(b *testing.B)   { benchmarkEqual(b, Strncmp, 1) }
func BenchmarkStrncmp1K(b *testing.B)  { benchmarkEqual(b, Strncmp, size1K) }
func BenchmarkStrncmp32K(b *testing.B) { benchmarkEqual(b, Strncmp, size32K) }
func BenchmarkStrncmp1M(b *testing.B)  { benchmarkEqual(b, Strncmp, size1M) }
func BenchmarkStrncmp1G(b *testing.B)  { benchmarkEqual(b, Strncmp, size1G) }

func BenchmarkBytesEqual1(b *testing.B)   { benchmarkEqual(b, bytes.Equal, 1) }
func BenchmarkBytesEqual1K(b *testing.B)  { benchmarkEqual(b, bytes.Equal, size1K) }
func BenchmarkBytesEqual32K(b *testing.B) { benchmarkEqual(b, bytes.Equal, size32K) }
func BenchmarkBytesEqual1M(b *testing.B)  { benchmarkEqual(b, bytes.Equal, size1M) }
func BenchmarkBytesEqual1G(b *testing.B)  { benchmarkEqual(b, bytes.Equal, size1G) }

func BenchmarkRuntimeMemequal1K(b *testing.B) {
	var b1 [size1K]byte
	var b2 [size1K]byte
	for i := 0; i < len(b1); i++ {
		b1[i] = 'a'
		b2[i] = 'a'
	}
	b.SetBytes(int64(len(b1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if b1 != b2 {
			panic("failed")
		}
	}
}

func BenchmarkStringEqual1M(b *testing.B) {
	s1 := string(make([]byte, size1M))
	s2 := string(make([]byte, size1M))
	b.SetBytes(int64(len(s1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if s1 != s2 {
			panic("failed")
		}
	}
}

type indexByteFunc func([]byte, byte) int

func benchmarkIndexByte(b *testing.B, indexByte indexByteFunc, size int) {
	b1 := make([]byte, size)
	for i := 0; i < len(b1); i++ {
		b1[i] = 'a'
	}
	b.SetBytes(int64(len(b1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if indexByte(b1, 'b') != -1 {
			panic("failed")
		}
	}
}

func BenchmarkBytesIndexByte1K(b *testing.B)  { benchmarkIndexByte(b, bytes.IndexByte, size1K) }
func BenchmarkBytesIndexByte32K(b *testing.B) { benchmarkIndexByte(b, bytes.IndexByte, size32K) }
func BenchmarkBytesIndexByte1M(b *testing.B)  { benchmarkIndexByte(b, bytes.IndexByte, size1M) }
func BenchmarkBytesIndexByte1G(b *testing.B)  { benchmarkIndexByte(b, bytes.IndexByte, size1G) }

func BenchmarkMemchr1K(b *testing.B)  { benchmarkIndexByte(b, Memchr, size1K) }
func BenchmarkMemchr32K(b *testing.B) { benchmarkIndexByte(b, Memchr, size32K) }
func BenchmarkMemchr1M(b *testing.B)  { benchmarkIndexByte(b, Memchr, size1M) }
func BenchmarkMemchr1G(b *testing.B)  { benchmarkIndexByte(b, Memchr, size1G) }
