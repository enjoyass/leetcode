package main

import (
	"fmt"
)

func main(){
	a:="10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b:="110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	if addBinary(a,b)== "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000" {
		fmt.Println("SUCCESS")
	}else{
		fmt.Println("FAILED")
	}
}
func addBinary(a string, b string) string {
	if len(a) <len(b) {
		a,b=b,a
	}
	var result []byte
	var flag bool = false
	for i:=0;i<len(b);i++  {
        tmepvalA:=len(a)-i-1
        tmepvalB:=len(b)-i-1
		if a[tmepvalA] =='1' && b[tmepvalB] =='1' {
			if flag {
				result=append(result,'1')
			}else{
				flag =true
				result=append(result,'0')
			}
		}else if (a[tmepvalA] =='0' && b[tmepvalB] =='1')||(a[tmepvalA] =='1' && b[tmepvalB] =='0') {
			if flag {
				flag =true
				result=append(result,'0')
			}else{
				result=append(result,'1')
			}
		}else{
			if flag {
                flag = false
                result = append(result, '1')
            } else {
                result = append(result, '0')
            }
		}
	}

	gap :=len(a)-len(b)

	if gap ==0 {
		if flag {
		  result = append(result, '1')
		}
	}else {
		for i:=gap-1;i>=0;i--{
			if a[i]=='1' && flag {
				result=append(result,'0')
			}else if flag {
				flag=false
				result=append(result,'1')
			}else{
				result = append(result, a[i])
			}
		}
		if flag {
            result = append(result, '1')
            flag = false
        }
	}
	for i:=0; i < len(result)/2 ;i++{
        li := len(result) - i -1
        result[i],result[li] = result[li],result[i]
	}
	return string(result)
}