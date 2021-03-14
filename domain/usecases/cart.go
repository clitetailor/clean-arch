package usecases

import (
	"clean-arch/domain/entities"
	"strings"
)

func NewCartService() entities.ICartService {
	return &CartService{
		Cart: &entities.Cart{
			CartItems: make([]*entities.CartItem, 0),
		},
	}
}

type CartService struct {
	Cart *entities.Cart
}

func (s *CartService) AddItemToCart(product *entities.Product) {
	item, _ := s.Cart.CartItems.First(func(ci *entities.CartItem) bool {
		return ci.Product.Name == product.Name
	})

	if item != nil {
		item.Count = item.Count + 1

		return
	}

	newCartItem := &entities.CartItem{
		Product: product,
		Count:   1,
	}

	s.Cart.CartItems = append(s.Cart.CartItems, newCartItem)
}

func (s *CartService) GetCartItems() []*entities.CartItem {
	return s.Cart.CartItems.SortBy(func(ci1, ci2 *entities.CartItem) bool {
		return strings.Compare(ci1.Product.Name, ci2.Product.Name) < 0
	})
}
