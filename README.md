# go-hatena

### Overview
[Hatena](http://developer.hatena.ne.jp/ja/documents/bookmark/apis/rest) API client by golang.  

<br/>
### Quick Start
```shell
$ go get github.com/takecy/go-hatena/hatena
```
```go
cli := hatena.NewHatena(nil)
count, err := cli.Bookmarks.Count("https://www.google.co.jp/")
```

and see [examples](examples)

<br/>
### Features
* Count API
 - [x] count bookmarks of URL
* REST APIs
 - [ ] Entries
 - [ ] Bookmarks
 - [ ] Tags
 - [ ] Users

<br/>
### Development
go version 1.5.1+
```shell
$ git clone git@github.com:takecy/go-hatena.git
$ go get github.com/tools/godep
$ godep restore
$ make test
```

### License
MIT
