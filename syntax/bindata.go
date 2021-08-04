// Code generated by go-bindata. DO NOT EDIT.
// sources:
// syntax/implicit_import.arrai (921B)
// syntax/stdlib-safe.arraiz (3.464kB)
// syntax/stdlib-unsafe.arraiz (1.419kB)

package syntax

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _syntaxImplicit_importArrai = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\x41\x8f\xdb\x2e\x10\xc5\xef\xfe\x14\x23\x65\xb5\x06\xc9\x7f\x7c\x27\xca\xfe\x73\xaa\xd4\x7b\x8f\x2b\xed\x52\x33\x5e\x53\x61\x70\x81\xa4\xb1\x5c\x7f\xf7\x6a\x70\x62\xd7\xad\xea\x93\x19\x7e\x0f\xde\x3c\xe0\x00\x5f\x3a\x13\x21\x36\xc1\x0c\x09\x3a\xe5\xb4\xc5\x08\xa9\x43\xe8\xd5\x30\x18\xf7\x01\xbe\x05\x8d\x8d\xd7\x18\x22\x18\x07\xa6\x1f\xac\x69\x4c\xa2\x1f\x1f\x52\x04\x8d\x03\x3a\x9d\x49\x57\x1c\xb2\x14\x6f\x09\x5d\x34\xde\x91\x98\x0a\x0b\x8b\x1a\x5a\x63\x51\x14\xc5\x61\x5b\xb2\x57\x43\xcc\x65\x68\x7d\xe8\x55\x8a\x90\x7c\xd6\xc4\xa4\x9c\x56\x41\x6f\xe8\x25\xa2\xa6\xd9\xa5\x90\x45\xd9\x52\xea\x54\xba\xab\x45\x61\x31\x6d\x8a\x13\x4c\x05\x00\x40\xd9\xc4\x6b\x29\xa1\xae\xd1\x35\x9e\xbc\x8a\x26\x5e\xc5\x82\x55\x0b\xf1\x2d\x7a\xb7\x47\xa8\xb2\x67\x6e\x36\xde\xf6\x0c\x55\xf6\xcc\xa8\x7a\x5b\x4a\x60\x77\x0f\x3b\x9a\xe6\x1e\x34\x28\x6b\x54\x94\x30\x95\x63\x6f\xcb\x99\x57\xc5\x7c\xa4\x60\x1a\xef\xae\x48\xb9\xe2\x15\xc3\x98\x3a\x0a\x36\x79\x50\xa0\x4d\x93\x8c\x77\x2a\x8c\x30\xad\x01\xcb\x47\xaf\x9f\xdc\xfc\x67\xeb\x2c\x1b\x5a\x0b\xad\xb1\x09\x03\x88\xf3\x55\xd9\x0b\xde\x83\xa1\x8f\xc9\x3b\x53\x81\xcc\xa6\xb8\x84\x89\x49\x71\xae\x60\x61\xd7\x4d\xf8\x0c\x3f\x81\x65\x06\x4e\x2f\xc0\xce\x52\xfc\xcd\xf0\x6a\x5d\xf9\x4d\xc2\x24\xe6\x65\x3c\x17\x1c\xfe\x7b\x81\xba\x0e\x68\xc5\xc5\x19\xef\x98\xe0\xc7\xa2\x78\xdd\x2e\xcb\xeb\xd7\x31\x61\x86\xa9\x93\xad\x7e\xa2\x50\x34\x4c\x75\x1d\xf1\xbb\xe8\x54\x7c\x1b\x02\xb6\xe6\xc6\x4a\x51\x56\x1b\xc7\x29\x6b\x22\x52\x30\xfd\x3f\x90\x8a\x3c\xad\xc3\xf9\xb8\x8b\x88\x6d\xdc\xff\x8c\xbc\x70\xc9\xd6\x5e\xea\xda\xfa\x0f\x31\x04\xe3\x12\x7b\x7a\x5f\x5f\x41\xd6\xd2\x19\xb5\xca\x58\xba\xe0\x3e\xfc\xe6\xfc\x69\x3b\xa9\x59\xc0\xe7\xfc\x0a\x08\xfe\x61\x52\x07\xb4\x43\x7c\x6c\xfe\xce\xe1\xf9\x19\xd6\x00\x78\xf1\x2b\x00\x00\xff\xff\x39\x95\x6f\x16\x99\x03\x00\x00")

func syntaxImplicit_importArraiBytes() ([]byte, error) {
	return bindataRead(
		_syntaxImplicit_importArrai,
		"syntax/implicit_import.arrai",
	)
}

func syntaxImplicit_importArrai() (*asset, error) {
	bytes, err := syntaxImplicit_importArraiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syntax/implicit_import.arrai", size: 921, mode: os.FileMode(0644), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0x61, 0xe5, 0x91, 0x8e, 0xa9, 0x63, 0x16, 0x96, 0x63, 0xc7, 0xb9, 0x92, 0x57, 0x54, 0xf6, 0x8, 0xc, 0xca, 0xf0, 0xea, 0x12, 0x5a, 0xfa, 0x4a, 0xb8, 0xfc, 0xa8, 0x37, 0x95, 0x6f, 0x1f}}
	return a, nil
}

var _syntaxStdlibSafeArraiz = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x79\x3c\xd4\x7b\xbf\xff\xd9\x86\xb1\x1c\x84\x44\x69\xd1\xb1\x86\x11\x65\xcb\x9e\x90\x21\xbb\xec\x0e\x66\xcc\x8c\x86\x91\x19\x59\x87\x09\x83\x08\x39\xd9\x65\x49\xf6\x24\xb2\x1d\x63\x22\x4b\xc8\x30\x23\x3b\x91\xb5\xac\x35\x3c\x9a\x2c\x59\xee\xeb\x79\xce\x3d\xf7\x76\xba\xaf\xfb\xdc\x7b\xcf\x7d\xbe\xff\x7c\x5f\xdf\x3f\x3e\xfb\xfb\xf3\xfe\xbe\x2d\xa0\x2c\xac\x82\x00\x07\xc0\x01\xfc\x78\xb8\x01\x00\xf0\xc4\xf8\x7a\xa1\x10\x0a\xee\xfe\xfe\xee\xa8\xc1\x1e\xb3\xbe\xbe\x04\x48\x2f\x0d\x52\x63\x65\x61\x42\x93\x7b\x23\x6d\x36\x08\xa5\xd2\x46\xcc\xe5\x64\x07\xcd\xfa\x58\x2d\xef\x65\x99\x68\x70\xfe\x62\x8b\xca\x52\x2f\xb7\xe3\x34\xd9\x6b\xda\x95\x29\x56\x2f\xdf\x53\xb3\x59\x13\xbe\xf9\x31\xed\x67\xbb\x65\xe1\x8c\xa2\x45\x98\xa4\x9e\xcc\xf4\x7c\x67\x9d\x8d\x56\x1c\x2a\x9b\x00\x00\xc0\xd1\x91\x05\x94\x9d\x03\x3f\xf4\x30\xf6\x26\x00\x00\x38\x00\x00\xfe\xfb\x84\x24\x00\x00\xf0\xc1\xc0\x02\xd0\x70\x08\x02\x85\x43\x06\x78\x28\x78\x62\x7c\x20\xee\xfe\xfe\xf2\xee\x28\xc8\x3f\x52\x84\x20\x30\x0a\x3e\x18\x58\x58\x6a\x3f\xe9\x89\x0a\x77\xcc\x58\xbf\x6d\x35\x5a\x78\xc7\xa8\x6b\x5f\x6e\xd2\x25\xdf\xf5\x95\xf3\xdd\x8b\x91\x41\xfa\x8a\x69\x66\x92\x8b\xab\x4b\xe7\xfa\xfa\xee\xa9\x63\xa9\xbd\x5c\xf8\xbd\x7d\x93\x17\x01\x8c\x69\x25\x0f\x95\xdf\x82\x5c\x1d\xbd\x0e\xd2\x4e\x16\x4e\xfd\xf6\x49\xf1\x33\x4e\x34\x4c\x32\xcb\xa6\x54\x60\x4e\x53\x34\x47\xba\x63\x0a\xfa\xda\x2c\x79\xb1\x36\x07\x16\xbf\xb2\x4e\x8f\x38\xf6\x74\x25\xa5\xd1\x8b\x44\x46\x11\x84\x72\x46\x8e\x12\xce\xf7\x1d\x6a\x1e\x6c\x8a\xdb\x15\xef\x04\xc6\xbb\x0d\x1c\x1a\x85\x67\x6d\x41\x32\x02\x2e\x75\x6e\xfa\xde\xc6\xf1\x2e\xaa\xf6\xf8\xbe\x12\x14\xb8\xfe\xab\xc8\xdb\xae\xda\xd6\xec\x57\x60\x44\xaf\x6f\x24\x58\xd2\x9f\xef\x6c\x34\xb3\x7c\x3c\x73\x41\x2f\x97\x83\xf0\x31\x03\x37\x3b\x62\xbf\x1a\x6d\x6a\x58\x5d\xff\xa4\xf5\xd3\xfc\xea\x2d\xf8\x38\x1e\xb0\x6a\x7b\x9f\x99\xb7\xcf\x52\x6a\x76\x2e\xe7\x7c\x7f\x95\x5a\x8e\xdf\xfb\x79\x06\xbd\x43\x79\x77\x79\xc6\x18\x2e\xcf\x09\x58\x0a\xa5\x8d\xaa\x7f\xc1\x27\x88\x35\xc4\xed\x35\x36\xcd\x4c\x96\xcf\x58\xd9\xde\xde\x68\x58\xea\x8e\xef\x04\x9b\xa9\x92\xc0\x02\xdf\xac\xa4\xcb\x0b\xa6\xd6\x73\x38\x1d\xa3\x0b\x2f\x78\x9b\x0e\xb2\xe7\x9f\x4f\x29\x7d\xb5\xed\x99\x93\x06\xfb\x5b\x81\x48\xff\xa5\xf5\x9d\x4f\x83\x98\x63\xfd\xb3\x1d\xe1\x3f\xe5\x4d\xa9\x89\x1e\xf2\xbc\x9b\xfd\xc4\xb6\x90\x6e\x51\x5d\x57\x4a\x35\x1d\x4d\x0e\xba\x2c\xc0\x9b\x52\x09\x36\xf7\x7f\x63\x53\x48\xca\x3e\x51\x9a\xf8\x96\x4d\x60\x81\x24\x2c\x65\x5d\xba\xab\xb4\xef\x21\x5c\xa9\x5e\x3d\x74\xb3\x8e\x57\x99\xc4\x95\x1f\x86\x3b\x20\xdd\xb2\x62\x77\xd4\x3f\x7b\x86\x19\x63\xe4\x09\xd1\x6a\xa2\x56\xea\xe4\xe5\xc2\xe6\x3b\x34\x58\xba\xb6\xfa\x07\x66\x55\x75\xdb\x8d\xd7\x5c\xf5\x35\xf2\x67\x91\x59\x2c\x37\xc2\xbe\xb0\xc8\xd3\x0e\x47\xd9\xbb\xb6\x67\x64\x78\x8c\xda\x91\xb4\x4e\xfc\x37\xc6\x61\x18\x4b\xbe\xde\x81\xa0\x11\x77\xe8\xcb\x5f\x6d\xd9\x8b\x88\x21\xd8\xab\xbe\x46\xdb\x4f\x4c\x82\x38\xf3\xac\xbd\xf2\x40\xe1\x4b\xe9\xa3\xbb\x8c\x67\x91\x7c\x97\x39\x1e\xbc\x38\xa6\x31\xd6\xc0\x6c\x3d\x56\x51\x5e\xa4\x4e\x20\x87\x98\xbe\x5b\x8c\x10\x31\x7e\xac\x97\x60\x75\x57\x01\x3b\x34\xa9\x7f\xe1\x01\x4f\xc3\x9e\x00\x0b\x26\x3d\x28\x86\xb5\x90\x1f\xb0\x12\xb2\x58\x22\xb1\x4c\x8b\xa0\x56\xab\x2c\xf2\x83\xa2\x2f\xff\xcc\x57\x7c\x4c\x1f\x69\x21\x2e\x8c\x1f\x3a\x29\x6d\xe6\x76\xb3\x7b\x7a\xc3\x9b\x6f\xf3\xf4\x76\x49\x86\x53\x27\xef\xf9\x17\xc2\x6d\x84\xa1\x39\x27\x4c\x06\xba\x7f\xd0\xd1\x26\xd8\xe9\xce\xc0\x5a\x83\xb0\xcb\x88\x18\x24\x18\x71\xad\x16\x4f\x8d\x70\x3b\x43\x9f\x7f\x7e\xc4\xdb\xeb\xa2\xec\x06\xfa\x28\xd9\x92\xbd\xd6\xe3\x2a\x0c\x17\xb5\x98\x30\x7b\xf9\xf6\xba\xe5\x1b\x90\x28\x81\x8c\x34\x71\x41\x5a\xcb\x5a\xd6\x83\xa7\xdd\x4a\xe9\x4c\x7f\x60\xdc\xf4\x73\xeb\xa3\x14\x66\x00\x38\x03\xfa\x67\x18\x57\xfd\x9f\x31\x8e\x0d\xf6\xc5\xb9\x07\x41\xb0\x38\x18\x1a\xe5\x01\xf1\x42\xbb\xff\xfb\x7e\x56\xda\xf7\x63\x0b\x8d\x05\x5b\x69\x6e\x9b\x25\x31\x9c\xfd\xf1\xb9\xa1\x88\x61\x1b\x9f\x3b\x49\x0f\x29\xa3\xae\x1e\x76\x8e\xa8\x34\x59\xd6\xa7\x4c\xca\x32\x31\x47\xf8\x35\x06\x7b\x01\x21\xf3\x19\x39\x4a\x9c\x1c\xee\xd6\xba\x4d\xa4\xc8\xf1\x2e\x20\x0f\xd8\x80\xb3\xc7\xd5\xf9\x24\x89\xbc\xa6\xc6\x59\x26\xfb\xda\x9b\xe3\xec\x26\xe5\xec\xca\x03\x95\xbc\x97\x31\x3d\xb5\x2f\x7b\x7e\x01\x59\x0d\x98\x85\xa8\x19\xd6\x9c\x6d\x6b\x1d\x5a\x11\xdf\x10\x2d\x71\x36\x8c\x95\xa2\x94\x99\x0f\x54\xda\x0b\x8c\x5d\x36\x9d\xd1\x8d\xcf\xdd\x61\x73\xca\x3a\x4d\x80\xbe\x39\xbf\xd7\x70\xa3\xa1\x9e\x20\xfd\xcd\xe1\xd4\xa8\xbe\xeb\xaa\xc3\x0d\x61\x58\x5b\xfa\x1c\xab\xce\x8a\xd2\x7e\x87\xd3\x7e\x17\xfd\x9e\xd1\xf1\xf5\x0f\x90\xcd\x3d\xa5\x5c\xb2\xc2\xc6\x76\x3b\x58\x33\x64\x76\xd7\x6b\x46\x69\xed\xa0\x4c\x4e\x76\x5a\x2e\xc3\xf2\x05\x16\xef\x37\x86\x12\x5d\x08\x2c\xfc\xd2\x94\xf6\x51\x72\x55\x00\x05\x99\x3f\x1f\xf5\x41\xd2\xac\xe1\x6d\x55\xcb\x2d\x70\xf1\xaa\x91\x62\xad\x70\x39\x37\xd1\xb9\x95\x05\x1f\xbd\x7b\x2a\x7d\x31\xa5\x4d\xe5\x63\x73\xfe\xfe\xd9\xb0\x20\xcd\xcd\xdf\xf8\xf6\xa8\xf9\xf1\x33\x6b\x1c\xd7\x26\x8b\xae\x8f\x7b\x44\xbe\xc8\xe0\xba\xa7\x88\xa3\xc4\xb2\x5d\x50\x6a\x11\xe8\x0b\xd8\xcc\xee\xe6\xa7\xf5\x05\x48\x37\x90\xe3\xb0\x0c\x2e\x4b\xb8\xae\xbd\x67\x8a\xf6\xc7\x28\xed\xf9\x28\x7e\x84\x79\xf0\x61\x81\x17\x8c\xfc\xa0\x98\xe6\xfd\x81\xb0\x57\x67\xb6\x3a\x98\xcd\x54\x75\xe5\x6c\xdd\xea\x14\x7e\x79\xea\x76\xe4\xfb\x27\xb4\x6e\x4d\x93\x87\xf7\x79\x44\xd7\x52\x41\x79\xbb\x61\xc6\x31\xc5\x2b\xc3\x32\xe5\x01\x15\xfe\x26\x15\xfb\x77\x9b\xaf\xf8\x31\x9f\x29\xf9\x6a\x28\x34\xab\xfc\xe5\x44\x0b\xd5\x54\xfc\xdc\x43\x2e\x05\x53\xc9\x82\x3c\x1b\xaa\xa2\x5f\x7a\x19\xd8\x74\xad\x52\xe0\x81\x86\xe3\x0d\x94\x88\x2c\x9d\x34\x3d\xbd\xc1\x48\xb1\xb6\x62\xeb\x0d\xf8\x4c\xa8\x81\x0f\xc3\xc1\xa2\xb3\x49\xce\x5e\xfa\x5f\x33\x40\xe9\xaf\x4d\x6e\x1a\x20\x29\x85\x72\x32\xc5\x69\x37\x88\x42\x2c\x23\xe8\x38\x81\xcc\xd2\x58\x71\x8a\xbb\x3c\xe1\xea\x09\x1f\x3e\x13\x85\x19\xbc\xf1\x20\x19\x54\xc5\x49\x5e\x7d\xfc\x04\x9e\xf3\x8a\x3b\x5d\xea\x74\x87\x3f\x24\x68\xdc\xa9\xa0\x5e\xa2\x87\x49\xe2\xf8\x62\x45\x79\x67\xda\x76\x96\x8a\x43\xae\x06\x9c\xfa\xe5\x6b\xd7\xc9\xc4\xeb\x54\x0d\x96\x1a\x62\x5c\xf1\x2a\x8f\x54\x9d\x77\xd7\xeb\x8b\x31\x22\x8f\x82\x2a\xad\x7f\x2d\x33\xc7\xb4\x35\x68\x52\x0d\xc6\x94\xa0\xbd\x15\x6b\x59\xf7\x0a\x73\xe2\x90\x81\xce\xa5\xc3\x59\x74\xef\x23\x03\x50\xac\x38\xf5\x89\xa7\x9c\xcd\xf3\xf6\xec\x40\x13\xcd\x7b\xa7\xae\xc7\xb5\x0e\x34\xf5\xf2\x54\xd5\xf1\x17\x09\x45\x0d\xe2\x78\x02\x7e\xc9\x6c\xd6\xf1\xf9\xb6\x6b\xe9\x9d\x5c\xc6\xd4\x93\xed\xfb\x2d\xe2\xae\x72\x80\xe3\xb6\x5e\xe2\xd9\xb2\x13\xd5\x24\x4b\x55\x74\x68\x9e\xd6\x84\x12\x0d\x7e\x8e\xcc\xd4\xfa\x0d\x97\x8b\x42\x2e\xeb\xdf\x74\x8a\xea\xbd\xf7\xce\xb6\x00\x48\x4e\x2c\xe6\xb4\x3c\xa0\x37\x7b\xd1\x48\xd4\xf7\x68\x00\x49\x46\xb7\xcb\xd4\x34\x00\x43\x06\xba\x24\xb6\xd2\x61\x16\xf2\xe9\xea\xb4\x67\x86\x23\xde\xb1\xfe\x68\x5d\x77\xa8\x88\x38\xe6\x9d\x4d\xea\xa3\xa0\x53\xc4\xf1\x40\xfe\x03\xd7\xf2\x91\x8e\x0c\xaa\x96\xc7\x60\x63\xd5\xb5\xcc\xbe\x36\x1b\xff\x0b\xb5\x97\xae\x5e\x19\x91\xdc\xfa\xba\xb5\xe5\x53\xf0\x8e\xbe\x0a\x5f\x38\x6c\x29\xda\xbe\x76\xfa\xb5\x50\x9a\xb8\x9c\x4a\x24\xce\x74\x52\x32\xdc\xc7\xce\xc1\x25\xd3\xd2\x2e\x6f\x66\xc2\x9e\xea\xc2\x50\x98\xda\x07\x45\xe8\xe8\x68\x29\xd7\x1f\x6b\xd3\x83\xa9\x5d\xd3\x00\x55\xd4\x10\xb4\xb2\x28\x2e\x2a\x56\x88\x1e\xaf\x25\x34\xc4\xf0\xfc\x1b\x84\x0a\xb4\xfe\x03\xff\x85\x65\x13\x4a\xf4\x68\x48\x08\xdf\xb9\x3e\xc7\x04\x6f\x6c\xf8\x6f\x3e\x93\x5b\x99\xf3\x8c\x16\x84\xc6\xf3\x94\x15\xf9\xcc\x26\x92\x42\x89\x1d\x82\x72\x03\x62\x6d\x08\xe3\x92\x89\x4c\x7e\x49\x9c\x94\xcc\x5c\xfb\x8a\x49\x50\x1c\x67\xd6\x43\x84\x40\xc5\xa8\x1a\xcf\xd6\xaf\x86\xda\x43\x52\xad\xd0\x97\x66\x75\xaa\x23\xb9\x97\xbd\x4f\xf1\x76\x94\xdb\xc7\x5c\xc4\xeb\x41\x6d\x3d\x29\xc1\x82\xad\x49\x51\x13\x13\x3a\x9b\xcf\x8a\xc1\x05\x7f\xdb\x07\x99\xce\x14\x86\x25\xbb\xee\x47\x4d\xd3\xd6\x05\x89\xaf\xbc\x5f\x73\x4a\x78\x91\x27\xe6\x6d\xe1\x4e\x2a\x7a\xb6\x7d\x23\x71\xa2\x1f\x5a\xbc\xd1\x0b\xd0\xc7\xae\x52\x0e\x02\x2e\x91\x73\x1b\xb4\xfe\x9f\xbd\x91\xf7\xef\xdb\x49\xf4\x93\xe9\xce\xc9\x25\x12\x3e\x52\xc9\xf3\x27\xd7\xeb\xbe\x14\x8a\x17\xf8\x3d\x14\xa7\x18\xf9\x35\x17\x75\x41\x35\xc3\x56\x8a\x8f\x7d\xe2\x41\x1c\x4a\x24\x76\x77\x58\xdd\x02\x2d\x73\x94\x77\xde\x55\x14\xca\xca\x96\x0b\x0d\x6f\x67\xd0\x4c\x0d\xc6\xca\xf9\xcf\xcc\x20\xe5\x5c\x44\x02\xd3\xd7\x9b\xa5\xbe\xf9\x2d\x5e\xb1\x3c\xb3\x96\x99\x01\x9b\xa9\x35\xbb\xff\x6e\xf7\x4e\x2d\x77\x91\xf6\x38\xeb\x71\x4a\x79\xe2\x6c\x90\xb4\x1c\x9e\x50\x4f\x7c\xac\x30\x61\xec\xd9\xad\x17\x6b\x3b\x57\xb2\xb6\xa5\xac\x17\x85\x1d\x71\x1c\x5b\x2a\xa1\xf1\xd8\x55\x92\xc4\xda\xef\xc8\xad\xd7\x7a\x89\xb5\x25\x5d\xc9\x5a\x1e\xdd\xd3\x00\x75\x6c\x84\xce\x8d\x80\x95\x87\xb7\x30\xc1\xa5\xcf\xb7\x41\xa7\xde\x24\xea\x20\x1b\x31\x62\xc1\xc1\x8c\x4f\x57\x0c\x1d\x82\xb5\x72\x1b\x74\xdc\xb3\xd2\x9f\x7a\x0c\x60\x5a\xf9\xbd\x77\x70\xd8\x23\xad\xe5\x15\xd5\x59\x77\x8e\xdd\xf7\xe4\x8d\xb2\xd4\xc6\x23\x34\x6e\x4c\x0c\xba\x5c\xb0\x23\x04\x0f\x4e\x5d\x0f\x2f\xc3\xd6\x62\x89\xb5\xea\x03\xdd\x73\x2b\x7a\xa4\x50\x26\xf2\x27\xa2\xa0\x2e\x83\x74\x23\x47\xeb\xe4\xe7\xcf\x04\xed\xf6\x31\xfa\xba\x73\x31\xfb\x56\xf0\x81\xdb\x55\xa3\xa1\x0f\xef\x83\xbf\x3e\x5c\x96\xf6\x3e\x4d\xa7\x44\xfc\x07\x09\x77\x10\x72\x65\xcf\xb2\x01\x80\x8e\xe0\x3f\x23\x61\xed\xff\x2b\x09\xff\x7e\xc9\x63\xdd\xbd\xe0\xbf\x73\x71\xf4\x03\x73\x5a\x8e\x22\x77\x17\x3d\x2c\xde\x28\x04\x52\x60\xb7\x92\x8b\x1a\xa9\x6a\xa6\xf7\xbc\xe6\x63\x79\x3b\xc8\x2a\x51\x12\x51\xe7\x9b\x8f\xe7\xc4\x09\x3a\xa2\x9e\xf3\xa6\x2c\x2c\xed\xbe\xbe\xa3\xbd\xf1\x44\x95\xbf\x06\xec\xc5\xc7\xdd\x96\x9f\xf8\x91\xa3\x9d\xe8\x8b\xc4\xeb\x7e\x91\x32\xbc\x42\x32\xbc\x1f\xc8\xdf\xa5\x17\x62\x28\xe8\x9f\x1d\x19\x73\x39\xa0\xc0\x2f\x2a\x68\xf8\x3e\x54\x51\x7c\xfd\xd7\x04\x7d\x0b\x98\x60\x0b\x57\x74\xb8\xff\xd4\x2f\x13\xdd\xcc\x55\x0f\xa8\x4e\xfe\x62\xd2\xf5\x19\xb7\xab\xb9\xfd\x45\xef\xc2\xb9\x15\x22\x1f\x65\xe4\x75\x5c\x6d\x3a\xb1\xa2\xe7\x98\x73\xe9\xf8\xc4\x48\x34\x5d\x19\x66\x2b\xe8\x63\x96\xcb\x44\x80\xeb\x1d\xf6\xab\xcd\xcd\xf6\xa4\xef\x64\x7c\x7a\x9c\x2a\xbb\x21\x54\xbb\x4d\x81\xe5\x61\x19\x4b\x2d\x40\xe0\x7a\x52\x8f\xd4\xc5\x6b\xb7\xca\x2f\xdc\x56\xdb\xb7\x8a\x4d\x3d\xa4\x22\x0a\xc7\x21\x3f\xdd\x49\x1c\xeb\x0a\x68\xf6\x6c\x9a\xb7\x72\x36\xff\x5a\x9f\xaa\x8f\x5d\xb8\x1c\x7b\xcc\x76\xc1\xe5\x96\xc6\x22\xb6\xd8\xc0\x76\x26\x3a\xe3\xee\x08\x03\x36\x22\xf7\xa1\x16\x9b\x0c\x93\xb6\xeb\x6f\xc4\x1c\xd2\x89\xe4\xa0\x01\xfa\x46\x9d\x5e\x64\xe2\x49\xc3\x89\xc1\xda\x74\x73\x93\x9a\x73\xf9\x8d\x0c\xc4\xa3\x37\x2d\x11\x06\x15\x49\x3f\x87\xbf\xc4\x56\xa4\x95\xed\x68\x44\x23\xba\x6d\x1e\x25\xc1\xe7\x8a\x5a\x15\xf6\xc1\xd9\x2a\xb9\xd5\x4b\x84\x09\xdc\x3e\xcf\x1f\x63\x6b\xdc\xd6\xb7\x84\x32\x01\x00\x17\xcb\xbf\xf4\xef\x0c\xc0\xa1\xd0\xbf\xcf\x0b\x9d\xd4\x16\x1f\xad\xc8\x67\x40\x3f\x10\x62\x0b\xb5\x7c\xb0\xee\x32\x01\x71\xd0\xaa\xcb\x65\x65\x7e\xab\x2b\xe5\x10\x0d\x3f\xc4\xc7\x29\x29\x73\x9b\xba\x90\x82\x19\x3f\x31\x88\xf2\x07\xea\x5c\xcf\x51\x51\x45\x2f\x91\xf6\xad\xb7\xc0\x6f\xab\x65\xd9\x53\xa3\xb3\xc5\xc2\xac\x16\x3d\x3c\xed\xfb\xa5\x9e\x7a\x65\x9c\xd1\xba\x76\xb3\x80\x6a\x77\x7b\xa7\x6d\x1d\x0c\x28\x48\x0e\xc4\xd8\xb3\x2a\x3f\xcd\x33\x4b\x28\x9a\x8e\xb4\xdd\x3c\x32\x90\x6f\xa5\x70\x1c\x17\x79\x3f\x77\xf1\xd5\x09\x6b\xb5\x4e\x78\xf4\x9d\x96\x5b\x9c\x49\xe1\xba\x9e\x07\x94\x96\xdc\x19\x83\x25\xcb\xec\x38\xd2\xe6\xe4\xc8\x90\xd6\x1f\xd5\xc7\x4a\x6a\xe4\x65\x01\x00\xb0\xfd\x0f\x75\xcc\xc4\x2c\x08\xfc\x67\xfd\xdf\x2b\x67\xee\xff\xd2\x91\xef\xe5\xfb\x8f\x96\xdf\xeb\x11\x89\x3f\x59\xa5\xfc\x6f\x75\xf6\x8f\x2e\xbf\xdf\x2e\xd5\x3f\xb9\x0c\x60\xf9\xcb\xb2\xe6\xc7\x20\xdf\x63\x41\xfb\x4f\x41\xf6\x38\xfe\xbf\x6b\xfb\x63\xac\xef\x3b\xff\xe7\x82\x6a\x38\xff\x32\xd6\x2c\xa0\x6c\xa0\xbf\xbb\x60\x03\xd8\x80\x0e\x26\x00\x78\xce\xf5\xf7\xd7\xbf\x05\x00\x00\xff\xff\x49\xf3\x58\x3b\x88\x0d\x00\x00")

func syntaxStdlibSafeArraizBytes() ([]byte, error) {
	return bindataRead(
		_syntaxStdlibSafeArraiz,
		"syntax/stdlib-safe.arraiz",
	)
}

func syntaxStdlibSafeArraiz() (*asset, error) {
	bytes, err := syntaxStdlibSafeArraizBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syntax/stdlib-safe.arraiz", size: 3464, mode: os.FileMode(0644), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0xa0, 0x71, 0xbc, 0xee, 0xe, 0xd5, 0xc1, 0x9d, 0x86, 0xd8, 0xd7, 0x6e, 0x97, 0x93, 0x69, 0x39, 0x82, 0xec, 0x90, 0xb, 0x69, 0xa0, 0x92, 0x6d, 0x5b, 0xcf, 0x2e, 0x2d, 0x33, 0x6e, 0xf1}}
	return a, nil
}

var _syntaxStdlibUnsafeArraiz = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\xf0\x66\x66\x11\x61\xe0\x60\xe0\x60\x40\x07\x3c\x0c\x0c\x0c\xc9\xf9\x79\x69\x99\xe9\x7a\x89\x45\x45\x89\x99\x97\x4e\xf8\x9e\x39\xd3\xad\x7f\xea\xbc\xfe\xc6\xa0\x00\xaf\xf3\x3a\x27\x35\x7c\x2f\x79\x9f\x3b\x7f\xd5\x5f\x47\xeb\x92\xef\x19\x96\xc0\xce\x99\x5e\x56\x5c\x09\xa1\x99\x33\x2d\x57\x86\x71\x79\xfd\xda\xf9\x53\x73\xa9\xe5\xca\x5f\x16\x21\xaf\xc5\xc3\x9f\x4d\x55\x09\x7b\x21\x3e\x7d\xc9\x93\x14\x35\x47\xcd\xbb\x97\xf6\x3c\x3a\xb2\x25\xc4\xb6\xa3\x72\x56\x03\x03\x03\xc3\xff\xff\x01\xde\xec\x1c\xfa\x49\xa6\xa6\x91\x0c\x0c\x0c\x65\x0c\x0c\x0c\xb8\x9d\xa4\xca\xc0\xc0\x90\x9b\x9f\x52\x9a\x93\xaa\x9f\x9e\x59\x92\x51\x9a\xa4\x97\x9c\x9f\xab\x9f\x58\x54\xa4\x9b\x98\xa9\x0f\x76\xa4\x7e\x7a\xbe\x5e\x6e\x7e\x4a\xcd\x94\xb3\xbb\x16\x9b\xf1\xb4\x5d\x3f\x1b\xba\x21\x47\xfc\x87\xfb\xb1\x3f\x3a\xb7\x62\xe7\xc7\x1d\x88\x69\x34\x6c\xaa\x70\x32\x98\xea\xab\xf6\xe4\xd5\x73\xc5\x33\x67\x3a\x2d\x8b\xcf\x9d\xe2\xae\xfd\xf5\xc7\x6b\x53\xe9\xd7\xbb\x46\x49\x66\x3b\x2a\xe2\xa2\xd2\xfe\x4e\x95\x5e\x74\x67\xc7\x5b\x83\x77\x25\x52\x35\x6a\x33\x43\x96\x0b\x3f\xb4\x91\x9a\xad\x71\xf8\x8e\xf7\x51\xdf\xfe\x27\x9b\x67\xa7\x74\xbd\x7c\xf3\xbe\x5e\x68\xf5\xcb\x89\xdb\xd3\x76\xed\xc9\x6c\x10\x9d\x7d\xf5\x7f\xb7\xf2\x99\x7f\x36\x7f\x3f\x2a\x85\x2d\xfd\x51\xde\x15\x7f\xe1\x9f\x7b\xdd\xcc\xcf\xfa\xd3\x4b\x4d\x8e\x7c\xcc\x2b\x2c\xe1\x7f\x62\x7e\x22\xef\x80\x88\xb0\xe7\x24\xc9\x8b\xc7\x36\xef\x9f\x75\x80\x33\xfd\x54\x5e\x13\xa7\x5a\x91\x80\x42\x0b\x93\x6e\x17\xd3\x82\x53\xdc\x91\xe2\x42\x2e\xf1\x61\xad\x67\x2d\xce\xdf\xb9\x62\xe9\x24\x1d\xbc\x7a\xfe\x86\xcf\xa9\x37\x6a\x19\x82\x0e\xde\x9b\x31\xef\x0f\xf3\x72\x5f\xc5\xd9\xca\x67\xd7\x59\xcc\x2e\xb8\xf7\xe8\xeb\xfb\xc3\xc6\x3f\x5f\xdc\xf7\x48\xd5\xe5\x62\x08\x14\x9d\x7a\xcd\xf2\x4b\x6d\xb7\xec\xb6\x8e\x5f\xdb\x77\xde\xbf\xb5\xf2\x7e\x50\x68\xe1\x87\x6d\xcf\x8f\x77\x1d\xe1\xf4\x35\xdf\xc5\x29\xfc\x3b\x48\x63\xe5\x82\x3b\x6f\x66\x73\x45\xb5\x2c\xd2\xce\xf2\xb9\xc4\x3e\x5f\x79\xe2\xf2\x03\xdf\x93\x67\x4f\x4d\xf9\xb4\x40\xf2\xac\xc9\x9b\x1f\x6f\x2f\xe5\x0b\x9d\x7d\x70\xb8\x8e\x6f\xde\x1d\x0b\xa9\x7f\xbc\xb7\x1f\xbc\x65\x7d\x3c\x2d\x60\xc3\x96\xe5\xe7\x7c\xae\xf5\x57\x98\x0a\xf3\x4f\x5c\xc3\xe9\x5f\x74\x32\x64\xd1\xae\x59\x12\xcb\x7b\x2f\xb2\x0a\x3f\xde\x25\xae\x1e\xbc\xfc\xa7\xd1\x9f\x24\xf1\x35\x96\x1b\x2e\x87\x6f\xe1\x37\xde\xc5\x3d\xbf\xa6\xe4\xef\xae\xec\x20\xf6\x28\x27\x05\x79\xa6\x7c\xf7\x64\x7d\xdb\x9d\xe7\xd6\xd8\xcf\x9b\x9b\xf2\xe8\xb0\x15\xf3\xb1\xcf\x67\x2f\x3c\x30\x77\x38\xe4\xf1\x3a\xce\xc9\x6a\xfe\x83\x8c\x99\xcc\x7e\x35\x5f\x98\x75\xcf\xff\xbb\xc6\x7e\xec\xfb\x7d\x4d\x5e\xf7\x43\x19\xe7\x8f\xd4\xfe\xfe\xfa\xaf\x86\x79\xbe\xe3\x5f\x11\x77\x9e\xea\xbd\x93\x42\xd9\x97\xb4\x56\x15\x3b\xe7\xb9\x7f\x5f\xec\x55\xc1\x35\x2f\x38\x6d\x1e\x5b\xdd\xf3\x69\xd7\x7e\x7e\x5d\xdb\x24\x60\xca\x31\x61\x93\x90\xd5\xf5\x6d\x4c\xc1\xd7\x57\xad\x5c\x62\xd9\xb0\xa7\xca\xe7\xf6\x93\x7a\x49\x8f\x85\x8e\xdd\x41\x8d\x7a\xc5\x97\x6f\x39\x69\x4f\xe0\xdd\xf6\x4b\x98\x39\x7f\x5a\x45\x1b\xcb\x22\x41\x86\x20\xd1\x80\xe7\xbb\x98\xef\x4a\x66\xbe\x5a\x17\x30\xbf\xa2\xc5\x54\x45\x60\xa9\x90\x53\x46\x80\x92\x78\xed\x65\x69\x0d\xdf\xf8\xf0\xe3\x77\x3f\x64\x09\x7c\x94\xfb\xbe\x6c\x7a\xf4\x11\x7e\xe5\x4d\xe2\x07\x1b\x2e\x3f\x8c\xce\x9f\x9e\x73\xf6\x52\x54\x48\x65\x74\xd9\x85\xd7\xdb\xc4\x63\xaf\xca\xea\x57\xa6\xbb\x6e\xae\x3d\x57\x1f\x2f\xff\xfe\xd1\xfa\xff\xfc\xa7\x62\x8d\xe3\xd9\x9e\xa9\xed\x9b\xf5\xfa\x44\x9c\x78\xaa\x54\xc0\x4d\xdf\xbd\x17\x3d\x03\x4f\xb2\x49\x35\xec\xc9\xf0\x8a\xcd\x08\xd6\x0a\xdc\xca\x79\x37\x7e\xf9\x7b\x46\x58\x1a\xf7\x79\xb7\x7f\xce\x44\x26\x06\x06\x79\x36\x7c\x69\xdc\x81\x70\x1a\x2f\xae\xcc\x2b\x49\xac\xd0\x2f\x2e\x49\xc9\xc9\x4c\x82\x52\xba\xa5\x79\xc5\x89\x69\xa9\x90\xac\x9a\xd2\xe7\xd8\xd5\x12\xc0\xe3\xf2\x7d\xde\xa2\x03\x7b\xbf\xcf\xf8\xab\x93\x2e\xfc\xf7\xa2\xfc\xb3\x72\x25\xcf\x9c\x8f\xdf\x7a\x15\x7a\x7d\xdf\x73\xad\x92\xfb\x5a\xf5\x27\x8c\xf3\xc8\x93\x4d\x6a\x61\xdb\x97\x8b\xf0\x09\x94\x59\x19\xcc\x56\x7e\x73\x70\x86\xf9\xdc\x20\x1f\xe6\x45\x27\x4e\x75\xa8\xe8\x3e\x61\xf5\x75\x53\xb7\x9f\xd0\x36\x71\xe3\xee\x94\x4d\xf3\xb4\x8b\xe4\x3d\xd8\xfe\xa8\x4c\xfa\x67\x69\xb7\x61\x5d\x7d\xb3\x79\xca\x46\xb7\x5d\x0d\xe7\x7e\xdf\xb9\x76\x4c\x41\xb7\xe0\xdc\x59\x53\x87\xac\xc8\x9b\xeb\xd3\xa6\xaf\x59\x10\xf4\xe4\xca\xa2\x32\xe3\x3f\x65\x2f\xfe\x29\x79\x4f\xde\xc8\xdc\xa0\xbd\x8f\x09\x16\x04\x93\xe7\xae\xd8\x38\x95\x81\x81\xe1\x07\x38\x9b\x33\x32\x89\x30\x20\x02\x01\xb9\x08\xe0\xc1\x08\x16\xe4\x92\x08\x5d\x27\x72\xc0\xaa\xa2\xe8\x9a\x4c\x6c\x81\x81\x6e\x24\xb2\x43\x1d\x50\x8c\x2c\x67\xa6\x3c\x7e\x02\xbc\x59\xd9\x40\x66\x31\x33\x30\x33\xfc\x66\x60\x60\xa8\x62\x01\xf1\x00\x01\x00\x00\xff\xff\xd5\x16\x01\x32\x8b\x05\x00\x00")

func syntaxStdlibUnsafeArraizBytes() ([]byte, error) {
	return bindataRead(
		_syntaxStdlibUnsafeArraiz,
		"syntax/stdlib-unsafe.arraiz",
	)
}

func syntaxStdlibUnsafeArraiz() (*asset, error) {
	bytes, err := syntaxStdlibUnsafeArraizBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syntax/stdlib-unsafe.arraiz", size: 1419, mode: os.FileMode(0644), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x48, 0x63, 0xfa, 0x21, 0xc7, 0xb3, 0xe6, 0x8f, 0x38, 0x2c, 0xdc, 0xe3, 0x8b, 0x9a, 0xad, 0x80, 0xf5, 0x16, 0x6d, 0xf5, 0x29, 0x8f, 0x75, 0xa, 0x4, 0xe1, 0x3, 0xfb, 0x7b, 0xa7, 0xd6, 0xa8}}
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
	"syntax/implicit_import.arrai": syntaxImplicit_importArrai,
	"syntax/stdlib-safe.arraiz":    syntaxStdlibSafeArraiz,
	"syntax/stdlib-unsafe.arraiz":  syntaxStdlibUnsafeArraiz,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"syntax": {nil, map[string]*bintree{
		"implicit_import.arrai": {syntaxImplicit_importArrai, map[string]*bintree{}},
		"stdlib-safe.arraiz":    {syntaxStdlibSafeArraiz, map[string]*bintree{}},
		"stdlib-unsafe.arraiz":  {syntaxStdlibUnsafeArraiz, map[string]*bintree{}},
	}},
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
