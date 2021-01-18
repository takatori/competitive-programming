type CacheItem struct {
	value int
	count int
  }
  
  type LRUCache struct {
	capacity int
	cache map[int]*CacheItem
	usage map[int]map[int]*CacheItem
  }
  
  
  func Constructor(capacity int) LRUCache {
	  
  }
  
  
  func (this *LRUCache) Get(key int) int {
	
	if item, ok := this.cache[key]; !ok {
	  return -1
	} else {
	  v := item.value
	  c := item.count
	  
	  // 回数のクラスに存在するか
	  if i, ok := usage[c]; ok {
		// 既に存在していたら消す
		if _, ok := i[v]; ok {
		  delete(i, v)
		}
	  }
	  
	  item.count++
	  c++
	  
	  // 追加
	  if i, ok := usage[c]; !ok {
		m := map[int]*CacheItem{}
		m[v] = item
		usage[c] = m
	  } else {
		
	  }
  
  
	  return item.Value
	}
  }
  
  
  func (this *LRUCache) Put(key int, value int)  {
	  
  }
  
  /**
   * Your LRUCache object will be instantiated and called as such:
   * obj := Constructor(capacity);
   * param_1 := obj.Get(key);
   * obj.Put(key,value);
   */