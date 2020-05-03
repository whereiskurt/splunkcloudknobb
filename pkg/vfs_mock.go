// +build !release

// This package is only used during development to load the resources from disk.

package pkg

import (
	"net/http"
	"path"
	"runtime"
)

// PackageEmbed implements the http filesystem, but is overridden when we
// build with tags (go build -tags release) this file won't be built, but
// templates_generate.go will be.
var PackageEmbed http.FileSystem

func init() {
	// This needs to be set to an absolute folder path, so we derive it. :-)
	_, filename, _, _ := runtime.Caller(0)

	PackageEmbed = http.Dir(path.Join(path.Dir(filename)))

	return
}
