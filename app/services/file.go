package services

import (
	"bytes"
	"io"
	"mime/multipart"
)

func CreateBufferFromFormFile(file multipart.File, header multipart.FileHeader) (multiPartBuffer bytes.Buffer, contentType string, err error) {
	writer := multipart.NewWriter(&multiPartBuffer)

	part, err := writer.CreateFormFile("file", header.Filename)

	if err != nil {
		return multiPartBuffer, contentType, err
	}
	_, err = io.Copy(part, file)

	if err != nil {
		return multiPartBuffer, contentType, err
	}

	err = writer.Close()

	if err != nil {
		return multiPartBuffer, contentType, err
	}

	contentType = writer.FormDataContentType()

	return multiPartBuffer, contentType, err
}
