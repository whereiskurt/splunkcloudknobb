package restore

import "github.com/whereiskurt/cloudcrawler/pkg/config"

//Cmd describes the command
type Cmd struct {
	Config *config.Config
	DTS    string
}
