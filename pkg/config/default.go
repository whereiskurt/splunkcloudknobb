package config

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func (c *Config) promptAndWriteDefault() {
	conffile := filepath.Join(c.HomeFolder, c.HomeFilename)
	conffile = conffile + "." + defaultConfigType

	c.CLI.Stderr("PromptNewConnfiguration", map[string]string{"Filename": conffile})

	// First run, try and get user inputted configuration
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		c.InputConfigFirstRun()
	}

	c.ValidateOrFatal()

	c.encrypt()
	c.writeFile(conffile)

	err := viper.MergeInConfig()
	if err != nil {
		log.Warnf("warning: couldn't viper MergeInConfig after default config: %s", err)
		return
	}
}

func (c *Config) shouldSaveKey(savekey string) bool {
	return strings.Contains(strings.ToLower(savekey), "y")
}

// InputConfigFirstRun reads STDIN for the default values, and stores in Config c.
func (c *Config) InputConfigFirstRun() {
	reader := bufio.NewReader(os.Stdin)
	prompt := c.CLI.Prompt
	nl := "\n"

	prompt(`1) Enter you Splunk Cloud instance URL. Must start with 'https://'` + nl)
	prompt(`   and end with a '/'.` + nl)
	prompt(`   [eg: https://instance2.splunkcloud.com/en-US/)]` + nl)
	prompt(nl)
	prompt(`> `)
	fqdn, _ := reader.ReadString('\n')

	defaultCookiePort := "8443"
	prompt("2) Enter the cookie port for your configuration [default:" + defaultCookiePort + "]: " + nl)
	prompt(`> `)
	cookiePort, _ := reader.ReadString('\n')
	cookiePort = strings.TrimSpace(cookiePort)
	if cookiePort == "" {
		cookiePort = defaultCookiePort
	}

	prompt(nl)
	prompt("3) Enter your Splunk Cloud username (someone@somewhere.com)" + nl)
	prompt(`> `)
	username, _ := reader.ReadString('\n')

	prompt(nl)
	prompt("4) Enter your Splunk Cloud password (invisible, won't show)" + nl)
	prompt(`> `)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("failed to read password from STDIN: %+v", err)
	}
	password := string(bytePassword)
	prompt(nl + nl)

	defaultkey := randomHex(16)[:16]
	prompt("5) Enter a cryptokey for configuration [default:" + defaultkey + "]: " + nl)
	prompt(`> `)
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)
	if key == "" {
		key = defaultkey
	}
	prompt(nl)

	prompt(nl)
	prompt(`6) Save cryptokey to configuration? [default: No, which means you wil pass --key=XYZ on each future execution]` + nl)
	prompt(`> `)
	savekey, _ := reader.ReadString('\n')
	savekey = strings.TrimSpace(savekey)
	prompt(nl)

	c.URL = strings.TrimSpace(fqdn)
	c.SaveCryptoKey = c.shouldSaveKey(savekey)
	c.Username = strings.TrimSpace(username)
	c.Password = strings.TrimSpace(password)
	c.CryptoKey = strings.TrimSpace(key)
	c.CookiePort = strings.TrimSpace(cookiePort)

	return
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func (c *Config) writeFile(conffile string) {
	file, err := os.Create(conffile)
	if err != nil {
		log.Warnf(fmt.Sprintf("Cannot create default configuration file '%s':%s", conffile, err))
		return
	}
	defer file.Close()

	c.CLI.Prompt(fmt.Sprintf("-> Begin writing configuration file to '%s'", conffile))

	fmt.Fprintf(file, "#################################################\n")
	fmt.Fprintf(file, "##Splunk Cloud KNowledge OBject Backup (scknobb) \n")
	fmt.Fprintf(file, "##scknobb configuration created at: %v\n", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Fprintf(file, "#################################################\n")
	fmt.Fprintf(file, "url: %s\n", c.URL)
	fmt.Fprintf(file, "username: %s\n", c.Username)
	fmt.Fprintf(file, "password: %s\n", c.Password)
	fmt.Fprintf(file, "cookiePort: %s\n", c.CookiePort)
	if c.SaveCryptoKey == true {
		fmt.Fprintf(file, "cryptoKey: %s\n", c.CryptoKey)
	} else {
		c.CLI.Prompt("\n")
		c.CLI.Prompt("---\n")
		c.CLI.Prompt("NOTE: Crypto key was NOT written to configuration. You will\n")
		c.CLI.Prompt("      need to pass '--key=" + c.CryptoKey + "' when running.\n")
		c.CLI.Prompt("---\n")
	}
	fmt.Fprintf(file, "verboseLevel: 3\n")

	c.CLI.Prompt(fmt.Sprintf("\n√ Successfully!\nWrote configuration: '%s\n\n'", conffile))

	return
}
