# Go Go Value Object!

[![Build Status](https://travis-ci.org/codeinabox/go-go-valueobject.svg?branch=master)](https://travis-ci.org/codeinabox/go-go-valueobject)

Inspired by my interest in domain driven design and [Nicolò Pignatell's value object library for PHP](https://github.com/nicolopignatelli/valueobjects/) I wanted 
to see if value objects could be implemented in Golang. Currently this is very much WIP with only two value objects: 1) a roman numeral 2) email address


### Installation

`go get github.com/codeinabox/go-go-valueobject`

### Usage

```go
	import (
		"fmt"
		"github.com/codeinabox/go-go-valueobject"
	)

	func main() {
		email, err := valueobject.NewEmailAddress("joe@blogs.com")
		fmt.Println(email.String())

		differentEmail, err := valueobject.NewEmailAddress("mandy@blogs.com")
		if email.Equals(differentEmail) == false {
			fmt.Println("As expected the two emails are different")
		}
	}
```

### Implementing your own value object
Implementing your own value object is easy, it just needs to adhere to the `Value` interface eg

```go
	import (
		"github.com/codeinabox/go-go-valueobject"
	)

	type Identifier struct {
		value string
	}

	func NewIdentifier(value string) (Identifier, error) {
		var n Identifier
		n.value = value
		return n, nil
	}

	func (n Identifier) String() string {
		return n.value
	}

	func (n Identifier) Equals(value valueobject.Value) bool {
		otherIdentifier, ok := value.(Identifier)
		return ok && n.value == otherIdentifier.value
	}
```
