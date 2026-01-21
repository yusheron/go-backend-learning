## 两数之和

### 思路
遍历数组，用 map 记录已经出现过的数字和下标。
每次判断 target - 当前值是否已经出现。

### Go 实现
```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        if idx, ok := m[target-v]; ok {
            return []int{idx, i}
        }
        m[v] = i
    }
    return nil
}
```
