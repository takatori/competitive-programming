func isAlienSorted(words []string, order string) bool {
  
	if len(words) < 2 {
	  return true
	}
	
	dict := make(map[byte]int, len(order))
	
	for i := 0; i < len(order); i++ {
	  dict[order[i]] = i
	}
	  
	for i := 0; i < len(words)-1; i++ {
	  if !isSort(words[i], words[i+1], dict) {
		return false
	  }
	}
	return true
  }
  
  
  func isSort(a string, b string, dict map[byte]int) bool {
	for i := 0; i < len(a); i++ {
	  if len(b) <= i {
		return false
	  }
	  if dict[a[i]] < dict[b[i]] {
		return true
	  }
	  if dict[a[i]] > dict[b[i]] {
		return false
	  }
	}
	return true
  }