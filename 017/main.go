package main

func main() {
	var (
		minecraft      = game{title: "minecraft", price: 20}
		tetris         = game{title: "tetris", price: 10}
		mobydick       = book{title: "moby dick", price: 10}
		thegreatgatsby = book{title: "the great gatsby", price: 30, publishing: 118281600}
		rubik          = toy{title: "rubik", price: 17}
		yoda           = toy{title: "baby yoda", price: 0}
	)

	var store list
	store = append(store, &minecraft, &tetris, &mobydick, thegreatgatsby, rubik, &yoda)
	store.discount(0.5)
	store.print()

}
