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

const (
	LOAD int = iota
	LOAD16
	LOADF
)

func Load(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(LOAD, path, flipVertical, desiredChannels)
}

func Load16(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(LOAD16, path, flipVertical, desiredChannels)
}

func Loadf(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(LOADF, path, flipVertical, desiredChannels)
}

func loadHelper(loadFunc int, path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	if flipVertical {
		C.stbi_set_flip_vertically_on_load(1)
	} else {
		C.stbi_set_flip_vertically_on_load(0)
	}

	wantedChannels := C.int(desiredChannels)
	var width, height, nrChannels C.int
	cString := C.CString(path)
	defer C.free(unsafe.Pointer(cString))

	var data unsafe.Pointer
	var failed bool
	switch loadFunc {
	case LOAD:
		p := C.stbi_load(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(data)
	case LOAD16:
		p := C.stbi_load_16(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(data)
	case LOADF:
		p := C.stbi_loadf(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(data)
	}
	if failed {
		failure := C.stbi_failure_reason()
		return nil, 0, 0, 0, func() {}, errors.New(C.GoString(failure))
	}

	cleanup := func() {
		C.stbi_image_free(unsafe.Pointer(data))
	}
	return unsafe.Pointer(data), int32(width),
		int32(height), int32(nrChannels), cleanup, nil
}

// setUnpremultiplyOnLoad
// convertIphonePngToRgb
