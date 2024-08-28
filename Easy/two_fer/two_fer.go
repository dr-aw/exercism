/*
Your task is to determine what you will say as you give away the extra cookie.
If you know the person's name (e.g. if they're named Do-yun), then you will say:

	'One for Do-yun, one for me.'

If you don't know the person's name, you will say you instead.

	'One for you, one for me.'
*/
package main

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func main() {
	fmt.Println(ShareWith(""))
}
