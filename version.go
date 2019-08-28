package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	_G_HASH string
	_G_REVS string
	_BUILT_ string
)

func verinfo() string {
	self := filepath.Base(os.Args[0])
	_, h, r := getGitInfo()
	return fmt.Sprintf("%s V%d.%s", self, r, h)
}

func getGitInfo() (branch, hash string, revisions int) {
	branch = "branch_unkown"
	hash = "hash_unkown"
	bran, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return
	}
	branch = strings.TrimSpace(string(bran))
	ha, err := exec.Command("git", "log", "-1", "--pretty=format:%h").Output()
	if err != nil {
		return
	}
	hash = string(ha)

	revs, err := exec.Command("git", "log", "--oneline").Output()
	revisions = len(strings.Split(string(revs), "\n")) - 1
	return
}
