package utils

import(
	"math/rand"
    "time"
)

func randomStringFromCharset(length int, charset string) string {
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}


func GetRandomString() string {
    rand.Seed(time.Now().UnixNano()) // 为随机数生成器设置种子
    charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
    return randomStringFromCharset(15, charset)
}

