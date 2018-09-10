package main

type MinStack struct {
	dataStatck []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.dataStatck=append(this.dataStatck,x)
}


func (this *MinStack) Pop()  {
    this.dataStatck=this.dataStatck[:len(this.dataStatck)-1]
}


func (this *MinStack) Top() int {
    return this.dataStatck[len(this.dataStatck)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.dataStatck)==0 {
		return 0
	}
	var minData int = this.dataStatck[0]
    for i:=1;i<len(this.dataStatck);i++ {
		if minData > this.dataStatck[i] {
			minData =this.dataStatck[i]
		}
	}
	return minData
}

func main(){
	
}