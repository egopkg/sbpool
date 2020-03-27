# sbpool
Easy to use strings builder pool

### Example

```go
package main

import (
	"strings"

	"github.com/ergopkg/sbpool"
)

func main() {
	sbOne := sbpool.AcquireStringsBuilder()
	sbOne.WriteString("hello")
	sbOne.WriteString(" ")
	sbOne.WriteString("world")
	sbTwo := sbpool.AcquireStringsBuilder()
    
	defer func() {
		sbpool.ReleaseStringsBuilder(sbOne)
		sbpool.ReleaseStringsBuilder(sbTwo)
	}()
}
```

### Concurrency supporting

As `bytes.Buffer`, the `strings.Builder` doesn’t support concurrency when writing and reading. So we should take care it if we need them.
We can try a little bit with `strings.Builder` to add `10000` character, at the same time.

```go
package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ergopkg/sbpool"
)

func main() {
	b := sbpool.AcquireStringsBuilder()
    defer sbpool.ReleaseStringsBuilder(b)
	var wait sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wait.Add(1)
		go func() {
			b.WriteString("1")
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(len(b.String()))
}

```

If you run it, you had different result’s lengths. But they aren’t enough `10000` as we add.

```
go run main.go => 7329
go run main.go => 7650
go run main.go => 7623
```

### io.Writer interface

The io.Writer interface is implemented on `strings.Builder` with `Write()` method `Write(p []byte) (n int, err error)`. 
So, we have a lot of useful case with `io.Writer`:

 - `io.Copy(dst Writer, src Reader) (written int64, err error)`
 - `bufio.NewWriter(w io.Writer) *Writer`
 - `fmt.Fprint(w io.Writer, a …interface{}) (n int, err error)`
 - `func (r *http.Request) Write(w io.Writer) error`
 - and other libraries that uses io.Writer
 
### Usefull articles

- [7 notes about strings.builder in Golang](https://medium.com/@thuc/8-notes-about-strings-builder-in-golang-65260daae6e9)

### Special thanks

Inspired by valyala/fasthttp ideas. 