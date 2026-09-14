package api

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"

	bcehttp "github.com/baidubce/bce-sdk-go/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// aes128EncryptWithFirst16Char encrypts content with AES-128/ECB/PKCS5Padding, using only the
// first 16 bytes of secretKey as the key, and returns the result hex-encoded. This mirrors the
// Java SDK's BesClientV3#aes128WithFirst16Char and is the protocol BES expects for password
// fields: the server looks up the same secretKey by the caller's AK (sent via the
// X-Bce-Accesskey header) and decrypts with the identical algorithm.
func aes128EncryptWithFirst16Char(content, secretKey string) (string, error) {
	if len(secretKey) < 16 {
		return "", errors.New("secretKey length should be at least 16")
	}
	key := []byte(secretKey[:16])

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plain := pkcs5Pad([]byte(content), blockSize)
	encrypted := make([]byte, len(plain))
	for start := 0; start < len(plain); start += blockSize {
		block.Encrypt(encrypted[start:start+blockSize], plain[start:start+blockSize])
	}
	return hex.EncodeToString(encrypted), nil
}

func pkcs5Pad(data []byte, blockSize int) []byte {
	padSize := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, padding...)
}

// createMultipartRequest sends a multipart/form-data request. RequestBuilder cannot be reused
// here because it always JSON-marshals the body; this path sets the body explicitly instead.
func createMultipartRequest(cli bce.Client, region, method, uri string, body *bce.Body,
	contentType string, resp interface{}) error {
	req := &bce.BceRequest{}
	req.SetMethod(method)
	req.SetUri(uri)
	req.SetHeader(bcehttp.CONTENT_TYPE, contentType)
	req.SetHeader(HEADER_REGION, region)
	req.SetBody(body)

	bceResp := &bce.BceResponse{}
	if err := cli.SendRequest(req, bceResp); err != nil {
		return err
	}
	if bceResp.IsFail() {
		return bceResp.ServiceError()
	}
	defer bceResp.Body().Close()

	if resp == nil {
		return nil
	}
	return bceResp.ParseJsonBody(resp)
}
