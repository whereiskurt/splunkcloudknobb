package restore

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/splunkcloudknobb/pkg/config"
)

//Cmd describes the command
type Cmd struct {
	Config *config.Config
	DTS    string
}

// NewRestore returns a new BackupCmd
func NewRestore(cfg *config.Config) (c *Cmd) {
	c = new(Cmd)
	c.Config = cfg
	c.DTS = time.Now().Format("20060102T150405")
	return c
}

// HelpHandler all of the knowledge objects (ie.  dashboards, reports, lookup files, ..)
func (c *Cmd) HelpHandler(cmd *cobra.Command, args []string) {
	c.Config.CLI.StderrHelpTemplate("UsageRestore", nil)
}

// RestoreListHandler attempts to upload the local CSV remotely
func (c *Cmd) RestoreLookupFileHandler(cmd *cobra.Command, args []string) {

}
