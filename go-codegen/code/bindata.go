package code

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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_list_go_t = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x57\xdf\x4f\xdb\x3a\x14\x7e\xcf\x5f\x61\x5e\xae\x9c\xab\x12\x2e\xaf\x88\x3e\x70\xa7\x55\x42\xea\x10\x82\x8d\x3d\xa0\x0a\x65\xa9\x53\x0c\xae\x53\xc5\x6e\x59\x15\xe5\x7f\xdf\xb1\x9d\x1f\x4e\x62\x87\x09\xd8\x86\xb4\x3e\x20\xe2\x9c\x1f\xdf\xf7\x1d\xfb\xf8\xa4\x28\xa2\xcb\xc7\x55\x59\x06\x81\xdc\x6f\x08\x82\xc7\x8b\x78\x4d\xca\xf2\x32\x27\x4b\x9a\xc4\x92\xa0\x74\xcb\x13\x84\xe1\xc5\xe7\xff\xca\x32\x44\xdf\xb2\x8c\xf5\x8d\x6f\xa8\xa0\x32\xcb\x2b\x53\x8a\x28\x97\x13\xb4\x8b\x19\x1a\xf5\x9a\x13\x21\x66\xca\xc3\xb8\xa5\x34\x17\xe0\x26\x48\x92\xf1\xe5\xb8\x27\x12\x32\xdf\x26\x12\x15\x01\x82\xdf\x1d\xa3\x42\xa2\xdb\x45\xe5\xa2\xd7\x98\x8a\xcd\x87\xb9\x02\xfd\xb6\x28\x68\x8a\xa2\x73\x71\xbd\xe7\x49\xe5\x70\xc7\xb2\xe4\x11\x09\x58\x88\xae\xbe\x7e\xda\x4a\xf2\xbd\xb2\x24\x7c\x09\x26\xa0\x8f\x46\x79\x41\x9e\x9a\x98\x38\x44\xff\xb6\x90\x0c\x96\x9c\xc8\x6d\xce\x3b\x66\xb3\x3c\x5b\x5f\x33\x9a\x10\xdc\x40\x2c\xca\xd0\x19\xb1\x35\xa5\x92\xac\x45\xcb\xc9\x9f\xe9\x9f\x66\xbd\x30\x42\x9c\x20\xe3\xdb\x62\xc6\xcc\xf2\x0e\xd1\x8c\x32\x49\x72\x9c\x3a\x6a\xed\x48\xe3\x90\x8a\x45\x5a\xac\xe8\x6a\x0e\x7f\x71\xa8\xd7\x96\x24\x25\x79\xfb\xe6\x0b\x67\xed\xbb\x5a\xc3\x0a\xb6\xd8\x32\x89\x4e\xa6\x2e\xe4\x96\x40\xc8\xe4\x4a\x61\x57\xdd\xc1\x5e\x52\x0e\x79\xcc\x57\x44\xe7\x50\xf5\x36\xf0\xd4\x0f\x00\xa6\x78\x17\x5a\x2b\x6d\xa2\xca\x78\x8a\xe2\xcd\x06\x50\x60\x7b\x15\xc2\x86\x8d\x8b\x49\x57\xda\xd2\x1a\x5b\x9f\x8c\xe7\xa3\x3a\x9e\x1e\x26\xf7\x31\xaf\x77\x31\x6a\x6a\xa6\x22\x0a\x45\x66\x1d\x3f\x12\x6c\xdb\x18\x28\xab\x4c\x9f\x06\x6c\xb3\x71\x94\xc0\x57\x86\xe7\x4a\x61\x97\xa3\x7e\xfe\x19\x89\xfd\x32\xdb\xb4\x4e\x0f\xd1\xae\xf3\xb2\x0c\x86\xff\x25\x2c\x13\xa4\xaa\x83\x30\xa8\xe0\x24\x05\x43\xe1\x85\x57\xf9\xbf\x4a\x76\xa7\xb8\x6f\x29\xe7\xc7\x38\xb9\xc7\xbb\xaa\x83\xf7\x5b\xfa\x6f\xe8\x08\x4a\x08\xfa\xfc\x11\x3f\xa8\x20\x62\x65\x3b\x3c\xec\x9a\x29\x1b\x3b\xce\x4c\x09\xe0\x54\xe0\x8c\xef\x3d\xfb\x49\x5d\x3e\xbf\x90\xf6\x4b\x3b\x9b\xe6\x03\x57\x20\x19\xe3\x9b\xc6\x4c\x10\x5f\xd1\xcf\x18\x7b\xb7\x94\x0f\xfc\x9c\x0d\xa7\x11\xd2\x5a\x14\xef\xc5\xc7\x97\x33\x35\x64\x38\x99\x4f\x14\x8f\x18\x0e\xca\x8d\x3d\xb7\xd4\x9d\xe4\x9d\xee\x81\x7e\x47\xe8\x68\xd1\xf2\x19\x53\x64\x1e\xbf\x2b\x41\xb6\x30\xfb\x81\x1a\x16\xf6\xd7\x09\x65\x22\x4e\xc7\x95\xd2\x46\xba\x3d\x1c\x1d\xa1\x19\xe4\x12\x59\x2e\xdd\x92\xfd\xbf\x57\x63\x24\xf6\xcd\x97\x76\xbf\x2c\x2a\x3d\x2a\xdb\x69\x35\x94\x0e\x7a\x92\x2b\x8d\x8a\xa7\x3a\xdd\x83\x9a\xa5\xcd\xa1\xec\x8c\x7d\x75\x54\x5c\xe9\x70\x4b\x17\x93\x5a\x93\xdb\x87\x45\xe8\xed\x75\x73\xc2\xe1\x96\x83\xa0\xdd\x31\x92\x91\x26\x54\xe8\xc3\x74\xfd\x14\x6f\x2c\x4c\x45\x53\xee\x61\x7e\x45\xb6\x79\x98\x58\x46\x5e\x5c\x1f\x58\xc6\x89\x6b\xa2\x7e\xeb\x2d\x96\x64\x1b\x4a\x96\xcd\x2c\xd0\x8c\x9c\x93\x8e\x06\x61\x6d\xbb\xc7\xc6\xa1\x21\xd1\xbd\x5b\xdd\xd3\xbb\x71\xd1\x42\xc2\x86\x82\xaf\x89\x18\x2e\x2e\xe1\xe9\xc4\x66\x30\x35\x43\x7b\x14\x45\xed\x49\x7b\x8e\xbe\x97\xbd\x9f\x3c\xeb\xcf\xc3\xac\x1e\x85\x75\x7e\x48\xef\xad\x3e\xf4\x85\x3f\x09\x54\xa7\x6d\x8a\x30\x0a\x34\xdb\xe0\x57\x75\xae\x17\xc0\x85\x58\xf6\xee\x41\xa7\xe8\xd8\xea\x43\xc3\x7e\x6c\x3a\x50\xef\x83\xa8\x3e\x25\x76\xa4\xc3\xe3\x45\x5f\x90\xda\xec\xc4\x69\xd7\xff\x76\x81\xfd\x07\x82\x5c\xd1\xd5\xbd\x44\xc1\x8f\x00\x00\x00\xff\xff\xf8\x06\x24\x1f\xed\x0f\x00\x00")

func templates_list_go_t_bytes() ([]byte, error) {
	return bindata_read(
		_templates_list_go_t,
		"templates/list.go.t",
	)
}

func templates_list_go_t() (*asset, error) {
	bytes, err := templates_list_go_t_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/list.go.t", size: 4077, mode: os.FileMode(420), modTime: time.Unix(1483615388, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_map_go_t = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x93\x4b\x4b\xc3\x40\x10\xc7\xef\xf9\x14\xe3\xa5\x6c\x4a\x08\x7a\x2d\xf4\x28\x22\x68\x91\x3e\xf4\x20\x45\x62\x32\xa9\x25\xaf\x92\x6c\xa2\x25\xec\x77\x77\x76\xbb\xbb\x49\x6c\x83\xa7\x16\x5a\xba\xf3\x60\x7e\xf3\xdf\xff\xb6\xad\xff\x92\xec\x84\x70\x1c\x7e\x3c\x20\xd0\x71\x11\x64\x28\xc4\xeb\xbe\xda\xf3\xa2\x84\xb8\xce\x43\x60\x14\x5e\xdf\x0a\xe1\xc9\xfc\xfa\x4e\x08\x17\x3e\x8b\x22\xb5\x3d\xa0\x9a\x40\x08\xa8\x78\x59\x87\x1c\x5a\x07\xe8\xf3\x91\x05\x07\xa0\xef\xbb\x6e\xdf\xea\x6e\x47\x65\xdb\x76\x1f\x83\xff\x58\xad\x8e\x79\x48\x31\xd5\x90\x16\x61\x02\xd3\x8a\x22\xfe\xf2\xed\xb9\xe6\xf8\xa3\x4b\x31\x8f\xa8\x86\x3a\x15\xcf\x02\xbf\x2d\x28\x73\x61\x6a\x0f\x7a\x70\x89\x15\xcc\xe6\x30\xb1\xf1\x53\xd8\x30\xcd\x2e\x41\xb5\xc2\xb3\x45\x17\xd0\x2c\xde\x0c\x26\x7d\xbe\x61\xdb\x09\x53\xfe\x17\x1a\x84\xd7\x65\x2e\x79\x2c\x3b\xcb\x7a\xbc\x2e\x3c\x20\x67\x09\x1e\x41\xd3\xb8\xc0\x1a\x23\xb2\x07\x71\x51\xe7\x91\x92\xda\xd5\x9b\x0d\xc9\x40\xc5\x32\x5f\x81\xf9\xcb\x27\xfa\x65\xae\x8a\x45\x18\x63\xd9\x65\x36\x79\xda\xe5\x0c\xa6\x3a\x34\x66\xca\x5c\x56\x4b\x5d\x88\x66\xdb\x83\x1f\x03\x5f\x0d\xc1\x3d\x68\x82\xb4\xb3\xc7\xff\xb4\xa3\xb0\xe3\xac\x3d\x40\xc2\xa5\x79\x63\x6c\xf7\x41\xf8\xc5\x1a\xed\xe0\xbf\x96\xbe\x9a\x94\x31\x0d\x4b\x48\x07\x69\xbd\x32\xc8\x77\xa8\x81\xa1\x73\x1f\x4d\xbc\xd1\x5c\x4c\x96\x12\x8b\xf1\x88\xb1\xcd\xc8\x4a\x9b\x43\x14\x70\x64\x55\x19\x5e\x32\x6f\xff\x0d\x5c\x45\xfa\xf3\xdd\x24\x49\xb7\x98\xb9\x1a\x75\x31\xe7\x0f\x20\x1b\x5b\x6b\x89\x59\xd1\xe0\xf0\x05\x5c\x65\x81\x08\x53\x24\xfd\x4e\x9c\x1e\xd0\x40\x97\x98\x7e\x03\x00\x00\xff\xff\xfe\x75\x61\x17\xfe\x04\x00\x00")

func templates_map_go_t_bytes() ([]byte, error) {
	return bindata_read(
		_templates_map_go_t,
		"templates/map.go.t",
	)
}

func templates_map_go_t() (*asset, error) {
	bytes, err := templates_map_go_t_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/map.go.t", size: 1278, mode: os.FileMode(420), modTime: time.Unix(1483615591, 0)}
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
	"templates/list.go.t": templates_list_go_t,
	"templates/map.go.t": templates_map_go_t,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"list.go.t": &_bintree_t{templates_list_go_t, map[string]*_bintree_t{
		}},
		"map.go.t": &_bintree_t{templates_map_go_t, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

