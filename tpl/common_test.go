package tpl

import "testing"

func TestCheckMD(t *testing.T) {
	tt := []struct {
		name  string
		filename  string
		rst     bool
	}{
		{"md file", "aa.md", true},
		{"jpg file", "aa.jpg", false},
		{"without ext", "aa", false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rst := check_md_file(tc.filename)
			if rst != tc.rst {
				t.Fatalf("%v should be %v, but get %v", tc.filename, tc.rst, rst)
			}
		})
	}

}


func TestGetSourcePath(t *testing.T) {
	filename := "aa.md"
	source_path := "content/aa.md"
	path := get_source_path(filename)
	print(path)
	if path != source_path {
		t.Fatalf("%v should be %v, but get %v", filename, source_path, path)
	}
}


func TestGetDistPath(t *testing.T) {
	filename := "aa.md"
	dist_path := "dist/aa.html"
	path := get_dist_path(filename)
	print(path)
	if path != dist_path {
		t.Fatalf("%v should be %v, but get %v", filename, dist_path, path)
	}
}




