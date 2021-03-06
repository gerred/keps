package settings

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/calebamiles/keps/pkg/sigs"
)

const (
	ContentRootEnv = "KEP_CONTENT_ROOT"
)

// FindContentRoot looks for a location with the following structure
// <some-containing-dir>/content/<all-known-sigs>
// returning `content`
func FindContentRoot() (string, error) {
	var foundRoot string

	startLocation, err := contentSearchRoot()
	if err != nil {
		return "", err
	}

	envRoot := os.Getenv(ContentRootEnv)
	cachedRoot, err := contentRoot()
	if err != nil {
		return "", err
	}

	switch {
	case hasDirForEachSIG(envRoot):
		foundRoot = envRoot
	case hasDirForEachSIG(cachedRoot):
		foundRoot = cachedRoot
	default:
		err = filepath.Walk(startLocation, func(path string, info os.FileInfo, wErr error) error {
			if foundRoot != "" {
				return nil
			}

			if wErr != nil {
				return wErr
			}

			if !info.IsDir() {
				return nil
			}

			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}

			if hasDirForEachSIG(path) {
				foundRoot = path
			}

			return nil
		})

		if err != nil {
			return "", err
		}

	}

	if foundRoot == "" {
		return "", fmt.Errorf("could not find KEP content under: %s", startLocation)
	}

	return foundRoot, nil
}

func contentRoot() (string, error) {
	settingsFileLocation, err := findSettingsFile()
	if err != nil {
		return "", err
	}

	s := &User{}
	err = readSettingsFile(settingsFileLocation, s)
	if err != nil {
		log.Warn("reading user settings file")
		return "", nil
	}

	return s.ContentRoot, nil
}

func hasDirForEachSIG(p string) bool {
	knownSIGs := sigs.All()
	for _, s := range knownSIGs {
		if _, sErr := os.Stat(filepath.Join(p, s)); os.IsNotExist(sErr) {
			return false
		}
	}

	return true
}

func contentSearchRoot() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Error("could not get current user information")
		return "", err
	}

	homeDir := usr.HomeDir
	return homeDir, nil
}
