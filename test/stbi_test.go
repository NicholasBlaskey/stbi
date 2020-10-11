package stbi

import (
	"testing"

	"github.com/nicholasblaskey/stbi"
	"unsafe"
)

func loadFuncTest(
	load func(string, bool, int) (
		unsafe.Pointer, int32, int32, int32, func(), error), t *testing.T) {
	// Try to load an image that doesn't exist
	_, _, _, _, _, err := load("./dne.png", false, 0)
	if err == nil {
		t.Error("Did not get an error upon file not existing")
	}

	// Load one that does exists
	_, width, height, nrChannels,
		cleanup, err := load("./gopher.jpg", false, 0)
	defer cleanup()
	if err != nil {
		t.Errorf("Incorrectly got error with %s", err.Error())
	}
	if width != 960 || height != 720 {
		t.Errorf("Got wrong dimensions on load "+
			"expected (960, 720) got (%d, %d)", width, height)
	}
	if nrChannels != 3 {
		t.Errorf("Got wrong number of channels expected 3 got %d",
			nrChannels)
	}
}

func TestLoad(t *testing.T) {
	loadFuncTest(stbi.Load, t)
}

func TestLoadf(t *testing.T) {
	loadFuncTest(stbi.Loadf, t)
}

func TestLoad16(t *testing.T) {
	loadFuncTest(stbi.Load16, t)
}
