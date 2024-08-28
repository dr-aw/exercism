package resistorcolor

type color struct {
	text string
	num  int
}

var clrs = []color{
	{"black", 0},
	{"brown", 1},
	{"red", 2},
	{"orange", 3},
	{"yellow", 4},
	{"green", 5},
	{"blue", 6},
	{"violet", 7},
	{"grey", 8},
	{"white", 9},
}

// Colors returns the list of all colors.
func Colors() []string {
	var list []string
	for _, c := range clrs {
		list = append(list, c.text) // only text
	}
	return list
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for _, c := range clrs {
		if c.text == color {
			return c.num
		}
	}
	return 0
}
