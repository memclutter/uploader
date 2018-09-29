# Uploader

[![Language Golang](https://img.shields.io/badge/language-golang-blue.svg)](https://img.shields.io/badge/language-golang-blue.svg)
[![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/memclutter/uploader)
[![Build Status](https://travis-ci.com/memclutter/uploader.svg?branch=master)](https://travis-ci.com/memclutter/uploader)

File uploader for golang projects.

## Usage

Use the structure of the `Uploader` to describe the download of files.

```go
upl := uploader.Uploader{
    BaseDir:     "/var/www/upload",
    BaseUrl:     "",
    Dir:         "avatars",
    Path:        "avatars",
    MimeTypes:   uploader.MimeTypes{"image/png", "image/jpeg"},
    MaxFileSize: MaxFileSize,
}
```

After that, you get a instance of `multipart.FormFile`  and pass it to the `Uploader.Upload` method

```go
_, multipartFile, err := r.FormFile("avatars")
// check err ...
result := upl.Upload(multipartFile)
// process result
```

The result is a struct 

```go
type Result struct {
	Code     int
	Path     string
	Name     string
	MimeType string
	Size     int64
}
```

## List of Codes

You can process the code returned as a result

- `CodeOk` upload successful  
- `CodeErrDetectMimeType` error detect mime type
- `CodeErrInvalidMimeType` invalid mime type
- `CodeErrExceedMaxFileSize` exceed max file size
- `CodeErrGenerateFileName` error generate file name
- `CodeErrCreateDir` error create directory
- `CodeErrCreateFile` error create new file
- `CodeErrOpenFile` error open uploaded file
- `CodeErrCopyFile` error copy src to dst file

## Example

TODO: