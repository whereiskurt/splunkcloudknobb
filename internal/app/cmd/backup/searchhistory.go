package backup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/cloudcrawler/pkg/service"
)

// SearchHistory will pull back
func (c *Cmd) SearchHistory(cmd *cobra.Command, args []string) {

	config := c.Config
	log := config.Log
	s := service.NewService(log)

	config.CLI.Prompt("\n")
	config.CLI.Prompt("Starting backup for Search History...\n\n")

	auth, err := s.Login(config.URL, config.Username, config.Password, config.CookiePort, config.Log)
	if err != nil {
		log.Fatalf("fatal: Splunk Cloud Authentication: %+v", err)
	}
	log.Infof("Successful login to instance:\n  %s", auth.URL)
	config.CLI.Prompt(fmt.Sprintf("√ Successful login with '%s' to instance:\n    %s", auth.Username, auth.URL) + "\n")

	chanr := make(chan service.SearchHistoryResult)
	go func() {
		log.Infof("Retrieving Search History for ALL searches ... ")
		err := s.ListSearchHistory(auth, chanr)
		if err != nil {
			log.Fatalf("fatal: couldn't retrieve search history lists: %s", err)
		}
		config.CLI.Prompt("\n√ Fetched search history listing for ALL search history ... ")
	}()

	folder := filepath.Join(c.Config.OutputFolder, "search")
	abs, _ := filepath.Abs(folder)
	log.Infof("Creating backup search history folder named '%s%c'", abs, os.PathSeparator)
	os.MkdirAll(abs, 0777)

	rawfolder := filepath.Join(c.Config.OutputFolder, "search", "raw")
	rawabs, _ := filepath.Abs(rawfolder)
	log.Infof("Creating backup search history folder named '%s%c'", rawabs, os.PathSeparator)
	os.MkdirAll(rawabs, 0777)

	config.CLI.Prompt(fmt.Sprintf("\n√ Beginning to write files folder '%s%c' .", abs, os.PathSeparator))

	var totalbytes = 0
	var count = 0
	for r := range chanr {
		config.CLI.Prompt(".")

		filenamexml := filepath.Join(folder, fmt.Sprintf("%s.splunk.txt", r.SearchID))
		c.writeContent(r.Search, filenamexml)
		c.touchFile(filenamexml, r.Time)

		filenamejson := filepath.Join(folder, "raw", fmt.Sprintf("%s.json", r.SearchID))
		c.writeContent(r.RawEntry, filenamejson)
		c.touchFile(filenamejson, r.Time)

		totalbytes = totalbytes + len(r.RawEntry)/1024
		log.Debugf("%s, %dkb, %s, %s", filenamexml, len(r.RawEntry)/1024, auth.Username, r.SearchID)
		count = count + 1
	}

	log.Info(fmt.Sprintf("Wrote '%d' search histories folder '%s%c'", count, abs, os.PathSeparator))
	config.CLI.Prompt(fmt.Sprintf("\n√ Wrote '%d' search histories \n√ Writing files to folder '%s%c'\n", count, abs, os.PathSeparator))

	log.Infof("Done. Wrote '%dkb' total of content. :-)", totalbytes)
	config.CLI.Prompt(fmt.Sprintf("  done! :-) \n\n√ Success! Wrote '%dkb' across '%d' search histories backup files.\n\n√ Congratulations you now have a local backup! :-)\n\n", totalbytes, count))

	return
}
