//+build mage

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/shurcooL/vfsgen"
)

//TODO: Completely redo this Mage file :-) Proper Go constructs

var (
	VERSION  = "v0.0.1-release"
	GOOS     = runtime.GOOS
	GOARCH   = runtime.GOARCH
	GIT_HASH = hash()
	TAGS     = "release"
	SRC      = "cmd/scknobb.go"
	LDFLAGS  = fmt.Sprintf(`-X "github.com/whereiskurt/splunkcloudknobb/internal/app.ReleaseVersion=%s" -X "github.com/whereiskurt/splunkcloudknobb/internal/app.GitHash=%s" -X "github.com/whereiskurt/splunkcloudknobb/internal/app.ReleaseDate=%s"`, VERSION, GIT_HASH, time.Now().Format(time.RFC3339))
)

// Outputs binaries for Windows and Linux into 'release/' folder
func Release() error {
	mg.Deps(Clean)
	mg.Deps(Generate)
	mg.Deps(Test)
	mg.Deps(goModule)
	mg.Deps(Build)
	mg.Deps(Binaries)

	return nil
}

// Run go test for the -tags TAGS
func Test() error {
	err := sh.Run("go", "test", "-tags", TAGS, "./...")
	return err
}

// Build a binary for local GOOS and GOARCH into 'release/' folder
func Build() error {
	log.Println("Making directory 'release' for the binaries ....")
	if err := os.MkdirAll("release", 0777); err != nil {
		return err
	}

	mg.Deps(Generate)

	ext := ""
	if GOOS == "windows" {
		ext = ".exe"
	}

	return build(GOOS, GOARCH, ext)
}

// Remove vfs_generated files and "log" folder
func Clean() error {
	//NOTE: File will grow and grow otherwise, because it includes itself during generation.
	os.Remove("internal/app/vfs.go")
	os.Remove("pkg/vfs.go")
	os.Remove("go.sum")
	os.RemoveAll("log/")
	return nil
}

// Run go generate and output code generated files
func Sanity() error {
	mg.Deps(Release)

	sh.RunV("git", "add", ".")
	sh.RunV("git", "commit", "-m", `"-Mage sanity commit (`+time.Now().Format("20060102T150405")+")")
	sh.RunV("git", "push")
	//sh.RunV("git", "gc")

	return nil
}

// Runs go mod download/vendor
func goModule() error {
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	if err := sh.Run("go", "mod", "vendor"); err != nil {
		return err
	}
	return nil
}

// Outputs various GOOS/GOARCH combos for Windows/Linux
func Binaries() error {
	mg.Deps(binaryLinuxAmd64, binaryWindows386, binaryWindowAmd64, binaryDarwinAmd64)
	return nil
}
func binaryLinuxAmd64() error {
	return build("linux", "amd64", "")
}
func binaryWindows386() error {
	return build("windows", "386", ".exe")
}
func binaryWindowAmd64() error {
	return build("windows", "amd64", ".exe")
}
func binaryDarwinAmd64() error {
	return build("darwin", "amd64", "")
}

func build(goos, goarch, ext string) error {
	output := fmt.Sprintf(`./release/scknobb.%s.%s%s`, goos, goarch, ext)

	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goarch)
	err := sh.Run("go", "build", "-tags", TAGS, "-ldflags="+LDFLAGS, "-o", output, SRC)

	return err
}

func hash() string {
	w := log.Writer()

	// Disable extra log statement from sh.Output
	log.SetOutput(ioutil.Discard)
	hash, _ := sh.Output("git", "rev-list", "-1", "master")
	log.SetOutput(w)

	return hash[:8]
}

func Generate() (err error) {
	mg.Deps(Clean)

	//NOTE: File will grow and grow otherwise, because it includes itself during generation.
	os.Remove("internal/app/cmd/vfs.go")
	os.Remove("pkg/vfs.go")

	err = vfsgen.Generate(http.Dir("internal/app/cmd/"), vfsgen.Options{
		Filename:     "internal/app/cmd/vfs.go",
		PackageName:  "cmd",
		BuildTags:    "release",
		VariableName: "CmdHelpEmbed",
	})
	if err != nil {
		return err
	}

	err = vfsgen.Generate(http.Dir("pkg/"), vfsgen.Options{
		Filename:     "pkg/vfs.go",
		PackageName:  "pkg",
		BuildTags:    "release",
		VariableName: "PackageEmbed",
	})
	if err != nil {
		return err
	}

	return err
}
