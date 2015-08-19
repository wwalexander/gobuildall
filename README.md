go-buildall
===========

A tool which builds Go programs for every OS and architecture

Building
--------

    go build

Usage
-----

Run the `go build` command for
[every supported OS and architecture combination](https://golang.org/doc/install/source#environment)
with `ARGUMENTS`, saving the output of each command to `PATH/[OS]/[arch]`:

    go-buildall [PATH] [ARGUMENTS]

For instance, running:

    go-buildall buildall -race

would run

    go build -race

for each OS/architecture combination, creating a hierarchy like

    buildall/
        darwin/
		386
	darwin/
		amd64
	darwin/
		arm
	darwin/
		arm64
	dragonfly/
		amd64
	...
