package app

import (
	"strings"

	"github.com/whereiskurt/splunkcloudknobb/internal/app/cmd/backup"
	"github.com/whereiskurt/splunkcloudknobb/pkg/config"

	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// CommandList entry[0] becomes default when a 'command' is omitted
	CommandList = []string{"backup", "restore", "list"}
)

//App is the top level application code
type App struct {
	Config       *config.Config
	RootCmd      *cobra.Command
	DefaultUsage string
}

var (
	// ReleaseVersion is set by a --ldflags during a build/release
	ReleaseVersion = "v0.0.1-development"
	// GitHash is set by a --ldflags during a build/release
	GitHash = "0x0123abcd"
	// ReleaseDate is set by a --ldflags during a build/release
	ReleaseDate = "2020-04-18"
)

// NewApp is handles the Cobra/Viper mapping from commandline
func NewApp(config *config.Config) (a *App) {
	a = new(App)

	a.Config = config

	a.RootCmd = new(cobra.Command)
	a.RootCmd.Run = a.AppHelp

	a.RootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("SCKNOBB")    //This is important to prevent collisons like "%USERNAME%" on Window.s :)
		a.Config.UnmarshalViper()        // copy values from cobra
		cmd.ParseFlags(os.Args)          // parse commandline for parameters
		a.Config.ValidateOrFatal()       // and validate.
		a.Config.SetDefaultLogFilename() //enablds --trace to STDOUT and logfiles
	}

	a.RootCmd.SetUsageTemplate(a.AppUsageHeader())
	flagS("Username", &a.Config.Username, []string{"user"}, a.RootCmd)
	flagS("CookiePort", &a.Config.CookiePort, []string{"port", "cookiePort"}, a.RootCmd)
	flagS("URL", &a.Config.URL, []string{"url"}, a.RootCmd)
	flagS("OutputFolder", &a.Config.OutputFolder, []string{"out", "target"}, a.RootCmd)
	flagS("CryptoKey", &a.Config.CryptoKey, []string{"key", "k"}, a.RootCmd)
	flagS("VerboseLevel", &a.Config.VerboseLevel, []string{"level"}, a.RootCmd)
	flagB("VerboseLevel1", &a.Config.VerboseLevel1, []string{"Q", "silent"}, a.RootCmd)
	flagB("VerboseLevel2", &a.Config.VerboseLevel2, []string{"q", "quiet"}, a.RootCmd)
	flagB("VerboseLevel3", &a.Config.VerboseLevel3, []string{"info"}, a.RootCmd)
	flagB("VerboseLevel4", &a.Config.VerboseLevel4, []string{"v", "debug"}, a.RootCmd)
	flagB("VerboseLevel5", &a.Config.VerboseLevel5, []string{"V", "trace"}, a.RootCmd)

	helpCmd := command("help", a.AppHelp, a.RootCmd)

	b := backup.NewBackup(a.Config)
	bCmd := command("backup", b.HelpHandler, a.RootCmd)
	_ = command("backup", b.HelpHandler, helpCmd)

	_ = command("all", b.All, bCmd)

	_ = command("dashboard", b.Dashboard, bCmd)
	_ = command("report", b.Report, bCmd)
	_ = command("history", b.SearchHistory, bCmd)

	return a
}

// InvokeCLI passes control over to the root cobra command.
func (a *App) InvokeCLI() {

	if len(os.Args) == 1 || strings.ToLower(os.Args[1]) == "--help" || strings.ToLower(os.Args[1]) == "-h" {
		a.Config.CLI.StderrHelpTemplate("AppUsage", nil)
		a.Config.CLI.StderrHelpTemplate("AppUsageExample", nil)
	} else {
		os.MkdirAll(a.Config.LogFolder, 0777)
		// os.MkdirAll(a.Config.OutputFolder, 0777)

		// Call Cobra Execute which will PreRun and select the Command to execute.
		_ = a.RootCmd.Execute()
	}

	return
}

// AppUsageHeader asdf
func (a *App) AppUsageHeader() string {
	versionMap := map[string]string{"ReleaseDate": ReleaseDate[:10], "ReleaseVersion": ReleaseVersion, "GitHash": GitHash}
	a.Config.CLI.StderrHelpTemplate("AppHeader", versionMap)
	return "\x00"
}

// AppHelp renders help for the Help command
func (a *App) AppHelp(cmd *cobra.Command, args []string) {
	a.Config.CLI.StderrHelpTemplate("AppUsage", nil)
	a.Config.CLI.StderrHelpTemplate("AppUsageExample", nil)
}

func contains(a []string, x string) bool {
	for i := range a {
		if x == a[i] {
			return true
		}
	}
	return false
}
func command(s string, run func(*cobra.Command, []string), parent *cobra.Command) *cobra.Command {
	alias := []string{fmt.Sprintf("%ss", s)} // Add a pluralized alias
	child := &cobra.Command{Use: s, Run: run, PreRun: parent.PreRun, Aliases: alias}
	parent.AddCommand(child)
	return child
}
func subcommand(s string, run func(*cobra.Command, []string), parent *cobra.Command) {
	command(s, run, parent)
	return
}
func flagB(name string, ref *bool, aliases []string, cob *cobra.Command) {
	cob.PersistentFlags().BoolVar(ref, name, *ref, "")
	_ = viper.BindPFlag(name, cob.PersistentFlags().Lookup(name))
	if len(aliases) > 0 && len(aliases[0]) == 1 {
		cob.PersistentFlags().Lookup(name).Shorthand = aliases[0]
	}
	for _, alias := range aliases {
		cob.PersistentFlags().BoolVar(ref, alias, *ref, "")
		cob.PersistentFlags().Lookup(alias).Hidden = true
		viper.RegisterAlias(alias, name)
	}

	return
}
func flagS(name string, ref *string, aliases []string, cob *cobra.Command) {
	cob.PersistentFlags().StringVar(ref, name, *ref, "")
	_ = viper.BindPFlag(name, cob.PersistentFlags().Lookup(name))
	if len(aliases) > 0 && len(aliases[0]) == 1 {
		cob.PersistentFlags().Lookup(name).Shorthand = aliases[0]
	}
	for _, alias := range aliases {
		cob.PersistentFlags().StringVar(ref, alias, *ref, "")
		cob.PersistentFlags().Lookup(alias).Hidden = true
		viper.RegisterAlias(alias, name)
	}

	return
}
