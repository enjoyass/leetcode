package main

func main(){

}
func countSegments(s string) int {
	res := 0
	n := len(s)
	for i:= 0; i < n; i++ {
		if s[i] == ' '{
			continue
		} 
		 res++
		for i < n && s[i] != ' '{
		 i++
		} 
	}
    return res
}