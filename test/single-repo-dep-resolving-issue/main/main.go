package main

import (
	"fmt"

	"github.com/radisvaliullin/test-golang/test/single-repo-dep-resolving-issue/pkg1"
	"github.com/radisvaliullin/test-golang/test/single-repo-dep-resolving-issue/pkg2"
)

func main() {

	fmt.Println(pkg1.Name())

	fmt.Println(pkg2.Name())
}
