package main

import (
	"branchkv-core/internal/activation"
	"branchkv-core/internal/embedding"
	"branchkv-core/internal/layernorm"
	"branchkv-core/internal/mixtureofexperts"
	"branchkv-core/internal/rope"
	"branchkv-core/internal/router_moe"
	"branchkv-core/internal/softmax"
	"branchkv-core/internal/transformer"
	"fmt"
)

func main() {

	tf := transformer.NewTransformer()

	tf.Forward()

	embed := embedding.NewEmbedding(
		32000,
		128,
	)

	vec := embed.Encode(42)

	ropeEngine := rope.NewRoPE()

	rotated := ropeEngine.Rotate(
		10,
		2,
	)

	soft := softmax.Compute(
		[]float64{
			1, 2, 3,
		},
	)

	norm := layernorm.Normalize(
		[]float32{
			1, 2, 3,
		},
	)

	gelu := activation.GELU(3.14)

	moe := mixtureofexperts.NewMoE()

	moe.AddExpert(1)

	router := router_moe.NewRouter(4)

	target := router.Route(42)

	fmt.Println(
		"layers:",
		tf.Layers(),
	)

	fmt.Println(
		"embedding:",
		len(vec),
	)

	fmt.Println(
		"rope:",
		rotated,
	)

	fmt.Println(
		"softmax:",
		len(soft),
	)

	fmt.Println(
		"norm:",
		len(norm),
	)

	fmt.Println(
		"gelu:",
		gelu,
	)

	fmt.Println(
		"moe:",
		moe.Size(),
	)

	fmt.Println(
		"expert:",
		target,
	)
}
