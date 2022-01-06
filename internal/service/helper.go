package service

import "math"

var ServiceHelper = &serviceHelper{}

type serviceHelper struct {
}

func (s *serviceHelper) RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func (s *serviceHelper) ArraySplit(arr []string, splitCount int) [][]string {
	if len(arr) == 0 || splitCount <= 0 {
		return nil
	}
	arrLen := int(math.Ceil(float64(len(arr)) / float64(splitCount)))
	i := 1
	var ret [][]string
	for i <= arrLen {
		// 计算开始裁剪位置
		left := (i - 1) * splitCount
		right := left + splitCount
		// 如果右边限定值超出总数 则修改到最右侧
		if right > len(arr) {
			right = len(arr)
		}
		ret = append(ret, arr[left:right])
		i++
	}
	return ret
}
