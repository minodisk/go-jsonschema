// Code generated by go-bindata.
// sources:
// assets/routing.go.tmpl
// assets/struct.go.tmpl
// assets/validator.go.tmpl
// DO NOT EDIT!

package generator

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

var _assetsRoutingGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x3c\x8f\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x0a\x53\xed\xc0\xa4\xe1\xde\x91\x38\x20\xb6\x03\x12\x30\x04\x7d\x81\xac\x31\x6d\xd8\x62\x57\x69\x7a\x40\x51\xde\x9d\x2c\xcb\x76\x72\xfd\xe9\xff\xea\x3f\x6d\x0b\x2f\x62\x08\x06\x62\xf2\x3a\x90\x81\xc3\x1f\xfc\xce\xc2\x73\x3f\x92\xd3\xa8\x72\x60\xbb\x87\x8f\x7d\x07\xbb\xed\x6b\x77\xa7\xd4\xa4\xfb\xa3\x1e\x08\x62\xc4\xfa\x99\x92\x52\xd6\x4d\xe2\x03\x34\x83\x0d\xe3\x72\xc0\x5e\x5c\x3b\xc8\x83\xd3\x3e\x58\xb6\x6d\x9d\x8d\x52\x3f\x0b\xf7\xf0\x25\x4b\xa0\x7b\x0f\x15\x63\xd9\xfd\x1a\xa2\x8a\xd1\x6b\xce\x3f\x5f\xb1\x76\xb4\x81\xd5\xa5\x05\x3c\x3e\x01\xd6\x42\x9f\x5e\x26\xca\x1a\xcd\xf8\x5d\xc8\x9c\xd2\xcd\xb2\x59\x39\x59\x3e\x9e\x85\xea\xe2\x5b\xde\x73\x06\xc0\xe3\xb3\x31\x97\xd3\x4d\x8c\x25\x87\xef\x14\x46\x31\x29\x35\x1b\xb8\xb1\x1d\x9b\x49\x2c\x87\x42\xaf\xb0\xb3\xe1\x94\x1f\xba\xce\x0d\x89\xcd\xf9\x64\x19\x49\xfd\x07\x00\x00\xff\xff\x14\x9f\x66\x2d\x41\x01\x00\x00")

func assetsRoutingGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsRoutingGoTmpl,
		"assets/routing.go.tmpl",
	)
}

func assetsRoutingGoTmpl() (*asset, error) {
	bytes, err := assetsRoutingGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/routing.go.tmpl", size: 321, mode: os.FileMode(420), modTime: time.Unix(1438848084, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsStructGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xb1\x4e\x87\x30\x10\xc6\xf7\x3e\xc5\xf9\x0f\xa3\x29\x3b\x89\x13\x18\xe3\x02\x26\xf6\x01\xa8\x70\x41\x04\xda\xa6\xad\x03\xb9\xf0\xee\x5e\x6d\x9d\x34\x6e\xd7\xaf\xdf\xfd\x7e\x69\xeb\x1a\x5a\x3b\x23\x2c\x68\xd0\xeb\x88\x33\xbc\x9d\xf0\x11\xac\x09\xd3\x3b\x1e\x5a\x0a\x2e\x74\x03\xf4\x83\x82\xc7\xee\x59\xdd\x09\xe1\xf4\xb4\xe9\x05\x81\x48\x96\xf1\xba\x84\x20\xf2\xda\x70\x5a\x19\x7d\xe0\x3d\x54\x79\x1d\x9a\x07\x90\x85\xf4\xe2\xad\x43\x1f\x57\x0c\xf2\xf5\x3b\x09\xbc\x47\x54\x9a\x3d\xaf\xa5\x76\x30\x7a\x43\x65\x5b\x3e\xee\x19\xc6\xad\x78\xba\xe4\x2b\x55\xa9\xd6\xb8\x73\x0c\x21\xfa\xcf\x29\x02\xfd\x92\xbb\xac\x3a\x13\xb0\xfa\x4f\x0f\x4c\xfd\xcb\x98\x64\x3f\x10\xa9\xd8\x2e\x9f\x2c\xa7\x63\xfa\x98\xe6\xc6\x97\xb9\x76\x1b\x89\xd0\xcc\x0c\x4a\x2f\xc9\xd3\x57\x00\x00\x00\xff\xff\x39\x22\x35\x44\x51\x01\x00\x00")

func assetsStructGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsStructGoTmpl,
		"assets/struct.go.tmpl",
	)
}

func assetsStructGoTmpl() (*asset, error) {
	bytes, err := assetsStructGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/struct.go.tmpl", size: 337, mode: os.FileMode(420), modTime: time.Unix(1438852763, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsValidatorGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x52\xc1\x6e\x9b\x40\x10\xbd\xf3\x15\x13\xe4\x83\x91\x2c\x72\x6f\xe5\x93\x93\x43\xa4\x36\x8e\x54\xda\x4b\xd5\xc3\x06\x8f\xf1\x36\xb0\xd0\xdd\xc5\xb2\xb5\xe2\xdf\x3b\xbb\x80\xc1\x4a\xc0\x28\x4e\x7c\xf1\x32\xf3\xe6\xcd\x7b\x33\x73\x7b\x0b\xab\x7c\x83\x90\xa0\x40\xc9\x34\x6e\xe0\xf9\x08\x7f\x55\x2e\x54\xbc\xc3\x8c\x85\x1e\x01\xee\xd6\xf0\xb8\x8e\xe0\xfe\xee\x21\xba\xf1\xbc\x82\xc5\x2f\x2c\x41\x30\x26\x6c\x9e\x55\xe5\x79\xc6\x48\x26\x28\x3a\x13\x2c\xc3\x05\xcc\xea\x72\xf8\xb2\x84\xb0\x61\x7a\x92\x79\x81\x52\x73\x54\xe1\x0f\x17\x51\x54\x67\x4c\x83\x7c\xa4\x32\x8b\x56\x82\xbd\x60\x94\xaf\xe8\x33\xad\xc9\x6a\x14\x17\x4a\x33\x11\x3b\xcc\x96\x4b\xa5\x57\x3b\x26\xdb\x36\x61\xc4\x75\x6a\x81\xdb\x52\xc4\x30\xef\xc1\xab\x0a\x4e\x2d\x5a\x54\x00\xbf\x58\xca\x37\x64\x76\x1e\xc0\x1c\xa5\x54\xf0\xfb\x0f\xfd\xe5\x32\x00\xe3\x01\xb8\x08\xb5\x69\x82\xa6\xa2\xd8\x9e\x9a\xed\xeb\xaa\x5c\xb6\xf5\xb9\x7c\x65\xbb\xa8\x4d\x1e\x6d\xfd\x6c\xc4\x38\x89\xb2\x94\x43\x86\xa1\xd7\x6c\xd9\xb5\x33\xfe\xd9\xbc\xaa\xca\x5f\x80\x5f\x53\xd1\xdb\x96\x19\xc3\xb7\x9d\x8a\x30\x3a\x16\x18\xae\x72\xa1\x19\x0d\x04\x7c\x2e\x34\x26\x28\x7d\xf0\x45\x99\x3d\xd3\xc3\xb5\x7a\x55\xf5\xbd\x4c\x35\x2f\x52\x5c\x6f\x9b\x3c\x65\x69\x14\x56\xed\x49\x56\x0f\x74\x3e\xef\xb0\xd5\xb3\xb0\x93\x7f\x93\x33\xf8\xea\xe8\x6e\x96\x20\x78\xea\x46\x6e\x7f\x6e\xec\x4b\x60\x45\x81\x62\xe3\xd6\xb2\xb0\xb1\xc0\xa5\x69\x60\x14\xb5\x73\x3b\x57\xca\x0e\x3c\x2b\xb3\x31\x99\x35\x62\x9a\xc6\x96\xed\x3c\x7c\x7f\x88\xd3\x52\xf1\x3d\x9e\xf2\x1f\x68\x80\x8b\x4b\x06\x6a\xc4\x34\x03\x2d\xdb\x90\x81\x36\x7f\x85\x01\x77\x64\xdd\x6b\xec\xdc\x94\x96\x5c\x24\x43\x47\xc6\x0e\xdf\x50\x24\x7a\x37\xbe\xbc\x1a\x33\x75\x7d\x2d\xe3\x87\x2e\xe8\xb2\xcc\x16\x33\x75\x49\x57\xc8\x1c\x14\xfa\xc4\xb4\x46\x29\x46\x64\x36\x88\x41\x91\x7e\x5f\xe5\x89\xce\xff\xa4\x5b\x79\xd0\x98\xa9\xc1\xd3\xe8\x67\x07\x6f\xc3\x81\xa6\x9e\x46\xc3\xf8\xb6\x9b\xf7\x0e\x9d\x76\x39\x41\x69\x03\x9a\x7a\x1d\x9f\xa2\xf4\xa7\xe0\xff\x4a\xbc\x2c\xb6\x87\x1b\xd0\x7b\xa5\xae\xde\x4d\x74\x11\x89\xba\x94\xc2\x51\x78\x5d\xfc\x7f\x00\x00\x00\xff\xff\xe9\xaa\x2c\x73\x9c\x08\x00\x00")

func assetsValidatorGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsValidatorGoTmpl,
		"assets/validator.go.tmpl",
	)
}

func assetsValidatorGoTmpl() (*asset, error) {
	bytes, err := assetsValidatorGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/validator.go.tmpl", size: 2204, mode: os.FileMode(420), modTime: time.Unix(1439198992, 0)}
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
	"assets/routing.go.tmpl": assetsRoutingGoTmpl,
	"assets/struct.go.tmpl": assetsStructGoTmpl,
	"assets/validator.go.tmpl": assetsValidatorGoTmpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"routing.go.tmpl": &bintree{assetsRoutingGoTmpl, map[string]*bintree{
		}},
		"struct.go.tmpl": &bintree{assetsStructGoTmpl, map[string]*bintree{
		}},
		"validator.go.tmpl": &bintree{assetsValidatorGoTmpl, map[string]*bintree{
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

