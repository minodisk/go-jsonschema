// Code generated by go-bindata.
// sources:
// fixtures/schema.md.tmpl
// DO NOT EDIT!

package doc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _fixturesSchemaMdTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x54\xc1\x8e\xd3\x30\x10\xbd\xfb\x2b\x2c\x75\x0f\x20\xb1\xa9\xf6\x5a\x89\x0b\xb0\xa8\x07\x10\x65\xdb\x1b\x42\x8a\xb5\x19\x36\x16\xad\x6d\xec\xa9\x44\x15\xe5\xdf\xb1\x63\x3b\xb1\x9d\x14\x96\xfb\x9e\x6a\x7b\x9e\xdf\xbc\x37\x7e\xcd\x8a\x76\x5d\x75\xe0\x78\x84\xbe\x27\x76\xf9\x01\xcc\xa3\xe6\x0a\xb9\x14\xf6\xc0\x9e\x68\x26\x9e\x80\xde\x08\x76\x82\x37\xf4\xc6\x3c\xb6\x70\x62\x74\xf3\x96\x56\x3b\x2d\x15\x68\xe4\x60\xaa\xfd\x70\x6a\xfa\xfe\x96\x7e\xeb\xba\x00\x8a\xac\xdf\x5f\x75\xdd\x59\x1f\xb7\xcc\xb4\xb4\x28\xbd\x9e\xf8\xb9\x25\x3f\x72\xf1\xd3\x51\x47\xd4\x27\xbb\xb7\xa4\x94\x7a\x5a\x57\xae\x3e\x03\xb6\xb2\xb9\x17\x8d\x92\x5c\x60\xc1\xbe\x88\x70\x4d\x40\x34\x7d\x1f\x7e\xfe\xd7\x13\x59\xb9\x19\x15\xca\xdd\x64\xf8\x8f\x51\x69\x36\xb5\x09\x3c\x1b\xe6\x20\x60\x65\x09\xa7\x46\xf3\x19\x2b\x5f\xbb\xa4\xa3\x58\x1e\xb6\xed\xe4\xee\x0c\x9e\x9c\x9a\x78\x33\x6f\xec\xe6\x67\x91\xcb\xc5\x69\x38\x19\xc1\xfd\x6f\x76\x52\xce\xa8\xbb\x1c\x36\x1b\x5a\xa7\x34\x23\xa4\xbe\xc2\x71\xb8\xa8\x40\xe0\x56\x9b\x4c\x83\xaf\x5d\xb9\xb8\x63\x88\xa0\x83\xf2\xb0\xb1\xcd\xd7\x29\xc1\x88\x59\x5f\x6b\xff\x51\xea\x13\x43\x4f\xe2\xd7\xb9\x84\x58\xbf\x72\xfd\x01\x58\xf3\x45\x1c\x2f\x9e\x20\xee\x72\x8a\x09\x53\x64\xec\x39\xb1\x26\x64\xe5\x93\xb5\x98\xda\x18\xb0\xa1\x58\xc6\x6b\x7e\x48\x72\x17\x03\xc0\x07\x25\x11\xb3\x94\xaf\x04\xf9\x12\xb2\x67\x85\x2c\xfb\x96\xd4\x75\xdd\x22\x2a\x92\x3f\xa3\x25\x8d\x07\xc1\xc0\xf4\xb0\xb1\xf0\xf5\x0c\xfa\xb2\x47\xcd\xc5\x93\x85\x6f\x0f\x87\xdd\xfa\xae\xba\x4b\x9e\xcf\x7e\xd2\x1e\xe0\xd7\x19\x0c\xbe\x93\x8d\x8b\xd8\x7b\x29\x10\x04\xde\x8e\x46\x07\x58\xc0\x84\xa2\xb7\x1c\xd5\x6d\xa5\xf1\xa1\x0f\xc9\x73\xfb\x2c\x22\xb3\x1e\xa4\xa0\xf5\xc7\x89\x5b\x32\x5a\x8e\x92\x13\x25\x46\x49\x61\x60\x8f\x0c\xcf\x26\x99\x41\x2c\xd8\xbf\x8b\x91\x62\xd7\x6a\x66\x60\x2e\xc3\x63\xfe\xe1\xd5\x83\x96\xcc\xfe\x95\x8e\x94\x0c\x33\x5f\xc5\xbb\xfe\x09\x00\x00\xff\xff\xc4\x58\xbe\x7d\x17\x07\x00\x00")

func fixturesSchemaMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_fixturesSchemaMdTmpl,
		"fixtures/schema.md.tmpl",
	)
}

func fixturesSchemaMdTmpl() (*asset, error) {
	bytes, err := fixturesSchemaMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/schema.md.tmpl", size: 1815, mode: os.FileMode(420), modTime: time.Unix(1435540919, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"fixtures/schema.md.tmpl": fixturesSchemaMdTmpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"fixtures": &bintree{nil, map[string]*bintree{
		"schema.md.tmpl": &bintree{fixturesSchemaMdTmpl, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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
