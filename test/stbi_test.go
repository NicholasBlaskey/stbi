package stbi

import (
	"testing"

	"github.com/nicholasblaskey/stbi"
)

func TestLoad(t *testing.T) {
	// Try to load an image that doesn't exist
	_, _, _, _, _, err := stbi.Load("./dne.png", false, 0)
	if err == nil {
		t.Error("Did not get an error upon file not existing")
	}

	_, width, height, nrChannels,
		cleanup, err := stbi.Load("./gopher.jpg", false, 0)
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
	cleanup()

}
