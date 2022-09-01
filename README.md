# go-rich

Go CLI output utils inspired by the amazing Python package [rich](https://github.com/Textualize/rich).

This library is currently in very early development.

## Printing in Color!

You can use `rich.Stylize` to add styles to your strings.

```go
package main

import (
	"fmt"
	"github.com/tmr232/go-rich"
)

func main() {
	fmt.Println(rich.Stylize("[green]Hello, [#5555dd underline]World!"))
}
```

Will print `Hello, ` in green, and `World!` in `#5555dd`.