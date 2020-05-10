package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/splunkcloudknobb/pkg/service"
)

// ListLookupFileHandler outputs a CSV list of all lookup files
func (c *Cmd) ListLookupFileHandler(cmd *cobra.Command, args []string) {
	config := c.Config
	log := config.Log
	s := service.NewService(log)

	config.CLI.Prompt("\n")
	config.CLI.Prompt("Starting Lookup File listing for Reports...\n\n")

	auth, err := s.Login(config.URL, config.Username, config.Password, config.CookiePort, config.Log)
	if err != nil {
		log.Fatalf("fatal: Splunk Cloud Authentication: %+v", err)
	}
	log.Infof("Successful login to instance:\n  %s", auth.URL)
	config.CLI.Prompt(fmt.Sprintf("√ Successful login with '%s' to instance:\n    %s", auth.Username, auth.URL) + "\n")

	chanr := make(chan interface{})
	go func() {
		log.Infof("Fetching Lookup File Listing ... ")
		s.ListLookupFiles(auth, chanr)
		config.CLI.Prompt("\n√ Fetched Lookup File listing.\n")
	}()

	//Output a CSV of all of the lookup files
	var i = 0
	fmt.Printf("%s\n", s.AppTemplate.RenderPackage("lookupFileCSVHeader", nil))
	for rec := range chanr {
		i = i + 1
		lkp := rec.(service.LookupFile)
		rec := make(map[string]string)
		rec["Filename"] = lkp.Filename
		rec["Path"] = lkp.Path
		rec["Owner"] = lkp.Owner
		rec["Sharing"] = lkp.Sharing
		rec["App"] = lkp.App
		rec["Status"] = lkp.Status

		fmt.Printf("%s\n", s.AppTemplate.RenderPackage("lookupFileCSVRow", rec))
	}

	config.CLI.Prompt(fmt.Sprintf("√ Done. %d lookup files list. :-)\n", i))
}
