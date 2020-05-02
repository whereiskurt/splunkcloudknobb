package ui

import (
	"bytes"

	"github.com/whereiskurt/splunkcloudknobb/internal/app/cmd"

	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// CLI makes the text output to the terminal.
type CLI struct {
	// Config        *config.Config
	HelpTemplates []string
	Log           *log.Logger
}

// NewCLI takes a configuration used for describing how to output.
func NewCLI(log *log.Logger) *CLI {
	// cli.Config = c
	var cli = new(CLI)
	cli.Log = log
	cli.HelpTemplates = append(cli.HelpTemplates, "scknobb.tmpl")
	cli.HelpTemplates = append(cli.HelpTemplates, "backup/backup.tmpl")
	return cli
}

// StderrHelpTemplate renders at template name to STDERR
func (cli *CLI) StderrHelpTemplate(name string, data interface{}) {
	fmt.Fprintf(os.Stderr, cli.renderHelp(name, data))
}

// Prompt prints to STDERR
func (cli *CLI) Prompt(line string) {
	fmt.Fprintf(os.Stderr, "%s", line)
	return
}

// Render will output the UI templates as per the config bind the data.
func (cli *CLI) renderHelp(name string, data interface{}) (usage string) {

	t := template.New("")
	for i := range cli.HelpTemplates {

		file, err := cmd.CmdHelpEmbed.Open(fmt.Sprintf("%s", cli.HelpTemplates[i]))
		if err != nil {
			cli.Log.Fatalf("fatal: err: %v: %s", err, cli.HelpTemplates[i])
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			cli.Log.Fatalf("fatal: couldn't load template file: %s: %s", fmt.Sprintf("%s", cli.HelpTemplates[i]), err)
		}

		t, err = t.Funcs(
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

	if err := t.ExecuteTemplate(&raw, name, data); err != nil {
		cli.Log.Fatalf("fatal: couldn't execute template: %v", err)
	}

	usage = raw.String()
	return
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
