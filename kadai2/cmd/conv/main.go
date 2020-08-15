package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	conv "github.com/sourjp/gopherdojo-studyroom/kadai2"
)

var (
	baseDir string
	srcExt  string = "jpg"
	dstExt  string = "png"
)

func init() {
	flag.StringVar(&srcExt, "se", srcExt, "source exetention to decode(jpg, jpeg, png, gif)")
	flag.StringVar(&dstExt, "de", dstExt, "destinate extension to encode(jpg, jpeg, png, gif)")
}

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Image convert has finished!")
}

// Run handle the flags and return the result of Convert.
func Run() error {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return errors.New("please specifiy a directory, conv find image files recursively")
	}
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return fmt.Errorf("failed to found directory: dir=%s", args[0])
	}
	return Converter(args[0], srcExt, dstExt)
}

// Converter handle the image convertng.
func Converter(d, se, de string) error {
	c := conv.New(d, se, de)
	if ok := c.IsValidatedExt(); !ok {
		return fmt.Errorf("failed to read specified extension: srcExt=%s, dstExt=%s", srcExt, dstExt)
	}

	paths, err := c.GetImagePaths()
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		return fmt.Errorf("failed to find image files:  dir=%s, ext=.%s", srcExt, baseDir)
	}

	for _, path := range paths {
		img, err := c.Decode(path)
		if err != nil {
			return err
		}

		if err = c.Encode(path, img); err != nil {
			return err
		}
	}
	return nil
}
