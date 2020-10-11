package stbi

/*
#cgo LDFLAGS: -lm

#define STB_IMAGE_IMPLEMENTATION
#include <stb_image.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

func Load(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	if flipVertical {
		C.stbi_set_flip_vertically_on_load(1)
	} else {
		C.stbi_set_flip_vertically_on_load(0)
	}

	var width, height, nrChannels C.int
	cString := C.CString(path)
	defer C.free(unsafe.Pointer(cString))

	data := C.stbi_load(cString, &width, &height, &nrChannels,
		C.int(desiredChannels))
	if data == nil {
		failure := C.stbi_failure_reason()
		return nil, 0, 0, 0, func() {}, errors.New(C.GoString(failure))
	}

	cleanup := func() {
		C.stbi_image_free(unsafe.Pointer(data))
	}
	return unsafe.Pointer(data), int32(width),
		int32(height), int32(nrChannels), cleanup, nil
}

// load
// load16
// loadf

// imageFree

// failureReason
// info

// setUnpremultiplyOnLoad
// convertIphonePngToRgb
// setFlipVerticallyOnLoad
