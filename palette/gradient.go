package palette

// https://colordesigner.io/gradient-generator

var Gradients map[string]*Gradient = map[string]*Gradient{
	"default": nil,
	"purple_to_blue": {
		r_init: 36,
		r_step: 12,

		g_init: 40,
		g_step: -1,

		b_init: 135,
		b_step: 4,

		max_iter: 17,
	},
	"heat_to_turquoise": {
		r_init: 36,
		r_step: 10,

		g_init: 135,
		g_step: -4,

		b_init: 121,
		b_step: -5,

		max_iter: 17,
	},
	"darkblue_to_white": {
		r_init: 80,
		r_step: 10,

		g_init: 85,
		g_step: 10,

		b_init: 150,
		b_step: 6,

		max_iter: 17,
	},
}

type Gradient struct {
	r_init int
	r_step int

	g_init int
	g_step int

	b_init int
	b_step int

	max_iter int
}
