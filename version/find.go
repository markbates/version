package version

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

// Find a semver version from a reader:
/*
package mypkg

const Version = "v0.12.6"
*/
func Find(in io.Reader, allowDev bool) (string, error) {
	re := regexp.MustCompile(`[const|var] [vV]ersion = "(.+)"`)

	bb, err := ioutil.ReadAll(in)
	if err != nil {
		return "", errors.WithStack(err)
	}

	var matches []string
	lines := string(bb)
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			continue
		}
		matches = re.FindStringSubmatch(line)
		if len(matches) > 1 {
			break
		}
	}
	if len(matches) < 2 {
		return "", errors.New("failed to find the version")
	}

	v := matches[1]
	if allowDev && strings.HasPrefix(v, "dev") {
		return v, nil
	}
	if _, err = semver.NewVersion(v); err != nil {
		return "", errors.WithStack(err)
	}
	return v, nil
}
