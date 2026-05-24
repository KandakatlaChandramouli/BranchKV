package router_moe

type Router struct {
	total int
}

func NewRouter(
	total int,
) *Router {

	return &Router{
		total: total,
	}
}

func (r *Router) Route(
	token uint64,
) int {

	return int(token) % r.total
}
