# stbi

A go wrapper of stb_image.h for loading images easily into OpenGL. 

### Install

A C compiler is required to be installed as of right now. Then just do

```
go get -u github.com/nicholasblaskey/stbi
```

### Example

```go
package main

import (
	"github.com/nicholasblaskey/stbi"
)

func main() {
	shouldFlipVertical := true
	path := "test/gopher.jpg"
	desiredChannels := 0 // leave 0 for it to decide for you

	data, width, height, nrChannels, cleanup, err := stbi.Load(path, shouldFlipVertical, desiredChannels)
	if err != nil {
		// Handle error
	}
	defer cleanup()
	...
}
```

[Here is a more complete example](https://github.com/NicholasBlaskey/go-learn-opengl/blob/master/src/6.pbr/2.1.1.ibl_irradiance_conversion/ibl_irradiance_conversion.go#L138).

### Why?

As far as I can tell there isn't a great way to load images in the standard library without alpha values easily or choosing the data format. In addition I wasn't having great luck loading .hdr files. 

### Future

Write this in pure go.
