package conv_test

import (
	"os"
	"testing"

	conv "github.com/sourjp/gopherdojo-studyroom/kadai1"
)

func TestIsValidatedExt(t *testing.T) {
	tests := []struct {
		name   string
		srcExt string
		dstExt string
		expect bool
	}{
		{name: "Support jpg, jpeg", srcExt: "jpg", dstExt: "jpeg", expect: true},
		{name: "Support png, gif", srcExt: "png", dstExt: "gif", expect: true},
		{name: "Unsupport one ext", srcExt: "jpg", dstExt: "none", expect: false},
		{name: "Unsupport both ext", srcExt: "none", dstExt: "none", expect: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New("", test.srcExt, test.dstExt)
			if got := c.IsValidatedExt(); got != test.expect {
				t.Errorf("IsValidatedExt() = %t, expect %t", got, test.expect)
			}
		})
	}
}

func TestGetImagePaths(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		srcExt  string
		dstExt  string
		expect  []string
	}{
		{name: "Support jpg, jpeg", baseDir: "testdata", srcExt: "jpg", dstExt: "jpeg", expect: []string{"testdata/t1.jpg", "testdata/testdata2/t2.jpg"}},
		{name: "Failed to find images", baseDir: "testdata", srcExt: "none", dstExt: "none", expect: []string{}},
		{name: "Failed to find dir", baseDir: "none", srcExt: "jpg", dstExt: "jpeg", expect: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New(test.baseDir, test.srcExt, test.dstExt)

			got, _ := c.GetImagePaths()
			for i := range test.expect {
				if got[i] != test.expect[i] {
					t.Errorf("GetImagePaths() = %s, expect %s", got, test.expect)
				}
			}
		})
	}
}

func TestEncodeAndDecode(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		srcExt  string
		dstExt  string
		paths   []string
		expect  []string
		fail    bool // If true, A test case expect to get err.
	}{
		{name: "Convert jpg to png", baseDir: "testdata", srcExt: "jpg", dstExt: "png", paths: []string{"testdata/t1.jpg", "testdata/testdata2/t2.jpg"}, expect: []string{"testdata/t1.png", "testdata/testdata2/t2.png"}, fail: false},
		{name: "Convert png to gif", baseDir: "testdata", srcExt: "png", dstExt: "gif", paths: []string{"testdata/t1.png", "testdata/testdata2/t2.png"}, expect: []string{"testdata/t1.gif", "testdata/testdata2/t2.gif"}, fail: false},
		{name: "Convert gif to jpeg", baseDir: "testdata", srcExt: "gif", dstExt: "jpeg", paths: []string{"testdata/t1.gif", "testdata/testdata2/t2.gif"}, expect: []string{"testdata/t1.jpeg", "testdata/testdata2/t2.jpeg"}, fail: false},
		{name: "Failed to read file", baseDir: "testdata", srcExt: "png", dstExt: "jpg", paths: []string{"testdata/noimage.png"}, expect: []string{}, fail: true},
	}

	// madeFiles save file name which are created to delete after testing.
	var madeFiles []string
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New(test.baseDir, test.srcExt, test.dstExt)

			for i, path := range test.paths {
				img, err := c.Decode(path)
				if err != nil {
					if test.fail {
						continue // If it's expect to fail, Good to go next case.
					}
					t.Errorf("Decode() got err: %s", err)
				}
				err = c.Encode(path, img)
				if err != nil {
					if test.fail {
						continue // If it's expect to fail, Good to go next case.
					}
					t.Errorf("Encode() got err: %s", err)
				}

				if _, err := os.Stat(test.expect[i]); os.IsNotExist(err) {
					t.Errorf("TestEncodeAndDecode() go err: %s", err)
				} else {
					madeFiles = append(madeFiles, test.expect[i])
				}
			}
		})
	}

	for _, f := range madeFiles {
		if err := os.Remove(f); err != nil {
			t.Errorf("TestEncodeAndDecode() couldn't remove files = %s, and got err = %s", f, err)
		}
	}
}
