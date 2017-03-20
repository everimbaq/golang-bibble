package utils

func CompareSlice(a,b []int) bool {
	if a == nil && b== nil{
		return true
	}else if a== nil || b== nil{
		return false
	}else if len(a) != len(b){
		return false
	}

	for i:= range a {
		if a[i] != b[i]{
			return false
		}
	}
	return true
}
