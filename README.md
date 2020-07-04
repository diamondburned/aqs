# aqs

Anime quotes generator, used for test data.

Source: https://github.com/jiashengc/anime-quote-generator-v2/

## Usage

```go
package main

import (
	"fmt"

	"github.com/diamondburned/aqs"
	_ "github.com/diamondburned/aqs/data"
)

func main() {
	var character = aqs.RandomCharacter()
	fmt.Printf("%s once said: %q\n", character.Name, character.RandomQuote())
}
```

## Disclaimer

This package contains possibly copyrighted content not owned by me.
