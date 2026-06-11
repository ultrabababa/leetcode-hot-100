package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 这里的思路通常是：
	// 遍历每个字符串，将其“排序”后作为 map 的 key。
	// 比如 "eat", "tea", "ate" 排序后都是 "aet"。
	// map[string][]string -> key是"aet", value是["eat", "tea", "ate"]

	keyToGroup := make(map[string][]string)
	for _, str := range strs {
		key := sortStr(str)
		keyToGroup[key] = append(keyToGroup[key], str)
	}
	result := make([][]string, 0, len(keyToGroup))
	for _, strings := range keyToGroup {
		result = append(result, strings)
	}
	return result
}

func sortStr(str string) string {
	strBytes := []byte(str)
	sort.Slice(strBytes, func(i, j int) bool {
		return strBytes[i] < strBytes[j]
	})
	return string(strBytes)
}
