package city

import (
	"strconv"
	"strings"
)

// 按照城市灰度发布
// 城市ID--是否发布
var cityIDOpen = [1200]bool{}

func init() {
	// 读取配置
	configs := "{....}"
	for i := 0; i < len(cityIDOpen); i++ {
		// 根据配置信息设置每个城市ID是否满足灰度
		if strings.Contains(strconv.Itoa(i), configs) {
			cityIDOpen[i] = true
		}
	}
}

func isPassed(cityID int) bool {
	return cityIDOpen[cityID]
}

// 如果cityID赋值比较大，可以考虑使用map来存储映射关系

var cityIDOpenMap = map[int]struct{}{}

func init() {
	// 读取配置文件
	configs := map[int]struct{}{}
	for city := range configs {
		cityIDOpenMap[city] = struct{}{}
	}
}

func isPassedOfMap(cityID int) bool {
	if _, ok := cityIDOpenMap[cityID]; ok {
		return true
	}
	return false
}
