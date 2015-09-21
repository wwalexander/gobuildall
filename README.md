gobuildall
===========

A tool that builds Go packages for every OS and architecture

Building
--------

    go build

Usage
-----

Run the `go build` command for
[every supported OS and architecture combination](https://golang.org/doc/install/source#environment)
with `ARGUMENTS`, saving the output of each command to `PATH/[OS]-[arch]`:

    gobuildall [PATH] [ARGUMENTS]

For instance, running

    gobuildall bin -race

in a Go package directory would run

    go build -race

for each OS/architecture combination, creating a hierarchy like:

    bin/
        darwin-386
        darwin-amd64
        darwin-arm
        darwin-arm64
        dragonfly-amd64
        ...

Notes
-----

*   Packages using `cgo` cannot cross-compile without rebuilding the toolchain.

*   If the `-o` flag is included in `ARGUMENTS`, the executables will not
    save in the correct location but will instead overwrite each other.

*   [mitchellh/gox](https://github.com/mitchellh/gox) is a more full-featured
    tool, which works with Go <1.5 and `cgo`. gobuildall is more similar to a
    shell script than a real program.
