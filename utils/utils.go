package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/lab42/semver/config"
)

const (
	GIT = "git"
)

func Run(args ...string) (string, error) {
	cmd := exec.Command(GIT, args...)
	out, err := cmd.CombinedOutput()

	return strings.TrimSuffix(string(out), "\n"), err
}

// Return all tgs for a repository
func LatestTags() string {
	out, err := Run("describe", "--tags", "--abbrev=0")
	if err != nil {
		return "0.0.0"
	}

	return out
}

// Return git log messages
func Logs() []string {
	out, err := Run("log", "--pretty=\"%s\"", fmt.Sprintf("%s..HEAD", LatestTags()))
	if err != nil {
		return make([]string, 0)
	}

	return strings.Split(out, "\n")
}

// Get current latest tag from repository
func Current() string {
	return semver.MustParse(LatestTags()).String()
}

// Get next tag to be created according to repository
func Next() semver.Version {
	incMajor := false
	incMinor := false
	incPatch := false

	for _, rule := range config.Config.Rules {
		r := regexp.MustCompile(rule.Regex)

		logs := Logs()

		if len(logs) == 0 {
			incPatch = true
			continue
		}

		for _, log := range logs {
			if loc := r.FindIndex([]byte(log)); loc != nil {
				switch rule.Bump {
				case "major":
					incMajor = true
				case "minor":
					incMinor = true
				case "patch":
					incPatch = true
				}
			}
		}
	}

	switch true {
	case incMajor:
		return IncMajor()
	case incMinor:
		return IncMinor()
	case incPatch:
		return IncPatch()
	default:
		return *semver.MustParse(Current())
	}
}

// Increase major version
func IncMajor() semver.Version {
	v := semver.MustParse(Current())
	return v.IncMajor()
}

// Increase minor version
func IncMinor() semver.Version {
	v := semver.MustParse(Current())
	return v.IncMinor()
}

// Increase patch version
func IncPatch() semver.Version {
	v := semver.MustParse(Current())
	return v.IncPatch()
}
