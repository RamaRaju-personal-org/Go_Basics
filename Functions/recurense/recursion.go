package recurense


func Recure(num int)int {
	if num == 0 {
		return 1
	}
	return num * Recure(num-1) 
}
