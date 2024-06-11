package kvraft

type MemoryKVStateMachine struct {
	KV map[string]string
}

func NewMemoryKVStateMachine() *MemoryKVStateMachine {
	// 创建一个集合，里面的元素是KV键值对
	return &MemoryKVStateMachine{
		KV: make(map[string]string),
	}
}

func (mkv *MemoryKVStateMachine) Get(key string) (string, Err) {
	if value, ok := mkv.KV[key]; ok {
		return value, OK
	}
	return "", ErrNoKey
}

func (mkv *MemoryKVStateMachine) Put(key, value string) Err {
	mkv.KV[key] = value
	return OK
}

func (mkv *MemoryKVStateMachine) Append(key, value string) Err {
	mkv.KV[key] += value
	return OK
}
