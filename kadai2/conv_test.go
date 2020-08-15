package conv_test

import (
	"os"
	"testing"

	conv "github.com/sourjp/gopherdojo-studyroom/kadai2"
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
				t.Fatalf("IsValidatedExt() = %t, expect %t", got, test.expect)
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
		{name: "Support jpg, jpeg", baseDir: "testdata/success", srcExt: "jpg", dstExt: "jpeg", expect: []string{"testdata/success/t1.jpg", "testdata/success/img/t2.jpg"}},
		{name: "Failed to find images", baseDir: "testdata/success", srcExt: "none", dstExt: "none", expect: []string{}},
		{name: "Failed to find dir", baseDir: "none", srcExt: "jpg", dstExt: "jpeg", expect: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New(test.baseDir, test.srcExt, test.dstExt)

			got, _ := c.GetImagePaths()
			for _, gv := range got {
				isContain := func() bool {
					for _, ev := range test.expect {
						if ev == gv {
							return true
						}
					}
					return false
				}()
				if !isContain {
					t.Fatalf("GetImagePaths() = %s, expect %s", got, test.expect)
				}
			}
		})
	}
}

func TestDecodeAndEncode(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		srcExt  string
		dstExt  string
		paths   []string
		expect  []string
	}{
		{name: "Convert jpg to png", baseDir: "testdata", srcExt: "jpg", dstExt: "png", paths: []string{"testdata/success/t1.jpg", "testdata/success/img/t2.jpg"}, expect: []string{"testdata/success/t1.png", "testdata/success/img/t2.png"}},
		{name: "Convert png to gif", baseDir: "testdata", srcExt: "png", dstExt: "gif", paths: []string{"testdata/success/t1.png", "testdata/success/img/t2.png"}, expect: []string{"testdata/success/t1.gif", "testdata/success/img/t2.gif"}},
		{name: "Convert gif to jpeg", baseDir: "testdata", srcExt: "gif", dstExt: "jpeg", paths: []string{"testdata/success/t1.gif", "testdata/success/img/t2.gif"}, expect: []string{"testdata/success/t1.jpeg", "testdata/success/img/t2.jpeg"}},
	}

	// madeFiles save file name which are created to delete after testing.
	var madeFiles []string
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New(test.baseDir, test.srcExt, test.dstExt)

			for i := range test.paths {
				err := testDecodeAndEncode(t, c, test.paths[i], test.expect[i])
				if err != nil {
					t.Fatalf("%s", err)
				}
				madeFiles = append(madeFiles, test.expect[i])
			}
		})
	}

	for _, f := range madeFiles {
		if err := os.Remove(f); err != nil {
			t.Errorf("TestEncodeAndDecode() couldn't remove files = %s, and got err = %s", f, err)
		}
	}
}

// testDecodeAndEncode may be uselss, but I added to test t.Helper().
func testDecodeAndEncode(t *testing.T, c *conv.Converter, path, expect string) error {
	t.Helper()

	img, err := c.Decode(path)
	if err != nil {
		t.Fatalf("Decode() got err %s, expect %s", err, expect)
	}
	if err := c.Encode(path, img); err != nil {
		t.Fatalf("Encode() got err %s, expect %s", err, expect)
	}
	if _, err := os.Stat(expect); os.IsNotExist(err) {
		t.Fatalf("IsNotExist() got err %s, expect %s", err, expect)
	}
	return nil
}

func TestDecodeAndEncodeFailuer(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		srcExt  string
		dstExt  string
		paths   []string
		expect  []string
	}{
		{name: "Failed to read file", baseDir: "testdata", srcExt: "jpg", dstExt: "png", paths: []string{"testdata/failure/unimg.jpg"}, expect: []string{}},
		{name: "Failed to read file", baseDir: "testdata", srcExt: "bmp", dstExt: "png", paths: []string{"testdata/failure/unimg.bmp"}, expect: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := conv.New(test.baseDir, test.srcExt, test.dstExt)

			for _, path := range test.paths {
				_, err := c.Decode(path)
				if err == nil {
					t.Errorf("Decode() should get err, but it passed: case = %s", path)
				}
			}
		})
	}
}
