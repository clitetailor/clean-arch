package entrypoints

import (
	"clean-arch/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func NewCartRoute(
	router fiber.Router,
	cartService entities.ICartService,
) *CartRoute {
	return &CartRoute{
		Router:      router,
		CartService: cartService,
	}
}

type CartRoute struct {
	Router      fiber.Router
	CartService entities.ICartService
}

func (r *CartRoute) Mount() {
	r.Router.Post("/cart-items", r.AddItemToCart)
	r.Router.Get("/cart-items", r.GetCartItems)
}

func (r *CartRoute) AddItemToCart(ctx *fiber.Ctx) error {
	product := new(entities.Product)

	ctx.BodyParser(product)

	r.CartService.AddItemToCart(product)

	return nil
}

func (r *CartRoute) GetCartItems(ctx *fiber.Ctx) error {
	items := r.CartService.GetCartItems()

	ctx.JSON(items)

	return nil
}
