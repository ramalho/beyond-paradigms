package main

type PositiveIterator struct {
	target       []float64
	index        int
	currentValue float64
}

func NewPositiveIterator(s []float64) *PositiveIterator {
	return &PositiveIterator{target: s}
}

func (p *PositiveIterator) Advance() bool {
	for p.index < len(p.target) {
		if v := p.target[p.index]; v > 0 {
			p.currentValue = v
			return true
		}
		p.index++
	}
	return false
}

func (p *PositiveIterator) Next() float64 {
	p.index++
	return p.currentValue
}
