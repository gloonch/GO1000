package main

func main() {
	var (
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 10}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)

	my := list(items)
	//my = nil
	my.print()
}
