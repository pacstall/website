package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^/api/packages[/#]?(\?[a-zA-Z0-9-=&]*)?$`)
	fmt.Printf("%#v\n", re.FindStringSubmatch("/api/packages?asd=2as&asd=2as"))
}
