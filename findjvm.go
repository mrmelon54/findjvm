package findjvm

import (
	"bufio"
	"errors"
	"github.com/Masterminds/semver/v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ErrMissingVersion = errors.New("missing version")

var ErrUnableToFindMatchingJVM = errors.New("unable to find matching jvm")

func FindPath(versionConstraint string) (string, error) {
	target, err := semver.NewConstraint(versionConstraint)
	if err != nil {
		return "", err
	}

	// search for jvm libraries
	jvmDir, err := os.ReadDir("/usr/lib/jvm/")
	if err != nil {
		return "", err
	}
	for _, i := range jvmDir {
		if !i.IsDir() {
			continue
		}
		jvmPath := filepath.Join("/usr/lib/jvm/", i.Name())
		version, err := GetVersion(jvmPath)
		if err != nil {
			return "", err
		}
		if target.Check(version) {
			return filepath.Join(jvmPath, "bin/java"), nil
		}
	}
	return "", ErrUnableToFindMatchingJVM
}

func GetVersion(javaPath string) (*semver.Version, error) {
	openRelease, err := os.Open(filepath.Join(javaPath, "release"))
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(openRelease)
	for sc.Scan() {
		t := sc.Text()
		n := strings.Index(t, "=")
		if n == -1 {
			continue
		}
		if t[:n] != "JAVA_VERSION" {
			continue
		}
		unquote, err := strconv.Unquote(t[n+1:])
		if err != nil {
			continue
		}
		return semver.NewVersion(unquote)
	}
	return nil, ErrMissingVersion
}
