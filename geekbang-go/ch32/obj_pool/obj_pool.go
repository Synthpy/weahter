package obj_pool

func c(int) int {
	return 0
}

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

type tt struct {
	New func() interface{}
	ll  int
}

func s()int{
	return 1
}

var ttt tt{
	New s
	ll 3
}
