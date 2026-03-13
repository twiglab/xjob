package cfas

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/imroc/req/v3"
)

func HmacSha256(data string, key string) string {
	// 1. 创建 HMAC 实例，指定哈希算法和密钥
	h := hmac.New(sha256.New, []byte(key))
	// 2. 写入数据
	h.Write([]byte(data))
	// 3. 计算哈希值
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type Config struct {
	AppKey    string
	AppSecret string
	BaseURL   string
}

func aksk(conf Config) req.RequestMiddleware {
	return func(client *req.Client, req *req.Request) error {
		s := sign(req.Method, conf.AppKey, req.RawURL, conf.AppSecret)
		req.SetHeader("X-Ca-Key", conf.AppKey)
		req.SetHeader("X-Ca-Signature", s)
		req.SetHeader(ContentType, JsonContentType)
		req.SetHeader("X-Ca-Signature-Headers", "x-ca-key")
		return nil
	}
}

func sign(meth, key, url string, secret string) string {
	var sb strings.Builder
	sb.WriteString(meth)
	sb.WriteByte('\n')
	sb.WriteString(JsonContentType)
	sb.WriteByte('\n')
	sb.WriteString("x-ca-key:")
	sb.WriteString(key)
	sb.WriteByte('\n')
	sb.WriteString(url)
	return HmacSha256(sb.String(), secret)
}
