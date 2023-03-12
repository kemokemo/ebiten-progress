package progress

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Progress struct {
	img     *ebiten.Image
	op      *ebiten.DrawImageOptions
	prevVal int
	length  float64
	delta   float64
}

func NewProgress(img *ebiten.Image, x, y, length float64) *Progress {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	return &Progress{img: img, op: op, length: length}
}

func (p *Progress) SetPercent(val int) {
	if val < 0 || 100 < val || val == p.prevVal {
		return
	}

	p.delta = p.length * float64(val-p.prevVal) / 100.0
	p.op.GeoM.Translate(p.delta, 0.0)

	p.prevVal = val
}

func (p *Progress) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.img, p.op)
}
