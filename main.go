/*
Copyright © 2023 anti-duhring
*/
package main

import (
	"github.com/anti-duhring/fit/cmd"
	"github.com/anti-duhring/fit/internal/config"
)

func main() {
	cmd.Execute(
		config.WithInitFolder(".fit"),
	)
}
