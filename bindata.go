// Code generated by go-bindata.
// sources:
// r.js
// DO NOT EDIT!

package shot

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

var _rJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\x3d\xcf\x9b\x30\x10\xc7\x77\x7f\x8a\x2b\xc3\x83\x51\x10\x21\xa9\xaa\x4a\xa0\xac\x9d\x3a\x25\xdd\xaa\x0c\x2e\x1c\x60\x09\x6c\xd7\x3e\x92\xb6\x11\xdf\xbd\xb6\x69\x5e\xba\x74\xea\x86\xff\x2f\xf8\x77\xe7\x8b\xb0\xe0\x7e\x3a\xc2\x09\x0e\x60\xf1\xfb\x2c\x2d\xf2\x74\x55\xd2\xac\x66\xc1\x17\xb6\x77\xde\x5d\xc5\x22\x9c\x6a\xc6\xa2\x33\xdb\xd1\x1b\xc9\x40\x64\xaa\xed\x36\xd9\x04\xef\xeb\xee\x9c\xb3\x4e\x8e\xa8\xc4\x84\xde\x8d\xda\xfe\xbc\x49\x0a\xa3\xfa\x24\x67\x24\x27\xd4\x33\xdd\x9d\xf7\x3e\x6d\x44\x8f\xaf\xd7\x5f\xf1\x5b\x90\xd2\xac\x68\x2c\x0a\x42\xee\x41\x62\xa8\xb8\x48\xbc\x1a\x6d\xe9\x24\x7f\x85\xc6\x0d\xae\xb2\xa5\xa1\x82\xdd\xbe\x2c\x73\x18\x50\xf6\x03\x55\xf0\xb1\x2c\x61\xa9\xd7\x46\x33\x4a\x73\xc4\x86\x62\x9a\xb4\xa9\xc0\x07\x47\xec\x28\x7e\xfc\xab\x1e\x06\xec\x7c\xad\x13\xa3\xc3\x3b\x80\x56\x9f\xb5\x68\x3f\x49\x25\xdd\x80\x6d\x70\x67\xd5\x90\xd4\x8a\x3b\x12\x34\xbb\x0c\x6e\x0c\xa0\xd1\xca\xe9\x11\x8b\x51\xf7\x3c\x3d\x45\xa3\x82\x14\x36\xf0\x27\x54\xfb\x0c\x38\xa4\x2f\xeb\x2e\xf8\xe3\x27\x59\x68\x03\xc4\xab\x2c\xaa\x16\x2d\xbf\xaf\x72\x2d\x81\x19\x84\x22\x3d\x15\xf8\x43\x12\x8f\xda\x92\xc3\xee\x43\x59\x96\x19\x5b\x9e\x94\x47\x74\x7a\xb6\x0d\xfa\xd1\x51\x5e\xfe\x26\xb5\xe8\x8c\x07\xc4\xc0\x2a\x3b\x78\x9c\x8b\xf8\x9c\x87\x43\x7c\xd6\xb7\x37\x78\xd7\xad\xd3\xfc\x7f\x50\x88\x8b\x25\x3b\x23\x5b\x5e\xa8\x0d\x2a\xee\xef\xce\x9f\xa8\x2b\x00\x00\x5b\xb2\xfa\x77\x00\x00\x00\xff\xff\xd9\x06\xa0\x6b\xab\x02\x00\x00")

func rJsBytes() ([]byte, error) {
	return bindataRead(
		_rJs,
		"r.js",
	)
}

func rJs() (*asset, error) {
	bytes, err := rJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "r.js", size: 683, mode: os.FileMode(436), modTime: time.Unix(1442780851, 0)}
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
	"r.js": rJs,
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
	"r.js": &bintree{rJs, map[string]*bintree{}},
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
