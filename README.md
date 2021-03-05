# go-random

## A fast, clear, and cryptographically-secure random data generator for Golang

go-random is a fast, clear, and cryptographically-secure random data generator for Golang. It was developed out of the frustration on finding simple ways to write random string and data generator in Golang without having to write complex functions every time.

Why? Security should be baked-in within frameworks for developers. Security should be the easy option.

go-random uses the standard Golang APIs to generate randomness (crypto/rand), with an abstracted API that is clear to use.

# Features

- Clear and simple API
- Extremely fast
- Cryptographically-secure (based on crypto/rand) Golang APIs
- Concurrency-safe: You can run millions of go-routines and not have a token collision as the randomness seed is different for each call.

# Installation

```shell
go get -u github.com/mazen160/go-random
```

# Usage

## Generating Random String

A secure string (token) with the length of 32 characters.

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	data, err := random.String(32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
```

Output:

```
> G83zSBEbKeCiV3ij0WcuqR0lIZRJMMc2
```

## Generating Random Bytes

Random bytes with the length of 32 bytes.

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	data, err := random.Bytes(32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

```

Output:

```
> [57 63 142 3 148 39 175 87 69 12 178 46 138 42 11 175 93 13 142 0 97 231 132 200 151 143 241 82 235 167 34 33]
```

## Generating Random Integer

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	data, err := random.GetInt(1024)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
```

Output:

```
> 981
```

## Finding a random integer between 50 to 2000.

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	data, err := random.IntRange(50, 2500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
```

Output:

```
> 1527
```

## Selecting a random set of characters from a given character set (charset)

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	charset := "abcde01235"
	length := 20
	data, err := random.Random(length, charset, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
```

Output:

```
> 0ee130e152dce0bb5123
```

## Select a random string from a slice of strings

```golang
package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	s := []string{"a", "b", "c", "d", "e", "f"}
	data := random.Choice(s)

	fmt.Println(data)
}
```

Output:

```
> d
```

### Use Pre-defined Charsets

go-random has a pre-defined charset:

* `Digits` Digits: `[0-9]`
* `ASCIILettersLowercase`: Asci Lowercase Letters: `[a-z]`
* `ASCIILettersUppercase`: Ascii Uppercase Letters: `[A-Z]`
* `Letters`: Ascii Letters: `[a-zA-Z]`
* `ASCIICharacters`: Ascii Charaters: `[a-zA-Z0-9]`
* `Hexdigits`: Hex Digits: `[0-9a-fA-F]`
* `Octdigits`: Octal Digits: `[0-7]`
* `Punctuation`: Punctuation and special characters.
* `Printables`: All printables.

To use, you can access it as: `random.Letters`.

# Contribution

Contributions are more than welcome. Please share your ideas by Github issues and pull requests.

### Wishlist

- Gopher Logo 🙏
- Reviews and thoughts on changes and improvements?

# License

The project is licensed under MIT license.


## Author
*Mazin Ahmed*
* Website: [https://mazinahmed.net](https://mazinahmed.net)
* Email: mazin [at] mazinahmed [dot] net
* Twitter: [https://twitter.com/mazen160](https://twitter.com/mazen160)
* Linkedin: [http://linkedin.com/in/infosecmazinahmed](http://linkedin.com/in/infosecmazinahmed)
