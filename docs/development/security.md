# 安全规范文档

## 1. 身份认证

### 1.1 密码安全 

```
// pkg/auth/password.go
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
if , err := rand.Read(salt); err != nil {
return "", err
}
// 使用Argon2id进行密码哈希
config := &PasswordConfig{
Time: 1,
Memory: 64 1024,
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
```


### 1.2 JWT令牌

```
// pkg/auth/jwt.go
package auth
import (
"github.com/golang-jwt/jwt/v4"
"time"
)
type Claims struct {
UserId int64 json:"user_id"
Username string json:"username"
jwt.RegisteredClaims
}
// 生成访问令牌
func GenerateToken(userId int64, username string, secret string, expire time.Duration) (string, error) {
claims := Claims{
UserId: userId,
Username: username,
RegisteredClaims: jwt.RegisteredClaims{
ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
IssuedAt: jwt.NewNumericDate(time.Now()),
NotBefore: jwt.NewNumericDate(time.Now()),
},
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString([]byte(secret))
}
// 验证令牌
func ParseToken(tokenString string, secret string) (Claims, error) {
token, err := jwt.ParseWithClaims(
tokenString,
&Claims{},
func(token jwt.Token) (interface{}, error) {
return []byte(secret), nil
},
)
if err != nil {
return nil, err
}
if claims, ok := token.Claims.(Claims); ok && token.Valid {
return claims, nil
}
return nil, fmt.Errorf("invalid token")
}
```


### 1.3 会话管理

```
// pkg/session/session.go
type Session struct {
ID string
UserId int64
LoginTime time.Time
ExpireAt time.Time
Device string
IP string
}
// 创建会话
func CreateSession(ctx context.Context, userId int64, device, ip string) (Session, error) {
session := &Session{
ID: uuid.New().String(),
UserId: userId,
LoginTime: time.Now(),
ExpireAt: time.Now().Add(24 time.Hour),
Device: device,
IP: ip,
}
// 存储会话信息
key := fmt.Sprintf("session:%s", session.ID)
if err := rdb.HMSet(ctx, key, map[string]interface{}{
"user_id": session.UserId,
"login_time": session.LoginTime.Unix(),
"expire_at": session.ExpireAt.Unix(),
"device": session.Device,
"ip": session.IP,
}).Err(); err != nil {
return nil, err
}
// 设置过期时间
rdb.Expire(ctx, key, 24time.Hour)
return session, nil
}
```


## 2. 数据安全

### 2.1 数据加密


```
// pkg/crypto/aes.go
package crypto
import (
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"encoding/base64"
"io"
)
// AES加密
func Encrypt(plaintext []byte, key []byte) (string, error) {
block, err := aes.NewCipher(key)
if err != nil {
return "", err
}
// 创建GCM模式
gcm, err := cipher.NewGCM(block)
if err != nil {
return "", err
}
// 创建随机数
nonce := make([]byte, gcm.NonceSize())
if , err := io.ReadFull(rand.Reader, nonce); err != nil {
return "", err
}
// 加密
ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
// Base64编码
return base64.StdEncoding.EncodeToString(ciphertext), nil
}
// AES解密
func Decrypt(ciphertext string, key []byte) ([]byte, error) {
data, err := base64.StdEncoding.DecodeString(ciphertext)
if err != nil {
return nil, err
}
block, err := aes.NewCipher(key)
if err != nil {
return nil, err
}
gcm, err := cipher.NewGCM(block)
if err != nil {
return nil, err
}
nonceSize := gcm.NonceSize()
if len(data) < nonceSize {
return nil, fmt.Errorf("ciphertext too short")
}
nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
if err != nil {
return nil, err
}
return plaintext, nil
}
```


### 2.2 敏感数据处理

```

// pkg/sensitive/mask.go
package sensitive
// 手机号脱敏
func MaskPhone(phone string) string {
if len(phone) != 11 {
return phone
}
return phone[:3] + "" + phone[7:]
}
// 邮箱脱敏
func MaskEmail(email string) string {
parts := strings.Split(email, "@")
if len(parts) != 2 {
return email
}
username := parts[0]
if len(username) <= 3 {
return username[:1] + "@" + parts[1]
}
return username[:3] + "@" + parts[1]
}
// 身份证脱敏
func MaskIDCard(idCard string) string {
if len(idCard) != 18 {
return idCard
}
return idCard[:6] + "" + idCard[14:]
}
```


## 3. 接口安全

### 3.1 防重放攻击

```
// pkg/middleware/replay.go
func AntiReplayMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
// 获取请求时间戳和签名
timestamp := r.Header.Get("X-Timestamp")
nonce := r.Header.Get("X-Nonce")
signature := r.Header.Get("X-Signature")
// 验证时间戳是否过期
ts, err := strconv.ParseInt(timestamp, 10, 64)
if err != nil || time.Now().Unix()-ts > 300 {
httpx.Error(w, errorx.NewReplayError("request expired"))
return
}
// 验证nonce是否重复使用
nonceKey := fmt.Sprintf("nonce:%s", nonce)
if exists, := rdb.Exists(r.Context(), nonceKey).Result(); exists == 1 {
httpx.Error(w, errorx.NewReplayError("nonce already used"))
return
}
// 存储nonce并设置过期时间
rdb.Set(r.Context(), nonceKey, "1", 5time.Minute)
// 验证签名
if !verifySignature(r, timestamp, nonce, signature) {
httpx.Error(w, errorx.NewReplayError("invalid signature"))
return
}
next(w, r)
}
}
```


### 3.2 SQL注入防护

```
// pkg/middleware/sql.go
func SQLInjectionMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
// 检查查询参数
for key, values := range r.URL.Query() {
for , value := range values {
if containsSQLInjection(value) {
httpx.Error(w, errorx.NewSecurityError("sql injection detected"))
return
}
}
}
// 检查POST数据
if r.Method == http.MethodPost {
if err := r.ParseForm(); err == nil {
for key, values := range r.PostForm {
for , value := range values {
if containsSQLInjection(value) {
httpx.Error(w, errorx.NewSecurityError("sql injection detected"))
return
}
}
}
}
}
next(w, r)
}
}
func containsSQLInjection(value string) bool {
// SQL注入特征检测
patterns := []string{
"(?i)\\b(and|or)\\b.+?(?i)\\b(union|select|insert|delete|update|drop|truncate)\\b",
"(?i)\\b(union|select|insert|delete|update|drop|truncate)\\b.+?(?i)\\b(and|or)\\b",
"(?i)\\b(union|select|insert|delete|update|drop|truncate)\\b",
"--",
";",
"/",
"/",
"@@",
}
for , pattern := range patterns {
if matched, := regexp.MatchString(pattern, value); matched {
return true
}
}
return false
}

```


### 3.3 XSS防护


```
// pkg/middleware/xss.go
func XSSMiddleware(next http.HandlerFunc) http.HandlerFunc {
return func(w http.ResponseWriter, r http.Request) {
// 设置安全响应头
w.Header().Set("X-XSS-Protection", "1; mode=block")
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-Frame-Options", "DENY")
// 过滤请求参数
for key, values := range r.URL.Query() {
for i, value := range values {
r.URL.Query()[key][i] = html.EscapeString(value)
}
}
if r.Method == http.MethodPost {
if err := r.ParseForm(); err == nil {
for key, values := range r.PostForm {
for i, value := range values {
r.PostForm[key][i] = html.EscapeString(value)
}
}
}
}
next(w, r)
}
}
```

## 4. 安全审计

### 4.1 操作日志

```
// model/ent/schema/operation_log.go
type OperationLog struct {
ent.Schema
}
func (OperationLog) Fields() []ent.Field {
return []ent.Field{
field.Int64("id"),
field.Int64("user_id"),
field.String("username"),
field.String("operation"),
field.String("method"),
field.String("path"),
field.String("params"),
field.String("ip"),
field.String("user_agent"),
field.Time("created_at"),
}
}
// 记录操作日志
func LogOperation(ctx context.Context, client ent.Client, log OperationLog) error {
return client.OperationLog.Create().
SetUserID(log.UserID).
SetUsername(log.Username).
SetOperation(log.Operation).
SetMethod(log.Method).
SetPath(log.Path).
SetParams(log.Params).
SetIP(log.IP).
SetUserAgent(log.UserAgent).
SetCreatedAt(time.Now()).
Exec(ctx)
}

```


### 4.2 安全审计

```
// pkg/audit/audit.go
type SecurityAudit struct {
client ent.Client
logger zap.Logger
}
func (a SecurityAudit) AuditLogin(ctx context.Context, userId int64, success bool, ip string) error {
return a.client.SecurityAudit.Create().
SetUserID(userId).
SetEventType("login").
SetSuccess(success).
SetIP(ip).
SetCreatedAt(time.Now()).
Exec(ctx)
}
func (a SecurityAudit) AuditPermissionChange(ctx context.Context, userId int64, targetUserId int64, action string) error {
return a.client.SecurityAudit.Create().
SetUserID(userId).
SetTargetUserID(targetUserId).
SetEventType("permission_change").
SetAction(action).
SetCreatedAt(time.Now()).
Exec(ctx)
}
func (a SecurityAudit) AuditSensitiveOperation(ctx context.Context, userId int64, operation string, detail string) error {
return a.client.SecurityAudit.Create().
SetUserID(userId).
SetEventType("sensitive_operation").
SetOperation(operation).
SetDetail(detail).
SetCreatedAt(time.Now()).
Exec(ctx)
}
```

