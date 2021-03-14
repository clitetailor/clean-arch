package entities

type Cart struct {
	CartItems CartItemSlice
}

// +gen * slice:"First,All,SortBy"
type CartItem struct {
	Product *Product `json:"product"`
	Count   int      `json:"count"`
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ICartService interface {
	AddItemToCart(*Product)
	GetCartItems() []*CartItem
}
