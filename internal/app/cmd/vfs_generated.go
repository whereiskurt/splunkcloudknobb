// Code generated by vfsgen; DO NOT EDIT.

// +build release

package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// CmdHelpEmbed statically implements the virtual filesystem provided to vfsgen.
var CmdHelpEmbed = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 4, 26, 23, 46, 57, 141359468, time.UTC),
		},
		"/backup": &vfsgen۰DirInfo{
			name:    "backup",
			modTime: time.Date(2020, 4, 26, 23, 29, 55, 505186519, time.UTC),
		},
		"/backup/backup.go": &vfsgen۰CompressedFileInfo{
			name:             "backup.go",
			modTime:          time.Date(2020, 4, 26, 23, 29, 39, 105125585, time.UTC),
			uncompressedSize: 1481,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x6f\x8b\xe3\x36\x10\xc6\x5f\x5b\x9f\x62\x6a\x38\xce\x3e\x7c\xb6\xf7\x7a\x9b\x42\x20\x2f\x52\x27\xcb\x16\x8e\xed\x92\xdd\x72\xd0\xe3\xe8\x29\xd2\xc8\x56\x23\x4b\x46\x92\x6b\x96\xf6\xbe\x7b\x91\x95\xec\xa6\x29\x85\x52\xe8\x2b\xeb\xcf\x68\xe6\xf9\x3d\x33\x1e\x28\x3b\xd0\x16\x61\x4f\xd9\x61\x1c\x08\x91\xfd\x60\xac\x87\x8c\x24\xa9\x71\x29\x49\x52\x2f\x7b\x4c\x09\x49\xd2\x56\xfa\x6e\xdc\x97\xcc\xf4\xd5\xd4\xa1\x45\xe9\x0e\xa3\xf5\x15\x53\x66\xe4\xcc\xd2\x49\xa1\xad\x86\x43\x5b\x31\xa3\x85\x6c\x2f\x9f\xb8\x41\x5c\x7d\x5b\x31\xb3\xb7\x34\x25\x39\x21\x55\x05\x4d\xcf\xa1\x33\x8a\x3b\xf0\x1d\x42\x7c\x36\x5a\xea\xa5\xd1\xc4\x3f\x0d\x38\x07\x38\x6f\x47\xe6\xe1\x77\x92\x34\x73\x00\xbc\x89\x81\x65\xdc\x92\x64\xf3\xf8\x00\x00\x21\x4e\xea\x96\x7c\x9d\x33\xdf\xe1\xf4\xfd\xcc\x03\x16\xfd\x68\xb5\x03\x0a\x1a\x27\x88\x87\x4d\xcf\x89\x18\x35\x7b\x09\xcb\x98\xb8\x4c\x9c\x43\xc6\xe0\x4d\xd3\xf3\x3c\xd4\x66\xb0\x0a\x09\xb2\xb0\x27\x09\x3b\xc6\xc0\x0a\x98\x68\xc3\x3e\x88\x58\x41\xb0\xaa\xbc\x33\x53\x96\x97\x37\xc6\xf6\xd4\x67\xe9\xbb\xba\x5e\xd4\x57\xf5\xbb\xc7\xab\xeb\xfa\x7d\x7d\x9d\xe6\x24\x89\x92\x80\x1d\xb5\xde\xa2\x1a\x6e\xa9\xe6\x0a\x2d\x50\xa5\xc0\x88\xd9\x8e\x83\x36\x93\x42\xde\x22\x98\xfd\xaf\xc8\xbc\x83\x4c\x62\x09\xc0\xa9\xeb\xf6\x86\x5a\xee\x0a\xb0\x18\x7a\xe5\x0a\x50\xc6\x04\x58\x21\x15\xba\x02\xca\x32\x8f\x7c\xcf\x04\x67\x35\x32\xd6\xf3\x80\xba\xb7\xb4\x6c\x4c\xdf\x53\xcd\x0b\xa0\xb6\x75\xf0\xe9\x73\xf4\x30\x02\x1f\x11\xcb\xe6\xc3\x0f\xe5\x83\xe7\x68\x6d\x96\xfe\xe4\x68\x8b\xd1\xb1\xb4\x00\x2d\x55\x7e\x64\x58\x2b\xf5\xbf\x69\x5f\x2b\xf5\xaf\x35\x6f\x4e\x05\xc2\x8b\x18\x32\xb7\x6b\x37\x17\xbb\x38\x7c\x40\x6a\x59\x77\x2b\x9d\x37\xf6\xe9\xfc\x6e\x66\xf2\x66\x64\xdd\x8d\x54\x08\x93\x54\x0a\xc6\x81\x53\x8f\x33\x5d\x5c\xc8\x1e\xc1\x79\xda\x0f\x27\xe8\x00\x00\x03\x75\x0e\xf9\x05\xc1\x73\xaa\x2c\xc4\x68\xda\x63\x01\xdc\x3b\x38\x93\xee\x0b\x40\x58\x1e\x27\xe8\x9e\x5a\x87\x71\x74\xde\xd6\x57\x6f\xe7\xe1\x59\xd6\xef\x97\xf5\xf5\xcf\xf5\x77\xcb\xba\x4e\xe7\xe7\x39\x49\xa4\x00\x84\x6f\x56\xa1\x13\x21\xc9\x4b\xd3\x3e\x98\xb6\xbc\xa1\x9e\x2a\x91\x7d\x11\xe1\xbb\x04\x46\xb5\x36\x1e\x86\x90\xfb\x0c\xe1\xf5\x2b\xf7\xfa\xcb\x29\xdf\x57\x92\x20\xac\xc0\xb8\xb2\xe9\xc2\xad\x3b\x13\xec\x0b\xf0\xff\xad\xa4\x43\x0f\x8a\x3a\x0f\xbd\xe1\x52\x48\xe4\x15\x65\x8c\xa1\x8b\x02\x8c\xf8\x9b\x86\xf8\x83\x84\x36\xfc\xd5\xc7\xc9\x4a\x8f\x8d\xd1\x1e\xb5\xcf\x58\xfc\x1e\x4d\x2c\xe0\xa4\xf4\xdc\x55\x51\x00\x5a\x1b\x7c\x35\xae\xfc\x71\x40\x7d\xd1\x84\x70\xfa\xcb\xfa\xfe\x7e\x7b\xb7\xf9\x63\x5e\x37\xbb\xed\xfa\x71\x1b\xd7\xbb\xcd\xc7\x5d\x01\xf5\x62\xb1\x38\x62\x5b\xfb\x8f\xe0\x5b\x6b\x8d\x15\x59\x8a\xe1\x0b\x66\x40\x2d\x75\x1b\x27\x22\xc0\x81\x30\x76\x56\x2f\x75\xbb\x84\x57\xbf\xa5\x2f\x72\x67\x85\x39\x49\x4e\xd0\x81\x9f\xa3\x40\x0b\xa2\x6c\x94\x71\x98\xe5\x84\x24\xa2\xfc\x18\xd8\xb3\x4f\x9f\xf7\x4f\x1e\x4f\xec\x79\xb8\x7a\x36\xeb\xcf\x00\x00\x00\xff\xff\x8d\x22\x0d\x2f\xc9\x05\x00\x00"),
		},
		"/backup/backup.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "backup.tmpl",
			modTime:          time.Date(2020, 4, 21, 5, 16, 48, 940566514, time.UTC),
			uncompressedSize: 178,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x0a\xc2\x30\x18\x45\xf7\xef\x29\x2e\xc5\xd1\xbe\x80\xe0\xd0\x6a\x07\xc1\xfe\x40\x2d\x08\xe2\x90\x34\x9f\xb5\x54\xd3\x90\xc4\x41\x42\xdf\x5d\x5a\x5d\x5c\xef\x39\x17\x0e\x85\x00\xc5\xb7\x5e\x33\xa2\xc6\x89\x8e\x53\xd1\x0e\x2f\x13\x61\x9a\x88\x42\x88\xe1\xf9\x69\x1e\xc2\x33\xa2\xc4\x98\xc5\x58\xd8\xae\xcc\xf3\xa4\xd8\x6f\x08\x00\xe4\xf2\x21\xaa\x9b\xf4\x6f\x57\xc2\xdd\xe5\x28\xac\x5a\xc3\xb2\x19\xad\x77\x44\x65\x75\x3a\x94\x45\xfd\x15\xe2\x78\xe0\xf7\xf6\xe2\xbc\xed\x75\x77\x05\x51\x76\x4e\xf2\xea\x98\xfd\xf0\x0a\xae\x1d\xf4\x28\x25\xe6\x16\xd6\x6a\x8e\xfa\x04\x00\x00\xff\xff\xc6\x92\x1d\xbe\xb2\x00\x00\x00"),
		},
		"/backup/dashboard.go": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.go",
			modTime:          time.Date(2020, 4, 26, 23, 29, 23, 813068742, time.UTC),
			uncompressedSize: 2727,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xdd\x6e\xf3\xb4\x1b\x3f\x4e\xae\xe2\x79\x23\x55\x4d\xde\xe5\xef\x6e\x7f\x90\x26\x55\xda\xc1\xe8\x98\xf4\xa2\x02\xd3\xca\xc4\x01\x43\xc8\x71\x9c\xc4\xc4\xb1\x23\xdb\x59\xde\x09\xed\x9c\x6b\xe0\xf2\xb8\x12\xf4\x38\x6d\x9a\x6e\xad\x86\x04\x9c\xb4\xf9\xb0\x9f\xe7\xf7\x65\xc7\x2d\x65\x35\x2d\x39\x64\x94\xd5\x5d\x1b\x86\xa2\x69\xb5\x71\x10\x87\x41\x54\x34\x2e\x0a\x83\x48\x5b\xfc\x6d\xa9\xab\x16\x85\x90\x1c\x2f\xa2\x30\x0c\xa2\x52\xb8\xaa\xcb\x08\xd3\xcd\xa2\xaf\xb8\xe1\xc2\xd6\x9d\x71\x0b\x26\x75\x97\x33\x43\x7b\xc9\xcd\xa2\xad\xcb\x85\xe5\xe6\x49\x30\xfe\x7a\x8e\x6d\x8b\x8b\x2f\x16\x4c\x67\x86\x46\x61\x12\x86\x8b\x05\xdc\x50\x5b\x65\x9a\x9a\x1c\xb2\x5e\x48\xb9\xc5\x04\x54\x4a\xd0\x05\xb8\x8a\x43\x3e\x8e\x70\x1a\x28\x48\xcd\xa8\x04\x44\x15\x16\x9d\x62\x10\x33\xf8\xb8\x6a\xf2\x64\x5f\x29\x66\x4d\x0e\x1f\x7d\x17\xb2\xd2\x4d\x43\x55\x9e\x02\x35\xa5\x85\x9f\x7e\xb6\xce\x08\x55\x26\xf0\x5b\x18\x30\xad\x0a\x51\xc2\xf2\x0a\x18\x59\xf9\xeb\x30\x90\x7a\x78\xe0\x6f\xc9\x5a\x97\xe1\x6e\x1c\x59\xad\x3f\x91\x3b\xa3\x9b\xd6\xc5\xd1\xa3\x8a\x92\xa3\x2f\x36\x8e\x1a\x27\x54\xb9\xa3\x51\x68\xb3\xc7\x6f\x09\x21\x8f\xca\xcf\x0d\x03\xda\xb9\x2a\x05\x6e\x0c\xf6\xdb\xca\x85\x0d\x85\x8a\xb7\x75\x1f\xee\xd7\xe9\x0e\xc9\x83\xe5\x46\xd1\x86\x8f\x0f\xee\xa8\xb5\xbd\x36\xf9\xf8\x60\xa5\x75\x2d\xf8\x9d\x36\x2e\x9d\xc0\x4f\xc2\x40\x14\xbe\xcb\x87\x2b\x50\x42\x22\x6f\x24\x49\x6e\xa9\xa3\xb2\x88\xa3\x02\xff\x97\xb0\x69\x65\xa7\x6a\x58\xa1\x8f\x70\xdd\xb9\x8a\x2b\x27\x18\x75\x42\xab\x25\xcc\xce\x9e\x22\x0f\x35\x09\x83\x17\xaf\x11\xf9\xa4\x0a\x5d\xc4\xd1\xa6\x63\x8c\x5b\x5b\x74\x12\x24\x42\x47\x87\x84\xb2\x8e\x2a\xc6\x97\x8f\x0a\x60\x66\xa3\x14\x90\x2a\xb2\x39\x2a\x59\xd1\x38\xb2\x69\x8d\x50\xae\x88\xa3\x3f\x7f\xff\x03\xde\xd4\xec\x85\xab\x60\x3e\xb3\xf3\x37\xd5\x0f\xea\x8f\x0a\x8d\xed\xe0\x0c\xa2\xad\xda\xac\xa2\x2a\x47\xa5\x1b\x5a\xf3\x18\xef\x46\xcd\xc7\xd8\x24\x61\x50\x6a\xc0\x4c\xc5\xc9\xa8\xd3\x96\xe8\x3d\x77\x46\xf0\x27\x74\x76\x1f\xd8\xb5\xb0\xde\x6b\x34\xf9\x7a\xbd\x9e\x18\x0d\x84\x10\xc0\x88\x04\xaf\x0d\x16\xd6\xed\x73\x3a\x64\xc0\x63\xc3\xb1\x6f\x9d\x3a\x66\x15\xd3\x9d\xcc\xd5\xdc\x81\x19\x30\x6d\x17\x08\xd5\xd8\x57\x0a\xeb\xec\x72\x90\x65\x30\x0c\x1d\x7b\x89\x93\x53\x39\x46\xc5\x6f\xb9\x63\x15\xcf\x27\x0b\x4d\xbe\xcb\x2c\x0c\x0a\x2d\x73\xee\xc9\xed\x76\x08\xf2\x8d\x16\xea\xd0\xd0\x99\x25\x96\xd5\x4a\x67\x59\x94\x02\x23\x37\x3f\x6c\x92\x14\xa2\xb1\x1c\x4a\x44\x33\x9b\xc2\x2f\x07\x75\xae\x33\x1b\x0f\xe5\x93\x83\xb8\xad\x0c\xa7\xd3\xd5\xb5\x07\xbc\x05\x83\x01\xc8\x31\x2b\x33\x36\xc7\x60\x60\x69\x6d\xc9\x1d\x75\xd5\x86\xb7\xd4\x50\xa7\xb1\xa4\xb6\xe4\xdb\x3a\x17\xe6\x5a\xca\xd8\x8f\x39\xbf\xbc\xbc\x44\x4e\x86\xf6\xff\x0a\xad\x14\x22\x43\x7b\x64\x67\x68\x7f\x82\xe0\xd8\xeb\x9f\x71\xdc\x35\x78\x87\xe6\x6e\xd8\x8e\xe9\x3b\x0b\x71\x08\xc6\x57\xbc\x14\x4a\x21\x18\xa7\xa1\x37\xc2\x71\x4f\xc1\xee\xa0\x0c\x20\x80\x9c\x94\x1a\x3b\x3d\x51\x03\x4e\x3b\x2a\xb3\x67\xc7\x2d\x5c\xc1\xf9\xf0\x8c\xe9\x4e\xb9\xe1\xd6\xef\x92\x28\x90\xa1\xaa\xe4\xc3\x82\xf0\xf9\x3f\x12\x59\xe2\xc3\x17\x20\x0e\x54\xe2\x73\x23\x8f\x78\xe5\xe1\xa5\xf0\xda\xb3\xcf\x8d\x8c\x52\xc8\xc9\x77\xb4\xe1\x09\xae\x0d\x46\x3c\xab\x95\x56\x8e\x2b\x17\xe7\x64\x7b\x95\xc2\xa4\xfe\x30\xd0\xe9\x8e\x55\xb7\x42\xf2\x78\xf2\x0a\x8b\x3d\xb4\x39\x75\x3c\x3f\x40\xf5\xab\xd5\xea\x34\x2c\x9f\x8d\xb7\xe8\x70\xd2\x3b\xf0\xee\x69\xff\xb5\x72\xe6\x79\x8f\x0f\x27\x9d\x02\x88\xef\x5e\x23\x3c\x70\x62\x72\x73\x06\x92\xab\x49\x87\x64\x71\x71\xfe\xff\x2f\xb7\x9b\xe0\x0d\xcf\xba\xd2\x83\x4c\x61\x96\xd7\x59\x0a\xfe\x0a\x37\x99\x03\x31\x8e\x95\x40\x00\xdf\xf7\x0a\x79\x4f\x90\xa0\xb5\x83\xfd\xc3\xff\x19\x5c\xe0\xb7\x65\xbf\x12\x8e\xc5\xf1\x47\xa3\x1d\x87\xf9\x2c\x9f\x4f\x37\xa4\xdd\x3b\x31\xec\x58\x3e\x9f\x4e\x1f\x46\xf4\x11\x85\xf5\x9d\x4e\x27\xf5\xef\x2d\x89\xff\x16\xc3\xc1\x4e\x70\xa3\x15\x27\xfb\x86\x75\x36\x1f\x0c\xc3\x73\x11\x1b\x22\x41\x60\xf9\xbf\x24\x4a\x27\x46\xbe\xcf\x03\x72\xad\xf8\x07\x9c\x08\x78\x18\x99\x7c\x72\x3f\x1c\x36\xa3\xcc\x68\x6b\x5f\x71\x1d\x0f\x36\xc8\x91\xec\x0a\xac\xb4\x2a\x0d\x75\x9d\xf4\x27\x06\x0b\xcf\xba\x03\xa5\x7b\xa8\xe8\x13\x1f\x8f\x6c\xc3\x4c\xdf\xd9\x9f\x82\xa6\xb0\xb7\xca\x78\x09\x0c\x77\x9d\x51\xe1\x4b\xf8\x57\x00\x00\x00\xff\xff\xc3\xdd\xdc\xd5\xa7\x0a\x00\x00"),
		},
		"/backup/report.go": &vfsgen۰CompressedFileInfo{
			name:             "report.go",
			modTime:          time.Date(2020, 4, 26, 23, 29, 2, 532989599, time.UTC),
			uncompressedSize: 2672,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xcd\x6e\xec\xb4\x17\x5f\x4f\x9e\xe2\x34\xd2\x68\x92\xdb\xfc\x3d\xed\x1f\xa4\x4a\x23\x75\x51\xa6\x54\xba\x68\x80\xaa\x43\xc5\x82\x22\xe4\x71\x9c\xc4\xc4\xb1\x23\xdb\x69\xa8\x50\xf7\x3c\x03\x8f\xc7\x93\xa0\xe3\x7c\x34\xb9\x9d\x52\xa4\x2b\x36\x33\x49\x6c\x9f\xf3\xfb\x72\xe2\x9a\xb2\x92\xe6\x1c\x0e\x94\x95\x4d\x1d\x04\xa2\xaa\xb5\x71\x10\x05\x8b\x30\xab\x5c\x18\x2c\x42\x6d\xf1\xb7\xa6\xae\x58\x67\x42\x72\xbc\x08\x83\x60\x11\xe6\xc2\x15\xcd\x81\x30\x5d\xad\x6d\x9d\x9d\x7f\xb1\x66\xfa\x60\x68\x38\x1f\x69\x0b\x6e\xb8\xb0\x65\x63\xdc\x9a\x49\xdd\xa4\xcc\xd0\x56\x72\xb3\xae\xcb\x7c\x6d\xb9\x79\x14\x8c\x87\x41\x1c\x04\xeb\x35\xdc\x71\xdf\xba\x15\x52\xf6\x70\x80\x4a\x09\x3a\x03\x57\x70\x30\x7e\xd4\x82\xd3\x40\x41\x6a\x46\x25\x20\x9c\x20\x6b\x14\x83\x88\xc1\x87\x6d\x95\xc6\x7d\x8d\x88\x55\x29\x7c\xf0\x78\xc8\x56\x57\x15\x55\x69\x02\xd4\xe4\x16\x7e\xfa\xd9\x3a\x23\x54\x1e\xc3\xef\xc1\x82\x69\x95\x89\x1c\x36\x97\xc0\xc8\xd6\x5f\x07\x0b\xa9\xbb\x07\xfe\x96\xec\x74\x1e\x0c\xf3\xc8\x76\xf7\x91\xdc\x1a\x5d\xd5\x2e\x0a\x1f\x54\x18\x1f\x1d\xd8\x3b\x6a\x9c\x50\xf9\xc0\x20\xd3\xa6\x07\x65\x09\x21\x0f\xca\x2f\x0c\x16\xb4\x71\x45\x02\xdc\x18\x6c\xd6\xeb\x80\xdd\x84\x8a\xfa\xa2\xf7\x77\xbb\x64\x80\x71\x6f\xb9\x51\xb4\xe2\xe3\x83\x5b\x6a\x6d\xab\x4d\x3a\x3e\xd8\x6a\x5d\x0a\x7e\xab\x8d\x4b\x26\xd8\xe3\x60\x21\x32\xdf\xe5\xe4\x12\x94\x90\x48\x1a\x19\x92\x1b\xea\xa8\xcc\xa2\x30\xc3\xff\x0d\xec\x6b\xd9\xa8\x12\xb6\x68\x10\x5c\x35\xae\xe0\xca\x09\x46\x9d\xd0\x6a\x03\xcb\xd3\xc7\xd0\x43\x8d\x83\xc5\xb3\x17\x88\x7c\x54\x99\xce\xa2\x70\xdf\x30\xc6\xad\xcd\x1a\x09\x12\xa1\xa3\x37\x42\x59\x47\x15\xe3\x9b\x07\x05\xb0\xb4\x61\x02\x48\x15\xd9\x1c\xd5\x2b\xab\x1c\xd9\xd7\x46\x28\x97\x45\xe1\x5f\x7f\xfc\x09\xaf\x6a\xb6\xc2\x15\xb0\x5a\xda\xd5\xab\xea\xb3\xfa\xa3\x42\x63\x3b\x38\x85\xb0\x57\x9b\x15\x54\x79\xa5\x2b\x5a\xf2\x08\xef\x46\xcd\x3b\x6f\xe2\x60\x91\x6b\xc0\x28\x45\xf1\x28\x52\xcf\xf2\x8e\x3b\x23\xf8\x23\x7a\xda\x1b\x09\x3b\x61\xbd\xc7\x68\xee\xd5\x6e\x37\x66\x93\x10\x02\x98\x8b\xc5\xa7\xc6\x0a\xeb\xfa\x64\x76\xc6\x7b\x40\x38\xf1\xb5\x3d\xc7\xfc\x61\xba\x91\xa9\x5a\x39\x30\x1d\x96\x61\x37\x80\x14\xd6\xd9\x4d\xa7\x43\xe7\x10\x5a\x74\x3c\xb0\xa8\xee\x0d\x77\xac\xe0\xe9\x74\xf9\x3f\xf1\x78\x8e\x50\xbd\x4c\xcb\x94\x7b\x3e\xc3\xfe\x27\xdf\x68\xa1\xe6\xde\x2d\x2d\xb1\xac\x54\xfa\x70\x08\x13\x60\xe4\xfa\x87\x7d\x9c\x40\xd8\x15\xc4\x52\xf4\x60\x13\xf8\x65\x56\xe4\xea\x60\xa3\xae\x76\x3c\x8b\xd5\xd6\x70\x3a\xdd\x42\x3d\xd8\x1e\x06\xba\x9c\x62\x20\x96\x6c\x85\xee\x63\x5d\x6d\xc9\x2d\x75\xc5\x9e\xd7\xd4\x50\xa7\xb1\x9e\xb6\xe4\xdb\x32\x15\xe6\x4a\xca\xc8\xcf\x39\xbb\xb8\xb8\x40\x36\x86\xb6\x9f\x4f\x08\xaf\x68\x8b\xbc\x0c\x6d\xdf\xa0\x36\x36\xfa\x0c\x76\x43\xf5\x77\x08\x0e\xd3\x06\x8e\xef\xec\xb3\x2e\x0b\x5f\xf1\x5c\x28\x85\x48\x9c\x86\xd6\x08\xc7\x3d\x7e\x3b\x40\xe9\x40\x00\x79\x53\x64\xec\xf4\x48\x0d\x38\xed\xa8\x3c\x3c\x39\x6e\xe1\x12\xce\xba\x67\x4c\x37\xca\x75\xb7\x18\x2e\x2f\xb6\xa1\x2a\xe7\x5d\xf4\x7d\xd2\x8f\xa4\x94\xf8\xed\xba\x40\x1c\xa8\xc4\x6f\x95\x3c\xe2\x92\x87\x97\xc0\x2b\xb7\x6a\x89\x8a\x91\xef\x68\xc5\x63\xdc\x09\x8c\x78\x56\x5b\xad\x1c\x57\x2e\x32\x64\xcf\xa9\x61\x45\x02\x93\xf2\xdd\x3c\xa7\x1b\x56\xdc\x08\xc9\xa3\xc9\x10\xd6\xba\xaf\x53\xea\x78\x3a\x03\xf5\xab\xd5\xea\x6d\x54\x3e\x17\xaf\xc1\xe1\xa2\x77\xd0\xdd\xd1\xf6\x6b\xe5\xcc\xd3\x0b\x3e\x5c\xf4\x16\x40\x1c\xfb\x14\xe1\xcc\x88\xc9\xcd\x29\x48\xae\x26\x1d\xe2\xf5\xf9\xd9\xff\xbf\xec\xdf\x72\xd7\xfc\xd0\xe4\x1e\x64\x02\xcb\xb4\x3c\x24\xe0\xaf\xf0\x8d\x32\x13\xe3\x58\x09\x04\xf0\x7d\xab\x90\xf7\x04\x09\x3a\xdb\xb9\xdf\xfd\x9f\xc2\x39\x7e\x39\x5e\x76\xc1\xb1\x34\xfe\x68\xb4\xe3\xb0\x5a\xa6\xab\xf1\x2d\x34\x0c\x88\xee\x1d\xe5\xb3\xe9\xf4\x3c\x9e\x0f\xa8\xaa\x6f\xf3\x76\x4a\xff\xdd\x76\xf8\x0f\x01\xcc\xf6\xff\xb5\x56\x9c\xbc\x74\x2b\x0f\xab\xce\x2a\x3c\xe7\xb0\x2e\x0c\x04\x36\xff\x8b\xc3\x64\x62\xe1\xfb\x24\x00\x52\xad\xf8\x09\xae\x04\x3c\x65\x4c\xbe\xa5\x27\xf3\x6e\x94\x19\x6d\xed\x94\xe9\x78\x56\x41\x86\x64\x58\xbd\xd5\x2a\x37\xd4\x35\xd2\x9f\x03\x2c\x3c\xe9\x06\x94\x6e\xa1\xa0\x8f\x7c\x3c\x82\x75\x2b\x7d\x5b\x7f\xb6\x99\x82\xee\x75\xf1\x02\x18\xee\x1a\xa3\x82\xe7\xe0\xef\x00\x00\x00\xff\xff\xc1\xa4\xe9\x63\x70\x0a\x00\x00"),
		},
		"/backup/searchhistory.go": &vfsgen۰CompressedFileInfo{
			name:             "searchhistory.go",
			modTime:          time.Date(2020, 4, 26, 23, 29, 55, 511000000, time.UTC),
			uncompressedSize: 2758,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xdd\x6a\xe4\x36\x14\xbe\xb6\x9f\xe2\xac\x61\x18\x7b\xe3\x6a\x76\xdb\x42\x60\x20\x17\xe9\xa4\xa1\x29\x29\x84\xcc\x2e\xbd\x68\x4a\x91\x65\xd9\x56\x2d\x4b\x46\x92\xe3\x86\x25\xf7\x7d\x86\x3e\x5e\x9f\xa4\x1c\xd9\x9e\x8c\x77\x66\xc8\x42\x4b\x6f\x66\x6c\xfd\x9c\xf3\x7d\xdf\xf9\x8e\xac\x96\xb2\x9a\x96\x1c\x32\xca\xea\xae\x0d\x43\xd1\xb4\xda\x38\x88\xc3\x20\x2a\x1a\x17\x85\x41\xa4\x2d\xfe\xb6\xd4\x55\xab\x42\x48\x8e\x0f\x51\x18\x06\x51\x29\x5c\xd5\x65\x84\xe9\x66\x65\xdb\xe2\xfd\x37\x2b\xa6\x33\x43\xa3\xf9\x4c\x5f\x71\xc3\x85\xad\x3b\xe3\x56\x4c\xea\x2e\x67\x86\xf6\x92\x9b\x55\x5b\x97\x2b\xcb\xcd\xa3\x60\x3c\x0a\x93\x30\x5c\xad\x60\xcb\xa9\x61\xd5\x0f\xc2\x3a\x6d\x9e\xa0\x17\x52\x42\xdb\x49\xe9\xa1\x85\x45\xa7\x18\xc4\x0c\xde\x6e\x9a\x3c\x99\x2f\x8d\x59\x93\xc3\x5b\x9f\x9d\x6c\x74\xd3\x50\x95\xa7\x40\x4d\x69\xe1\x97\x5f\xad\x33\x42\x95\x09\x7c\x0a\xc3\x80\x69\x55\x88\x12\xd6\x17\xc0\xc8\xc6\x3f\x87\x81\xd4\xc3\x80\x7f\x25\xb7\xba\xdc\xad\x23\x9b\xdb\x1b\x72\x67\x74\xd3\xba\x38\x7a\x50\x51\x72\x74\x62\xeb\xa8\x71\x42\x95\xa3\x7e\x50\x68\x33\x82\x83\x11\x1d\x21\xe4\x41\xf9\xfd\x61\x40\x3b\x57\xa5\xc0\x8d\xc1\x9c\x23\x7b\x4c\x2a\x54\x3c\xc6\xfe\x78\x7f\x9b\x4e\x68\x3e\x5a\x6e\x14\x6d\xf8\x6e\xe0\x8e\x5a\xdb\x6b\x93\xef\x06\x36\x5a\xd7\x82\xdf\x69\xe3\xd2\x3d\x0a\x49\x18\x88\xc2\x67\x79\x73\x01\x4a\x48\xf8\x14\x06\x48\x94\x5c\x53\x47\x65\x11\x47\x05\xfe\xaf\x61\xdb\xca\x4e\xd5\xb0\xc1\xb2\xc0\x65\xe7\x2a\xae\x9c\x60\xd4\x09\xad\xd6\xb0\x38\x7b\x8c\x3c\xd4\x24\x0c\x9e\xbd\x4e\xe4\x46\x15\xba\x88\xa3\x6d\xc7\x18\xb7\xb6\xe8\x24\x48\x84\x0e\x4e\x83\x50\xd6\x51\xc5\xf8\xfa\x41\x01\x2c\x6c\x94\x02\x52\x45\x36\x47\x65\x2b\x1a\x47\xb6\xad\x11\xca\x15\x71\xf4\xf7\x9f\x7f\xc1\x41\xcc\x5e\xb8\x0a\x96\x0b\xbb\x3c\x88\x3e\x8b\xbf\x53\x68\x97\x0e\xce\x20\x1a\xd5\x66\x15\x55\x5e\xe9\x86\xd6\x3c\xc6\xb7\x9d\xe6\x33\xff\xdc\x73\xdb\x49\x97\x84\x41\xa9\x01\x6d\x16\x27\x3b\xc5\x46\xca\xf7\xdc\x19\xc1\x1f\xb1\xce\xf3\xe2\xfa\x7a\x5f\xde\xde\x82\xf5\xc3\xdc\x02\x21\x04\xd0\x2b\xc1\xe7\x55\x16\xd6\xcd\x5d\x3b\x98\xc1\x83\xc4\xf5\x87\x25\x3b\x56\x33\xa6\x3b\x99\xab\xa5\x03\x33\x40\xe2\x63\x66\xa8\x46\x40\x52\x58\x67\xd7\x83\x46\x43\xf5\xb0\x7c\xc7\x3d\x8d\xca\x5f\x73\xc7\x2a\x9e\x1f\x0b\x83\x74\xe7\xfc\x76\xd3\x13\xcb\xe7\x18\x85\x2e\xb4\xcc\xb9\x67\x3b\x1d\x10\xe4\x47\x2d\xd4\xbc\xcc\x0b\x4b\x2c\xab\x95\xce\xb2\x28\x05\x46\xae\x3e\x6c\x93\x14\xa2\x21\x2e\x86\xa2\x99\x4d\xe1\xb7\x59\x90\xcb\xcc\xc6\x43\xec\x64\xe6\xc0\x8d\xe1\x74\xbf\xe9\x3e\x03\x37\xc2\x41\x63\xe4\xe8\xa1\x05\x5b\xa2\x61\x30\xbe\xb6\xe4\x8e\xba\x6a\xcb\x5b\x6a\xa8\xd3\x18\x57\x5b\xf2\x53\x9d\x0b\x73\x29\x65\xec\xd7\xbc\x3b\x3f\x3f\x47\x56\x86\xf6\xff\x9e\x58\x0a\x91\xa1\x3d\xf2\x33\xb4\x3f\x41\x71\x97\xe8\x3f\x60\x39\x65\x79\x85\xe8\xb4\x6c\xe2\xfa\x4a\x8b\x0e\x56\xf9\x8e\x97\x42\x29\x44\xe4\x34\xf4\x46\x38\xee\x79\xd8\x09\xca\x00\x02\xc8\x49\xb1\x31\xd3\x23\x35\xe0\xb4\xa3\x32\x7b\x72\xdc\xc2\x05\xbc\x1b\xc6\x98\xee\x94\x1b\x5e\xd1\x73\x5e\x74\x43\x55\xc9\x87\x0e\xf1\x0d\x71\xc4\xc4\xc4\x77\x7a\x80\x38\x50\x89\x3f\x1a\x79\xa4\x5a\x1e\x5e\x0a\x07\x55\x6b\x25\x2a\x36\x1e\x06\x37\x57\x09\x36\x0b\x23\x9e\xd9\x46\x2b\xc7\x95\x8b\xa7\xd9\x14\xf6\x52\x0c\xeb\x9c\xee\x58\x75\x2d\x24\x8f\xf7\xa6\x30\xde\x07\xd1\xf0\x19\xaa\xdf\xad\x56\xa7\x61\x79\x83\x1c\xa2\xc3\x4d\x5f\x00\xef\x9e\xf6\xdf\x2b\x67\x9e\x5e\x00\xe2\xc6\x53\x08\x71\x6e\x06\x71\x56\x8a\xbd\x97\x33\x90\x5c\xed\x85\x4f\x56\xef\xdf\x7d\xfd\xed\x78\x2a\x5e\xf1\xac\x2b\x3d\xca\x14\x16\x79\x9d\xa5\xe0\x9f\xf0\xc8\x99\x49\x71\x2c\xc4\xc1\xc9\xbd\x47\xd0\x97\x78\xb0\xc1\xf0\x7f\x06\xef\xf1\xeb\xf3\xd2\x16\xc7\x6c\xf9\xb3\xd1\x8e\xc3\x72\x91\x2f\xe7\x2d\x22\xb8\x85\x69\x85\x18\x0e\x33\xef\x56\xa7\xe7\x86\x7d\x40\x99\x7d\xbe\xd3\xbe\xfd\xb2\x06\xf9\x3f\x90\xcc\x8e\x88\x2b\xad\x38\x79\x49\x5b\x67\xcb\xa1\x86\xa0\x0b\xbc\x0f\xa0\x45\x08\xac\xbf\x4a\xa2\x74\xaf\xb6\xaf\xb3\x01\xc8\xb5\xe2\x6f\x70\x27\xe0\xdd\x65\xef\x0b\xfd\x66\x9e\x8d\x32\xa3\xad\x3d\x41\x79\xba\x12\x21\x57\x32\xc5\xd9\x68\x55\x1a\xea\x3a\xe9\xef\x19\x16\x9e\x74\x07\x4a\xf7\x50\xd1\x47\x0e\x14\xa4\x66\x54\x8e\x3b\x3d\x00\x7f\x77\xda\x87\x3f\x2a\xe4\xa5\x30\xdc\x75\x46\x85\xcf\xe1\x3f\x01\x00\x00\xff\xff\x14\xce\xe6\xca\xc6\x0a\x00\x00"),
		},
		"/restore": &vfsgen۰DirInfo{
			name:    "restore",
			modTime: time.Date(2020, 4, 18, 4, 49, 17, 392213227, time.UTC),
		},
		"/restore/restore.go": &vfsgen۰CompressedFileInfo{
			name:             "restore.go",
			modTime:          time.Date(2020, 4, 25, 3, 48, 10, 811136991, time.UTC),
			uncompressedSize: 160,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcb\xc1\x0d\x83\x30\x0c\x85\xe1\x33\x9e\xc2\xe2\xd8\x03\x19\x82\x6e\xd0\x2e\x10\x1c\x37\x44\x90\x38\x72\x1c\xa1\xaa\xea\xee\x15\xe5\xf8\xf4\xfe\xaf\x7a\xda\x7c\x64\x54\x6e\x26\xca\x00\x29\x57\x51\xc3\x31\x26\x5b\xfb\x32\x91\x64\x77\xac\xac\x9c\xda\xd6\xd5\x1c\xed\xd2\x03\xa9\x3f\x76\x56\x57\xb7\xe8\x48\xca\x2b\xc5\x11\xc0\xb9\x39\x07\x0c\xdc\x48\xd3\xc2\x0d\x6d\x65\x24\xc9\xd9\x97\x00\xf6\xae\x8c\xe7\xdd\x4c\x3b\x19\x7e\x60\x98\xff\x0e\x6f\x97\x9f\xae\x09\xc3\xfd\xf9\x40\xc4\xb3\x4b\x25\xc2\x17\x7e\x01\x00\x00\xff\xff\xc2\x49\x53\xa4\xa0\x00\x00\x00"),
		},
		"/restore/restore.tmpl": &vfsgen۰FileInfo{
			name:    "restore.tmpl",
			modTime: time.Date(2020, 4, 18, 4, 49, 17, 392213227, time.UTC),
			content: []byte(""),
		},
		"/scknobb.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "scknobb.tmpl",
			modTime:          time.Date(2020, 4, 21, 5, 15, 51, 256443067, time.UTC),
			uncompressedSize: 1848,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x6d\x6f\x1a\x39\x10\xfe\xee\x5f\xf1\xa4\xaa\x94\x22\x05\x56\x6d\x7a\x5f\xd0\xbd\x28\x21\xaf\xba\x04\xa2\x42\x7a\x3a\x01\x5a\x99\xdd\x61\xd7\xc7\xae\xbd\xf1\x0b\x1c\x02\xfe\xfb\xc9\xde\x0d\x24\x39\x55\xed\x87\x3a\xd2\x64\xec\xf5\xf3\xcc\xd8\xf3\x78\xd8\x6c\x90\xd2\x5c\x48\xc2\xbb\xb3\xaa\xba\x21\x9e\x92\x7e\x87\xdd\x0e\xfb\x11\xe3\xcd\x88\x5f\x58\x86\x38\x8e\xe1\xcd\x16\xdb\xe0\xfb\x69\x70\xb6\xd8\xee\xad\xdf\x19\x21\x8e\xb7\xc1\x60\x1b\x21\xc2\x71\x8c\x09\x22\xc4\x98\x6c\x83\x5f\x5b\x60\xb3\xe9\x7c\xa1\x82\xb8\xa1\x0b\x6e\x69\xb7\x63\x93\xd8\xaf\x7f\xf0\x40\xe0\x57\x1f\xc7\xff\x7d\x88\x5b\x81\xfc\x60\x0f\xc0\xaf\xa4\x8d\x50\x72\xb7\x03\xdb\xc6\x71\x1c\x4d\x7c\x7e\xf1\x76\x12\x4f\x7c\x70\xef\xc4\x71\xb4\x8d\x3b\x7b\x1b\xa2\x5e\x0b\x7b\xc3\x4d\xee\x51\x6f\x8f\xfc\xcd\xc1\x7e\xc7\xb0\x2a\x9c\x5c\xa0\x57\x28\x97\xe2\xcf\xbe\x5a\x15\x94\x66\x84\xc1\xf9\x3f\x94\x58\x9c\xf3\x64\xe1\x2a\x30\x76\x26\xd1\xbb\xbb\x85\x90\x96\xf4\x9c\x27\xe4\x3d\x05\x9b\xd3\x6b\x02\x67\x84\xcc\x70\xad\x8e\x18\xbb\x12\x32\x45\xa9\xb4\xdf\x3a\x57\xba\xe4\x56\x28\x09\x6e\xbb\x21\xbd\xdc\xda\xca\x74\xa3\x28\x13\x36\x77\xb3\x4e\xa2\xca\x68\x95\x93\x26\x61\x16\x4e\xdb\xc8\x24\x0b\xa9\x66\xb3\x88\x6d\x36\x20\x99\x62\xb7\x63\xec\x75\xad\x1f\x0d\xcf\xc8\x97\x9a\x05\xaf\x66\x6d\x60\x18\x5f\xdf\x0d\xce\xcf\xee\xa6\x18\xf7\x06\xf7\xf7\x67\xfd\x8b\x29\xc6\xc3\xc7\xf3\xc3\xe4\xac\x37\xba\x1d\xf4\xd1\xe9\x74\xa6\x18\x0f\x1e\xfc\x64\x38\x65\xec\xba\x50\x33\x5e\x60\x50\xf9\x5c\x4d\xcd\xf9\x95\xf4\x4c\x19\x61\xd7\xdd\xe6\x5e\xdb\xed\x82\x96\x54\xfc\x76\xba\xbf\xc6\x21\x59\x13\xee\x42\x39\x5b\x39\x8b\xe5\x33\x04\x61\x27\xa4\x2b\x49\x8b\x84\x17\xc5\x1a\xe3\x94\xe6\xdc\x15\x76\xba\x67\x33\xa2\x20\x69\x4f\x80\xb6\x79\x66\x43\xa1\xb2\x4c\xc8\x2c\x6a\x08\x6b\x9a\x71\xf8\xf7\xf1\x80\x7c\x72\x82\x3c\x10\xed\xa7\x1f\x41\x7e\x3a\x20\x7d\x49\x4e\x5e\x08\xe1\x3b\xc8\xd3\xf6\xff\xb2\x4e\x69\xe6\xb2\x10\x7b\xf9\x23\x0c\x9f\x0f\x48\xab\x79\x42\x35\x32\x40\x07\xf5\x6e\xab\x30\x1c\x5d\x0c\x1e\x47\xe0\x32\xf5\xb3\x42\x65\x98\x8b\x82\x1a\x86\x5f\xa6\xdf\x11\xc3\xe5\xbf\xbc\xac\x8a\x5a\x13\x3d\x55\x96\x5c\xa6\x4d\x09\x67\xb5\x8a\xfd\x68\x04\xbd\x90\xcf\x42\x57\x33\x2f\x74\xc3\x58\x03\x6f\x20\xef\xf7\x5a\x6a\xc0\x29\x37\xf9\x4c\x71\x9d\x1a\xc6\xae\x94\xae\x95\x9d\x53\x51\xbd\xdd\xef\xd7\x1a\x50\xf8\xf2\x8d\xac\x1f\xb4\x2a\x2b\xdb\xa7\x55\x4f\x49\x39\x17\x99\xd3\xe1\x81\x84\xec\xdb\x3f\x77\xb0\xbe\x42\x4f\xbd\x08\x82\x2b\x7f\xad\x57\xca\xc9\xb4\x8b\xe3\xcd\xa6\xe3\xe7\x92\x97\xb4\xdb\x1d\xff\xf4\xe0\xe3\x5b\xff\x36\x84\xc1\x5a\x39\x8d\xb9\xd0\xc6\x42\x3b\x79\x82\x8a\x74\xce\x2b\xf3\x07\xfe\xba\xbc\xeb\x0d\xee\x2f\x8f\x8e\xd0\x6d\x4d\x19\x1b\x29\x24\x4a\x4a\xdf\x7e\xd6\xca\xa1\x74\xc6\xa2\xd2\x6a\x29\x52\xaa\x39\x3e\xb6\x5e\x37\x1d\x21\x8d\xe5\x32\x21\x3c\x7e\xb9\xc3\xa7\x16\x9c\x21\xed\x4f\x13\x84\xc4\x4e\x5b\xa8\xb8\x31\x2b\xa5\xd3\x0e\x46\x39\x19\xc2\x92\x17\x8e\x0c\xb8\x26\x90\x4c\xf4\xba\xb2\x94\x62\x25\x6c\x8e\xcf\xad\xf0\x8e\xc3\x9a\xc2\x82\xd6\x58\xe5\x24\x61\xac\xd2\xbe\xb1\x09\xc9\xfc\x67\x1f\xc0\xe0\xfd\x8d\x2a\xe9\x42\xe8\x0e\x63\xfd\xc1\xe8\xb2\x8b\xbf\x95\xc3\x4a\x14\x05\x72\xbe\xa4\xba\x1f\x84\x46\xe2\xc5\x2c\x95\x85\xf1\xcb\x35\xb5\x67\xf6\xd9\x3d\x9f\x4b\x58\xa4\x2e\x84\x68\x9e\x49\xa5\x55\xa6\x79\x09\x21\x97\x2a\xa9\x8b\x16\x12\x3c\x6e\xb7\x17\xb4\x3e\x06\x37\xe0\xa8\xb8\xe6\x25\x59\xf2\x29\x3c\x84\x9f\x8e\x3d\xa1\x8f\xae\xe9\xc9\x09\x4d\x69\x73\xde\x2e\x3b\x48\xf1\xbf\x00\x00\x00\xff\xff\xe9\x66\x0e\x15\x38\x07\x00\x00"),
		},
		"/vfs_mock.go": &vfsgen۰CompressedFileInfo{
			name:             "vfs_mock.go",
			modTime:          time.Date(2020, 4, 21, 3, 17, 35, 308041023, time.UTC),
			uncompressedSize: 685,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\xcf\x6e\x9c\x30\x1c\x84\xcf\xeb\xa7\x98\x72\xc9\xa2\x12\xb6\xe7\x48\x39\xa5\x8d\xaa\x5e\x9b\x9e\x57\x06\x0f\x60\xc5\x7f\x90\xfd\x63\x57\xab\xaa\xef\x5e\x19\x96\xaa\xb9\xd8\xc6\x1a\x7f\x8c\x3f\x9f\x4e\xf8\xdc\x2d\xd6\x19\x7c\x4a\x74\xd4\x99\x4a\x9d\x4e\x78\x9b\x6c\xc6\xac\xfb\x77\x3d\x12\x36\x23\x06\x77\xc3\x92\x69\x60\x96\x64\xc3\x08\xc3\x0b\x5d\x9c\x3d\x83\x40\x22\x5c\xd4\x06\x32\x11\x89\x39\x2e\xa9\x67\xc6\x90\xa2\x87\xb1\xf9\xbd\x2d\xc0\x5f\xb9\x9c\x1a\x23\x46\x06\x26\x2d\xc4\xd5\x3a\x87\x3e\xb1\xac\x35\x1e\x84\x7e\x76\x5a\x78\xde\x03\xed\x18\x1f\xb6\x50\xa6\xac\xec\xb7\x7b\xe4\x35\x3a\xc3\x04\x63\x87\x81\x89\x41\xdc\x4d\xa9\xbd\x6b\xef\x8d\x52\xd6\xcf\x31\x09\x8e\xea\x50\x05\xca\x69\x12\x99\x2b\x75\xa8\x66\x2d\x53\x99\xd3\x12\xc4\x7a\x56\xaa\x5e\xef\xfa\xe2\xcd\x77\xba\xf9\x9b\xef\x68\x60\xfd\xec\x58\x6e\x95\xd7\x5f\x96\xa3\x18\xac\x63\xbe\x65\xa1\x6f\xd0\x2d\xb2\xfa\xb8\x30\x25\x6b\x0c\x03\xae\x53\x19\x58\x48\x9b\xc8\xab\x95\x09\xa2\xc7\x8c\xe3\x18\xef\x7b\x8f\xeb\xf7\xdd\x70\x0d\x29\x7a\x0b\x16\xd7\x18\x1e\x04\x1d\xd7\x9c\xac\xfc\x42\xda\x6d\xe4\xff\x75\x6c\x36\x3a\xb6\xea\xa2\xd3\xc7\xda\xa5\x67\xfb\x6a\x1d\x7f\xae\x3d\x95\x1a\x96\xd0\xc3\x06\x2b\xc7\x1a\xbf\xd5\x61\x7f\xd2\x40\x9a\x5c\x1e\xac\xe3\xa6\x35\x42\x07\xe8\x2e\x47\xb7\x08\x31\x6c\x66\x8b\xa8\x06\x39\xe2\x4a\x18\x26\x7b\x21\xac\xb4\x78\x7a\xac\xd5\xe1\xdc\xac\xc5\x83\xf6\x6c\x70\x6e\x70\xc6\xd3\x33\xee\x46\xdb\x17\xed\x1c\xd3\xf1\x4b\xad\xd4\xe1\x43\xbf\xe7\xad\xe1\x57\x9b\x8e\x05\xde\xfe\x88\x36\x6c\xab\xb2\xb5\x03\xeb\x06\x55\xdb\x56\xff\x46\x3d\xcf\x65\xea\xbd\xa9\xea\x5a\xfd\x51\x7f\x03\x00\x00\xff\xff\x6a\xb4\xa6\x09\xad\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/backup"].(os.FileInfo),
		fs["/restore"].(os.FileInfo),
		fs["/scknobb.tmpl"].(os.FileInfo),
		fs["/vfs_mock.go"].(os.FileInfo),
	}
	fs["/backup"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/backup/backup.go"].(os.FileInfo),
		fs["/backup/backup.tmpl"].(os.FileInfo),
		fs["/backup/dashboard.go"].(os.FileInfo),
		fs["/backup/report.go"].(os.FileInfo),
		fs["/backup/searchhistory.go"].(os.FileInfo),
	}
	fs["/restore"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/restore/restore.go"].(os.FileInfo),
		fs["/restore/restore.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
