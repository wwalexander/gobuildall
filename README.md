gobuildall
===========

Builds Go packages for every OS and architecture

Building
--------

    go build

Usage
-----

    gobuildall [path] [arguments]

gobuildall runs the go build command for every supported OS and architecture
combination. If arguments are specified, they are passed to the go build
command. The output flag of each go build command is set to the given path;
the output's file name is of the form os-architecture.

If the -o flag is specified in the given arguments, the files will not save in
the correct location but will instead overwrite each other. Packages using cgo
cannot be built.
