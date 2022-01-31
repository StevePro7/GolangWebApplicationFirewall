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

var filenamesX []string

const defaultRulesetDirectory = "/etc/waf/"

func Check(directory string) error {

	DefineRulesSetDirectory(directory)

	err := CheckRulesSetDirectoryExists()
	if os.IsNotExist(err) {
		return err
	}

	err = ExtractRulesSetFilenames()
	if err != nil {
		return err
	}

	log.Info("WAF is Enabled!")
	wafIsEnabled = true
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
	var files []string
	items, err := ioutil.ReadDir(rulesetDirectory)
	if err != nil {
		return err
	}

	// todo remove
	for _, item := range items {
		log.Printf("'%s'\n", item.Name())
	}

	// Sort files descending to ensure lower cased files like crs-setup.conf are loaded first.
	// This is a requirement for Core Rules Set and REQUEST-901-INITIALIZATION.conf bootstrap.
	sortFileNameDescend(items)

	// todo remove
	for _, item := range items {
		log.Printf("'%s'\n", item.Name())
	}

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
		files = append(files, file)
		filenamesX = append(filenamesX, file)
		log.Infof("WAF Found Rules File[%d]('%s')", count, file)
		count++
	}

	log.Infof("WAF Total Core Rules Set files: %d", len(files))
	return nil
}

// CheckRulesSetExists
// invoke this WAF function first passing in the Core Rule Set directory or /etc/waf/ as per default;
// if this directory does not exist OR zero *.conf Core Rule Sets files exist then do not enable WAF.
func CheckRulesSetExists() {

	if _, err := os.Stat(rulesetDirectory); os.IsNotExist(err) {
		log.Printf("WAF Core Rules Set directory: '%s'  does not exist.  WAF not enabled!", rulesetDirectory)
		return
	}

	//items, err := ioutil.ReadDir(rulesetDirectory); os.IsNotExist(err) {
	//}

	log.Printf("WAF Core Rules Set directory: '%s'  Total Core Rules Set files: 0.  WAF not enabled!", rulesetDirectory)
	wafIsEnabled = true
}

// IsEnabled helper function used by client calling cod
func IsEnabled() bool {
	return wafIsEnabled
}

func GetFilenames() []string {
	return filenamesX
}

func sortFileNameDescend(items []os.FileInfo) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Name() > items[j].Name()
	})
}
