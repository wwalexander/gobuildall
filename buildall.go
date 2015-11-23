package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	osdarwin    = "darwin"
	osdragonfly = "dragonfly"
	osfreebsd   = "freebsd"
	oslinux     = "linux"
	osnetbsd    = "netbsd"
	osopenbsd   = "openbsd"
	osplan9     = "plan9"
	ossolaris   = "solaris"
	oswindows   = "windows"
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
	osdarwin: {
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

// Build builds a Go package for the given operating system and architecture
// with the given arguments, saving the output to the folder given by root.
func Build(osname string, arch string, args []string, root string) error {
	if err := os.Setenv("GOOS", osname); err != nil {
		return err
	}
	if err := os.Setenv("GOARCH", arch); err != nil {
		return err
	}
	outPath := path.Join(root, osname+"-"+arch)
	if osname == oswindows {
		outPath += ".exe"
	}
	args = append([]string{
		"build",
		"-o",
		outPath,
	}, args...)
	cmd := exec.Command("go", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, out)
		return err
	}
	return nil
}

const usage = `usage: gobuildall [path] [arguments]

gobuildall runs the go build command for every supported OS and architecture
combination. If arguments are specified, they are passed to the go build
command. The output flag of each go build command is set to the given path;
the output's file name is of the form os-architecture.

If the -o flag is specified in the given arguments, the files will not save in
the correct location but will instead overwrite each other. Packages using cgo
cannot be built.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	root := args[0]
	buildArgs := args[1:]
	if err := os.Mkdir(root, 0755); err != nil {
		log.Fatal(err)
	}
	for osname, archs := range osarchs {
		for _, arch := range archs {
			if err := Build(osname, arch, buildArgs, root); err != nil {
				fmt.Fprintf(os.Stderr, "building %s/%s failed: %v\n", osname, arch, err)
				continue
			}
			fmt.Fprintf(os.Stderr, "building %s/%s succeeded", osname, arch)
		}
	}
}
