package main

const ArraySize = 100

type Hash struct {
	table [ArraySize]*LinkedList
}

// Insert
func (h *Hash) Insert(word string) {
	index := hash(word)
	h.table[index].Insert(1, word)
}

// Search
func (h *Hash) Search(word string) bool {
	index := hash(word)
	return h.table[index].Search(word)
}

// Delete
func (h *Hash) Delete(word string) {
	index := hash(word)
	h.table[index].DeleteWithValue(word)
}

// hash function
func hash(word string) int {
	sum := 0
	for _, w := range word {
		sum += int(w)
	}

	return sum % ArraySize
}

// Init
func Init() *Hash {
	htable := &Hash{}
	for i := range htable.table {
		htable.table[i] = &LinkedList{}
	}
	return htable
}

// func main() {
// 	htable := Init()
// 	fmt.Println(htable)

// 	list := []string{
// 		"ERIC",
// 		"MARTIN",
// 		"BORJA",
// 		"ADRIAN",
// 		"DANIEL",
// 		"LUIS",
// 		"URKO",
// 		"CRISTIAN",
// 	}

// 	for _, s := range list {
// 		htable.Insert(s)
// 	}

// 	fmt.Println(htable)

// 	fmt.Println(htable.Search("MARTIN"))
// 	htable.Delete("LUIS")
// 	fmt.Println(htable.Search("MARTIN"))
// }
