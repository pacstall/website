package types

type Equaller interface {
	Equals(other Equaller) bool
}
