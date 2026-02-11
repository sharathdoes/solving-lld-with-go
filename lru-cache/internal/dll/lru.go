package dll

type LRU struct {
	Dll *DLL
	Mp map[string]*Node
}

func NewLRU(cap int) *LRU {
	return &LRU{
		Dll: NewDLL(cap),
		Mp:  make(map[string]*Node),
	}
}


func(l *LRU)PushValue(key string, value string){

	l.Dll.mu.Lock()
	defer l.Dll.mu.Unlock()
	
	if node, ok := l.Mp[key]; ok {
		l.Dll.RemoveNode(node)
		node.Value=value
		l.Dll.InsertAtLast(node)
		
	} else {
		if len(l.Mp)==l.Dll.capacity {
			node:=l.Dll.RemoveBeginning()
			delete(l.Mp, node.Key)

		}
		new_node:=NewNode(key,value)
		l.Dll.InsertAtLast(new_node)
		l.Mp[key] = new_node
	}
}

func (l *LRU) GetValue(key string) (string, bool) {
	l.Dll.mu.Lock()
	defer l.Dll.mu.Unlock()
    node, ok := l.Mp[key]
    if !ok {
        return "", false
    }

    l.Dll.RemoveNode(node)
    l.Dll.InsertAtLast(node)

    return node.Value, true
}


