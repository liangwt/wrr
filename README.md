## Weighted Round-Robin

Implementation of weighted round-robin and smooth weighted round-robin

Normal weighted round-robin was referenced by [LVS weighted round-robin](http://kb.linuxvirtualserver.org/wiki/Weighted_Round-Robin_Scheduling)

Smooth weighted round-robin was referenced by [nginx smooth weighted round-robin balancing](https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35)

As the name implies, smooth weighted round-robin algorithm avoids the choice of the same point at a certain moment.

Well known, time complexity of both is O(n) when looking for the next element. In this implementation, cache will set while first looping, O(1) time complexity after that

### USAGE

```bash
go get github.com/liangwt/wrr
```

```go
package main

import (
	"fmt"
	"github.com/liangwt/wrr"
)

func main() {
	points := []*wrr.Point{
		{Entry: "A", Weight: 5},
		{Entry: "B", Weight: 2},
		{Entry: "C", Weight: 3},
	}

	iter := wrr.NewWrr(points)
	for i := 0; i < 12; i++ {
		fmt.Printf("%s ", iter.Next().Entry)
	}

	fmt.Println()

	smIter := wrr.NewSmoothWrr(points)
	for i := 0; i < 12; i++ {
		fmt.Printf("%s ", smIter.Next().Entry)
	}
}
```

### FEATURE

- O(1) time complexity, benefit from cache
- test cover
