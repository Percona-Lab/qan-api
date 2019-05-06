// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/sql/01_init.down.sql (20B)
// migrations/sql/01_init.up.sql (9.61kB)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __01_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x6b\xf9\xb4\xa3\x14\x00\x00\x00")

func _01_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initDownSql,
		"01_init.down.sql",
	)
}

func _01_initDownSql() (*asset, error) {
	bytes, err := _01_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0xb9, 0x22, 0xf, 0xca, 0x89, 0xa5, 0x5, 0xf2, 0xed, 0x48, 0xd7, 0xe0, 0x9d, 0xec, 0x2, 0xab, 0x87, 0xf3, 0x9d, 0x3d, 0x5a, 0x71, 0x39, 0xe, 0x4a, 0x88, 0xde, 0x1d, 0x0, 0x1d, 0x4e}}
	return a, nil
}

var __01_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x5a\x4d\x73\xdb\x38\x12\xbd\xfb\x57\xf4\xcd\x76\x95\x9c\xda\x9d\xec\x61\x92\xd4\x1c\x64\x8b\x4e\xb4\xa5\x0f\x47\x52\x76\xc6\x27\x06\x22\x5b\x22\xd6\x24\xc0\x00\x60\x64\xe5\xd7\x4f\x01\xe0\x07\x48\x82\x92\x92\x39\x25\x66\x3f\xbc\xc6\xeb\x6e\x00\x4d\x88\x0f\xab\x60\xbc\x09\x60\x33\xbe\x9f\x05\x30\x7d\x84\xc5\x72\x03\xc1\x5f\xd3\xf5\x66\x0d\x19\x2a\x41\x23\x79\x75\x73\x05\x00\xf0\xad\x40\x71\xa4\x31\xac\x95\xa0\x6c\x0f\x0f\xcb\xf9\x3c\x58\x6c\xe0\x3a\x21\x32\x01\xbe\xb3\x76\xd8\x51\xb6\x47\x91\x0b\xca\xd4\xf5\xc8\x8c\x8b\x43\x89\xe2\x3b\x8a\xde\xc0\x98\x67\xc8\x14\xe5\xec\x3d\x4c\x9f\x80\x0b\x48\xb8\x54\x8c\x64\xa8\xd9\x26\xf7\x60\x87\xd5\x2c\x31\x51\x64\x4b\x24\xf6\x78\x9e\xb8\x54\x7b\x81\xeb\xcf\xb3\xf7\x50\x81\x1a\xdf\x51\x82\x19\xe9\x8d\x99\x1f\x5b\xf0\x0f\xe0\x92\xd8\x31\x35\x45\x21\x51\x98\x69\x75\x49\xa2\x94\x22\x53\xa0\xed\xa0\x01\xf5\x08\x6b\x08\xb5\x9e\xa1\x41\x6d\xc5\xe5\xc8\xbb\x3b\x80\xb5\x22\x2c\x26\x22\x86\x94\x6c\x31\x95\xe6\xb9\xc0\x3c\xa5\x11\xd1\xb1\x0a\x25\xf6\x39\x17\x65\xd0\x1c\x1c\x48\xac\x12\x10\xa5\x85\x54\x9e\xf8\x3f\x94\xcf\x9d\x09\xe8\x90\xd3\x08\x43\x75\xcc\xfb\x72\x37\xfa\x21\xdf\x55\xa0\x72\x08\xb2\xef\x54\x70\xa6\x53\xd9\x1b\x11\x38\x36\xc7\x0b\xf9\xd1\x43\x8e\xbf\x13\x9a\x92\x2d\x4d\xa9\x3a\xc2\x0f\xce\x2a\xa8\xc0\xbd\x16\xd3\x85\xaf\xec\x63\x87\x93\xf1\x18\xc3\x8c\xc7\x98\xf6\xa3\xc3\x63\x04\x63\xaa\x02\xc2\x99\x22\x94\xa1\x08\xbd\x69\x7d\xa8\xcc\xd0\x4e\xcd\x43\x21\x15\xcf\xdc\xbc\x7c\xb5\xff\x7f\xf3\x82\xc7\xaf\x30\x16\x82\x1c\x6f\x2c\xd9\xad\xc3\xe6\x8e\x32\x8c\xb2\xa4\xac\x46\x7f\x27\x69\x81\x17\x8e\x37\xd8\x8a\x80\xec\x75\x95\x15\x85\x67\x51\x4e\x63\xbd\xb4\x76\x14\x85\xce\x98\x01\x82\x4a\x88\x82\x88\xa7\x29\x46\x0a\x08\x8b\x41\x22\x8b\xab\x45\x5e\x52\x96\x7f\x85\x92\x17\x22\x42\x08\x58\x91\xfd\x7e\x73\x3d\x0f\x36\xab\xe9\xc3\x3a\x5c\x2f\xbf\xac\x1e\x82\x70\xba\xf8\xdf\x78\x36\x9d\x5c\xc3\x1f\xf0\xaf\x11\x5c\xcf\x9f\xd7\x9f\x67\xe1\x7a\xb6\xfc\x73\xb6\xfc\xa8\x1f\xfe\xbb\x7e\xf8\x14\xac\x1e\xd7\x0f\x9f\x82\xf9\x58\x3f\xff\x4d\x3f\x5f\x2e\x3e\x2e\x27\xf7\xe1\xd3\x6a\xf9\x38\x9d\x05\x2b\xfd\xfc\xad\x23\x77\x6d\x1d\xf3\x5d\x35\x95\xf7\x20\x53\x7e\x48\xf9\x7e\x04\x39\x8a\x5d\xb9\x38\x47\x80\x2a\x7a\x53\x4e\x3a\x47\x41\x79\x1c\x4a\x45\x84\x82\x09\x51\xb8\xa1\x19\x3a\x75\xab\xff\x3a\x24\xc8\x2a\xf1\xba\x74\xf8\x0e\xb6\x45\xf4\x82\x0a\xcc\x30\x8c\xdb\x5c\x29\xb2\xbd\x4a\xe0\xcb\x94\xa9\xb7\xbf\x35\x54\x93\x42\x90\x6a\xb8\x43\x66\x99\x4a\x06\x67\x07\xec\xa5\x25\x3b\xca\x6f\x29\xc4\x74\x8f\x52\x85\x0a\x5f\xd5\x87\x72\xdb\x3c\x50\x95\xf0\x42\x99\x0d\xa9\x5a\x58\xaf\x24\xcb\xd3\x7e\x79\x2e\x19\x36\xdb\x6d\x05\xda\x09\x9e\xe9\x25\x0f\x3b\x5e\xb0\x18\x68\x67\x4e\x25\x2c\xdc\x71\x91\x11\x55\xa5\x35\xf8\x6b\x3c\x7f\x9a\x05\xe1\xe3\x72\x35\x1f\x6f\xba\x69\x2d\xad\x55\x42\x1f\xa7\x8b\x8f\xc1\xea\x69\x35\x5d\x6c\x4c\x2e\x9d\x9c\x4d\x59\xac\x77\x1d\x94\xed\x12\x13\x48\xd2\xf6\x34\x25\x50\x09\xb9\xe0\x09\xdd\xd2\x26\xe4\x54\x86\x4a\x14\x4c\x33\xc4\x26\xe2\xbf\xfb\xa8\xe9\xce\xc3\xa5\x38\x87\x94\xb3\xbd\x29\xe7\x03\x91\x50\x13\x75\x84\x9b\x0d\xad\x23\x7b\xf3\xfc\xd4\xab\xe5\xd5\x78\x31\x59\xce\x2b\xcd\xba\xa6\x83\xf5\xa6\xaa\xdd\xc7\xf1\x7a\x53\xfe\xf9\x76\x04\xd7\x7f\x4e\x37\x9f\xc2\x60\xb5\x5a\x9a\x22\xfe\x8f\x37\x20\x07\x1d\x90\x76\xaa\xf4\x34\x73\x1a\xbd\x60\x0c\x45\xde\x99\x66\x59\xf3\xfd\xe3\xaa\x7c\xde\xcb\x3b\x65\xf0\xdf\xf5\x72\x01\x36\xb1\xd5\x82\x60\x45\x16\x6a\x1c\x45\x19\xea\xca\x0a\x0f\x44\x30\xca\xf6\x12\x1e\x53\x4e\x5a\x15\xfd\x89\x1f\x20\x23\xec\x08\x25\xde\x4c\x4f\x8f\x81\x7a\x4c\xb7\x98\xbe\x56\x96\x37\x11\x8f\xeb\x8d\xcb\x2e\x15\x27\x08\x33\x2a\x95\x9e\x70\x85\xf6\x8c\x2e\x98\xaa\x86\x97\x13\x73\xc6\x8f\x33\x6d\xd7\x0c\x48\xa2\x13\xd3\xe9\x89\x45\x21\xb8\xf8\x29\xa9\x66\x44\x5f\xa7\x25\xba\x4c\xe5\x8c\x48\xa5\x3d\x33\xde\x1b\xfd\x33\x2a\x1b\x9a\x53\x3a\xfb\xda\x4a\x9a\xca\x4e\x19\xa8\x84\xca\x36\xc1\xdd\x5d\xdd\xd1\x99\xad\xde\xb0\x1d\x43\x45\x33\x0c\x23\xa6\xfa\xa4\x9b\x04\xf5\xee\xa8\xd0\x1c\xe0\xf8\x8a\x51\x61\xb6\x3b\x3d\x42\xbb\x90\x18\x71\x16\xdb\x38\x66\x58\x57\x5f\x8b\x58\x16\xd9\x2f\x13\x7b\x09\x33\xca\xfa\x84\xeb\x8c\xa4\x29\x4a\x65\xcf\xc7\x7a\x99\x84\x15\x61\x2b\x0e\x6d\x3a\xf2\xda\xa7\xbb\xa7\xfb\xfd\x2f\xb1\xe5\xef\xde\xf5\xd9\xde\xbd\xd3\xa7\x4a\xa4\x0f\xe4\xd4\xb0\x5d\x46\x9b\xf2\xe8\xa5\x97\x9c\xbe\x71\x30\xc0\x86\x56\x71\x20\xd1\xb7\x82\x0a\x04\x3d\x44\x7a\xa3\xdb\x90\x39\xc1\xf5\x18\x9b\x50\xf5\x8d\x8e\xf2\xca\x28\xf8\x41\x86\x52\xb7\x28\x1e\x01\x8d\x71\x50\x00\x2b\xb2\xad\xed\x5f\x34\x58\x77\x2b\x4a\xeb\x51\x09\x82\x6d\xa2\x1b\x05\x0d\x9b\x47\x81\x63\xec\x2b\x68\x8c\x43\x0a\xf4\x56\x4b\x19\xc6\x83\x2a\x6a\x80\x57\xc9\xa2\xa3\x22\x22\x8c\x61\x0c\x77\xb0\x0e\x66\xc1\xc3\xa6\xa3\xa1\xe6\x1a\xd2\xd1\x00\x06\xb4\xd4\x80\x21\x3d\x64\xb7\xc3\x48\x9d\xd0\x53\x03\x2e\xd1\x13\x25\x84\xed\x8d\x9e\x2f\x4f\x93\xf1\x26\x18\xc1\x24\x98\x05\xfa\xdf\xe9\x62\x1d\xac\xba\xfa\x6a\xee\x21\x7d\x0d\x60\x40\x5f\x0d\x18\xd2\x27\x90\x0c\x6b\x33\xc6\x4b\x2b\x4e\x83\x6d\x67\xa5\xc8\x36\x45\xd9\xd1\x62\xb8\x86\x74\x58\xe3\x80\x06\x63\xf4\xcc\x3f\x43\xb1\xc7\x30\x27\x52\xa2\xf4\x49\x68\xd9\x2f\x50\x61\xf0\x60\xf1\xb6\x37\xd3\x8b\x47\x72\xa1\x80\xa4\x7b\x2e\xa8\x4a\x32\x48\x88\x84\x84\xc4\x7a\x69\xc5\xbc\x91\xd8\xf2\xe5\x51\xd9\xb6\xf7\x85\xb6\xec\x1e\xad\x94\x31\x1e\x6f\x43\xca\x43\x11\xf2\xdc\x2b\xb7\x0b\xf1\x2a\x7e\xd0\x07\x9f\x34\xc2\x1a\xe1\x39\xd9\xa3\x4d\x1f\xcf\xd1\xb6\xec\xd2\xbc\x39\xc4\x45\x8a\x71\x23\xb2\xeb\xc1\xa3\xb3\x07\xe9\x4b\xed\x42\xce\xa8\xdd\x1e\x95\x3f\xbd\x7d\x90\x57\xf1\x9a\x66\x34\x25\x42\x27\xac\x1c\x31\x5d\x5a\xcf\x23\xd8\x16\x36\xc7\x05\xa3\x4a\x37\xc8\x86\xc6\xaf\xd7\x7a\x38\xa3\xb8\x04\x9d\xd6\x6c\x41\x67\x54\x1f\x08\xf5\x1e\x04\x3d\x8c\x5f\x73\xa2\x17\x64\xc2\x0f\xb6\xdd\xbf\x69\x8e\xb1\x5b\xa0\xfa\x58\xe0\x2f\x30\x65\x8c\x4f\xee\xed\x91\xa7\x0a\x92\xa6\x47\x5b\x02\x3a\x1e\xfa\xbd\xaa\x7c\x49\x52\x5c\x90\x3d\xfa\x63\x62\x26\x70\x26\x24\x16\x73\x3a\x22\x06\x33\x1c\x10\x81\x91\x3d\x3a\xcf\x04\xa5\x8d\xfb\xe9\xc0\x68\xe1\xe5\xeb\x25\xd1\xef\x5b\xfa\x2d\x41\xef\x6c\xb6\x15\xe8\x45\xa0\xed\x6d\x38\x0a\x1d\xdc\x60\x24\xda\xb8\xe1\x68\x7c\x2b\xb0\xc0\x73\xa1\x70\x40\xff\x20\x0e\x32\x37\x0d\x27\x55\x09\x0a\x13\x14\xfd\x96\xa5\x38\x20\x53\x28\x0c\xae\x2c\x22\xe3\x0e\xcc\x4b\x81\xa4\x31\xda\xcd\xd3\x3e\xac\x86\xe9\x60\xd6\xad\x6b\x2f\x98\xce\x7c\x87\x23\xe9\x82\x06\xc3\xe8\x80\x86\x63\xa8\x37\x3c\x19\xc6\x54\x2a\xca\xa2\x53\x71\xec\x00\x4f\x6d\xa9\x24\xcf\x05\x7f\xa5\x19\x51\x98\x1e\x3b\x1b\x6c\xc1\xe8\xb7\x02\xcd\x3e\x2b\x9d\xf8\x92\x28\x42\x29\x3d\x3b\x6c\xc7\xed\x70\x48\xba\xc0\xc1\xb0\x74\x80\x9e\xd0\xd8\xee\xda\xde\xe4\xf8\x22\xd2\xb2\x5f\x52\x54\x8d\x4e\xda\x7b\x37\x29\x69\x3c\xc2\xda\xf6\xbe\x9e\x96\xdd\x23\xa3\x3c\x09\x06\xfa\x68\xc7\x7a\x41\x43\x60\xd0\x75\x27\x4d\xd2\xb4\xec\xa4\x1d\x39\x0e\xa1\x47\x8c\x6b\xed\x4b\x71\xac\x1e\x21\x2a\xcb\x43\xdb\x48\xf9\x84\x38\xd6\x33\x7d\xa7\xc2\x2c\xe7\x82\x88\x63\xd9\x96\x41\x24\xd0\xdc\x1d\x71\x06\x19\x66\x5c\x1c\xcd\xda\xac\xb3\xd5\x68\x73\x7c\x78\xb4\xb9\xd6\xbe\x36\xc7\x3a\xa0\x2d\xa6\xf2\xe5\x8c\x40\x17\xf2\x0f\x54\x6a\x9a\x53\x1a\x5d\x37\x03\x42\x5b\x10\xbf\x5a\x17\x72\x2a\x9d\xa1\xa4\x3f\xce\xe4\xb4\x84\xf8\x2b\x94\x2b\x92\xc2\x9a\xfe\xb0\x2f\xc1\xa6\x42\xb5\x36\x5d\x9d\xbd\x20\x14\x12\x63\x7b\xbb\x31\x9c\xdc\xd2\xd9\xa9\x0c\x57\x90\x13\x69\x2e\x21\x3d\xe1\x77\x77\x70\xcf\x79\x8a\x84\x75\x6f\x53\xa2\x30\xf1\x1f\x61\xa5\xc5\xab\xfe\xb3\xd9\x4f\x1e\x48\x94\x20\x24\xd4\x5d\x88\xbb\x22\x4d\x43\xfd\xba\xe8\xa3\x6c\x8c\x83\xab\xde\xee\x54\x39\x8a\x1d\x17\x19\xc6\x40\x40\x0f\xb2\x81\x34\xaf\xa1\x1d\x57\xff\xe7\x74\xd8\x95\x31\xfe\xac\x2b\x3d\x08\x6e\x88\xfd\xb7\xba\xe1\xa6\x2c\xc6\x57\x94\xb7\xbe\xc4\x9d\x2e\xa1\xd3\xde\xab\xe5\x41\x18\xd0\x2c\x4f\x69\xa4\x1b\x60\x7d\xac\x33\xd2\xab\x23\x9f\x6f\xce\x6c\xc1\x9f\x9c\x43\x05\x3a\x3d\x17\xd9\xf5\x67\x6e\xc8\x74\xe3\xd9\x2c\x5f\x27\xf8\x34\x45\xfd\x4a\xe6\x8d\x7d\x65\x3b\x23\xde\x2c\x0b\x02\x15\xdc\x43\x7e\x42\x5e\x0f\x33\xe8\xac\x42\xda\x3b\xec\x3a\xdd\x3d\x49\x12\x53\x8c\x94\xad\x1c\x41\xd8\x1e\x07\x8b\x6b\x00\x79\xc1\x51\xa6\x71\xe5\x4b\x6d\xa9\xde\x8c\x07\x89\x44\x44\x89\x9e\x13\x01\x81\x3b\x14\xc8\x22\xec\xa6\xbd\x74\x6b\x3d\x0e\x4f\xcb\xda\x7f\x61\x32\x66\xa0\xd4\x93\x50\x26\x6c\x42\xaa\xd3\x53\x48\x30\xf2\xe6\xc6\x83\xba\x78\x3a\xd5\x8a\x7b\xc1\x63\xf5\xcb\x8c\x26\x30\x9b\xeb\x0b\xea\xa2\xd1\x6f\xc9\x64\xa7\x5b\x5f\x73\xf9\x2c\xf8\xc1\x99\x9f\xae\x88\xe1\x00\x35\xd6\x0b\xe6\xa3\xd1\xe5\x14\x0e\x28\x10\x62\xce\x10\x0a\xa9\xbb\x68\x1b\xa9\xae\x5b\x7e\xf0\x9e\x27\x8d\xf1\x42\xa7\x3a\x15\xfc\xd0\xa5\x1f\xda\x54\x1b\xe3\xaf\x6a\xda\x1e\xed\x05\x1f\x2d\xbb\xc5\x4e\xce\x19\x0f\xcd\xf6\x17\xea\x22\xf1\xcd\xa0\x0d\xb8\x60\x16\xf5\x4f\x19\xee\xee\xda\x72\xb8\xe7\x3c\x3e\xef\xb5\x8b\xfa\x05\xd7\x9a\xa2\xed\xff\xee\x0e\x32\xce\xf6\xbc\x73\x50\xc6\x3c\x92\xa1\x40\x55\x88\x81\x5b\xd5\x36\xe0\x92\xdb\xba\x12\x0b\x31\x8f\x8a\xac\xdd\xcf\xb6\xc9\x3c\x4d\x41\x07\xd0\x6f\x09\xda\x00\xdf\xb5\x23\xca\x9c\x33\x89\x27\xde\x35\xba\x90\x41\x51\x15\x10\xca\x5f\xa0\x75\x27\x58\x6f\xf3\x02\x65\x91\xaa\xba\x4d\x72\xee\x23\x3b\xf4\xbe\x5b\xc9\x2e\xc4\x73\x37\xd9\x81\x78\xa4\x9a\x58\x94\x77\xd8\x83\xa9\xab\xec\x97\xac\xa2\xf2\x3a\x7c\x28\x71\x15\xd5\x50\xde\x6a\xfb\x40\xda\x2a\xbb\x23\xe5\xea\xf6\x2a\x58\x7c\x9c\x2e\x02\xf8\x03\xe6\x28\xf6\xb8\x11\x88\x57\x4f\xe3\xd5\x66\xba\x99\x2e\x17\x70\xff\x0c\x8a\x3f\x3f\x3f\x3f\xcf\xe7\x93\xc9\x8d\xfb\x6d\xc1\xed\xd5\x72\x35\x09\x56\x1a\x71\x53\x7e\x0d\x35\xaa\x3f\x6f\x1a\x39\x9f\x28\x8d\xea\x0f\x8f\x46\xce\xf7\x43\xa3\xf6\x97\x41\xa3\xd6\x77\x0b\xb7\x1f\xae\xfe\x0e\x00\x00\xff\xff\x73\x4a\x59\x21\x8a\x25\x00\x00")

func _01_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initUpSql,
		"01_init.up.sql",
	)
}

func _01_initUpSql() (*asset, error) {
	bytes, err := _01_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9d, 0xcf, 0x18, 0x5e, 0xc3, 0xec, 0x18, 0x4, 0x7f, 0x42, 0x60, 0xa1, 0xd6, 0x9, 0x66, 0x7f, 0x61, 0xbf, 0x78, 0xfd, 0x3d, 0xe9, 0x5, 0x17, 0x15, 0xec, 0x46, 0x8d, 0xaa, 0xbf, 0x3f, 0x2d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"01_init.down.sql": _01_initDownSql,

	"01_init.up.sql": _01_initUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"01_init.down.sql": &bintree{_01_initDownSql, map[string]*bintree{}},
	"01_init.up.sql":   &bintree{_01_initUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
