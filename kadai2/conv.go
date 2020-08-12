package conv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Define supported extentions
const (
	JPG  string = "jpg"
	JPEG string = "jpeg"
	PNG  string = "png"
	GIF  string = "gif"
)

// Converter defines structure of converting informations.
type Converter struct {
	BaseDir string
	SrcExt  string
	DstExt  string
}

// New construct and return Converter.
func New(d, se, de string) *Converter {
	return &Converter{BaseDir: d, SrcExt: se, DstExt: de}
}

// IsValidatedExt validate extensions.
func (c *Converter) IsValidatedExt() bool {
	var supportExt []string = []string{JPG, JPEG, PNG, GIF}
	var src, dst bool
	for _, e := range supportExt {
		if c.SrcExt == e {
			src = true
		}
		if c.DstExt == e {
			dst = true
		}
	}
	return src && dst
}

// GetImagePaths search specified extensionis image files recusively and returns them.
func (c *Converter) GetImagePaths() ([]string, error) {
	var paths []string
	err := filepath.Walk(c.BaseDir, func(path string, info os.FileInfo, err error) error {
		ext := "." + c.SrcExt
		if filepath.Ext(path) == ext {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("failed to find image files:  dir=%s, ext=.%s", c.SrcExt, c.BaseDir)
	}
	return paths, nil
}

// Decode returns ImageData which compose file path and decoded image data.
func (c *Converter) Decode(p string) (img image.Image, err error) {
	f, err := os.Open(filepath.Clean(p))
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: file=%s, err=%v", p, err)
	}
	defer func() {
		if rerr := f.Close(); rerr != nil {
			err = fmt.Errorf("failed to close file: %v, the original error: %v", rerr, err)
		}
	}()

	switch c.SrcExt {
	case JPG, JPEG:
		img, err = jpeg.Decode(f)
	case PNG:
		img, err = png.Decode(f)
	case GIF:
		img, err = gif.Decode(f)
	default:
		return nil, fmt.Errorf("an extension doesn't support to decode: ext=%s", c.SrcExt)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode iamge file: file=%s, err=%v", p, err)
	}

	return img, nil
}

// Encode encodes image file.
func (c *Converter) Encode(p string, img image.Image) (err error) {
	fname := strings.TrimSuffix(p, filepath.Ext(p)) + "." + c.DstExt
	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("failed to create image file: file=%s, err=%v", p, err)
	}
	defer func() {
		if rerr := f.Close(); rerr != nil {
			err = fmt.Errorf("failed to close file: %v, the original error: %v", rerr, err)
		}
	}()

	switch c.DstExt {
	case JPG, JPEG:
		err = jpeg.Encode(f, img, nil)
	case PNG:
		err = png.Encode(f, img)
	case GIF:
		err = gif.Encode(f, img, nil)
	default:
		return fmt.Errorf("an extension doesn't support to encode: ext=%s", c.DstExt)
	}
	if err != nil {
		return fmt.Errorf("failed to encode iamge file: file=%s, err=%v", p, err)
	}

	return nil
}
