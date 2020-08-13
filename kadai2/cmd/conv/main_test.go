package main_test

import (
	"os"
	"testing"

	main "github.com/sourjp/gopherdojo-studyroom/kadai2/cmd/conv"
)

func TestConverter(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		srcExt  string
		dstExt  string
		expect  []string
		fail    bool // If true, A test case expect to get err.
	}{
		{name: "Convert jpg to png", baseDir: "../../testdata/success", srcExt: "jpg", dstExt: "png", expect: []string{"../../testdata/success/t1.png", "../../testdata/success/img/t2.png"}, fail: false},
		{name: "Convert png to gif", baseDir: "../../testdata/success", srcExt: "png", dstExt: "gif", expect: []string{"../../testdata/success/t1.gif", "../../testdata/success/img/t2.gif"}, fail: false},
		{name: "Convert gif to jpeg", baseDir: "../../testdata/success", srcExt: "gif", dstExt: "jpeg", expect: []string{"../../testdata/success/t1.jpeg", "../../testdata/success/img/t2.jpeg"}, fail: false},
		{name: "Failed to read file", baseDir: "../../testdata/failure", srcExt: "jpg", dstExt: "png", expect: []string{}, fail: true},
	}

	// madeFiles save file name which are created to delete after testing.
	var madeFiles []string
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := main.Converter(test.baseDir, test.srcExt, test.dstExt); err != nil {
				if !test.fail {
					t.Errorf("Converter() got err: %s", err)
				}
			}

			for i := range test.expect {
				if _, err := os.Stat(test.expect[i]); os.IsNotExist(err) {
					t.Errorf("TestConverter() go err: %s", err)
				} else {
					madeFiles = append(madeFiles, test.expect[i])
				}
			}
		})
	}

	for _, f := range madeFiles {
		if err := os.Remove(f); err != nil {
			t.Errorf("TestConverter() couldn't remove file = %s, and got err = %s", f, err)
		}
	}
}
