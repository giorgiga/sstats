# sstats

A package with (very) simple statistics utilities in go.

License: [0BSD](https://spdx.org/licenses/0BSD.html).

## Install:

```
go get github.com/giorgiga/sstats
```

## Using the Summary struct

`Summary` collects statistics on the data points it meets.
```go
package main

import (
	"fmt"
	"github.com/giorgiga/sstats"
)

func main() {
	var summary sstats.Summary = sstats.MakeSummary()
	// meet some datapoints
	summary.Meet(8)
	summary.Meet(10)
	summary.Meet(12)
	// have some statistics
	fmt.Printf("mean: %v\n", summary.Mean())
	fmt.Printf("variance: %v\n", summary.Variance())
}```

Output:
```
mean: 10
variance: 2.6666666666666665
```
