package main

func main() {
	var (
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 10}
		mobydick  = book{title: "tetris", price: 10}
		rubik     = book{title: "rubik", price: 17}
		yoda      = toy{title: "baby yoda", price: 0}
	)

	var store list
	store = append(store, &minecraft, &tetris, &mobydick, rubik, &yoda)
	store.discount(0.5)
	store.print()

}
