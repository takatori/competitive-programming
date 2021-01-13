func minWindow(s string, t string) string {
 
	var top, last int
	
	flags := make(map[byte]bool)
	
	count := 0
	min := len(s)+1
	j := 0
	
	for i := 0; i < len(t); i++ {
	  flags[t[i]] = false
	}
	
	for i := 0; i < len(s); i++ {
	  
	  // 追いつくまで
	  for j < i {
	  
	  // まだ見つけていない文字を見つけた
	  if val, ok := flags[s[i]]; ok && !val {
		
		flags[s[i]] = true // 見つけた文字をマーク
		count++ // 見つけた数増やす
		
		
		// 全部見つけた
		if count == len(t) {
		  // 一番小さい場合更新
		  if i - j < min {
			min = i - j
		  }
		  
		  flags[s[top]] = false
		  count--
		}
	  }
	  
	}
	
  }