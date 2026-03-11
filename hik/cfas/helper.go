package cfas

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json/v2"
	"time"

	"github.com/google/uuid"
	"github.com/imroc/req/v3"
)

func ContentMD5(v any) ([]byte, string, error) {
	var bf bytes.Buffer
	if err := json.MarshalWrite(&bf, v); err != nil {
		return nil, "", err
	}

	h := md5.New()
	md5byte := h.Sum(bf.Bytes())

	b64str := base64.StdEncoding.EncodeToString(md5byte)

	return bf.Bytes(), b64str, nil
}

func HttpDate(d time.Time) string {
	return d.Format(time.RFC1123)
}

func Nonce() string {
	return uuid.NewString()
}

func Timestamp(d time.Time) int64 {
	return d.UnixMilli()
}

func M(client *req.Client, req *req.Request) error {
	return nil
}
