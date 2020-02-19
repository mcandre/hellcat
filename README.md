# Hellcat: fs swiss army knife

![demon kitty](https://raw.githubusercontent.com/mcandre/hellcat/master/hellcat.png)

# examples/

```console
$ hh
-rw-rw-rw-    1   andrew    staff  95B 2020-02-19T16:44:00Z 1x1.png
Lrwxr-xr-x    1   andrew    staff   8B 2020-02-19T17:02:14Z b.txt -> test.txt
Lrwxr-xr-x    1   andrew    staff  10B 2020-02-19T18:39:49Z c.txt -> nosuchfile
-rw-r--r--    1   andrew    staff  29B 2020-02-19T15:34:37Z test.txt
Lrwxr-xr-x    1   andrew    staff   7B 2020-02-19T19:23:56Z z.png -> 1x1.png

$ hh test.txt
all ur base are belong to us

$ hh 1x1.png
00000000 89 50 4e 47 0d 0a 1a 0a
00000008 00 00 00 0d 49 48 44 52
00000016 00 00 00 01 00 00 00 01
00000024 01 03 00 00 00 25 db 56
00000032 ca 00 00 00 03 50 4c 54
00000040 45 00 00 00 a7 7a 3d da
00000048 00 00 00 01 74 52 4e 53
00000056 00 40 e6 d8 66 00 00 00
00000064 0a 49 44 41 54 08 d7 63
00000072 60 00 00 00 02 00 01 e2
00000080 21 bc 33 00 00 00 00 49
00000088 45 4e 44 ae 42 60 82
```

See `hh -h` for more options.

# ABOUT

Hellcat is a combination of familiar file system diagnostics, such as `ls`, `cat`, and `od`. Hellcat is convenient for quickly navigating large project directories. You can zoom through nested projects, previewing contents right within the terminal!

## Notable Features

* Performant, low memory requirement
* Accepts directories and regular files
* Easy for light typists
* Accepts symlinks, including broken symlinks
* Accepts text and binary files
* Can dump hex file contents
* Can recurse

# DOWNLOAD

https://github.com/mcandre/hellcat/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/hellcat

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/hellcat/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/hellcat.git $GOPATH/src/github.com/mcandre/hellcat
$ cd $GOPATH/src/github.com/mcandre/hellcat
$ git submodule update --init --recursive
$ go install ./...
```

# TEST REMOTELY

```console
$ go test github.com/mcandre/go-ios7crypt/...
```

# TEST LOCALLY

```console
$ go test
```

# COVERAGE

```console
$ mage coverageHTML
$ karp cover.html
```

# PORT

```console
$ mage port
```

# LINT

Keep the code tidy:

```console
$ mage lint
```

# CREDITS

Hellcat operates under an 80% utility principle. When in doubt, look to the classic tools:

* [ls](https://linux.die.net/man/1/ls)
* [cat](https://linux.die.net/man/1/cat)
* [od](https://linux.die.net/man/1/od)
* [mount](https://linux.die.net/man/8/mount)
* [fsck](https://linux.die.net/man/8/fsck)
* [grep](https://linux.die.net/man/1/grep)
* [find](https://linux.die.net/man/1/find)
