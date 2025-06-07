# urlpath

[![urlpath release (latest SemVer)](https://img.shields.io/github/v/release/foohq/urlpath?sort=semver)](https://github.com/foohq/urlpath/releases)
[![Go Reference](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/foohq/urlpath)

Module urlpath implements utility functions for manipulating paths. Paths can be represented as Unix, Windows paths, and URLs.

For example, the following paths are accepted by the module:

* `/home/user/file.txt`
* `file:///home/user/file.txt`
* `http://www.example.com/assets/file.js`
* `C:\Windows\System32\cmd.exe` (DOS path)
* `\\10.10.0.55\COMP\customers.xls` (UNC path)

> [!NOTE]
> DOS and UNC paths are only supported when `GOOS=windows`.

## Installation

```
go get github.com/foohq/urlpath
```

## Usage

```go
import "github.com/foohq/urlpath"

dir, base, err := urlpath.Split("http://www.example.com/assets/file.js")
if err != nil {
	panic(err)
}

println(dir)
// Prints "http://www.example.com/assets"
println(base)
// Prints "file.js"

pth, err := urlpath.Join(dir, "another.js")
if err != nil {
	panic(err)
} 

println(pth)
// Prints "http://www.example.com/assets/another.js"
```

## License

This module is distributed under the Apache License Version 2.0 found in the [LICENSE](./LICENSE) file.
