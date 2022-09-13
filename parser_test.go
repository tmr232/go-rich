package rich

import (
	"fmt"
	"testing"
)

func Test_tokenize(t *testing.T) {
	fmt.Println(tokenize("Hello, [green]Wo[[[r\\[ld!"))
	s := "abc"
	r := []rune(s)
	fmt.Println(&s, &r)

	parts, err := parse(tokenize("Hello, [green]World!"))
	if err != nil {
		panic(err)
	}
	fmt.Println(parts)
}
