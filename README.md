# shot
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

## Reference

[Handle inter-process communication between PhantomJS and Golang processes via hixie-76 websockets](http://studygolang.com/articles/3659)