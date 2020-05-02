package backup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/splunkcloudknobb/pkg/service"
)

// Report will backup all of the reports to a local file
func (c *Cmd) Report(cmd *cobra.Command, args []string) {
	config := c.Config
	log := config.Log
	s := service.NewService(log)

	config.CLI.Prompt("\n")
	config.CLI.Prompt("Starting backup for Reports...\n\n")

	auth, err := s.Login(config.URL, config.Username, config.Password, config.CookiePort, config.Log)
	if err != nil {
		log.Fatalf("fatal: Splunk Cloud Authentication: %+v", err)
	}
	log.Infof("Successful login to instance:\n  %s", auth.URL)
	config.CLI.Prompt(fmt.Sprintf("√ Successful login with '%s' to instance:\n    %s", auth.Username, auth.URL) + "\n")

	chanr := make(chan service.Report)
	go func() {
		log.Infof("Retrieving Reports Listing for ALL reports ... ")
		err := s.ListReport(auth, chanr)
		if err != nil {
			log.Fatalf("fatal: couldn't retrieve report lists: %s", err)
		}
		config.CLI.Prompt("\n√ Fetched report listing for ALL reports ... ")
	}()

	folder := filepath.Join(c.Config.OutputFolder, "report")
	abs, _ := filepath.Abs(folder)
	log.Infof("Creating backup report folder named '%s%c'", abs, os.PathSeparator)
	os.MkdirAll(abs, 0777)

	rawfolder := filepath.Join(c.Config.OutputFolder, "report", "raw")
	rawabs, _ := filepath.Abs(rawfolder)
	log.Infof("Creating backup report folder named '%s%c'", rawabs, os.PathSeparator)
	os.MkdirAll(rawabs, 0777)

	config.CLI.Prompt(fmt.Sprintf("\n√ Beginning to write files folder '%s%c' .", abs, os.PathSeparator))

	var totalbytes = 0
	var count = 0
	for r := range chanr {
		config.CLI.Prompt(".")

		filenamexml := filepath.Join(folder, fmt.Sprintf("%s.splunk.txt", r.Name))
		c.writeContent(r.Search, filenamexml)
		c.touchFile(filenamexml, r.Updated)

		filenamejson := filepath.Join(folder, "raw", fmt.Sprintf("%s.json", r.Name))
		c.writeContent(r.RawEntry, filenamejson)
		c.touchFile(filenamejson, r.Updated)

		totalbytes = totalbytes + len(r.RawEntry)/1024
		log.Debugf("%s, %dkb, %s, %s", filenamexml, len(r.RawEntry)/1024, r.Owner, r.Updated)
		count = count + 1
	}

	log.Info(fmt.Sprintf("Wrote '%d' reports to folder '%s%c'", count, abs, os.PathSeparator))
	config.CLI.Prompt(fmt.Sprintf("\n√ Wrote '%d' reports \n√ Writing files to folder '%s%c'\n", count, abs, os.PathSeparator))

	log.Infof("Done. Wrote '%dkb' total of content. :-)", totalbytes)
	config.CLI.Prompt(fmt.Sprintf("  done! :-) \n\n√ Success! Wrote '%dkb' across '%d' report backup files.\n\n√ Congratulations you now have a local backup! :-)\n\n", totalbytes, count))

	return
}
