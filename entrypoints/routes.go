package entrypoints

func NewRoutes(cart *CartRoute) *Routes {
	return &Routes{
		CartRoute: cart,
	}
}

type Routes struct {
	CartRoute *CartRoute
}

func (r *Routes) Mount() {
	r.CartRoute.Mount()
}
