package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	home "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/whereiskurt/cloudcrawler/pkg/ui"
)

// These defaults are needed to configure Viper/Cobra
const defaultLogFolder = "log/"
const defaultHomeFilename = ".scknobb.v1"

// defaultConfigType the file extension for the configuration files (using YAML instead of XML)
const defaultConfigType = "yaml"
const defaultConfigFilename = "default.scknobb.v1"

// Config holds common parameters for all invocations
type Config struct {
	Context        context.Context
	CLI            *ui.CLI
	Log            *log.Logger
	Username       string
	Password       string
	CookiePort     string
	CryptoKey      string
	SaveCryptoKey  bool
	URL            string
	DTS            string
	HomeFolder     string
	HomeFilename   string
	OutputFolder   string
	ConfigFilename string
	LogFolder      string
	VerboseLevel   string
	VerboseLevel1  bool
	VerboseLevel2  bool
	VerboseLevel3  bool
	VerboseLevel4  bool
	VerboseLevel5  bool
}

// NewConfig holds common parameters
func NewConfig() (c *Config) {
	c = new(Config)
	c.Log = log.New()
	c.DTS = time.Now().Format("20060102T150405")
	c.Context = context.Background()
	c.CLI = ui.NewCLI(c.Log)

	cobra.OnInitialize(func() {
		// Only read configuration when not invoked with 'help'
		if len(os.Args) == 1 || strings.ToLower(os.Args[1]) == "help" {
			return
		}

		if err := c.readWithViper(); err != nil {
			c.promptAndWriteDefault()
		}
	})

	c.SetupDefaults()

	return
}

// SetupDefaults hard defaults that aren't read from configruation folders
func (c *Config) SetupDefaults() {

	c.LogFolder = defaultLogFolder

	folder, err := home.Dir()
	if err != nil {
		log.Fatalf("fatal: failed to detect home directory: %v", err)
	} else {
		c.HomeFolder = folder
	}
	c.HomeFilename = defaultHomeFilename

	c.OutputFolder = filepath.Join(fmt.Sprintf("%s.scknobb", c.DTS))

	c.SetDefaultLogFilename()
	return
}

// SetDefaultLogFilename will set the ServerConfig log and duplicate to STDOUT
func (c *Config) SetDefaultLogFilename() {
	filename := fmt.Sprintf("knobb.%s.log", c.DTS)
	path := filepath.Join(c.LogFolder, filename)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	mw := io.MultiWriter(f, &fatalStderrSplitter{})
	if c.Log.IsLevelEnabled(log.TraceLevel) {
		mw = io.MultiWriter(f, os.Stdout, &fatalStderrSplitter{})
	}
	c.Log.SetOutput(mw)
}

// Inspect log entries and
type fatalStderrSplitter struct{}

// Write is called when logger is evaluating the fatalStderrSplitter
func (s *fatalStderrSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("fatal")) {
		os.Stderr.WriteString("\n**************\nFatal Error:")
		os.Stderr.Write(p)
		os.Stderr.WriteString("\n")
	}
	return len(p), nil
}

// UnmarshalViper copies all of the cobra/viper config data into our Config struct
// This is the delineation between cobra/viper and using our Config struct.
// Called by app just any cobra command runs (ie. prerun)
func (c *Config) UnmarshalViper() {

	err := viper.MergeInConfig()
	if err != nil {
		c.Log.Warnf("No configuration found: %s", err)
		return
	}

	// Copy everything from the Viper into our Config
	err = viper.Unmarshal(&c)
	if err != nil {
		c.Log.Fatalf("%s", err)
	}

	c.decrypt()

	return
}

func (c *Config) readWithViper() error {
	var err error

	viper.AutomaticEnv()

	curdir, _ := os.Getwd()

	defaultFilename := filepath.Join(curdir, defaultConfigFilename+"."+defaultConfigType)
	f, err := os.OpenFile(defaultFilename, os.O_RDONLY, 0222)

	defer f.Close()

	err = viper.ReadConfig(f)
	if err != nil {
		log.Fatalf("fatal: couldn't read default config: %s", err)
	}

	filename := filepath.Join(c.HomeFolder, c.HomeFilename)
	filename = filename + "." + defaultConfigType

	viper.AddConfigPath(c.HomeFolder)
	viper.SetConfigName(c.HomeFilename)

	err = viper.MergeInConfig()

	c.SetDefaultLogFilename()
	return err
}
