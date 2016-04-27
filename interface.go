package valueobject

type Value interface {
	String() string
	Equals(value Value) bool
}
