package common

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// GenHashedPwd 密码加密
func GenHashedPwd(pwd, salt string) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s+%s", pwd, salt)))

	return hex.EncodeToString(hash[:])
}

// GenPwdSalt 生成密码盐
func GenPwdSalt() string {
	return GenRandomString(8)
}

// GenPwd 密码加密
func GenPwd(pwd string) string {
	h := md5.New()
	_, _ = io.WriteString(h, fmt.Sprintf("%s+%s", "LDX_SECRET", pwd))

	return fmt.Sprintf("%x", h.Sum(nil))
}
