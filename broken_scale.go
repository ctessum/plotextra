package plotextra

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
)

type BrokenColorMap struct {
	Base, OverFlow palette.ColorMap
}

// SetHighCut sets the cut point for switching between the base
// and overflow color maps.
func (c *BrokenColorMap) SetHighCut(v float64) {
	c.Base.SetMax(v)
	c.OverFlow.SetMin(v)
}

// At returns the color associated with the given value.
// If the value is not between Max() and Min(), an error is returned.
func (c *BrokenColorMap) At(v float64) (color.Color, error) {
	if v > c.Base.Max() {
		return c.OverFlow.At(v)
	}
	return c.Base.At(v)
}

// Max returns the current maximum value of the ColorMap.
func (c *BrokenColorMap) Max() float64 {
	return c.OverFlow.Max()
}

// SetMax sets the maximum value of the ColorMap.
func (c *BrokenColorMap) SetMax(v float64) {
	c.OverFlow.SetMax(v)
}

// Min returns the current minimum value of the ColorMap.
func (c *BrokenColorMap) Min() float64 {
	return c.Base.Min()
}

// SetMin sets the minimum value of the ColorMap.
func (c *BrokenColorMap) SetMin(v float64) {
	c.Base.SetMin(v)
}

// Alpha returns the opacity value of the ColorMap.
func (c *BrokenColorMap) Alpha() float64 {
	return (c.Base.Alpha() + c.OverFlow.Alpha()) / 2
}

// SetAlpha sets the opacity value of the ColorMap. Zero is transparent
// and one is completely opaque. The default value of alpha should be
// expected to be one. The function should be expected to panic
// if alpha is not between zero and one.
func (c *BrokenColorMap) SetAlpha(v float64) {
	c.Base.SetAlpha(v)
	c.OverFlow.SetAlpha(v)
}

// Palette creates a Palette with the specified number of colors
// from the ColorMap.
func (c *BrokenColorMap) Palette(colors int) palette.Palette {
	panic("not implemented")
}

// BrokenScale can be used as the value of an Axis.Scale function to
// set the axis to a linear scale with a break.
type BrokenScale struct {
	HighCut         float64
	HighCutFraction float64
}

// Normalize returns the fractional broken distance of
// x between min and max.
func (b BrokenScale) Normalize(min, max, x float64) float64 {
	if x > b.HighCut {
		return b.HighCutFraction + (1-b.HighCutFraction)*(x-b.HighCut)/(max-b.HighCut)
	}
	return b.HighCutFraction * (x - min) / (b.HighCut - min)
}

type BrokenTicks struct {
	HighCut float64
}

func (b BrokenTicks) Ticks(min, max float64) []plot.Tick {
	ticks := plot.DefaultTicks{}.Ticks(min, b.HighCut)
	return append(ticks, plot.Tick{
		Value: max,
		Label: fmt.Sprintf("%.0f", max),
	})
}
