func hasDuplicate(nums []int) bool {
    dup:=make(map[int]bool,0)
    for _,v:=range nums {
        if _,ok:=dup[v];ok{
            return true
        }
            dup[v]=true
    }
    return false
}
