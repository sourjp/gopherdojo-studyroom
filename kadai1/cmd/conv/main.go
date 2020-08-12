package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	conv "github.com/sourjp/gopherdojo-studyroom/kadai1"
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
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Image convert was suceeded!")
}

func run() error {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return errors.New("please specifiy a directory, conv find image files recursively")
	}
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return fmt.Errorf("failed to found directory: dir=%s", args[0])
	}

	c := conv.New(args[0], srcExt, dstExt)
	if ok := c.IsValidatedExt(); !ok {
		return fmt.Errorf("failed to read specified extention: srcExt=%s, dstExt=%s", srcExt, dstExt)
	}

	paths, err := c.GetImagePaths()
	if err != nil {
		return err
	}

	for _, path := range paths {
		img, err := c.Decode(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		err = c.Encode(path, img)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
	return nil
}
