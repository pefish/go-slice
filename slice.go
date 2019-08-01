package go_slice

type SliceClass struct {

}

var Slice = SliceClass{}

func (this *SliceClass) IncludesBySliceInterface(strs []interface{}, str string) bool {
	for _, s := range strs {
		if s.(string) == str {
			return true
		}
	}
	return false
}

func (this *SliceClass) IncludesBySliceString(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func (this *SliceClass) GetLastOfSliceString(strs []string) string {
	return strs[len(strs) - 1]
}

func (this *SliceClass) RepeatSliceString(str string, num int) []string {
	result := []string{}
	for i := 0; i < num; i++ {
		result = append(result, str)
	}
	return result
}

func (this *SliceClass) RemoveOneFromStringSlice(slice_ []string, target string) []string {
	result := []string{}
	for _, v := range slice_ {
		if v == target {
			continue
		} else {
			result = append(result, v)
		}
	}
	return result
}
