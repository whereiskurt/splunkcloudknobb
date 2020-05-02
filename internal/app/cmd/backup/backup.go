package backup

import (
	"os"
	"time"

	"github.com/whereiskurt/splunkcloudknobb/pkg/config"

	"github.com/spf13/cobra"
)

// Cmd holds the configuration
type Cmd struct {
	Config *config.Config
	DTS    string
}

// NewBackup returns a new BackupCmd
func NewBackup(cfg *config.Config) (c *Cmd) {
	c = new(Cmd)
	c.Config = cfg
	c.DTS = time.Now().Format("20060102T150405")
	return c
}

// HelpHandler all of the knowledge objects (ie.  dashboards, reports, lookup files, ..)
func (c *Cmd) HelpHandler(cmd *cobra.Command, args []string) {
	c.Config.CLI.StderrHelpTemplate("UsageBackup", nil)
}

// All all of the knowledge objects (ie.  dashboards, reports, lookup files, ..)
func (c *Cmd) All(cmd *cobra.Command, args []string) {
	c.Dashboard(cmd, args)
	c.Report(cmd, args)
	c.SearchHistory(cmd, args)
}

//touchFile will update the date time stamp of the file passed
func (c *Cmd) touchFile(filename, dts string) {
	t, e := time.Parse("2006-01-02T15:04:05Z07:00", dts)
	if e != nil {
		c.Config.Log.Fatalf(`fatal: cannot parse date time '%s'`, dts)
	}
	e = os.Chtimes(filename, t, t)
	if e != nil {
		c.Config.Log.Fatalf(`fatal: cannot set last modified/accces time of '%s'`, dts)
	}
	return
}

func (c *Cmd) writeContent(content string, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		c.Config.Log.Errorf("error opening file '%s' for writing: %v", filename, err)
		return
	}
	defer f.Close()

	f.Write([]byte(content))

	return
}
