package stringx

import (
	"fmt"
	"sort"
	"strings"
)

// JoinStringsInASCII 按照规则，参数名ASCII码从小到大排序后拼接
// data 待拼接的数据
// sep 连接符
// onlyValues 是否只包含参数值，true则不包含参数名，否则参数名和参数值均有
// includeEmpty 是否包含空值，true则包含空值，否则不包含，注意此参数不影响参数名的存在
// exceptKeys 被排除的参数名，不参与排序及拼接
func JoinStringsInASCII(data map[string]string, sep string, onlyValues, includeEmpty bool, exceptKeys ...string) string {
	var list []string
	var keyList []string
	m := make(map[string]struct{})
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = struct{}{}
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if !includeEmpty && value == "" {
			continue
		}
		if onlyValues {
			keyList = append(keyList, k)
		} else {
			list = append(list, fmt.Sprintf("%s=%s", k, value))
		}
	}
	if onlyValues {
		sort.Strings(keyList)
		for _, v := range keyList {
			list = append(list, data[v])
		}
	} else {
		sort.Strings(list)
	}
	return strings.Join(list, sep)
}
