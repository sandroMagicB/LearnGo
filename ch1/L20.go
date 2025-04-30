package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

  msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent %s", name, openRate, "\n")

  // don't edit below this line

	fmt.Print(msg)
}
