package main

func main() {
	var (
		minecraft      = game{product{title: "minecraft", price: 20}}
		tetris         = game{product{title: "tetris", price: 10}}
		mobydick       = book{product{title: "moby dick", price: 10}, nil}
		thegreatgatsby = book{product{title: "the great gatsby", price: 30}, 118281600}
		rubik          = toy{product{title: "rubik", price: 17}}
		yoda           = toy{product{title: "baby yoda", price: 0}}
	)

	var store list
	store = append(store, &minecraft, &tetris, &mobydick, &thegreatgatsby, &rubik, &yoda)
	store.discount(0.5)
	store.print()

}
