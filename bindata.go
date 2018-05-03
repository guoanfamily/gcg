// Code generated by go-bindata.
// sources:
// template/controller/controller.go.tmpl
// template/controller/gen_controller.go.tmpl
// template/main.go.tmpl
// template/model/database/db.go.tmpl
// template/model/database/table.go.tmpl
// template/router/gen_router.go.tmpl
// template/router/router.go.tmpl
// DO NOT EDIT!

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

var _controllerControllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x4a\x3b\x31\x14\xc5\xd7\x13\xc8\x3b\x5c\xba\x6a\xa1\x90\xfe\xff\x4b\x61\x56\xea\xb2\x1b\x7d\x82\x98\xde\x86\xc1\x69\x66\x48\x32\x48\x29\x05\x07\x2c\x48\x57\x2e\xaa\xa0\x82\x1f\xa0\x3b\x61\x28\x8a\xe0\xaa\x2f\xe3\x7c\xf8\x16\x12\x67\x2a\xad\xab\x5c\xce\xcd\x39\x27\xbf\x30\x06\x12\x15\x6a\x6e\x11\x8e\xc6\x20\x03\x15\x5b\x4a\x62\x2e\x8e\xb9\x44\x10\x91\xb2\x3a\x0a\x43\xd4\x94\x50\xc2\x58\x3e\x7b\xc9\x67\xcb\xe2\xea\xa3\x78\xbd\x2c\x6f\xce\xaa\xd5\x22\xbf\xbd\xfb\x5c\x3d\x16\x69\x46\x89\x1d\xc7\x08\x07\x68\x92\xd0\x82\xb1\x3a\x11\x16\x26\x94\x78\xbb\xd1\x00\x21\x50\x96\x12\xaf\x6f\xa4\xdb\x04\x4a\x52\xe2\xed\x71\xcb\x9d\x8e\x7a\xc8\x05\x4e\xa6\x94\x4c\xeb\x92\xaf\xc5\x75\x95\x65\xe5\xfc\xbd\x38\x4d\xcb\x87\x94\x12\x11\x29\x63\xe1\x30\x11\x02\x8d\x01\xf0\xe1\x7f\xaf\xb7\x56\xf7\xb5\x8e\x74\xad\xfd\xdb\xb4\xaf\x1f\xb5\xe5\x75\xfd\x3e\xb4\x8a\xf3\x8b\x7c\x7e\xdf\xda\x8a\x68\x56\xf9\xd3\xb2\x7a\x7b\x6e\xb9\xa4\x61\xa2\x44\x83\xd3\x37\xb2\x2d\x1a\x0c\xe8\x8e\x7e\x29\xa0\x3b\xf8\x03\xd1\x69\x1c\x0e\x9c\x31\x8d\x06\x76\x7c\x50\x78\xd2\xae\xe5\x0e\x25\x9e\x46\x9b\x68\xb5\x71\xcf\x73\xd1\x5d\x37\x8c\x8c\xfc\x39\x5d\xaa\x1b\xdc\x9f\x7c\x07\x00\x00\xff\xff\x66\x15\x43\x9b\xa1\x01\x00\x00")

func controllerControllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerControllerGoTmpl,
		"controller/controller.go.tmpl",
	)
}

func controllerControllerGoTmpl() (*asset, error) {
	bytes, err := controllerControllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/controller.go.tmpl", size: 417, mode: os.FileMode(438), modTime: time.Unix(1525240335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGen_controllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x6b\xdb\x40\x10\x85\xef\x02\xfd\x87\xe9\x1e\x8a\x04\x66\x15\x7a\x0c\xf8\x92\x84\x16\x5a\x9a\xa4\x75\xa1\xe7\xd1\x6a\x2c\xaf\xbd\xda\x15\xb3\x23\xd3\x20\xf4\xdf\x8b\x2a\x19\x62\xa8\x1b\xd7\xb9\xe4\xa8\x9d\xd1\x7b\xf3\xde\x57\x14\x50\x93\x27\x46\x21\x28\x9f\xa0\xb6\xbe\x95\x34\x69\xd1\xec\xb0\x26\x30\xc1\x0b\x07\xe7\x88\xd3\x24\x4d\x6c\xd3\x06\x16\xc8\xd2\x04\x40\x79\x92\x62\x23\xd2\xaa\x3f\x5f\xb5\x95\x4d\x57\x6a\x13\x9a\xc2\x61\x19\x05\xcd\xae\x20\xb3\x09\xd3\xb4\xef\xf5\x23\x87\x2d\x19\xb9\xc7\x86\x86\xa1\x68\x42\x45\xae\xe8\xfb\xce\x57\xc4\xce\x7a\x02\x7d\x77\x33\xcd\x54\x9a\xe4\xa3\xd9\xba\xf3\x06\xfa\xbe\x6b\x5b\xe2\x5b\x6c\xc8\x81\xfe\x81\xa5\xa3\x69\x6b\x85\x7b\xca\x0c\x8c\x16\xfa\x36\x78\xa1\x5f\x92\x03\x31\x07\x86\x7e\xb4\x04\xd8\x23\x43\x28\xb7\x10\x89\xf7\xd6\x90\x3e\x29\x35\xad\x1b\x7d\x63\x7d\x95\xbd\x0f\xe5\x36\x9f\x5e\x98\xa4\x63\x0f\x46\x7f\x5e\x3d\xdc\x67\x63\x54\xbd\x12\x94\x2e\x3e\x7c\x59\x7c\xa7\xd8\x39\xf9\x1a\xeb\xec\xc3\xd5\xd5\x42\x85\x9d\x5a\x84\x72\xab\xc7\xab\xf2\x3c\x4d\x86\x97\x03\xdc\x91\x23\x79\x83\x11\xa6\xbb\x8e\x43\x7c\x22\x39\x69\xfd\x73\x43\xfc\x42\x0c\x13\x7c\x05\xd7\x4b\x30\xfa\x5b\x47\xfc\xf4\x88\x8c\x4d\xa6\xc6\x57\x35\xdf\x89\x5c\xc3\xf5\xf2\x78\x8e\x5c\x1f\xc6\x63\x0f\x4c\x11\xac\x17\xe2\x35\x1a\xea\xe7\xc4\x76\x3d\xfe\xfa\x6e\xa9\xd4\x6c\x35\x85\x8e\xb0\x3c\x54\x76\xc6\xe9\xc1\x57\x0b\xe4\x7a\xf6\x1a\xc8\x45\x82\xd7\xc9\x1d\xa4\x2e\x82\xc0\x14\xcf\x6f\xff\xa3\xe5\x28\xaf\x6d\xff\x0c\x02\xff\xa4\x70\x92\xc4\x7f\xd5\x37\x47\x39\xa6\xf1\x57\x22\x17\xca\x3e\x97\x7c\x8e\xf7\x42\x3a\xbf\x03\x00\x00\xff\xff\x48\x7c\x48\xba\x39\x05\x00\x00")

func controllerGen_controllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGen_controllerGoTmpl,
		"controller/gen_controller.go.tmpl",
	)
}

func controllerGen_controllerGoTmpl() (*asset, error) {
	bytes, err := controllerGen_controllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/gen_controller.go.tmpl", size: 1337, mode: os.FileMode(438), modTime: time.Unix(1525245094, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe5\xe2\xe5\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\xe0\xe5\x52\x50\x50\xaa\xae\xd6\x0b\x28\xca\xcf\x4a\x4d\x2e\xf1\x4b\xcc\x4d\xad\xad\xd5\x2f\xca\x2f\x2d\x49\x2d\x52\xe2\xe5\xd2\x04\xa9\x4d\x2b\xcd\x4b\x06\x6b\xd4\xd0\xac\xe6\xe5\xe2\x84\x48\xea\x05\x81\x29\x0d\x4d\x5e\xae\x5a\x5e\x2e\x40\x00\x00\x00\xff\xff\xf4\x0b\xd4\xcc\x5f\x00\x00\x00")

func mainGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_mainGoTmpl,
		"main.go.tmpl",
	)
}

func mainGoTmpl() (*asset, error) {
	bytes, err := mainGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go.tmpl", size: 95, mode: os.FileMode(438), modTime: time.Unix(1525254979, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseDbGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x8b\x1c\x45\x18\x3e\x4f\x43\xff\x87\x37\x85\xb3\x54\xbb\x93\xea\xd5\x80\xc8\x48\xa3\x6c\x86\xc4\x80\x9a\xe0\x6e\x4e\xc3\x10\xaa\xbb\xaa\x7b\xcb\xed\xae\xea\xad\xaa\xde\xcd\x30\xcc\xc1\x83\xc4\x8b\x92\x83\xe8\x2d\x22\xa2\x5e\xfc\x00\x0f\x22\x8a\xe4\xcf\x6c\x66\xe3\xbf\x90\xaa\xae\x1d\x7b\x23\x04\x9c\x4b\x57\xbd\x9f\xcf\xfb\x3c\x6f\x4d\x9a\x42\xc5\x25\xd7\xd4\x72\x06\xf9\x12\x2a\x21\x5b\x1b\x47\x2d\x2d\x8e\x69\xc5\x61\xb5\xea\x24\xe3\xba\x16\x92\x03\xf9\x80\x36\x7c\xbd\x8e\xa3\xd5\x0a\x5e\xb1\x34\xaf\x39\x4c\x33\x20\x87\xfe\xe4\xec\x71\x24\x9a\x56\x69\x0b\x38\x8e\x46\xa8\x12\xf6\xa8\xcb\x49\xa1\x9a\xb4\xea\x14\x95\x25\x6d\x44\xbd\x4c\xcd\x49\xfd\x10\xc5\xd1\xe8\x01\x5c\x89\x50\xd7\xcd\x49\x7d\x9d\x69\x71\xca\x75\xda\x2c\xcd\x49\xed\x82\x50\xd9\x58\xff\xad\x55\xe5\xbf\x95\x6a\x8f\x2b\x22\x64\xba\xa4\x4d\x4d\x4e\x5f\xf7\x46\xa1\x52\xa1\x3a\x2b\xfa\x14\x65\x50\x1c\x25\x0e\xcd\x29\xd5\x1e\x0a\xcb\xef\xc8\x52\x41\xff\x7b\xd5\x01\x20\xb3\xfd\x10\x53\x76\xb2\x80\xdb\xdc\xce\xf6\x71\xb2\xf5\xc1\x2a\x8e\x46\xa2\x84\x90\x78\x2d\x03\x29\x6a\x6f\x1c\x69\x6e\x3b\x2d\x83\x27\x8e\x46\xeb\x38\x1a\x99\xc9\x83\x69\xa6\x0c\xb9\xcd\xed\x19\xc3\x49\x1c\xb9\x3e\x2c\x37\x30\xcd\x5a\xaa\x0d\x9f\xed\xdf\x12\x35\xc7\x66\x17\xa5\x85\x92\xa5\xa8\x88\x43\x8f\x42\xa0\x51\x9d\x2e\x1c\x95\x15\x97\x07\xfe\x8c\x59\x6e\xe6\x7b\x0b\x82\x1b\xda\xce\x85\xb4\x5c\x97\xb4\xe0\xab\xf5\x62\x70\x4e\xb6\x6d\x1c\x90\x09\x70\xad\x9d\x1a\x7e\x80\x9b\x4a\x4a\x5e\x58\x8c\x7a\x1e\x27\xa1\xc5\x2e\xda\xf1\x70\x0e\x45\xc3\xb3\x43\xdd\xf1\x9d\x5a\x15\xd9\x7b\xaa\xa0\x1e\x8b\x1b\xd8\x55\x19\x4e\xdb\x52\x29\x0a\xcc\xb5\x4e\xc2\xa8\x7d\x3b\x72\xc0\xed\xfb\xf4\xe1\x1d\x56\x73\xd7\xcb\xe0\xd7\xf6\x92\x17\x9d\x77\x5b\x2e\x2f\x9d\x43\xef\x3d\x21\x2b\xc7\xd1\x8b\x4c\xae\xb7\x6a\xdc\xac\x95\xe3\x0c\x27\x2f\x91\x21\x14\x9b\xed\x13\x1f\x8d\xaf\xe2\x03\x1f\x19\x6a\xa6\x29\x0c\x54\x80\xe7\x3f\x7c\xbb\x79\xf2\xf8\xd9\xa3\x3f\x58\xfe\xf7\x27\x9f\x5d\xfc\xf5\xf3\xe6\xcb\x47\xe7\x7f\xfe\x76\xfe\xf4\x9b\xcd\xc7\xbf\x04\x04\x43\xd5\x58\xee\xd3\x8c\xd5\x42\x56\x09\xcc\x87\x22\x78\x2c\x6e\xcb\x72\x03\xf3\x45\xbe\xb4\x3c\xdc\x1d\x8f\x5c\x6b\xa5\xc3\xdd\x2d\xc3\x20\xaf\x1f\x2b\x37\xbd\x6c\x19\xf4\xdb\x4b\x3e\xe4\x94\x0d\x7a\x26\x6f\xfd\x47\x8f\x5a\x55\xe4\x16\xb5\xb4\x2e\x31\x1a\x1b\xe4\xf3\x2f\x47\x0f\xf2\x65\xe0\x9f\xc6\x7d\xd9\x50\x6d\x8e\x68\x8d\x5d\x9b\x1d\x96\x9b\xff\x59\x6e\x2b\x8f\x21\xf8\xca\xd0\xc9\xbf\xbc\x6e\x37\x16\x2e\xbe\xf8\x7a\xf3\xe9\x63\x96\x3f\x7f\xfa\x64\xf3\xf9\x77\xcf\x7e\xfa\xea\xe2\xc7\xef\xcf\x7f\xff\x35\xf0\x39\xdc\x6c\x78\xd9\x52\x07\x9a\x1d\xb8\xf6\x8c\xb9\x85\x66\xf9\x1c\xdd\xa3\xc6\x9c\x29\xcd\xd0\x82\xe0\xa0\x43\x3f\xaf\x8b\xb9\x96\x01\x42\x61\x5b\xcf\x18\x64\x80\xa6\x08\x76\x9d\x2b\x0c\xe2\xff\x91\xa6\x19\xdc\xb8\xb1\xf7\x46\xb8\xde\xcd\x3f\x9a\x80\x3a\xde\xd6\x57\xda\xa2\x45\x5f\x52\x1d\x87\x5a\x2e\x2b\x83\x10\x4d\xb0\x90\xf6\x92\x99\xed\x83\x85\xb2\xb1\xe4\xa0\xd5\x42\x5a\x4f\xe0\xd8\xbc\x63\x8b\x16\x8f\xcd\x74\xcc\x92\x74\x6c\xde\x2e\x8e\xdc\x22\xd9\xac\xb3\xe5\x9b\x68\xe2\x7b\xdd\x37\x5c\x4b\xda\x70\xb4\x98\x38\x8c\xbd\xf1\x5d\x65\xac\x37\x28\x6d\x7b\xcb\x8c\x5a\x9a\x53\xc3\xd1\x62\xf0\x54\xfa\xc6\x71\xb4\xfe\x27\x00\x00\xff\xff\xa5\x6a\xcf\x6e\xb5\x05\x00\x00")

func modelDatabaseDbGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseDbGoTmpl,
		"model/database/db.go.tmpl",
	)
}

func modelDatabaseDbGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseDbGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/db.go.tmpl", size: 1461, mode: os.FileMode(438), modTime: time.Unix(1525245701, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseTableGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x51\x6f\xda\x48\x10\x7e\x3e\x4b\xfe\x0f\xd3\x55\x54\xd9\x09\x5a\xd2\xd7\x48\x56\xa5\x26\x24\x45\x22\xe4\x0e\x38\xf5\xa1\xaa\xca\xda\x3b\x70\x4b\xed\xb5\xbb\xbb\x26\x45\x16\xff\xfd\x34\x6b\x48\xe1\x08\xa4\xb9\x7b\xb9\x97\x28\x9e\x9d\xf9\x76\xbe\xef\x9b\x1d\xba\x5d\x98\xa3\x46\x23\x1c\x4a\x48\x57\x30\x57\xba\x72\x61\x50\x89\xec\x9b\x98\x23\x34\x4d\xad\x25\x9a\x5c\x69\x04\x7e\xf3\x61\x28\x0a\x5c\xaf\xc3\x20\x0c\x54\x51\x95\xc6\x41\x14\x06\x00\xcc\xa9\x02\x99\xff\x6f\x56\x38\x16\x06\x31\x65\x2c\x85\x81\xaf\x90\x00\x1d\xf2\x61\xf9\x48\x31\xb7\xaa\x3c\x66\x55\xa1\xb9\x16\x05\xe6\xc0\x27\x22\xcd\xb1\xc5\x05\xeb\x4c\x9d\x39\x68\x08\xaa\x69\xc0\x08\x3d\x47\x38\x53\x1d\x38\x5b\xc2\x55\x02\xfc\xba\xcc\xeb\x42\x5b\xa0\x16\x60\x1f\xe7\x6c\xc9\x6f\x15\xe6\x72\xbd\x86\xa6\x39\x5b\xf2\xbb\x72\xb2\xaa\x08\x73\x2a\xd3\x2b\xe6\x43\x9b\x73\x06\x0b\x5b\xea\x7f\xc4\xa6\xd0\x34\xa8\x25\x01\x7b\x7a\xdd\xf3\x30\xb0\x62\x89\x61\x70\xde\x0d\x83\x59\xad\x33\x88\xca\x74\x01\xe7\x47\x9b\x8f\x61\x2c\x96\x18\xc5\xa0\xb4\xf3\x04\x88\xbf\xfd\x9e\x43\x02\x8c\x6d\xbf\x95\xa4\x63\xfa\x92\xc2\x89\x54\x58\x24\x5e\x77\xe8\x6e\x3e\x44\x31\x85\xd5\xcc\xdf\xc3\xfb\x12\x92\x04\x2e\xe3\x56\x0b\x68\x81\x58\x7f\x38\xee\x8d\x26\xd0\x1f\x4e\x1e\x60\xda\x34\xbb\xd7\x4f\x21\x6a\x9a\x39\x3a\x1f\xf2\xac\x28\x6e\x9f\x34\xa3\x06\x97\x22\xaf\xd1\x1e\x24\x5e\x97\xb5\x76\x7b\x99\xac\xbd\xd4\xa0\xa5\xf6\xb6\xad\xf2\xfb\xda\xba\xde\x0f\xcc\x22\xfb\x3d\xef\x78\x90\x87\x74\xd1\x56\xed\x54\x43\xbc\xad\x76\x4a\x76\xe0\x2b\x41\x18\xb4\x7c\x20\xac\xeb\x6b\x8b\xc6\xf5\x65\xb4\xc9\x51\x12\x12\x52\x24\xf2\xc9\x3e\xb8\xc6\xdc\xe2\x86\xf5\x96\x77\x5d\x49\xe1\xf0\x90\xb2\x45\x07\xbe\x8f\x3f\x7d\xc2\x61\x2b\x8f\x7f\xa1\x41\x50\x32\x79\xcf\xb6\x88\xff\x8e\x54\x67\x63\x4a\xfc\x13\xe6\x25\x76\xbf\x3d\x4f\x8f\xfe\x18\x74\xb5\xd1\xa0\xe4\xce\xb4\x49\xcc\xd1\xbd\x6a\xde\x6e\x7c\x45\xf4\x34\x23\x47\x46\xaa\x15\xf1\x2a\x81\x93\x32\xb6\xd7\x53\xbf\xef\xe0\x50\xb6\xe7\xd5\x82\x56\x94\xd8\xb3\x20\x0e\x77\xe8\x60\xa0\xac\xa3\x45\xf2\x89\x30\x76\xe8\xdc\xa1\x3b\xca\xc5\xe7\x46\x59\xa9\x25\xad\x00\xa5\xe7\x1d\x10\x66\x6e\x81\x73\xae\xb4\x43\x33\x13\x19\x36\xeb\x18\x3e\x7f\x39\xae\x47\xab\x42\x99\x2e\xbc\xbd\xa7\x32\x1b\xef\x41\xb7\x4b\x0f\xd2\xe7\x27\x50\x88\x6f\x18\x7d\xfe\xf2\x31\xbd\x41\x27\x54\xde\x81\xcb\xf8\xc4\x23\x6d\xf5\x64\xe3\xde\xa0\x77\x3d\x81\x17\x9e\x1d\xdc\x8e\x1e\xee\x0f\x35\xff\xf4\xb1\x37\xea\xed\xa8\x7e\x09\x6c\xf3\xfe\xbd\x0c\x6f\x68\x6b\xfc\x7c\xfb\x17\x09\x03\xa1\x25\x30\xb8\x68\xcf\x2f\x58\xf2\xbe\xad\xf0\x64\xd0\x98\xbd\xa1\x1e\x63\x8e\x99\x8b\xde\x12\xbd\x0e\x78\xaf\x48\x50\xce\xf9\x76\xcb\x50\xc5\x9b\x04\xb4\xca\xb7\xe3\x33\x2b\x1c\xff\xdd\x28\xed\x72\x1d\xa1\x31\x07\xd3\x4a\x58\x3b\xf3\x4a\x5e\xdf\x2a\xf3\x7a\xb3\x7d\xd1\x2f\x98\x7d\x62\xf6\x9f\xbc\x26\xd2\x6f\x5f\x30\xfa\xff\x61\xe2\x2f\xd9\xb8\x49\x19\xf4\xef\xfb\x13\x78\xc7\x9e\x73\xf6\x8f\x1a\xcd\x6a\x54\x3e\xfe\x88\xf6\x5c\xe5\x63\xff\xd3\x39\xce\x84\xa6\xbd\xf1\x9f\x5c\xf6\x26\xff\x1d\x00\x00\xff\xff\x2f\xd0\x8f\x95\x18\x08\x00\x00")

func modelDatabaseTableGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseTableGoTmpl,
		"model/database/table.go.tmpl",
	)
}

func modelDatabaseTableGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseTableGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/table.go.tmpl", size: 2072, mode: os.FileMode(438), modTime: time.Unix(1525230406, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerGen_routerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xc1\x4a\xc3\x40\x14\x85\xe1\x75\x07\xe6\x1d\x2e\xb3\x4a\x44\x66\x9e\x41\xb4\x76\xa7\x85\x16\x5c\x4f\xa6\xc7\x24\x76\x92\x09\xb7\x37\x82\x0c\x79\x77\x89\x59\x88\x8b\xda\xec\xff\xf3\x71\x9c\xa3\x1a\x3d\xd8\x0b\xa8\xfa\xa2\xba\xed\x07\xd1\x6a\xf0\xe1\xec\x6b\x10\xa7\x51\xc0\x5a\x69\xd5\x76\x43\x62\xa1\x42\x2b\x22\x22\x93\xb3\xdd\x73\xfa\x40\x90\x17\xdf\x61\x9a\x5c\x48\xbd\x70\x8a\x11\x6c\xe6\x64\x63\xea\x56\x9a\xb1\xb2\x21\x75\x2e\xfa\xea\x22\x3e\x9c\x1d\x42\x93\x8c\x56\xe5\x0c\xbe\x8f\x7d\xa0\x9c\xc7\x61\x00\x3f\xfa\x0e\x91\xec\xd1\x57\x11\x8b\xf7\x70\x3a\x15\xa0\xbb\x79\x60\xb7\xa1\x49\x65\xd6\x6a\x03\xbb\x7f\x3d\x1c\x0b\xe3\xae\xce\x0e\xfe\x13\xe6\xfe\xf7\x8b\xfd\xb7\x2c\x57\x99\x4f\x88\x90\x95\xea\xd2\x2e\xee\x6e\x3b\xb3\x3b\xc8\xd5\xfa\xad\x01\xff\x85\x6f\xd6\x6b\xe9\xe7\x96\x2f\xb2\x9a\xfe\xa9\x4b\xad\x26\xad\xbe\x03\x00\x00\xff\xff\x97\xdf\x9c\x2e\x10\x02\x00\x00")

func routerGen_routerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerGen_routerGoTmpl,
		"router/gen_router.go.tmpl",
	)
}

func routerGen_routerGoTmpl() (*asset, error) {
	bytes, err := routerGen_routerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router/gen_router.go.tmpl", size: 528, mode: os.FileMode(438), modTime: time.Unix(1525253852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerRouterGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x4d\x8b\xdb\x30\x10\x86\xcf\x16\xe8\x3f\x4c\x45\x0e\x0a\x04\xe9\x1e\x48\xa1\x84\x7e\x40\xcb\x6e\x89\xb7\xf4\xd2\x8b\x22\x4f\x15\xb3\xb2\x64\xc6\x63\x6f\x8b\xd1\x7f\x2f\x72\x0a\xbd\xb4\xf4\x22\x18\x9e\x77\x98\x47\xaf\xb5\x10\x30\x21\x39\x46\xb8\xfe\x84\xd0\xa7\x91\xa5\x18\x9d\x7f\x76\x01\x81\xf2\xcc\x48\x52\x48\xd1\x0f\x63\x26\x06\x2d\x05\x00\x80\x4a\xc8\xf6\xc6\x3c\xaa\x3a\x37\x2a\xf4\x7c\x9b\xaf\xc6\xe7\xc1\x46\x77\x9d\xd8\xf9\x67\x8b\xfe\x96\xff\x83\xed\xd0\x77\x5d\xc4\x17\x47\xa8\xa4\xd8\xd7\x3b\xdf\xe7\xe4\xe1\xb2\x9d\xd5\xfb\x55\x8a\x06\x8f\xa7\x1a\x35\x0f\xf8\xa2\xf7\x75\x36\x5f\x26\xd4\x7f\x16\xcd\xa7\x1c\x42\x0d\xff\x15\x5e\xd0\xe7\xe5\x9f\xf4\xfc\x78\x69\x37\x24\x45\x63\xed\xfd\x2c\x9c\x5e\xc3\xcd\xa5\x2e\xd6\x7f\x37\x68\xde\xbf\x7d\xd2\xca\xaa\x03\x54\x33\xed\x61\x93\x39\xe7\xc4\xf8\x83\xf7\x80\x44\x99\xa0\x7a\x36\x84\x3c\x53\x02\x6f\x5a\xa6\x3e\x05\x5d\xeb\x31\x2d\x3b\x9e\xa7\xc7\x8f\x07\x50\x1f\x30\xc6\x7c\x80\xaf\x99\x62\xf7\xea\x5b\x52\xd5\xa8\xd4\x67\x5d\x81\x5c\x0a\x08\xbb\xfe\x00\xbb\x05\x8e\x27\x30\x4f\xee\x1a\xf1\xc1\x0d\x38\x95\xb2\x45\xe6\x71\x44\x3a\xbb\x01\x23\xec\x96\x52\xde\x74\x9d\xc6\xfb\x32\xa6\x6e\xcb\x58\x0b\x2d\x3b\x62\x98\x90\x96\xdf\xf2\xf7\x6e\xcc\x3b\xc7\x2e\x6a\x34\x1b\xd7\xea\xb8\xae\xa6\x45\x5a\x7a\x8f\x9f\x33\x71\x29\xaa\x76\x50\xa4\xf8\x15\x00\x00\xff\xff\x6b\x07\x92\xc2\x0d\x02\x00\x00")

func routerRouterGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerRouterGoTmpl,
		"router/router.go.tmpl",
	)
}

func routerRouterGoTmpl() (*asset, error) {
	bytes, err := routerRouterGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router/router.go.tmpl", size: 525, mode: os.FileMode(438), modTime: time.Unix(1525252909, 0)}
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
	"controller/controller.go.tmpl": controllerControllerGoTmpl,
	"controller/gen_controller.go.tmpl": controllerGen_controllerGoTmpl,
	"main.go.tmpl": mainGoTmpl,
	"model/database/db.go.tmpl": modelDatabaseDbGoTmpl,
	"model/database/table.go.tmpl": modelDatabaseTableGoTmpl,
	"router/gen_router.go.tmpl": routerGen_routerGoTmpl,
	"router/router.go.tmpl": routerRouterGoTmpl,
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
	"controller": &bintree{nil, map[string]*bintree{
		"controller.go.tmpl": &bintree{controllerControllerGoTmpl, map[string]*bintree{}},
		"gen_controller.go.tmpl": &bintree{controllerGen_controllerGoTmpl, map[string]*bintree{}},
	}},
	"main.go.tmpl": &bintree{mainGoTmpl, map[string]*bintree{}},
	"model": &bintree{nil, map[string]*bintree{
		"database": &bintree{nil, map[string]*bintree{
			"db.go.tmpl": &bintree{modelDatabaseDbGoTmpl, map[string]*bintree{}},
			"table.go.tmpl": &bintree{modelDatabaseTableGoTmpl, map[string]*bintree{}},
		}},
	}},
	"router": &bintree{nil, map[string]*bintree{
		"gen_router.go.tmpl": &bintree{routerGen_routerGoTmpl, map[string]*bintree{}},
		"router.go.tmpl": &bintree{routerRouterGoTmpl, map[string]*bintree{}},
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
