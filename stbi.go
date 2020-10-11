// A go wrapper of stb_image.h for loading files into OpenGL easier.
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
	load int = iota
	load16
	loaff
)

func Load(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(load, path, flipVertical, desiredChannels)
}

func Load16(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(load16, path, flipVertical, desiredChannels)
}

func Loadf(path string, flipVertical bool, desiredChannels int) (
	unsafe.Pointer, int32, int32, int32, func(), error) {

	return loadHelper(loadf, path, flipVertical, desiredChannels)
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
	case load:
		p := C.stbi_load(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(p)
	case load16:
		p := C.stbi_load_16(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(p)
	case loadf:
		p := C.stbi_loadf(cString, &width, &height, &nrChannels, wantedChannels)
		failed = p == nil
		data = unsafe.Pointer(p)
	}
	if failed {
		failure := C.stbi_failure_reason()
		return nil, 0, 0, 0, func() {}, errors.New(C.GoString(failure))
	}

	cleanup := func() {
		C.stbi_image_free(data)
	}
	return data, int32(width),
		int32(height), int32(nrChannels), cleanup, nil

}

// for image formats that explicitly notate that they have premultiplied alpha,
// we just return the colors as stored in the file. set this flag to force
// unpremultiplication. results are undefined if the unpremultiply overflow.
func setUnpremultiplyOnLoad(shouldSet bool) {
	if shouldSet {
		C.stbi_set_unpremultiply_on_load(1)
	} else {
		C.stbi_set_unpremultiply_on_load(0)
	}
}

// indicate whether we should process iphone images back to canonical format,
// or just pass them through "as-is"
func convertIphonePngToRgb(shouldSet bool) {
	if shouldSet {
		C.stbi_convert_iphone_png_to_rgb(1)
	} else {
		C.stbi_convert_iphone_png_to_rgb(0)
	}
}
