package ui

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/whereiskurt/splunkcloudknobb/internal/app/cmd"
	"github.com/whereiskurt/splunkcloudknobb/pkg"
)

// AppTemplate are either Cmd or Package text/tmpls and are then rendered by this
type AppTemplate struct {
	HelpFilename    []string
	PackageFilename []string
	Log             *log.Logger
}

// NewTemplate manages the renderings of all templates
func NewTemplate() *AppTemplate {
	tt := new(AppTemplate)
	return tt
}

// AddHelp registers templates under the internal/app/cmd/* folders
func (t *AppTemplate) AddHelp(filename string) (err error) {
	t.HelpFilename = append(t.HelpFilename, filename)
	return
}

// AddPackage register templates under the pkg/* folders
func (t *AppTemplate) AddPackage(filename string) (err error) {
	t.PackageFilename = append(t.PackageFilename, filename)
	return
}

// RenderHelp will output the UI templates as per the config bind the data.
func (t *AppTemplate) RenderHelp(name string, data interface{}) (usage string) {
	usage, err := render(cmd.CmdHelpEmbed, t.HelpFilename, name, data)
	if err != nil {
		t.Log.Fatalf("fatal: cannot load template: %v", err)
	}

	return usage
}

// RenderPackage will output the UI templates as per the config bind the data.
func (t *AppTemplate) RenderPackage(name string, data interface{}) (usage string) {
	usage, err := render(pkg.PackageEmbed, t.PackageFilename, name, data)
	if err != nil {
		t.Log.Fatalf("fatal: cannot load template: %v", err)
	}

	return usage
}

func render(filesystem http.FileSystem, filenames []string, name string, data interface{}) (usage string, err error) {

	render := template.New("")
	for i := range filenames {

		file, err := filesystem.Open(fmt.Sprintf("%s", filenames[i]))

		if err != nil {
			return "", fmt.Errorf("fatal: err: %v: %s", err, filenames[i])
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return "", fmt.Errorf("fatal: couldn't load template file: %s: %s", fmt.Sprintf("%s", filenames[i]), err)
		}

		render, err = render.Funcs(
			template.FuncMap{
				"StringsSplit": strings.Split,
				"ToUpper":      strings.ToUpper,
				"ToLower":      strings.ToLower,
				"Contains":     strings.Contains,
				"Base64":       Base64,
				"Sumsha1":      Sumsha1,
			},
		).Parse(string(content))
	}

	var raw bytes.Buffer

	if err := render.ExecuteTemplate(&raw, name, data); err != nil {
		return "", fmt.Errorf("fatal: couldn't execute template: %v", err)
	}

	return raw.String(), nil

}

// Base64 takes a raw string and Base64 encodes it
func Base64(raw string) (encoded string) {
	encoded = string(base64.StdEncoding.EncodeToString([]byte(raw)))
	return
}

// Sumsha1 returns sha1 of raw string for templates
func Sumsha1(raw string) string {
	hash := sha1.Sum([]byte(raw))
	return hex.EncodeToString(hash[:])
}
