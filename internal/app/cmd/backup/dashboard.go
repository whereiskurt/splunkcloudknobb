package backup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/whereiskurt/cloudcrawler/pkg/service"

	"github.com/spf13/cobra"
)

// Dashboard bwill backup all of the dashboard to a local file
func (c *Cmd) Dashboard(cmd *cobra.Command, args []string) {
	config := c.Config
	log := config.Log
	s := service.NewService(log)

	config.CLI.Prompt("\n")
	config.CLI.Prompt("Starting backup for dashboards...\n\n")

	auth, err := s.Login(config.URL, config.Username, config.Password, config.CookiePort, config.Log)
	if err != nil {
		log.Fatalf("fatal: Splunk Cloud Authentication: %+v", err)
	}
	log.Infof("Successful login to instance:\n  %s", auth.URL)
	config.CLI.Prompt(fmt.Sprintf("√ Successful login with '%s' to instance:\n    %s", auth.Username, auth.URL) + "\n")

	chand := make(chan service.Dashboard)
	go func() {
		log.Infof("Retrieving Dashboard Listing for ALL dashboards ... ")
		err := s.ListDashboard(auth, chand)
		if err != nil {
			log.Fatalf("fatal: couldn't retrieve dashbaords lists: %s", err)
		}
	}()

	config.CLI.Prompt("\n√ Fetched dashboard listing for ALL dashboards ... ")

	folder := filepath.Join(c.Config.OutputFolder, "dashboard")
	abs, _ := filepath.Abs(folder)
	log.Infof("Creating backup dashboard folder named '%s%c'", abs, os.PathSeparator)
	os.MkdirAll(abs, 0777)

	rawfolder := filepath.Join(c.Config.OutputFolder, "dashboard", "raw")
	rawabs, _ := filepath.Abs(rawfolder)
	log.Infof("Creating backup dashboard folder named '%s%c'", rawabs, os.PathSeparator)
	os.MkdirAll(rawabs, 0777)

	config.CLI.Prompt(fmt.Sprintf("\n√ Beginning to write files folder '%s%c' .", abs, os.PathSeparator))

	var totalbytes = 0
	var count = 0
	for d := range chand {
		config.CLI.Prompt(".")

		filenamexml := filepath.Join(folder, fmt.Sprintf("%s.xml", d.Name))
		c.writeContent(d.Content, filenamexml)
		c.touchFile(filenamexml, d.Updated)

		filenamejson := filepath.Join(folder, "raw", fmt.Sprintf("%s.json", d.Name))
		c.writeContent(d.RawEntry, filenamejson)
		c.touchFile(filenamejson, d.Updated)

		totalbytes = totalbytes + len(d.RawEntry)/1024
		log.Debugf("%s, %dkb, %s, %s", filenamexml, len(d.RawEntry)/1024, d.Owner, d.Updated)
		count = count + 1
	}

	log.Info(fmt.Sprintf("Wrote '%d' dashboards to folder '%s%c'", count, abs, os.PathSeparator))
	config.CLI.Prompt(fmt.Sprintf("\n√ Wrote '%d' dashboards \n√ Writing files to folder '%s%c'\n", count, abs, os.PathSeparator))

	log.Infof("Done. Wrote '%dkb' total of content. :-)", totalbytes)
	config.CLI.Prompt(fmt.Sprintf(" done! :-) \n\n√ Success! Wrote '%dkb' across '%d' dashboard backup files.\n\n√ Congratulations you now have a local backup! :-)\n\n", totalbytes, count))

	return
}
