package main

func main() {
	l := list{
		{title: "minecraft", price: 20},
		{title: "tetris", price: 10},
		{title: "moby dick", price: 10, released: toTimestamp(nil)},
		{title: "the great gatsby", price: 30, released: toTimestamp(118281600)},
		{title: "rubik", price: 17},
		{title: "baby yoda", price: 0},
	}

	//l.discount(0.5)
	l.print()

}
