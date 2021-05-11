func groupAnagrams(strs []string) [][]string {
    
    memo := make(map[string][]string, 0)
    
    for _, s := range strs {
        key := []byte(s)
        sort.Slice(key, func(i int, j int) bool { return key[i] < key[j] })
        memo[string(key)] = append(memo[string(key)], s)
    }
    
    result := make([][]string, 0)
    
    for _, v := range memo {
        result = append(result, v)
    }
    
    return result
}
