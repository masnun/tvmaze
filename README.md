## TVMaze Search API for Go

Install the package first: 

```
go get -u github.com/masnun/tvmaze
```

Then start searching TV Shows: 

```go
package main

import "github.com/masnun/tvmaze"
import "fmt"

func main() {
	results := tvmaze.Search("doctor who")
	for _, result := range results {
		fmt.Println(result.Show.Name)
	}
}

```

You can check out the [Types](https://github.com/masnun/tvmaze/blob/master/types.go)
to know which fields are available. 