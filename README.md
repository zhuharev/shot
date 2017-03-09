# shot [![Go Report Card](https://goreportcard.com/badge/github.com/zhuharev/shot)](https://goreportcard.com/report/github.com/zhuharev/shot) [![GoDoc](https://godoc.org/github.com/zhuharev/shot?status.svg)](http://godoc.org/github.com/zhuharev/shot)
site screenshot microservice with go and phantomjs

## Usage 

```
package main

import (
	"github.com/zhuharev/shot"
)

func main() {
	srv, e := shot.New("cfg")
	if e != nil {
		panic(e)
	}
	srv.Run()
	srv.Serve()
}
```

## Screenshots
<table>
  <tr>
    <td width="50%"><img src="https://raw.githubusercontent.com/zhuharev/shot/master/example/github.com.png">
    <td width="50%"><img src="https://raw.githubusercontent.com/zhuharev/shot/master/example/google.com.png">
  </tr>
</table>
