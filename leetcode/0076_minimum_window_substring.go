func minWindow(s string, t string) string {
	if len(s) < len(t) { return "" }
	m := make(map[byte]int)
	for _, v := range []byte(t) {
	  m[v]++
	}
	start, end, index := 0, 0, 0
	counter := len(t)
	minLen := math.MaxInt32
	
	for end < len(s) {
	  // 拡大フェーズ
	  if m[s[end]] > 0 {  // 初めて見つけた場合(マイナスの場合見つけている？)
		counter--
	  }
	  m[s[end]]-- // 見つけた文字をメモ
	  end++
	  
	  // 全部見つかった場合
	  // ここから縮小フェーズ
	  for counter == 0 {
		// ウィンドウサイズが最小の場合更新
		if minLen > end - start {
		  minLen = end - start
		  index = start // 最小のウィンドウの一番左の位置を保存
		}
		m[s[start]]++ // ウィンドウから外れるので、見つけた数を減らす
		if m[s[start]] == 1 { // 一個しかない場合、ウィンドウの中の文字が消えるのでカウントも増やす
		  counter++
		}
		start++
	  }
	}
	if minLen == math.MaxInt32 {
	  return ""
	}
	return s[index:index+minLen]
  }