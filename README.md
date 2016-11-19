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

