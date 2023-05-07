package main

type person struct {
	name string
}

func main() {
	m := make(map[int]person)
	m[0] = person{
		name: "TDN",
	}
	// m[0].name = "ABC" // viet tat
	a := m[0]
	a.name = "ABC"
}
