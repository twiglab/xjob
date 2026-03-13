package cfas

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func ContentAny(a any) ([]byte, string, error) {
	body, err := json.Marshal(a)
	if err != nil {
		return nil, "", err
	}
	h := md5.New()
	h.Write(body)
	md5byte := h.Sum(nil)
	b64str := base64.StdEncoding.EncodeToString(md5byte)
	return body, b64str, nil
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

type Config struct {
	AppKey    string
	AppSecret string

	Host string
	Port int

	BaseURL string
}

func aksk(conf Config) req.RequestMiddleware {
	return func(client *req.Client, req *req.Request) error {
		fmt.Println(req.Body)
		fmt.Println(req.RawURL)
		fmt.Println(client.BaseURL)
		req.SetHeader("X-Ca-Key", conf.AppKey)
		return nil
	}
}
