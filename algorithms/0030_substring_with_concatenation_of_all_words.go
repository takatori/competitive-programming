func findSubstring(s string, words []string) []int {
    
    var result []int
    
    if len(words) == 0 || len(s) < len(words) * len(words[0]) {
        return result
    }

    N := len(s)        // stringLen
    wordCount := len(words)    // wordCount
    wordLen := len(words[0]) // wordLen

    wordsMap := make(map[string]int)
    currentMap := make(map[string]int)
    
    for _, w := range words {
        if v, ok := wordsMap[w]; ok {
            wordsMap[w] = v+1
        } else {
            wordsMap[w] = 1
        }
    }
    
    var str, tmp string
    
    for i := 0; i < wordLen; i++ {
        count := 0
        left, right := i, i
        for ;right + wordLen <= N; right += wordLen {
            str = s[right:right+wordLen]
            if _, ok := wordsMap[str]; ok {
                
                if v, ok := currentMap[str]; ok {
                    currentMap[str] = v+1
                } else {
                    currentMap[str] = 1
                }
                
                // あるwordの数が本来あるべき数よりまだ少ないか同じであれば、wordを追加できる
                if currentMap[str] <= wordsMap[str] {
                    count++
                }
                // 
                for currentMap[str] > wordsMap[str] {
                    tmp = s[left:left+wordLen]
                    currentMap[tmp] = currentMap[tmp] - 1
                    left += wordLen
                    if currentMap[tmp] < wordsMap[tmp] {
                        count--
                    }
                }
                
                
                if count == wordCount {
                    result = append(result, left)
                    tmp = s[left:left+wordLen]
                    currentMap[tmp] = currentMap[tmp]-1
                    left += wordLen
                    count--
                }
            } else {
                currentMap = make(map[string]int)
                count = 0
                left = right + wordLen
            }
        }
        currentMap = make(map[string]int)
    }
    return result
}
