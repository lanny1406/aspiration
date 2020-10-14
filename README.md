Quickstart 
==========
#### Install [Go] https://golang.org/doc/install

```go
go run github.com/lanny1406/aspiration
```

or 

#### Import packages in your code:
```
import (
	"fmt"

	"github.com/lanny1406/aspiration/mapper"
)
```

```
func main() {
	// result asPirAtiOn.cOm
   fmt.Println(mapper.Every3rdChar("Aspiration.com"))
}

```
# test

go test -v  -count=1  github.com/lanny1406/aspiration/...

# run

go run github.com/lanny1406/aspiration