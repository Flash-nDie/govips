package vips

// #cgo CFLAGS: -std=c99
// #include "iofunc.h"
import "C"

func vipsImageKill(in *C.VipsImage) error {
	incOpCounter("imageKill")

	C.image_set_kill(in)

	return nil
}

func vipsImageIsKilled(in *C.VipsImage) error {
	incOpCounter("imageIsKilled")

	if err := C.image_iskilled(in); err != 0 {
		return handleImageError(in)
	}

	return nil
}
