package api

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	bcehttp "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

// MAX_UPLOAD_FILE_SIZE is the maximum file size accepted by the multipart upload APIs.
const MAX_UPLOAD_FILE_SIZE = 5 * 1024 * 1024 * 1024 // 5GB

// multipartField is a single non-file form field to be encoded into a multipart/form-data body.
type multipartField struct {
	name  string
	value string
}

// openUploadFile opens the file at path and validates it is a regular file within the size limit.
func openUploadFile(path string) (*os.File, os.FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, nil, err
	}
	if fileInfo.IsDir() {
		file.Close()
		return nil, nil, fmt.Errorf("the file %s is a directory, not a regular file", path)
	}
	if fileInfo.Size() > MAX_UPLOAD_FILE_SIZE {
		file.Close()
		return nil, nil, fmt.Errorf("the file %s exceeds the maximum allowed size of %d bytes",
			path, MAX_UPLOAD_FILE_SIZE)
	}
	return file, fileInfo, nil
}

// buildMultipartBody encodes formFields and the given file into a multipart/form-data body.
// The file content is streamed via io.MultiReader and never buffered into memory as a whole.
func buildMultipartBody(formFields []multipartField, fileFieldName string, file io.Reader,
	fileName string, fileSize int64) (*bce.Body, string, error) {
	if file == nil {
		return nil, "", errors.New("upload file should not be nil")
	}

	boundary, err := newMultipartBoundary()
	if err != nil {
		return nil, "", err
	}

	var preamble bytes.Buffer
	for _, field := range formFields {
		writeMultipartFormField(&preamble, boundary, field.name, field.value)
	}
	writeMultipartFilePartHeader(&preamble, boundary, fileFieldName, fileName)

	trailer := []byte("\r\n--" + boundary + "--\r\n")

	size := int64(preamble.Len()) + fileSize + int64(len(trailer))
	reader := io.MultiReader(bytes.NewReader(preamble.Bytes()), file, bytes.NewReader(trailer))

	body, err := bce.NewBodyFromReader(reader, size)
	if err != nil {
		return nil, "", err
	}
	return body, "multipart/form-data; boundary=" + boundary, nil
}

func newMultipartBoundary() (string, error) {
	raw := make([]byte, 16)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}

func writeMultipartFormField(buf *bytes.Buffer, boundary, name, value string) {
	buf.WriteString("--" + boundary + "\r\n")
	buf.WriteString("Content-Disposition: form-data; name=\"" + name + "\"\r\n\r\n")
	buf.WriteString(value)
	buf.WriteString("\r\n")
}

func writeMultipartFilePartHeader(buf *bytes.Buffer, boundary, name, fileName string) {
	buf.WriteString("--" + boundary + "\r\n")
	buf.WriteString("Content-Disposition: form-data; name=\"" + name + "\"; filename=\"" +
		sanitizeFilename(fileName) + "\"\r\n")
	buf.WriteString(bcehttp.CONTENT_TYPE + ": " + guessContentType(fileName) + "\r\n\r\n")
}

// guessContentType infers a file's MIME type from its extension, falling back to the generic
// binary type when the extension is unknown.
func guessContentType(fileName string) string {
	const defaultContentType = "application/octet-stream"
	dot := strings.LastIndex(fileName, ".")
	if dot == -1 {
		return defaultContentType
	}
	ext := fileName[dot:]
	if contentType, ok := util.GetMimeMap()[ext]; ok {
		return contentType
	}
	return defaultContentType
}

// sanitizeFilename escapes characters that could break the Content-Disposition header structure
// or inject additional multipart fields.
func sanitizeFilename(filename string) string {
	filename = strings.ReplaceAll(filename, "\\", "\\\\")
	filename = strings.ReplaceAll(filename, "\"", "\\\"")
	filename = strings.ReplaceAll(filename, "\r", "")
	filename = strings.ReplaceAll(filename, "\n", "")
	return filename
}
