// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package main generated by go-bindata.// sources:
// shell/getsysteminfo.sh
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _shellGetsysteminfoSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4f\x6f\xd3\x30\x18\xc6\xef\xfe\x14\x0f\xae\xa5\x6e\x9a\x82\xd7\xb1\x53\xaa\x5e\x10\x68\x27\x38\x71\xab\x4a\x94\x3a\xce\x62\x2d\x71\x2c\xff\x29\xda\x68\x25\x2e\xf0\x0d\x90\x90\x00\x09\xee\x20\xed\xc8\x17\xca\xf6\x35\x90\xd3\x64\xaa\x60\x3e\xbe\xf2\xfb\xfb\x3d\x7e\x3c\x79\xc2\x83\xb3\x7c\xad\x34\x97\x7a\x03\x57\x91\x78\x26\x26\xac\x6b\x25\x32\x65\x16\xec\x48\x04\x5b\x23\x71\xa8\xbc\x37\x29\xe7\xaa\x14\xad\x2e\xd5\xe5\x53\xd1\x72\x65\x8e\xc9\x84\x48\x51\xb5\xa0\xca\x74\xdf\x6e\xbb\xef\x1f\x52\x50\x36\x8b\x10\x61\x42\x26\x5a\x2b\x23\x22\xf7\xe0\xc6\xb6\x82\x0b\x13\x94\x2e\x5b\x6c\x71\x69\xa5\x01\x15\x26\x20\x5e\x72\x14\x5b\x54\x32\x2f\x90\xcc\xb0\x45\xfe\xee\x0a\xd3\xf7\xc6\x2a\xed\xc1\xce\x77\xd3\x63\x42\x06\x8d\x30\xe1\xee\xc7\x9f\xbb\xcf\xb7\xd1\x33\x2a\x08\x69\x64\xd3\xda\xeb\xcc\xa9\x9b\xe8\x2b\xad\x94\x48\x9a\xd1\xf2\x4a\x36\xff\x32\xcf\x22\xb3\x27\x26\x1a\xb4\xfb\xf4\xb1\xfb\xfd\x25\x05\xdd\x8f\xd8\x01\xed\x61\x11\xeb\x6b\x2f\x17\x6c\xc6\x67\xa7\x67\xe7\x73\xec\x39\x71\x06\x7a\xf1\x9c\x62\x37\xed\x9b\xab\x37\xcd\x90\xa1\x28\x91\xbc\xa9\xc6\x08\x6f\x97\xcb\xd4\x99\x5c\xc8\x74\xb5\x7a\x40\xba\xd0\xe0\x64\x11\xc3\xcc\x5f\xbe\x7e\x81\x21\x9b\x0b\x4d\xff\x60\x55\x62\x89\xe4\x06\x94\x0d\x50\x8a\xd5\x1c\xbe\x92\x9a\x00\xc0\x68\x3a\x25\xa5\x22\x84\x38\x9f\xeb\x22\xb7\xc5\xe3\x76\x5e\xc8\xcd\x7f\xda\x67\x8f\x6a\x49\xa1\xdc\x55\x66\xa5\x0b\xb5\x5f\xb0\xa3\xbe\x90\xfd\x26\x1d\xcb\x1b\xd4\x27\xec\x50\xba\xa3\x71\x79\xf8\xa5\xfb\x9f\xbf\xee\xbf\xc6\x42\xd9\x01\x8d\x5e\x50\xf2\x37\x00\x00\xff\xff\x7c\x09\x0c\xb5\x70\x02\x00\x00")

func shellGetsysteminfoShBytes() ([]byte, error) {
	return bindataRead(
		_shellGetsysteminfoSh,
		"shell/getsysteminfo.sh",
	)
}

func shellGetsysteminfoSh() (*asset, error) {
	bytes, err := shellGetsysteminfoShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shell/getsysteminfo.sh", size: 624, mode: os.FileMode(420), modTime: time.Unix(1583293595, 0)}
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
	"shell/getsysteminfo.sh": shellGetsysteminfoSh,
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
	"shell": &bintree{nil, map[string]*bintree{
		"getsysteminfo.sh": &bintree{shellGetsysteminfoSh, map[string]*bintree{}},
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
