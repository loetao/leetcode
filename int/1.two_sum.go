/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

func twoSum(nums []int, target int) []int {
    var r []int
    for i:=0;i<len(nums);i++    {
        for j:=0;j<len(nums);j++        {
            if i!=j && nums[i]+nums[j]==target{
                fmt.Printf("[%d,%d]",i,j)
                r=append(r,i)
                r=append(r,j)
                return r
            }
            
        }
    }
    return r
    
}
