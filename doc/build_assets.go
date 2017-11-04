// Code generated by go-bindata.
// sources:
// schemas/v1/errors.json
// DO NOT EDIT!

package doc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemasV1ErrorsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x3d\x6f\xdb\x30\x10\xdd\xfd\x2b\x08\x26\x53\xe1\x44\xcd\xd2\xc1\x28\x02\x18\xf0\xd2\xa9\x19\x82\x2c\x81\x5b\xd0\xd2\xc9\x66\x4a\x91\x0a\x79\x32\x60\x04\xfe\xef\x05\x29\xea\xdb\x92\x5c\xd9\xe9\x66\xdc\xbd\xbb\xe3\x7b\xef\x44\xf3\x63\x46\x08\xbd\x35\xe1\x0e\x12\x46\x17\x84\xee\x10\xd3\x45\x10\xbc\x19\x25\xef\xf2\xe8\xbd\xd2\xdb\x20\xd2\x2c\xc6\xbb\xaf\xdf\x82\x3c\x76\x43\xe7\xb6\x2e\x02\x13\x6a\x9e\x22\x57\xd2\xd6\x2e\x89\xe0\x06\x89\x8a\x09\x68\xad\xb4\xc9\x41\x78\x48\xc1\x66\xd5\xe6\x0d\x42\xcc\x63\x1a\xde\x33\xae\x21\xa2\x0b\xf2\x4a\xf7\xa0\x8d\xed\x30\x27\x34\x52\x09\xe3\xee\x97\x81\xd0\xb6\x35\x74\xed\x0a\x52\xad\x52\xd0\xc8\xc1\xd0\x05\xb1\x67\x26\xa4\xac\x2b\x02\xb5\x59\x32\x4b\x36\xa0\xdd\x2c\x17\x07\x99\x25\x76\xd6\xc3\xda\x45\x8e\x79\xa2\x18\xb7\x20\x1f\xf4\x56\x43\x6c\x2b\x6f\x82\x08\x62\x2e\xb9\x1b\x1e\x78\x40\x81\x2f\x0f\xd5\x5f\x51\x42\x8e\x33\x3f\x87\xb2\x28\x72\x49\x26\x9e\xea\x2c\x62\x26\x0c\x78\x19\xcb\xfa\x8a\x5d\x75\xb6\x36\xb9\x9a\x90\x2e\xde\x10\x13\x39\x0a\xb0\x02\xfe\x81\x43\xae\x9d\xc3\x9c\xd0\x2f\xef\xe9\xe0\x96\x4e\xd1\xdd\xa0\xe6\x72\x5b\x50\x76\x20\xdb\xea\x14\xc4\x23\x4a\xe8\x20\xd3\x86\xf0\x75\x21\xc7\xe8\xa5\x0c\x11\xb4\x7c\xea\x61\xf0\xeb\xfe\xcb\xed\xb8\x1f\x57\x38\xec\x74\x2b\xfc\xd7\x70\x25\x37\x7c\xb7\x7e\xca\x1e\x70\x01\xe3\x6a\xc4\xe7\x9a\xe3\xe6\x5c\x7c\xd0\xe9\xc6\xd4\xef\xaf\x2b\xb9\xd3\xbc\x12\x07\x89\xaf\x6a\xd0\x7a\x0b\xa6\xb7\x59\x02\x12\x47\x3d\x5e\x96\xc0\x4b\x25\x5c\x35\x4f\x5d\x36\x91\x87\x9f\x76\xf8\x6b\x79\xb8\x21\xee\x95\x54\xfd\x46\xb8\x5c\xc2\x9b\x0b\xf3\xd0\xc8\xf6\xc8\xef\x72\xb1\xe6\x20\x23\x71\xfa\x3a\x9a\x37\xb1\x08\xe1\x4e\xf2\x90\x89\xa1\xbb\xab\x21\xda\x79\xc2\xb9\x12\xff\x6b\xdd\x95\x72\x59\x77\xef\x3f\x7c\x3e\xc5\xbc\x8b\x77\xa0\x6c\x74\xe6\x02\xc8\x4c\x88\x49\xf6\x0f\x19\xec\xeb\xce\x62\xfc\x6c\xb1\x6d\xd7\xdb\xdf\xdf\xc8\x92\xec\x99\xe0\x11\xc3\x33\x2e\xd4\x62\xea\x4b\x55\xd1\x9d\x1d\xb3\x4c\xe0\xb9\x9d\x56\x1e\xfe\xf9\xbb\xf8\xec\x55\x6d\xaf\xa3\x17\xa5\xf3\x4c\xaa\x36\xb0\x85\x20\x9d\xa7\x15\x21\x94\x4b\x84\x6d\x33\xb4\x51\x4a\x00\x93\xf5\x90\x7d\x1a\x7e\xcf\xdb\x3d\x76\xe2\x79\xd3\x6e\xdc\xb7\xee\x26\xfc\x80\x47\x3a\x4e\xfe\xa5\x61\x71\x5b\x02\xa6\x35\x3b\x54\x0a\x70\x84\xe4\xdf\x57\xa1\xda\x04\x9a\x49\xfe\x9e\xc1\x0f\xdf\x06\x75\x36\xf0\xa5\x55\xf5\x93\xbc\xe9\x3a\x91\x2a\xc3\x91\xef\xe1\x77\x37\x25\x95\x94\xb0\x65\x3d\xd9\xc2\xc2\x71\x35\x57\xd5\x8e\x4f\xba\x24\x86\xfe\x3f\xda\xcf\xf7\x93\xb9\x62\xb5\x5a\x6b\x6f\x5f\xdb\xb3\xe3\xec\x6f\x00\x00\x00\xff\xff\xf1\x40\xa5\xf7\xcf\x0c\x00\x00")

func schemasV1ErrorsJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemasV1ErrorsJson,
		"schemas/v1/errors.json",
	)
}

func schemasV1ErrorsJson() (*asset, error) {
	bytes, err := schemasV1ErrorsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schemas/v1/errors.json", size: 3279, mode: os.FileMode(420), modTime: time.Unix(1509760526, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"schemas/v1/errors.json": schemasV1ErrorsJson,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"schemas": &bintree{nil, map[string]*bintree{
		"v1": &bintree{nil, map[string]*bintree{
			"errors.json": &bintree{schemasV1ErrorsJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

