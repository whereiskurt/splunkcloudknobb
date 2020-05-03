package ui

import (
	"github.com/whereiskurt/splunkcloudknobb/pkg/tmpl"

	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// CLI makes the text output to the terminal.
type CLI struct {
	// Config        *config.Config
	HelpTemplates []string
	Template      *tmpl.UITemplate
	Log           *log.Logger
}

// NewCLI takes a configuration used for describing how to output.
func NewCLI(log *log.Logger) *CLI {
	// cli.Config = c
	var cli = new(CLI)
	cli.Log = log

	cli.Template = tmpl.NewTemplate()

	cli.Template.RegisterHelpFile("scknobb.tmpl")
	cli.Template.RegisterHelpFile("backup/backup.tmpl")

	return cli
}

// StderrHelpTemplate renders at template name to STDERR
func (cli *CLI) StderrHelpTemplate(name string, data interface{}) {
	fmt.Fprintf(os.Stderr, cli.Template.RenderHelp(name, data))
}

// Prompt prints to STDERR
func (cli *CLI) Prompt(line string) {
	fmt.Fprintf(os.Stderr, "%s", line)
	return
}
