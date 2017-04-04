gobuildall
===========

Builds Go packages for every OS and architecture

Building
--------

    go build

Usage
-----

    gobuildall [arguments]

gobuildall runs the go build command for every supported OS and architecture
combination. If arguments are specified, they are passed to the go build
command. Executables for each combination are output to files named as
os-architecture.