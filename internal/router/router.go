package router

import (
	"hash/fnv"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(
	key string,
	nodes int,
) int {

	h := fnv.New32a()

	h.Write([]byte(key))

	return int(
		h.Sum32(),
	) % nodes
}
