package ui

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// CLI makes the text output to the terminal.
type CLI struct {
	// Config        *config.Config
	HelpTemplates []string
	Template      *AppTemplate
	Log           *log.Logger
}

// NewCLI takes a configuration used for describing how to output.
func NewCLI(log *log.Logger) *CLI {
	// cli.Config = c
	var cli = new(CLI)
	cli.Log = log

	cli.Template = NewTemplate()
	cli.Template.Log = log

	cli.Template.AddHelp("scknobb.tmpl")
	cli.Template.AddHelp("backup/backup.tmpl")
	cli.Template.AddHelp("restore/restore.tmpl")

	return cli
}

// StderrHelpTemplate renders at template name to STDERR
func (cli *CLI) StderrHelpTemplate(name string, data interface{}) {
	fmt.Fprintf(os.Stderr, cli.Template.RenderHelp(name, data))
}

// Prompt prints to STDERR
func (cli *CLI) Prompt(line string) {
	if cli.Log.Level == log.WarnLevel {
		return
	}

	fmt.Fprintf(os.Stderr, "%s", line)
	return
}
