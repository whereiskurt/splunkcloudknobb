package remove

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/splunkcloudknobb/pkg/config"
	"github.com/whereiskurt/splunkcloudknobb/pkg/service"
)

//Cmd describes the command
type Cmd struct {
	Config *config.Config
	DTS    string
}

// NewRemove returns a new BackupCmd
func NewRemove(cfg *config.Config) (c *Cmd) {
	c = new(Cmd)
	c.Config = cfg
	c.DTS = time.Now().Format("20060102T150405")
	return c
}

// HelpHandler all of the knowledge objects (ie.  dashboards, reports, lookup files, ..)
func (c *Cmd) HelpHandler(cmd *cobra.Command, args []string) {
	c.Config.CLI.StderrHelpTemplate("UsageRemove", nil)
}

// RemoveLookupFileHandler deletes a remote lookup file
func (c *Cmd) RemoveLookupFileHandler(cmd *cobra.Command, args []string) {

	config := c.Config
	log := config.Log
	s := service.NewService(log)

	config.CLI.Prompt("\n")

	auth, err := s.Login(config.URL, config.Username, config.Password, config.CookiePort, config.Log)
	if err != nil {
		log.Fatalf("fatal: Splunk Cloud Authentication: %+v", err)
	}
	log.Infof("Successful login to instance:\n  %s", auth.URL)
	config.CLI.Prompt(fmt.Sprintf("√ Successful login with '%s' to instance:\n    %s", auth.Username, auth.URL) + "\n")

	if c.Config.Filename == "" && len(args) > 0 {
		c.Config.Filename = args[0]
	}

	s.SessionMap["Filename"] = c.Config.Filename

	config.CLI.Prompt(fmt.Sprintf("\n• Will attempt to remove remote lookup file '%s' from Splunk Cloud.\n", c.Config.Filename))

	err = s.RemoveLookupFile(auth)
	if err != nil {
		config.CLI.Prompt(fmt.Sprintf("\nX Failed to remove file '%s' from Splunk Cloud.\n", c.Config.Filename))
		return
	}

	config.CLI.Prompt(fmt.Sprintf("\n√ Successfully removed file '%s' from Splunk Cloud.\n\n", c.Config.Filename))
}
