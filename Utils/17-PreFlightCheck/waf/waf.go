package waf

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Assume WAF is not enabled by default.
var wafIsEnabled = false

// Directory where the Core Rules Set are stored.
var rulesetDirectory string

// Slice of filenames read from Core Rules Set directory.
var filenames []string

const defaultRulesetDirectory = "/etc/waf/"

// CheckRulesSetExists
// invoke this WAF function first passing in the Core Rule Set directory or /etc/waf/ as per default;
// if this directory does not exist OR zero *.conf Core Rule Sets files exist then do not enable WAF.
func CheckRulesSetExists(directory string) error {

	DefineRulesSetDirectory(directory)

	err := CheckRulesSetDirectoryExists()
	if os.IsNotExist(err) {
		return err
	}

	err = ExtractRulesSetFilenames()
	if err != nil {
		return err
	}

	wafIsEnabled = len(filenames) > 0
	return nil
}

func DefineRulesSetDirectory(directory string) {
	rulesetDirectory = directory

	// Defend against root access "/".
	if len(rulesetDirectory) < 2 {
		rulesetDirectory = defaultRulesetDirectory
	}
	log.Printf("WAF Core Rules Set directory: '%s'", rulesetDirectory)

	// Ensure rules directory ends with trailing slash.
	if !strings.HasSuffix(rulesetDirectory, "/") {
		rulesetDirectory = rulesetDirectory + "/"
	}
}

func CheckRulesSetDirectoryExists() error {

	_, err := os.Stat(rulesetDirectory)
	if os.IsNotExist(err) {
		return err
	}

	return nil
}

func ExtractRulesSetFilenames() error {

	// Read all core rule set file names from rules directory.
	items, err := ioutil.ReadDir(rulesetDirectory)
	if err != nil {
		return err
	}

	// Sort files descending to ensure lower cased files like crs-setup.conf are loaded first.
	// This is a requirement for Core Rules Set and REQUEST-901-INITIALIZATION.conf bootstrap.
	sortFileNameDescend(items)

	count := 1
	for _, item := range items {

		// Ignore files starting with ".." that link to the parent directory.
		filename := item.Name()
		if strings.HasPrefix(filename, "..") {
			continue
		}

		// Only load *.conf configuration files.
		if !strings.HasSuffix(filename, ".conf") {
			continue
		}

		file := rulesetDirectory + filename
		filenames = append(filenames, file)
		log.Infof("WAF Found Rules File[%d]('%s')", count, file)
		count++
	}

	log.Infof("WAF Total Core Rules Set files: %d", len(filenames))
	return nil
}

// IsEnabled helper function used by client calling code.
func IsEnabled() bool {
	return wafIsEnabled
}

// GetRulesSetFilenames helper function used for unit tests.
func GetRulesSetFilenames() []string {
	return filenames
}

func sortFileNameDescend(items []os.FileInfo) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Name() > items[j].Name()
	})
}
