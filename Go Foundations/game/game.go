package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Item is an item in the game
type Item struct {
	X int
	Y int
}

func main() {
	var item1 Item

	fmt.Println(item1)
	fmt.Printf("item1: %#v\n", item1)

	item2 := Item{1, 2}
	fmt.Printf("item2: %#v\n", item2)

	item3 := Item{
		Y: 10,
		X: 20,
	}
	fmt.Printf("item3: %#v\n", item3)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	item3.Move(100, 200)
	// item3 is value not pointer but is valid for move function
	// because it is take care by go compiler, go compiler is smart enough.

	fmt.Printf("item3 (move): %#v\n", item3)

	p1 := Player{
		Name: "Max",
		Item: Item{500, 300},
	}

	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&item1,
		&p1,
		&item2,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(17)) // calling spring method

	// time.Time import json.Marshaler interface
	//json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1)
	fmt.Println(p1.Key)

	p1.FoundKey(Crystal)
	fmt.Println(p1)
	fmt.Println(p1.Key)

	p1.FoundKey(5) // greater then Invalidkey in Key
	fmt.Println(p1)
	fmt.Println(p1.Key)

}

/*
type T struct {
	X int
}
*/

// i is called "the receiver"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item {
// func NewItem(x, y int) *Item {
// func NewItem(x, y int) (Item, error) {
const (
	maxX = 1000
	maxY = 600
)

// zero vs missing value
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// The Go compiler does "escape analysis" and will allocation i on the heap
	return &i, nil
}

type Player struct {
	Name string
	Item // Embed Item
	Key  []Key
	//T
}

//	type T struct {
//		X int
//	}
type mover interface {
	Move(x, y int)
	//Move(int, int)
}

// In most of the cases
// Rule of thumbs: Accepts interfaces as a return types
func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Go's version of enum
// iota in const always incremented by 1,
// so, Jade is 1, Copper is 2, Crystal is 3
const (
	Jade Key = iota + 1
	Copper
	Crystal
	Invalidkey // has value of 4
)

type Key byte

// Implement fmt.Stringer interface
// before this func jade is 1, copper is 2, crystal is 3
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

// Exercise

// FoundKey is method
func (p *Player) FoundKey(k Key) error {
	if k > Invalidkey || k < Jade {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !slices.Contains(p.Key, k) {

		p.Key = append(p.Key, k)
	}

	return nil

}

/* require go >= 1.18
func NewNumber[T int | float64](kind string) T {
	if kind == "int" {
		return 0
	}
	return 0.0
}
*/
