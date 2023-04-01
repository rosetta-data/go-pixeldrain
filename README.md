# A Pixeldrain client
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![CircleCI](https://circleci.com/gh/jkawamoto/go-pixeldrain.svg?style=svg)](https://circleci.com/gh/jkawamoto/go-pixeldrain)
[![Go Reference](https://pkg.go.dev/badge/github.com/jkawamoto/go-pixeldrain.svg)](https://pkg.go.dev/github.com/jkawamoto/go-pixeldrain)
[![codecov](https://codecov.io/gh/jkawamoto/go-pixeldrain/branch/master/graph/badge.svg?token=ppX3MVIqWA)](https://codecov.io/gh/jkawamoto/go-pixeldrain)
[![Release](https://img.shields.io/badge/release-0.5.2-brightgreen.svg)](https://github.com/jkawamoto/go-pixeldrain/releases/tag/v0.5.2)


## Usage
### Upload files
```shell
pd upload <path[:name]>...
```

`upload` command uploads files specified by the given `path`s to Pixeldrain and shows a URL to download them.
Each `path` can have an optional `name`. If a name is given, uploaded file will be renamed with it.

For example, this command reads `img.png` and uploads it as `another.png`:

```shell
pd upload img.png:another.png
```

If `path` is `-`, the uploading file is read from stdin. In this case, it's recommended to give a file name.
For example, this command reads data from stdin and uploads it as `output.log`:

```shell
pd upload -:output.log
```

If multiple files are given, an album consists of them will be created. By default, the album has a random name.
`--album` flag can specify the name.
For example, this command uploads two files and creates an album named `screenshots`:

```shell
pd upload --album screenshots img1.png img2.png
```

### Download a file
`download` command downloads a file from Pixeldrain and writes it to STDOUT.

```shell-session
$ pd download <file ID | URL>
```

If `-o` option is given with a directory path, the downloaded file is stored in
the directory instead of writing to STDOUT.

For example, this command downloads a file `abcdefg` in `~/Download`:
```shell-session
$ pd download abcdefg -o ~/Download
```

### Upload/Download a directory
Since this application supports uploading a file from STDIN and downloading a file to STDOUT,
it's also able to upload/download directories with `tar` command,.
For example, this command uploads `~/Documents` directory:

```shell-session
$ tar zcf - ~/Documents | pd upload -:documents.tar.gz
```

and this command downloads the file:

```shell-session
$ pd download <file id> | tar zxf - -C ~/Downloads
```



## Installation
If you're a Homebrew or Linuxbrew user, you can install this app by the following commands:

```
$ brew tap jkawamoto/pixeldrain
$ brew install pixeldrain
```

To build the newest version, use go get command:

```
$ go get github.com/jkawamoto/go-pixeldrain
```

Otherwise, compiled binaries are also available in [Github](https://github.com/jkawamoto/go-pixeldrain/releases).


## License
This software is released under the MIT License, see [LICENSE](LICENSE).
