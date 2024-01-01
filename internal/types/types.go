package types

type NavElement struct {
	Title string
	Value string
	Href  string
}

type NavElements struct {
	Values []NavElement
}

func (s *NavElements) Init() {
	vals := []NavElement{
		{Title: "Item 1", Value: "item1", Href: "/values/item1"},
		{Title: "Item 2", Value: "item2", Href: "/values/item1"},
		{Title: "Item 3", Value: "item3", Href: "/values/item1"},
		{Title: "Item 4", Value: "item4", Href: "/values/item1"},
	}
	s.Values = vals
}
