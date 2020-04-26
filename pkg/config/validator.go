package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ValidateOrFatal will validate the string values inside of the Config after copying from Unmarshal or self-setting.
func (c *Config) ValidateOrFatal() {

	// Only when we're not getting help or running with no arguments, do we valiate the config object
	if len(os.Args) == 1 || contains(os.Args, "help") {
		return
	}

	if len(c.CryptoKey) == 0 {
		c.Log.Fatalf("fatal: crypto key (--cryptokey=ABCD) not set")
	}

	//Shorter than length 16 keys get padded up to that length.
	c.CryptoKey = strings.Repeat(c.CryptoKey, (16/len(c.CryptoKey) + 1))[:16]

	if !(len(c.URL) > 5 && strings.HasPrefix(c.URL, "https")) {
		c.Log.Fatalf("fatal: URL for Splunk cloud must start with https: %s", c.URL)
	}
	if !strings.HasSuffix(c.URL, "/") {
		c.Log.Fatalf("fatal: URL must end with '/' for Splunk Cloud")
	}

	if c.Username == "" {
		c.Log.Fatalf("fatal: username not set (--username=ABCD)")
	} else if c.URL == "" {
		c.Log.Fatalf("fatal: URL not set (--url=ABCD) not set")
	}

	c.validateVerbosity()
	return
}
func contains(a []string, x string) bool {
	for i := range a {
		if x == a[i] {
			return true
		}
	}
	return false
}

func (c *Config) validateVerbosity() {
	if c.hasVerboseLevel() {
		switch {
		case c.VerboseLevel1:
			c.VerboseLevel = "1"
		case c.VerboseLevel2:
			c.VerboseLevel = "2"
		case c.VerboseLevel3:
			c.VerboseLevel = "3"
		case c.VerboseLevel4:
			c.VerboseLevel = "4"
		case c.VerboseLevel5:
			c.VerboseLevel = "5"
		}
	} else {
		c.VerboseLevel = "3"
	}

	switch c.VerboseLevel {
	case "5":
		c.VerboseLevel5 = true
		c.Log.SetLevel(log.TraceLevel)

	case "4":
		c.VerboseLevel4 = true
		c.Log.SetLevel(log.DebugLevel)
	case "3":
		c.VerboseLevel3 = true
		c.Log.SetLevel(log.InfoLevel)
	case "2":
		c.VerboseLevel1 = true
		c.Log.SetLevel(log.WarnLevel)
	case "1":
		c.VerboseLevel1 = true
		c.Log.SetLevel(log.ErrorLevel)
	}

	if !c.hasVerboseLevel() {
		log.Fatalf("invalid VerboseLevel: '%s'", c.VerboseLevel)
	}
	return
}

func (c *Config) hasVerboseLevel() bool {
	return c.VerboseLevel1 || c.VerboseLevel2 || c.VerboseLevel3 || c.VerboseLevel4 || c.VerboseLevel5
}
