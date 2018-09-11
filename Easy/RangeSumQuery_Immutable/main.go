package main

type NumArray struct {
    Arr []int
}


func Constructor(nums []int) NumArray {
    return NumArray{Arr:nums}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum :=0
    for k,v :=range this.Arr {
		if k>=i && k<=j {
			sum+=v
		}
	}
	return sum
}
func main(){

}