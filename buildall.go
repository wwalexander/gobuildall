package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	osdarwin = "darwin"
	osdragonfly = "dragonfly"
	osfreebsd = "freebsd"
	oslinux = "linux"
	osnetbsd = "netbsd"
	osopenbsd = "openbsd"
	osplan9 = "plan9"
	ossolaris = "solaris"
	oswindows = "windows"
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
	osdarwin: []string{
		arch386,
		archamd64,
		archarm,
		archarm64,
	},
	osdragonfly: {
		archamd64,
	},
	osfreebsd: {
		arch386,
		archamd64,
		archarm,
	},
	oslinux: {
		arch386,
		archamd64,
		archarm,
		archarm64,
		archppc64,
		archppc64le,
	},
	osnetbsd: {
		arch386,
		archamd64,
		archarm,
	},
	osopenbsd: {
		arch386,
		archamd64,
		archarm,
	},
	osplan9: {
		arch386,
		archamd64,
	},
	ossolaris: {
		archamd64,
	},
	oswindows: {
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
	root := args[0]
	if err := os.Mkdir(root, 0755); err != nil {
		log.Fatal(err)
	}
	for osname, archs := range osarchs {
		if err := os.Setenv("GOOS", osname); err != nil {
			log.Fatal(err)
		}
		for _, arch := range archs {
			if err := os.Setenv("GOARCH", arch); err != nil {
				log.Fatal(err)
			}
			outPath := path.Join(root, osname+"-"+arch)
			if osname == oswindows {
				outPath += ".exe"
			}
			buildArgs := append([]string{
					"build",
					"-o",
					outPath,
				}, args[1:]...)
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
