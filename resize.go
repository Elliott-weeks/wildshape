package wildshape

import "image"

// ReSampleMethod represents the method used for resampling during image resizing.
type ReSampleMethod int

const (
	// NearestNeighbour is a resampling method that picks the nearest pixel without interpolation.
	NearestNeighbour ReSampleMethod = iota
)

func (m ReSampleMethod) String() string {
	switch m {
	case NearestNeighbour:
		return "NearestNeighbour"
	default:
		return "Unknown"
	}
}

// Resize returns a resized copy of the input image with the specified width and height,
// using the given resampling method.
func Resize(img image.Image, newWidth, newHeight int, method ReSampleMethod) image.Image {
	switch method {
	case NearestNeighbour:
		return resizeNearestNeighbor(img, newWidth, newHeight)
	default:
		panic("unsupported resampling method")
	}
}

func resizeNearestNeighbor(img image.Image, newWidth, newHeight int) image.Image {
	srcBounds := img.Bounds()
	srcWidth, srcHeight := srcBounds.Dx(), srcBounds.Dy()

	// New image
	dst := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))

	// These are the scaling factors
	scalex := float64(srcWidth) / float64(newWidth)
	scaley := float64(srcHeight) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := int(float64(x) * scalex)
			srcY := int(float64(y) * scaley)

			// Limit the bounds
			if srcX >= srcWidth {
				srcX = srcWidth - 1
			}
			if srcY >= srcHeight {
				srcY = srcHeight - 1
			}

			c := img.At(srcBounds.Min.X+srcX, srcBounds.Min.Y+srcY)
			dst.Set(x, y, c)
		}
	}
	return dst
}
