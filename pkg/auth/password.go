package auth
import (
"crypto/rand"
"encoding/base64"
"golang.org/x/crypto/argon2"
)
type PasswordConfig struct {
Time uint32
Memory uint32
Threads uint8
KeyLen uint32
}
// 密码加密
func HashPassword(password string) (string, error) {
// 生成随机盐值
salt := make([]byte, 16)
if v, err := rand.Read(salt); err != nil {
return "", err
}
// 使用Argon2id进行密码哈希
config := &PasswordConfig{
Time: 1,
Memory: 64*1024,
Threads: 4,
KeyLen: 32,
}
hash := argon2.IDKey(
[]byte(password),
salt,
config.Time,
config.Memory,
config.Threads,
config.KeyLen,
)
// 编码存储格式: $argon2id$v=19$m=65536,t=1,p=4$salt$hash
encoded := fmt.Sprintf(
"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
config.Memory,
config.Time,
config.Threads,
base64.RawStdEncoding.EncodeToString(salt),
base64.RawStdEncoding.EncodeToString(hash),
)
return encoded, nil
}
// 密码验证
func VerifyPassword(password, encodedHash string) (bool, error) {
// 解析存储的哈希值
parts := strings.Split(encodedHash, "$")
if len(parts) != 6 {
return false, fmt.Errorf("invalid hash format")
}
// 解析参数
var config PasswordConfig
, err := fmt.Sscanf(
parts[3],
"m=%d,t=%d,p=%d",
&config.Memory,
&config.Time,
&config.Threads,
)
if err != nil {
return false, err
}
salt, err := base64.RawStdEncoding.DecodeString(parts[4])
if err != nil {
return false, err
}
decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
if err != nil {
return false, err
}
config.KeyLen = uint32(len(decodedHash))
// 使用相同参数重新计算哈希值
comparisonHash := argon2.IDKey(
[]byte(password),
salt,
config.Time,
config.Memory,
config.Threads,
config.KeyLen,
)
return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}