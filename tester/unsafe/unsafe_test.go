package tester

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_Unsafe(t *testing.T) {
	a := "aaa"
	tmpa := []byte(a)

	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))

	t.Logf("len(b)=%d,cap(tmpa)=%d,cap(b)=%d\n", len(b), cap(tmpa), cap(b))
}
