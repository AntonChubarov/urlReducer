package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x41\x0a\xc2\x40\x0c\x46\xe1\x75\x06\x72\x87\x2c\x15\xbc\x81\x87\x29\x71\x08\x4c\x68\x4c\x4b\xfa\x8f\xd8\xdb\x4b\x77\x6e\x1f\x8f\xaf\x97\x29\x4c\xa0\xaf\x30\x99\x15\xdc\x6e\xdc\x68\xe8\x31\xe4\xa3\xd5\x87\x96\xe4\x06\xc9\x19\xc1\x8d\xa8\x6f\x79\xa0\xd4\x13\xd7\xbc\xec\xeb\x15\x69\x2f\x7f\x6b\x9d\xb2\xda\xf9\xe0\x46\x9e\x0e\xd7\x58\x66\x85\xc0\xbe\xf8\x13\xee\xcf\x5f\x00\x00\x00\xff\xff\x62\x9d\x07\x8e\x71\x00\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var _generate_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xb1\x0d\xc5\x20\x0c\x04\xd0\x9e\x29\x6e\x01\xa0\xff\xdb\xdc\x4f\xac\x13\x42\xb1\x11\xf1\xfe\x4a\x93\x22\xf5\x7b\x8b\xc7\xa4\x0c\xd7\xd0\x66\x8e\xf0\xbb\x94\xde\x15\x3f\x99\xdb\x66\x1a\x14\xf5\x3f\xfc\x64\x12\x75\x4d\x7d\x26\x6a\xe0\xa5\xa6\x00\x5a\x79\x02\x00\x00\xff\xff\xe6\x24\xd8\x86\x4e\x00\x00\x00")

func generate_go() ([]byte, error) {
	return bindata_read(
		_generate_go,
		"generate.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"000001_init.up.sql": _000001_init_up_sql,
	"generate.go": generate_go,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.up.sql": &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{
	}},
	"generate.go": &_bintree_t{generate_go, map[string]*_bintree_t{
	}},
}}
