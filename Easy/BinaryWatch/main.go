package main

import (
	"fmt"
)

func main(){
	fmt.Println(readBinaryWatch(2))
}
func readBinaryWatch(num int) []string {
	var result []string
	if num==0 {
		result=append(result,"0:00")
		return result
	}
	if num <1 || num >9 {
		return result
	}
	hour:=map[int][]string{
		0:[]string{"0:"},
		1:[]string{"1:","2:","4:","8:"},
		2:[]string{"3:","5:","6:","9:","10:"},
		3:[]string{"7:","11:"},
	}
	min:=map[int][]string{
		0:[]string{"00"},
		1:[]string{"01","02","04","08","16","32"},
		2:[]string{ "03","05","06","09","10","12","17","18","20","24","33","34","36","40","48"},
		3:[]string{"07","11","13","14","19","21","22","25","26","28","35","37","38","41","42","44","49","50","52","56"},
		4:[]string{"58","57","54","53","51","46","45","43","39","30","29","27","23","15"},
		5:[]string{ "31","47","55","59" },
	}
	for i:=0;i<=num&& i<4;i++{
		j:=num-i
		fmt.Println(i,j)
		for _,vH:=range hour[i] {
			for _,vM:=range min[j] {
				result=append(result,vH+vM)
			}
		}
	}
	return result
}