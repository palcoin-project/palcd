palcd
====
palcd is a full node palcoin implementation written in Golang.


Palcd does *NOT* include wallet functionality This means you can't make or receive payments
directly with palcd.  That functionality is provided by the
[palcwallet](https://github.com/palcoin-project/palcwallet).

## Requirements

[Go](http://golang.org) 1.14 or newer.

## Installation

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain palcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/palcoin-project/palcd
$ GO111MODULE=on go install -v . ./cmd/...
```

- palcd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update palcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/palcoin-project/palcd
$ git pull
$ GO111MODULE=on go install -v . ./cmd/...
```
