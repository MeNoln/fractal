package palette

func Recolor(r, g, b int, p *Gradient) (uint8, uint8, uint8) {
	if p == nil {
		return uint8(r), uint8(g), uint8(b)
	}

	if r <= p.r_init {
		return uint8(p.r_init), uint8(p.g_init), uint8(p.b_init)
	}

	r_init := p.r_init
	for i := 1; i <= p.max_iter; i++ {
		temp := r_init + p.r_step
		if i == p.max_iter {
			temp = 255
		}

		if in(r, r_init, temp) {
			r = r_init
			g = p.g_init + (p.g_step * (i - 1))
			b = p.b_init + (p.b_step * (i - 1))
			break
		}

		r_init += p.r_step
	}

	return uint8(r), uint8(g), uint8(b)
}

func in(v, f, t int) bool {
	return v > f && v <= t
}
