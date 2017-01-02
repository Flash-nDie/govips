package govips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"

import (
	go_debug "github.com/tj/go-debug"
	"unsafe"
)

var STRING_BUFFER = fixedString(4096)

func byteArrayPointer(b []byte) unsafe.Pointer {
	return unsafe.Pointer(&b[0])
}

func freeCString(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func toGboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func fromGboolean(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func fixedString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = '0'
	}
	return string(b)
}

var debugFn = go_debug.Debug("govips")

func debug(fmt string, values ...interface{}) {
	debugFn(fmt, values)
}