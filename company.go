package tickr

// Company is a Trie
// It'll hold company name with ticker as value
type Company struct {
	ticker   string
	children map[rune]*Company
}

func NewCompany() *Company {
	return new(Company)
}

func (c *Company) Get(key string) string {
	node := c
	var ticker string
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return ticker
		}
		ticker = node.ticker
	}
	return ticker
}

func (c *Company) Put(key string, value string) bool {
	node := c
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			if node.children == nil {
				node.children = map[rune]*Company{}
			}
			child = NewCompany()
			node.children[r] = child
		}
		node = child
	}
	isNewVal := node.ticker == ""
	node.ticker = value
	return isNewVal
}
