package restore

import "github.com/whereiskurt/splunkcloudknobb/pkg/config"

//Cmd describes the command
type Cmd struct {
	Config *config.Config
	DTS    string
}
