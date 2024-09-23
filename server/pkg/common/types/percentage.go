package types

import "fmt"

type Percent float64

func (p Percent) String() string {
	return fmt.Sprintf("%v%%", float64(int(p*1000))/10)
}
