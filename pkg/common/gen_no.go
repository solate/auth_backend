package common

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

func getTomorrowMidnight() int {
	now := time.Now()
	tomorrowMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	secondsToTomorrowMidnight := tomorrowMidnight.Sub(now).Seconds()
	return int(secondsToTomorrowMidnight)
}

// 将十进制数转换为62进制字符串
func decimalToBase62(n int64) string {
	if n == 0 {
		return "0"
	}
	digits := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	for n > 0 {
		remainder := n % 62
		n /= 62
		result = string(digits[remainder]) + result
	}
	return result
}

func getYTo62() string {
	now := time.Now()
	// 获取年月日字符串，格式为 YYMMDD
	yearMonthDay := fmt.Sprintf("%02d%02d%02d", now.Year()%100, int(now.Month()), now.Day())
	// 将年月日字符串转换为十进制数
	decimalValue, err := strconv.ParseInt(yearMonthDay, 10, 64)
	if err != nil {
		return ""
	}
	base62Hash := decimalToBase62(decimalValue)
	return base62Hash
}

// GenNo 按需生成编号
// t : S代表商家，F代表粉丝，D代表代理商与服务商，U代表用户
func GenNo(t string, r *redis.Redis) string {
	// S 商家  F 粉丝 D 服务商、代理商 U 用户
	s := []string{"S", "F", "D", "U"}
	exist := lo.Contains(s, t)
	if !exist {
		return ""
	}
	redisKey := fmt.Sprintf("%s%s%s", "hdz:gen_no:", t, ":number")
	dayNo, err := r.Incr(redisKey)
	if err != nil {
		return ""
	}
	err = r.Expire(redisKey, getTomorrowMidnight())
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s%s%s%05d", "ldx", t, getYTo62(), dayNo)
}
