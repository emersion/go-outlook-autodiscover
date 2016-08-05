package autodiscover

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

const (
	officePath = "Software\\Microsoft\\Office"
	autoDiscoverPath = "Outlook\\AutoDiscover"
)

func getVersions() (names []string, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, officePath, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return
	}
	defer k.Close()

	names, err = k.ReadValueNames(-1)
	return
}

// Check if an autodiscover file can be set up. Usually false is returned when
// Microsoft Outlook is not installed.
func CanSetup() bool {
	versions, err := getVersions()
	if err != nil {
		return false
	}

	return len(versions) > 0
}

// Setup an autodiscover file locally.
func Setup(filepath, domain string) error {
	versions, err := getVersions()
	if err != nil {
		return err
	}

	if len(versions) == 0 {
		return errors.New("autodiscover: cannot find installed Outlook version")
	}

	p := officePath + "\\" + versions[0] + "\\" + autoDiscoverPath
	k, err := registry.OpenKey(registry.CURRENT_USER, p, registry.QUERY_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	return k.SetStringValue(domain, filepath)
}
