package main

func main() {
	func Items() iter.Seq[Item] {
		return func(yield func(Item) bool) {
		items := []Item{1, 2, 3}
		for _, v := range items {
		if !yield(v) {
		return
	}
	}
	}
	}
}