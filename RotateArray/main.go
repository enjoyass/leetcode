package main

func main(){
	rotate([]int{1,2,3,4,5,6,7},3)
}
func rotate(nums []int, k int)  {
	k=k%len(nums)
    if k!=0{
        copy(nums,append(nums[len(nums)-k:],nums[:len(nums)-k]...))
    }
}