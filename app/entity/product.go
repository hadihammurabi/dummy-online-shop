package entity

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint

	Categories []*Category
}
