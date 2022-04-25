/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package main

import (
	_ "embed"
	"fmt"

	"github.com/imhshekhar47/ops-admin/cmd"
)

//go:embed LICENSE
var license string

func main() {
	fmt.Printf("\n%s\n\n", license)

	cmd.Execute()
}
