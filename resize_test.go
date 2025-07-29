package wildshape_test

import (
	ws "github.com/Elliott-weeks/wildshape"
	"image"
	"image/color"
	"testing"
)

func TestReSampleMethod_String(t *testing.T) {
	tests := []struct {
		method   ws.ReSampleMethod
		expected string
	}{
		{ws.ReSampleMethod(ws.NearestNeighbour), "NearestNeighbour"},
		{ws.ReSampleMethod(999), "Unknown"},
		{ws.ReSampleMethod(-1), "Unknown"},
	}

	for _, tt := range tests {
		result := tt.method.String()
		if result != tt.expected {
			t.Errorf("ReSampleMethod(%d).String() = %q; want %q", tt.method, result, tt.expected)
		}
	}
}

func TestResizeNearestNeighbor(t *testing.T) {
	// Create a 2x2 source image with distinct colors
	src := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	src.Set(0, 0, color.RGBA{255, 0, 0, 255})   // Red
	src.Set(1, 0, color.RGBA{0, 255, 0, 255})   // Green
	src.Set(0, 1, color.RGBA{0, 0, 255, 255})   // Blue
	src.Set(1, 1, color.RGBA{255, 255, 0, 255}) // Yellow

	// Resize to 4x4 using Nearest Neighbor
	dst := ws.Resize(src, 4, 4, ws.NearestNeighbour)

	// Check dimensions
	if dst.Bounds().Dx() != 4 || dst.Bounds().Dy() != 4 {
		t.Fatalf("expected output size 4x4, got %dx%d", dst.Bounds().Dx(), dst.Bounds().Dy())
	}

	// Check one known pixel color
	c := dst.At(0, 0)
	expected := color.RGBA{255, 0, 0, 255} // should match src(0, 0)

	if !colorsEqual(c, expected) {
		t.Errorf("expected color %v at (0,0), got %v", expected, c)
	}
}

func TestResizeUnknownMethod(t *testing.T) {
	src := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	dst := ws.Resize(src, 4, 4, ws.ReSampleMethod(99))

	if !dst.Bounds().Empty() {
		t.Errorf("expected empty image, got %v", dst.Bounds())
	}
}

// Helper: Compare colors ignoring type differences (e.g. RGBA vs NRGBA)
func colorsEqual(a, b color.Color) bool {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
