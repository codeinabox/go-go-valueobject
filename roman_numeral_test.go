package valueobject_test

import (
	"fmt"
	"github.com/codeinabox/go-go-valueobject"
	"testing"
)

var integerToRomanNumeralTests = []struct {
	integer int
	numeral string
}{
	{1, "I"},
	{2, "II"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{10, "X"},
	{50, "L"},
	{100, "C"},
	{257, "CCLVII"},
	{500, "D"},
	{1000, "M"},
	{2000, "MM"},
	{2257, "MMCCLVII"},
}

// Create a different type of value object used in Equals() check
type NotRomanNumeral struct {
	value string
}

func (n NotRomanNumeral) String() string {
	return n.value
}

func (n NotRomanNumeral) Equals(value valueobject.Value) bool {
	return false
}

func ExampleString_RomanNumeral() {
	numeral, _ := valueobject.NewRomanNumeral(5)
	fmt.Println(numeral.String())
	// Output: V
}

func TestConvertIntegerToRomanNumeral(t *testing.T) {
	for _, example := range integerToRomanNumeralTests {
		n, err := valueobject.NewRomanNumeral(example.integer)
		if err != nil {
			t.Fatal(err)
		}
		if n.String() != example.numeral {
			t.Fatalf("string representation should be %s, was %s", example.numeral, n.String())
		}
	}
}

func TestShouldntAcceptInvalidString(t *testing.T) {
	_, err := valueobject.NewRomanNumeral("B")
	if err == nil {
		t.Fatal("We expected an error with A")
	}
}

func ExampleEquals_RomanNumeral() {
	a, _ := valueobject.NewRomanNumeral(5)
	b, _ := valueobject.NewRomanNumeral("V")

	fmt.Println(a.Equals(b))
	// Output: true
}

func TestShouldNotBeEqualIfNotRomanNumeral(t *testing.T) {
	var notRomanNumeral NotRomanNumeral
	numeral, _ := valueobject.NewRomanNumeral("I")

	if numeral.Equals(notRomanNumeral) == true {
		t.Fatal("Different value object types can not be equal")
	}
}

func TestShouldBeEqualIfSameRomanNumeral(t *testing.T) {
	a, _ := valueobject.NewRomanNumeral("I")
	b, _ := valueobject.NewRomanNumeral("I")
	if a.Equals(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldBeEqualIfIntegerEquivalent(t *testing.T) {
	a, _ := valueobject.NewRomanNumeral(5)
	b, _ := valueobject.NewRomanNumeral("V")
	if a.Equals(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldCompareTwoRomanNumeralsAsNotEqual(t *testing.T) {
	a, _ := valueobject.NewRomanNumeral("I")
	b, _ := valueobject.NewRomanNumeral("X")
	if a.Equals(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}
