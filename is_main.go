package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	if TagIsOnMainHead(os.Getenv("GITHUB_REF")) {
		fmt.Println("Is on main")
		return
	}
	fmt.Println("Is not on main")
}

func TagIsOnMainHead(tag string) bool {
	// Fetch main.
	_ = git("fetch", "--depth", "1", "origin", "main")
	branch := git("branch", "-r", "--contains", tag)
	fmt.Printf("{%v}\n", branch)
	return regexp.MustCompile("(?m)/main$").MatchString(branch)
}

func git(args ...string) string {
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if err.ProcessState.ExitCode() == 129 {
				log.Printf("INFO: Ignoring %v\n", err)
				return ""
			}
		}
		log.Fatalf("ERROR: git failed: %q: %q (%q)", err, out, args)
	}

	return string(out)
}
