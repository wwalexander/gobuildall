package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	arch386     = "386"
	archamd64   = "amd64"
	archarm     = "arm"
	archarm64   = "arm64"
	archppc64   = "ppc64"
	archppc64le = "ppc64le"
)

// https://golang.org/doc/install/source#environment
var osarchs = map[string][]string{
	"darwin": []string{
		arch386,
		archamd64,
		archarm,
		archarm64,
	},
	"dragonfly": {
		archamd64,
	},
	"freebsd": {
		arch386,
		archamd64,
		archarm,
	},
	"linux": {
		arch386,
		archamd64,
		archarm,
		archarm64,
		archppc64,
		archppc64le,
	},
	"netbsd": {
		arch386,
		archamd64,
		archarm,
	},
	"openbsd": {
		arch386,
		archamd64,
		archarm,
	},
	"plan9": {
		arch386,
		archamd64,
	},
	"solaris": {
		archamd64,
	},
	"windows": {
		arch386,
		archamd64,
	},
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for osname, archs := range osarchs {
		osPath := path.Join(args[0], osname)
		if err := os.MkdirAll(osPath, 0755); err != nil {
			log.Fatal(err)
		}
		if err := os.Setenv("GOOS", osname); err != nil {
			log.Fatal(err)
		}
		for _, arch := range archs {
			if err := os.Setenv("GOARCH", arch); err != nil {
				log.Fatal(err)
			}
			buildArgs := append([]string{"build", "-o", path.Join(osPath, arch)}, args[1:]...)
			cmd := exec.Command("go", buildArgs...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("building %s/%s failed:\n%s", osname, arch, out)
			} else {
				log.Printf("building %s/%s succeeded", osname, arch)
			}
		}
	}
}
