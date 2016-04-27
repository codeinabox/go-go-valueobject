package valueobject_test

import (
	"fmt"
	"github.com/codeinabox/go-go-valueobject"
	"testing"
)

// Create a different type of value object used in Equals() check
type NotEmailAddress struct {
	value string
}

func (n NotEmailAddress) String() string {
	return n.value
}

func (n NotEmailAddress) Equals(value valueobject.Value) bool {
	return false
}

func ExampleString_EmailAddress() {
	numeral, _ := valueobject.NewEmailAddress("joe@blogs.com")
	fmt.Println(numeral.String())
	// Output: joe@blogs.com
}

func TestShouldntAcceptInvalidEmailAddress(t *testing.T) {
	_, err := valueobject.NewEmailAddress("invalid")
	if err == nil {
		t.Fatal("We expected an error")
	}
}

func TestShouldntAcceptNonStringValue(t *testing.T) {
	_, err := valueobject.NewEmailAddress(1234)
	if err == nil {
		t.Fatal("We expected an error")
	}
}

func ExampleEquals_EmailAddress() {
	a, _ := valueobject.NewEmailAddress("joe@blogs.com")
	b, _ := valueobject.NewEmailAddress("joe@blogs.com")

	fmt.Println(a.Equals(b))
	// Output: true
}

func TestShouldCompareTwoEmailAddresssAsNotEqual(t *testing.T) {
	a, _ := valueobject.NewEmailAddress("joe@blogs.com")
	b, _ := valueobject.NewEmailAddress("mandy@blogs.com")
	if a.Equals(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}

func TestShouldNotBeEqualIfNotEmailAddress(t *testing.T) {
	var notEmailAddress NotEmailAddress
	numeral, _ := valueobject.NewEmailAddress("joe@blogs.com")

	if numeral.Equals(notEmailAddress) == true {
		t.Fatal("Different value object types can not be equal")
	}
}
