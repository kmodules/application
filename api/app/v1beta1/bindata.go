// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package v1beta1 generated by go-bindata.// sources:
// app.k8s.io_applications.yaml
package v1beta1

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

var _appK8sIo_applicationsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\xcb\x72\xeb\xb6\x92\x7b\x7f\x45\x97\xef\xc2\x49\x95\x1e\x37\x27\x93\x79\x78\xe7\xb1\x93\x8c\x27\xe7\xe1\x39\xf6\x39\xb3\x48\x65\x01\x91\x2d\x09\x11\x09\x30\x00\x68\x59\x49\xe5\xdf\xa7\xba\x01\x50\xa4\x44\x8a\xb4\xe5\x64\x6e\x6e\x09\x1b\x5b\x24\xd0\xe8\x6e\xf4\x13\x20\x9b\xa2\x90\x9f\xd1\x58\xa9\xd5\x25\x88\x42\xe2\x93\x43\x45\xbf\xec\x64\xf5\xef\x76\x22\xf5\xf4\xf1\xab\xb3\x95\x54\xe9\x25\x5c\x97\xd6\xe9\xfc\x23\x5a\x5d\x9a\x04\x6f\x70\x2e\x95\x74\x52\xab\xb3\x1c\x9d\x48\x85\x13\x97\x67\x00\x42\x29\xed\x04\x5d\xb6\xf4\x13\x08\xe6\x58\x14\x85\xd1\x8f\x98\x4e\x56\xe5\x0c\x8d\x42\x87\x04\xf9\x12\x96\xce\x15\xf6\x72\x3a\x5d\x48\xb7\x2c\x67\x93\x44\xe7\xd3\x6d\x8f\xb1\x95\x0b\x3b\x15\x45\x91\xc9\x84\x01\x4e\x8b\x32\xcb\xa6\x6f\x18\x6a\xa2\x95\x33\x3a\xcb\xd0\x8c\x17\xa8\x18\xee\xac\x94\x59\x8a\x86\x51\x8e\x04\x3d\xfe\x7d\xf2\xcd\xe4\xef\x67\x00\x89\x41\x86\xf1\x20\x73\xb4\x4e\xe4\xc5\x25\xa8\x32\xcb\xce\x00\x94\xc8\x91\x28\xaf\xa6\xb1\x13\x51\x14\x81\xf8\x33\x5b\x60\x42\x74\x2c\x8c\x2e\x0b\xee\x16\xef\xf8\x91\x81\xc8\x44\x38\x5c\x68\x23\xe3\xef\x31\x08\x06\x0e\xe0\x59\x77\xb5\x05\xcf\x57\x33\x69\xdd\x0f\xbb\x77\xde\x4a\xeb\xf8\x6e\x91\x95\x46\x64\x4d\xa4\xf8\x86\x5d\x6a\xe3\xde\x6f\xa7\x1d\x53\x17\x7f\x47\xaa\x45\x99\x09\xd3\x18\x74\x06\x60\x13\x5d\xe0\x25\xf0\x98\x42\x24\x98\x9e\x01\x04\xee\x30\x8c\x31\x88\x34\xe5\x55\x14\xd9\x9d\x91\xca\xa1\xb9\xd6\x59\x99\xab\x6a\x86\x14\x6d\x62\x64\xe1\x98\x9f\x0f\x4b\x04\xb7\x29\x10\xf4\x1c\xdc\x12\x77\x26\xa3\xf6\xb3\xd5\xea\x4e\xb8\xe5\x25\x4c\x88\x77\x93\x38\x5c\x9b\x09\x0d\x0c\xbd\x3c\xd3\x1f\xb6\x17\xe8\xde\x25\x58\x67\xa4\x5a\x74\x4d\x1c\xd7\x10\x52\xe1\xb0\x7f\xba\x40\x66\x63\xc6\xcf\x8d\x6b\x43\x26\xad\x51\x08\x7a\xf6\x33\x26\x0e\xf4\x5a\x59\xa6\x3e\x17\x2e\x59\x62\x0a\x26\xa8\x84\xed\x40\x4a\xa4\xe9\x87\xb5\x42\xf3\x11\xe7\x0d\x6c\xf8\x62\x03\x97\x99\xd6\x19\x0a\xd5\x86\xcc\xfb\x32\x9f\xa1\xb1\xc4\xf9\x44\xe7\x85\x56\xa8\x9c\x05\x83\x22\xdd\xb4\x4c\xeb\x84\x2b\xed\x64\xdb\xf1\x63\xad\x9f\x9f\xbc\x7e\xe5\x15\xb8\x1f\x6d\xc0\x64\x4f\xd3\x1a\xb3\x5e\x2d\x9a\x2b\x5e\x01\xf3\xb7\x1f\xbf\x9a\xa1\x13\x5f\x79\x91\x4e\x96\x98\x8b\xcb\xd0\x5d\x17\xa8\xae\xee\x6e\x3f\x7f\x7d\xdf\xb8\x0c\x4d\x44\x6b\xda\x04\xd2\x2f\x92\xef\x0f\x73\x6d\x76\x25\xd6\xc2\xd5\xdd\x6d\x05\xa7\x30\xba\x40\xe3\x2a\x15\xf6\xad\x66\x1d\x6b\x57\x77\x66\xbd\x20\xc4\x7c\x2f\x48\xc9\x2c\xa2\x9f\x3a\x48\x20\xa6\x81\x16\xaf\x35\x92\x56\xad\x30\x68\x51\xb9\xba\xe6\xc4\xa6\xe7\x20\xa2\xac\x4d\xe0\x1e\x0d\x81\x21\xd5\x2f\xb3\x94\xec\xde\x23\x1a\x07\x06\x13\xbd\x50\xf2\xd7\x0a\xb6\x05\xa7\x79\xd2\x4c\x38\x0c\x76\x64\xdb\x58\xb1\x95\xc8\xe0\x51\x64\x25\x8e\x40\xa8\x14\x72\xb1\x01\x83\x2c\xd1\xa5\xaa\xc1\xe3\x2e\x76\x02\xef\xb4\x41\x90\x6a\xde\xb4\xd3\xd1\x2b\x24\x3a\xcf\x4b\x25\xdd\x66\xca\xa6\x58\xce\x4a\xa7\x8d\x9d\xa6\xf8\x88\xd9\xd4\xca\xc5\x58\x98\x64\x29\x1d\x26\xae\x34\x38\x25\x27\xc0\xa8\x2b\x6f\x61\xf3\xf4\x6f\x95\xd2\x5c\x34\x70\xdd\x13\x45\xdf\xd8\x8c\x1e\x58\x01\x32\xa6\xb4\xe2\x22\x0c\xf5\x54\x6c\x19\x4d\x97\x88\x3b\x1f\xbf\xbd\x7f\xa8\xf4\x95\x17\x63\x97\xfb\x5e\xc7\xab\x81\x76\xbb\x04\xc4\x30\xa9\xe6\x68\xfc\x22\xce\x8d\xce\x19\x26\xaa\xb4\xd0\x52\x39\xfe\x91\x64\x12\xd5\x2e\xfb\x6d\x39\xcb\x25\x6b\xeb\x2f\x25\x5a\x47\x6b\x35\x81\x6b\x76\x95\x30\x43\x28\x0b\x52\x83\x74\x02\xb7\x0a\xae\x45\x8e\xd9\xb5\xb0\xf8\x87\x2f\x00\x71\xda\x8e\x89\xb1\xc3\x96\xa0\xee\xe5\x77\x3b\x7b\xae\xd5\x6e\x44\xbf\xd9\xb1\x5e\x35\x3d\xbd\x2f\x30\x69\x68\x0d\x0d\x95\xf3\xa8\xc5\xa4\xb7\x42\xd5\x07\x4c\x1a\x60\xdb\xd5\x96\x5a\xcd\xea\xee\xde\xda\xc5\x66\xdb\x33\x50\x62\x61\x0c\xf3\x4c\x2c\x48\xa9\xa4\x4a\x69\x66\x04\x39\x87\x35\x82\x42\x4c\xc1\xe9\x3d\x88\x3c\x21\x44\x38\xac\x8e\xec\x1f\x48\xf2\x22\xd0\x77\xf1\x82\xb4\x90\x6a\x85\x30\xdb\x40\x49\xae\x1b\xee\x31\xc3\xc4\x69\xd3\x02\xd6\x69\xf8\xa5\x44\xb3\xa1\x90\x02\xae\xa3\x45\xff\x9e\x02\x12\x92\xfa\x5d\x01\x6e\x73\x25\x35\x1c\xad\xc5\x7c\x96\x6d\xee\x96\xc2\x62\x1f\x5b\xea\x7d\x6b\x1a\xe1\xe5\xbc\x34\x06\x95\x83\x82\x6f\xee\x87\x03\x17\xfb\x78\x6d\x67\x9f\xc0\x95\x02\xcc\x0b\xb7\x09\x7a\x2a\x2d\xe0\x2f\xa5\x7c\x14\x19\xc1\x74\x1a\xce\xef\xcb\x24\x41\x4c\x31\x3d\x9f\x74\x10\xb8\x27\xa0\xc0\x71\x61\xe0\x0f\xb3\xa6\x87\xc0\x16\x66\x7a\x1b\x42\xe1\x19\xd1\xe4\x2f\x91\x08\x5e\xf5\x50\x56\x73\xc9\x5f\xe0\x64\x31\x81\x1b\x2c\x32\xbd\xc9\xe9\xca\x08\xee\x74\x6a\x47\x6c\x47\x64\x82\x76\x04\xd7\x1f\x6f\xec\x97\x13\xb8\x75\x90\x08\xc5\x16\xc0\x72\x64\xb6\xdb\xa4\x22\x83\xff\x73\xa9\x12\xd6\x85\xb5\x74\x4b\xe6\x73\x03\x9b\x4a\x76\x88\x6f\x1e\x73\x03\x6b\x92\xb4\x36\x61\x6a\x8e\xb6\x35\xc4\xf7\xf9\x2c\x1d\xe6\x2d\x3c\xdc\xe1\x62\xc5\xbc\xa8\xb9\x48\x2c\xe4\xab\xec\x67\x04\xb3\x71\x04\xb3\xd2\x41\xaa\xd1\x82\xd2\xbb\xf6\xd1\xb7\xb9\x26\x9b\x2c\xa2\xdb\x9c\x00\x3c\x90\x9d\x95\x96\xd8\x33\x2f\x33\x5e\x08\x99\x92\x09\x9b\x6f\x48\x69\x12\xad\x12\x2c\x9c\x85\xb4\x6c\x11\x05\xdf\x32\xad\x57\x65\x01\xd6\x89\x05\x5a\xe6\xa0\x2e\x1d\x2c\xc5\x23\x01\x28\x84\x71\x52\x64\x19\x8b\xa1\x4c\x59\xb0\xda\x56\xb7\xdb\xc8\xf8\xe6\xb3\x83\xd6\x5b\x07\xa5\xd5\xb7\x7d\xdf\x36\x78\x30\x39\x14\x69\xb0\x75\xf8\xd8\xa3\xd5\x7a\x87\xa6\x6c\x93\x8e\x76\x63\x5e\xbf\x29\x8c\x11\x9b\x9d\x7b\xdb\xa0\xbb\x47\xe1\x6e\xaa\x8e\x60\x90\xb1\xb3\xec\xe0\x4c\xee\x8d\x3d\x87\x25\xc1\xcd\x80\x98\xd1\x4a\xed\x59\x30\xe0\xa8\xac\xc3\x1b\x40\xef\x62\xd5\xf1\x69\xe5\x79\x2b\xc2\x21\x9e\x14\x30\x33\x12\xe7\x31\xc6\xa8\x75\x0d\x16\xb0\x63\x15\x3b\xbd\x57\x93\xb7\x9d\xcb\x2c\x93\x2a\x9d\x3e\x88\xee\x2d\xf5\x63\x44\x15\xe8\xc2\x67\x75\x95\x31\x63\x20\xd1\x9f\x1e\xe4\xa0\x6f\x04\xac\xb1\x3a\x52\x25\x59\x99\x46\x2f\xcd\x11\xd4\x08\xac\xfc\x35\xc6\x93\x32\xf7\xc9\x61\x3b\xc0\x4e\x63\xb2\x4f\x45\x2e\x16\xc8\x61\x01\x45\x36\x42\xaa\x1d\x21\x09\x72\x01\x92\xfa\x75\x99\x4e\xdf\x04\xb3\x82\x28\xef\xa2\xb2\x4f\xb5\x81\x93\xeb\x5f\x5b\x7c\x65\x07\xf6\x5f\x44\xc6\x7f\xc9\x69\x13\x8d\x8d\xee\xd1\x23\x2c\x15\x14\xf2\x09\xb3\x36\x43\xb3\x6d\xec\x45\x46\xf0\xe6\x9b\xa7\x37\xdf\x7c\xd9\x85\x3c\x0c\x31\x2f\x4c\x81\x49\x06\x13\xf0\x50\xad\xae\x37\xb8\x8c\x74\xe5\xfd\x31\x25\xa6\xa2\x74\x4b\x6c\x0b\x56\xb6\x8d\x84\x6c\x66\x75\x56\x3a\x84\x4f\x1f\xdf\xc6\xf4\xc4\x83\x23\x21\x84\x1b\xd2\x72\xba\x15\x96\xf9\x10\x01\x10\x7c\x17\x0f\x9f\x54\x43\x2d\x08\x83\x21\x7a\x4c\x89\xb3\x1f\xbf\xbb\x86\x37\x5f\xff\xc7\xbf\x1d\xcd\x31\xee\xf4\xc2\x35\xcf\xa5\x6a\x6e\x93\x78\xa2\xfd\x92\x1e\x24\xf1\x9c\x7b\x4e\x0b\xb5\x38\x3f\x76\xd1\x0f\xb9\x06\x60\x27\x60\x4d\x72\xd0\xeb\x74\xf8\x01\x38\xe8\x0b\xa8\xad\x70\xb3\xd6\xa6\x2d\xfc\x82\x5d\xce\xfd\x10\xba\x76\x19\xad\x15\x6e\xc0\x77\x10\xd6\xea\x44\x52\x9e\xd4\x81\x74\x15\x1d\xd5\xb7\x6c\x7c\x30\xf6\x6e\x73\xff\x3f\x6f\x47\xf0\xf1\xe6\x3f\xdf\xdd\x8f\x80\xbc\xcb\x4c\x58\xec\xe0\xf0\x41\x43\xd5\xcb\xf8\xc3\x9c\xc9\xa4\x5a\x0d\x61\xcb\x5b\xea\xc7\xb2\xbd\x0d\x46\xab\x0e\x8f\xe8\x45\x9f\x92\x7a\x95\x76\x65\x23\xd4\x42\x64\x49\xba\x67\x4b\x33\x17\x14\x5e\x55\xdb\x7d\x90\xea\xa4\xcc\xe3\x26\x04\xb1\xc5\x2e\x67\x5a\x18\x0a\x55\xd1\x25\xc7\x5a\x71\x22\xa1\xc7\x80\x07\xb3\x10\x50\xeb\x14\xf7\x97\xe0\x39\xcc\xb0\xf7\x06\x03\x2d\x74\xed\x84\x04\xcb\x32\x17\x8a\xb7\xe1\xc4\x2c\x43\x26\x98\x52\x18\x7c\x2a\xb2\x81\x26\xad\x28\x4d\xa1\xb7\x39\x14\x89\xc8\xd1\xd6\xab\x34\xd9\x60\x82\x3e\x99\x8c\x80\xca\x84\x03\x61\xde\xc9\xb0\x20\x1c\x08\x58\xe3\xcc\x4a\xc7\x32\x63\xd0\xb6\x24\x08\xcf\xc2\xea\x48\xa3\x92\x0b\xc9\xb2\x84\x66\x88\x02\xbd\xdb\xf6\xee\x32\x2d\x35\x80\xa0\xe7\x5d\x48\x37\xed\xc9\xc4\xdb\xf7\x3a\x70\xe5\x77\x83\x18\x6a\xbc\xc1\xc3\xba\xc3\xc1\xe0\x5d\x13\x9d\xe2\xc8\xfb\x06\xeb\x43\xa8\x42\x24\x2b\xf2\x14\x2d\x9b\x95\xc7\xaa\xe3\x35\x69\x62\xe2\xd8\x71\xf6\x84\x55\x2a\x95\x8f\x32\x2d\x45\xd6\xb9\xde\xda\x80\x36\x0b\xa1\xe4\xaf\x07\x23\xc8\x21\x2a\x88\xb9\x90\xc3\x65\xf5\x5b\xea\x1d\x77\x76\x79\xe8\x2b\x09\x67\xd8\x80\x1e\x8a\xc7\x7b\x91\x63\x44\xa3\x6e\x97\x09\xc8\x9f\xae\xbc\x09\xef\x0c\x6f\x55\x78\x86\xff\x68\xba\xab\xb4\xeb\x92\x81\x26\x5b\xa9\x5f\x14\x4f\x10\xbb\xc6\xd5\x2a\x59\x14\xe8\xb6\x9e\xaf\x03\x5b\x8a\xfa\xe1\x97\x52\x26\x2b\xca\xfb\x8d\xab\xf4\xa9\xb4\xe1\x04\x65\x67\x1b\x64\x02\xd7\x3a\xcf\xb5\x7a\x27\xcc\xaa\x03\x66\x2e\xcc\x2a\xd5\x6b\x05\x76\xa3\x9c\x78\xe2\xcd\xe0\xe8\x5e\x09\xbc\x91\xc9\x12\x1c\x3e\xb9\x9d\xcd\xfd\x17\xe5\x7b\x7a\x3d\xd0\xcc\xf1\x76\x63\xa7\x85\x23\x2a\x3d\xa8\x2a\x10\x55\xd6\x89\x2c\xeb\xe6\xdc\xae\xa9\x6b\x0e\xaf\x47\x56\xe1\x40\x62\xe6\x7d\x9e\x48\xba\x03\xb3\x60\x0f\xf1\x91\x1c\xa3\x9e\x83\x80\x22\x13\x8a\x02\x76\x6d\xa0\x54\xf1\x47\x2a\xad\x29\xbd\x5f\x15\xf3\x39\x26\xee\x80\x38\x9e\x6c\xe3\x3e\x15\x27\xdb\xf8\x97\xb4\x8d\xdd\x89\x66\x33\x2b\xa7\x74\x32\x70\xb5\xe3\x04\x3e\x24\x3b\xff\xab\x4d\x7a\x47\x74\x75\xa5\x9a\x21\x1b\xba\x16\xd6\x0a\x95\x1a\xd1\x91\x07\xf5\x90\xfe\xd8\x76\x2c\xda\x8a\x7a\x3c\x1a\xdd\xb1\x52\x8f\xf1\xb2\x3f\x5f\xd1\x26\x9a\xe9\x3f\x64\x13\xed\xc0\x2a\xf1\x11\xdb\xe1\x8d\xcb\x5b\x35\xd7\x5b\xd3\xb0\xe3\x96\x56\xb8\x19\xf9\x63\x8c\x42\x48\x63\x0f\x90\x71\x90\x84\x61\xbb\xed\x84\xc9\xad\xc3\xdc\x6f\x47\x1e\xc4\xa4\x6f\x5f\x45\xe6\x85\x36\x4e\x28\xd7\x62\xeb\x96\x7a\x4d\xa9\x99\x48\x12\xb4\x76\xcf\x5d\xbe\x60\x8b\xfc\x90\x09\x69\x35\x1f\x7b\xc4\x39\xe9\xb2\x18\x1a\x4b\x0b\x85\xc4\x04\xbb\x03\x77\xa8\x13\xd5\xa5\xe1\xc3\xb2\xf9\x21\x58\x3f\xd4\x54\xd2\xaf\x40\x85\x68\x5c\xb1\x17\x23\xc1\xf0\x06\x61\xf1\x39\x1e\xa6\xb5\x67\xa5\xc7\x61\xf0\x9d\xd1\xf9\x70\x2c\xa8\x77\x75\xac\x2b\xc0\xe0\x1c\x0d\xaa\x04\x49\xaa\x52\x34\xe4\x10\x2a\x5e\x75\x2e\x21\x1f\xb1\x0b\xa5\xdd\x12\x4d\xc8\x9a\x8e\xf1\xa7\x89\x56\x73\xb9\x78\x27\x8a\x1f\x70\xd3\x7a\x2a\xdc\x41\x90\x3f\x60\x23\x32\x56\xb8\xf1\x11\xcc\x75\x04\x75\xc8\x79\x0c\xc1\x09\x0e\x3c\x64\xd2\x83\xd6\xd5\xdd\x6d\x65\x44\x83\xe8\x05\x36\x77\xae\x74\x6c\x83\x7c\x2d\xb5\xb9\xc4\x2c\xe5\x47\x7c\x9e\x83\xd9\xc5\xed\xdc\xa3\xc2\xc7\x24\x64\x47\x2a\x75\xdd\x3e\xdb\xc2\xc1\x28\x8a\xee\xcd\xfb\xd8\xfc\x20\x54\x4e\x1a\x0c\x63\x47\x5e\xb5\xc2\x39\xcc\xf6\xd9\x18\x9f\x37\xf4\x42\xf4\xa7\x7d\xff\x7d\xff\xe1\xfd\xf4\x7b\xed\x69\x8c\xa6\xce\x3a\xe1\x30\x47\xe5\x46\x60\xcb\x64\x49\xb9\x44\x8a\x56\x1a\x4c\xef\xe9\xce\x24\x17\x4a\xce\xd1\xba\x49\x98\x0d\x8d\xfd\xf1\xcd\x4f\x7d\xfc\x06\xf8\x4e\x1b\xc0\x27\x91\x17\x19\xe5\xfc\x21\x36\x8f\xcf\x9a\x44\xdd\x90\xd6\x33\xab\x82\xdd\x0b\x76\x2d\xdd\x92\x53\xa5\x42\xa7\x81\x29\x6b\x1f\xf2\x88\x15\x02\x99\xf3\x60\x8e\x32\xb9\xc2\x4b\x38\xe7\xe7\xd1\xb6\xa8\xff\x46\x66\xf9\xf7\xf3\xde\x69\xbe\x58\x2f\xd1\x20\x9c\x53\xf7\x73\x8f\x70\xf5\xc0\x11\x5d\x8b\xf2\x37\x1c\x71\xb7\x14\x0e\x9c\x91\x8b\x05\x1a\x4c\xb7\x09\xc2\x97\x14\x40\xcb\x39\x28\xbd\x05\xc6\x53\xf4\x42\x24\x79\x08\x07\xd0\xe9\x1e\xa1\x3f\xbe\xf9\xe9\x1c\xbe\xd8\x42\x24\xbe\xf5\x43\x54\x29\x3e\xc1\x9b\x6a\x8f\xa7\xd0\xe9\x97\x13\x7f\x28\x1d\xb2\x40\x69\x21\x59\x6a\x8b\x0a\xb4\xca\x36\xdd\x1b\xb1\xdb\xb6\x14\x8f\x08\x56\xe7\x08\x6b\xcc\xb2\x71\x3c\xba\x58\x0b\x36\x2c\x51\x14\x48\xaa\x05\x1f\x4f\x1f\xf2\x70\xb1\xd5\x1e\x17\x7b\xf8\x70\xf3\xe1\xd2\x63\x4b\x62\xbb\xe0\xa8\x4b\x69\x07\x73\x49\x31\x97\x50\x69\xdb\x43\x4f\xfb\xcd\xeb\x04\x11\x5a\x7a\x21\x75\x1a\x92\xa5\x50\xfe\xfc\x8a\x56\x6b\x5e\xba\xd2\xe0\xe4\xe2\xb5\x2c\xcd\x0a\x37\xcf\xb2\x31\x94\x9d\x92\x39\x76\x1a\x2c\xdb\xe7\x57\xb3\x79\x87\xce\xe3\x5b\x51\xf1\x4f\x9f\xed\x9a\xe0\xff\xb7\xe7\xb7\x8e\x20\xbd\x2f\xd1\xdb\x27\xfd\x7d\x4d\xfb\x0f\x92\xde\x78\xf4\x7b\x9a\xea\xc4\x4e\xe3\xf3\x1b\x53\xfd\x88\xe6\x51\xe2\x7a\xba\xd6\x66\x25\xd5\x62\x4c\xea\x39\x0e\x4f\x4d\x4d\xf9\x39\xeb\xe9\xdf\xf8\xcf\xab\x52\xca\x4f\x44\x3f\x9f\x5c\x1e\x76\x88\xe6\x7e\x23\xf0\x3a\x3c\x21\x3c\xec\xf4\xd5\x58\x12\x1f\x8b\x7c\x49\x2c\x72\x71\x1f\x9e\xda\xdb\x85\x42\xfa\xb9\x5e\xf2\xe6\xd8\x10\xbb\xd3\xf0\x85\xb9\x48\xbd\xb3\x14\x6a\xf3\x87\xab\x13\x31\x9e\x9f\x69\x4b\x36\xe3\xf0\x76\xc1\x58\xa8\x94\xfe\xb7\xd2\x3a\xba\xfe\x6a\x9c\x2e\xe5\x33\x0d\xcc\xa7\xdb\x9b\x3f\x47\xc9\x4a\xf9\x6a\xd6\xa4\x77\x5f\x04\xd8\xc5\x2e\x0c\x5a\xfb\xfc\x68\x9c\x9c\xde\xad\x1f\x7c\x0a\xc1\x5b\x30\x3b\x85\xe0\x7b\xed\x14\x82\x37\xdb\x29\x04\x3f\x85\xe0\x4d\xc6\x68\xeb\x9e\x1d\x83\x57\xfb\xb9\x34\xfa\x14\x8d\x9f\xa2\xf1\x97\x52\x7a\x8a\xc6\x1b\xad\x78\xae\xc7\x6f\x28\xe3\x7f\x3d\x3c\xdc\x31\x88\x57\xd3\xc3\xc2\x68\xa7\x13\x7d\xf0\xf8\x6d\x0f\xa7\xbb\x30\xa8\x3a\x81\x0f\xf1\xde\x29\x65\x39\xa5\x2c\x7f\xd5\x94\xc5\x62\x62\xd0\x1d\x7f\x84\x70\xcf\x70\x4e\xc9\x4b\x0b\x66\xa7\xe4\x65\xaf\x9d\x92\x97\x66\x3b\x25\x2f\xa7\xe4\xa5\xd1\x4e\xe7\x07\xa7\x8c\xe5\x94\xb1\x9c\xce\x0f\x4e\xc1\x78\x0b\x77\xff\x89\x83\x71\x7e\xa7\xfe\x25\xe7\x07\xf1\x7d\xfc\x53\x04\xde\x82\xd9\x29\x02\xdf\x6b\xa7\x08\xbc\xd9\x4e\x11\xf8\x29\x02\x6f\xb4\x53\xdc\xfb\x0c\xd2\x4f\x71\xef\x3f\x4d\xdc\xfb\x0f\xb7\x53\xaf\xcd\x11\xc7\x78\x34\x7a\x78\x52\xec\x1f\xb1\xbf\x04\xa9\xdc\xd7\x6f\x06\x11\x20\x95\xc3\x45\x8f\x7b\x79\x9d\xb3\x86\x10\x1b\x9e\xd2\x9b\x53\x7a\xf3\x57\x4d\x6f\x9e\x55\x85\x24\xbe\x83\x72\xf8\x45\x89\x81\x28\xf6\xa0\xf7\xc2\x4a\x51\x36\x94\x2a\xeb\x79\xdd\xea\xa2\x2a\x69\xe6\xcb\xb1\x89\x19\x66\xa1\x08\x1e\x2d\x13\x47\x1a\xd6\x07\xa0\x5c\x6e\x14\x53\x98\xb5\xbe\x5b\xb7\xfb\xf6\xfe\xad\x83\xbc\xe4\x17\xf5\x5d\xb2\x0c\xf1\x6e\xa8\x7f\x16\x8b\xf4\x5d\x5c\xf8\xf9\x42\x05\xcc\x16\xa8\xaf\x25\x53\x7e\x9a\xe9\xdf\xf8\xef\x38\xb2\xa6\x45\xbe\x0e\x67\xa1\x4c\xcb\xb7\x4f\x85\x41\x6b\xe5\xb0\x3a\x51\xbb\x43\x9a\x55\xef\x3c\xbb\x23\x3e\xb1\x90\x4c\xde\x5e\x23\xce\xb7\x07\xd6\xbc\x6d\x3f\x2e\x5d\x72\xf5\xfe\x06\xd3\x63\xdf\xcd\xbd\x3a\x80\x4e\x28\xf8\x59\x95\xbf\x5b\x8a\x6e\x4d\xaa\x5e\xde\xf3\xe5\x4d\x47\xfe\x24\xce\x97\x5d\xe0\xb7\x12\xd1\x88\x08\x04\x0c\x72\x15\x55\x16\x90\x15\xb6\x49\x96\x6f\x34\x38\x94\x4b\xed\xe8\x33\x64\x03\xa1\x67\xe7\xba\xc1\x8e\x15\x6e\xe2\x5b\xa0\x9e\x2f\xbc\x85\xcd\x99\x18\x7b\xbc\xc0\x0a\x16\xfa\xd6\x52\x76\xdb\xe6\xf4\xd1\x6f\xe3\x46\xae\x0d\x46\xbf\x62\x73\xad\x9a\x24\x2f\xc4\x85\xf5\x4c\x27\x69\x5c\xca\xa2\x2f\x0f\xa3\x55\x67\x59\x8d\xc5\x6a\x3f\xf3\x7e\x40\x04\xef\xe5\xef\x56\x8d\xe0\xbd\x76\xf4\xe7\xdb\x27\x69\xdd\x61\x76\xd0\x5a\xde\x68\xb4\xef\xb5\xe3\xde\x47\x33\xc7\xa3\x36\x98\x35\xbe\x7b\x78\x49\x96\x6d\x27\xdb\xf2\x5a\x35\x5b\x3b\x81\xdb\x43\x55\xe6\x7c\xab\x58\xcc\xaf\x1e\x52\x2e\x1e\x78\x50\xbd\x6e\x67\x03\x78\xb6\x84\x33\x04\xa5\xd5\x98\x8b\x71\x1e\x0e\xf6\x6e\xc3\x96\x47\x0d\xbe\x67\x2b\xcd\x51\xe7\x5c\x7d\xaa\xc3\x2c\x6f\xa0\xe1\x51\xf0\xc9\xb9\xbf\xe3\x2b\x25\x67\x22\xc1\x34\x94\x78\xf4\xf5\x7d\x85\xc3\x85\xec\xaa\x5d\xe5\x5b\x8e\x66\x81\x14\x56\x27\x07\xe3\xea\x83\x76\xc8\xb7\x67\xb9\xf3\xae\x97\xbc\x61\x50\x41\xae\x6e\x53\x33\xae\xd8\xfe\x32\x77\xdd\x5f\x5c\xc7\x25\xcb\xb7\xec\x8f\xda\xd1\xab\x57\x89\xef\xb3\x68\xbd\x1c\xdb\xf7\x45\x7e\x6a\x6f\xcf\x73\x51\x90\xe4\xff\x56\xbd\xcc\xfc\xbb\x7f\xaf\x7a\x02\x57\x5c\xe6\x3e\xeb\x92\xff\xfa\x88\xb0\xc5\x51\x07\x4e\x70\xf7\xea\xcc\x0a\x05\x98\xb1\x33\xe9\x00\xca\x85\x82\x9a\xde\x72\x04\xeb\xa5\xb6\xfe\xf4\xb0\xda\x65\x39\x5f\xe1\xe6\x7c\xd4\xd0\x90\x0e\x88\xd4\xf9\x56\x9d\x8f\xc2\x56\xce\x8e\x52\x56\x7e\x8a\xb7\xa3\xce\xf9\xde\xf9\x64\xcf\xc1\x76\xc0\xee\x71\xbb\x03\xa2\xb6\x96\x9b\x5d\x25\x9e\xb9\xc6\xfc\xc0\x22\xcf\xdc\xb7\x7a\x1f\x78\xfb\xf5\x86\x0b\x1b\xf6\x51\x29\x59\xc3\xd4\x6f\xde\x9e\xed\x2d\xc1\xee\xf7\x13\x62\xeb\x76\xaf\xdb\xba\xb6\x3d\xb1\xe6\x07\xbf\x3b\xe6\xa9\x09\x8b\xc0\x45\x32\xb3\x6c\xaf\x80\xf3\x1e\xa4\x61\xef\xec\xfb\x29\x02\x0f\x58\xc6\x17\xa8\xd0\xc8\x24\xce\xba\xd4\x59\x8a\xbe\xfe\x41\xd7\x4c\x47\x16\xa1\x6d\x23\xb9\xab\x3c\x2c\x1c\x5b\xb4\xb6\x59\xe1\x2f\xec\xf1\x1d\xcc\x74\x7a\xa7\xcb\xa4\x5a\x0d\x2f\x3b\xe7\xf4\x91\xd3\x3d\xaf\x5a\xc1\xd1\xd4\xb5\xa9\x52\xc7\x84\x17\x5e\x8c\x26\xfe\x2d\x7b\x7b\x09\xb7\xea\xce\x68\x7e\xa8\x72\xe4\xbf\xe9\x30\x82\x4f\x6a\xa5\xf4\x5a\x75\x25\xaa\x3d\xe8\xbc\x30\xb5\xdb\xf9\xd6\x44\x5f\x86\x77\xbd\xd3\x3d\xaa\x42\x75\x0c\x51\xd5\xd3\x8e\xbb\xd5\xbc\xd3\xd4\x82\x2f\x7f\x02\x63\xea\xb4\x13\xd9\x3e\xc5\x07\x4b\x85\x2b\xef\xd3\xfa\xcb\x84\xc7\x8e\xbb\x45\xd0\xfd\xb7\x16\xbc\xd1\xaa\x0a\x2e\xb5\x33\x6e\x98\xa5\xa8\xa6\x0a\x97\x67\xb1\xda\x6e\x9c\x61\x7b\xf6\xc6\x15\xfa\x12\x34\xe4\x2d\x5a\x17\x9a\x2b\xf9\xbd\xa4\x54\x47\x26\xac\x7b\x30\x42\x59\xc6\xe4\x41\x0e\x54\x85\xb7\xc2\x3a\x70\x5c\x04\xd8\x9f\x23\x05\x4a\x5c\x05\x0a\x53\x5f\xd1\x41\x2b\x0c\xcb\xdd\xbd\xf1\xa1\x63\xdd\x87\xae\xf0\x2d\xee\x3c\xa6\xc2\xe1\x98\xa6\x7d\xb1\x69\x11\xd6\x7d\xe2\x2f\x30\x1c\x4d\xea\x5a\x58\xe2\xee\xac\xb3\x14\xd6\xab\x21\x9d\xa3\xb5\x62\x31\x0c\xdb\xab\xdd\x72\x20\x61\x70\xac\xbc\xe3\x4b\x59\x3b\x21\x33\xeb\xab\xbf\x74\xaf\xca\x12\x6b\xcb\xf9\xe2\x72\x22\x06\x85\xed\xde\x40\xdd\xdb\x9a\xf6\xdd\xab\x6d\xdd\x8a\xdd\x17\x96\xd7\xee\x35\x30\x7a\x86\x01\xbe\xdf\xb1\x53\x01\x99\x11\x4b\xb5\x9e\xc3\x83\x29\x71\x04\xdf\x89\xcc\x62\x65\x89\xff\xbc\xfa\x33\x15\x3e\x2f\x9c\xf2\x70\x41\xf9\x4e\xa5\x1d\x43\xed\xe3\x4e\xfb\xd3\x3d\xdb\x99\xc4\x78\xf0\x7b\x8a\x93\x38\xec\xeb\x8d\xe2\x76\x07\xc4\x6d\x92\x5c\x5b\xfe\x5c\x0e\x05\xfc\x8b\xed\xdd\x38\x43\x1b\x9f\x6e\x1d\x24\xda\x18\xb4\x85\xe6\x7d\x46\x7f\x4a\xee\xa3\xa6\x0b\x5b\x83\x32\x0a\x7b\xf7\xd2\xc6\x6f\xb8\x80\x56\x90\x97\xad\xdf\xf5\xa1\x36\xdb\xf8\x5a\x4a\x77\xb7\xe1\xab\x32\xfb\xd3\xd7\x4e\x54\xfe\xf5\x5f\x3a\x58\xd6\x76\x86\xd2\xca\xe9\xbd\x8b\x9e\xe8\x4b\x70\x26\x94\xdd\xb1\x4e\x1b\x32\x24\xb5\x2b\xe5\xac\xfa\x36\x4f\xe4\x7a\xd0\x0f\xf8\xed\xf7\xb3\xad\xaa\x88\x24\xc1\xc2\x61\x5a\xfb\x20\x99\xff\xda\xd9\xb9\x7f\x1c\x21\x7e\xc6\x8c\x7f\xd6\xbc\x2d\xfc\xf8\xd3\x99\x9f\x18\xd3\xcf\xf1\x5b\x64\x74\xf1\xff\x02\x00\x00\xff\xff\xa7\x15\x41\x92\x8b\x6e\x00\x00")

func appK8sIo_applicationsYamlBytes() ([]byte, error) {
	return bindataRead(
		_appK8sIo_applicationsYaml,
		"app.k8s.io_applications.yaml",
	)
}

func appK8sIo_applicationsYaml() (*asset, error) {
	bytes, err := appK8sIo_applicationsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "app.k8s.io_applications.yaml", size: 28299, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"app.k8s.io_applications.yaml": appK8sIo_applicationsYaml,
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
	"app.k8s.io_applications.yaml": &bintree{appK8sIo_applicationsYaml, map[string]*bintree{}},
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
