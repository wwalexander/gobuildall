package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	osandroid   = "android"
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
	arch386      = "386"
	archamd64    = "amd64"
	archarm      = "arm"
	archarm64    = "arm64"
	archmips     = "mips"
	archmipsle   = "mipsle"
	archmips64   = "mips64"
	archmips64le = "mips64le"
	archppc64    = "ppc64"
	archppc64le  = "ppc64le"
)

// https://golang.org/doc/install/source#environment
var osarchs = map[string][]string{
	osandroid: {
		archarm,
	},
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
		archmips,
		archmipsle,
		archmips64,
		archmips64le,
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
// with the given arguments.
func Build(osname string, arch string, args []string) error {
	if err := os.Setenv("GOOS", osname); err != nil {
		return err
	}
	if err := os.Setenv("GOARCH", arch); err != nil {
		return err
	}
	path := osname + "-" + arch
	if osname == "windows" {
		osname += ".exe"
	}
	cmd := exec.Command("go", append([]string{
		"build",
		"-o",
		path,
	}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

const usage = `usage: gobuildall [arguments]

gobuildall runs the go build command for every supported OS and architecture
combination. If arguments are specified, they are passed to the go build
command. The output flag of each go build command is set to the named path;
the file name is of the form os-architecture.`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		if arg == "-o" {
			log.Fatal("-o is not allowed")
		}
	}
	for osname, archs := range osarchs {
		for _, arch := range archs {
			if err := Build(
				osname,
				arch,
				args,
			); err != nil {
				fmt.Fprintf(
					os.Stderr,
					"building %s/%s failed: %v\n",
					osname,
					arch,
					err,
				)
				continue
			}
			fmt.Fprintf(
				os.Stderr,
				"building %s/%s succeeded\n",
				osname,
				arch,
			)
		}
	}
}
