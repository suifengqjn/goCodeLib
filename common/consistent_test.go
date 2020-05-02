package common

import (
	"fmt"
	"testing"
)

func TestNewConsistent(t *testing.T) {

	 cons := NewConsistent()

	 // 添加3个节点
	 cons.Add("node0")
	 cons.Add("node1")
	 cons.Add("node2")

	 // 查看1000条数据在3个节点中的数据分布
	 res := make(map[string]int)
	 for i := 0; i< 1000 ; i++  {
		node,err := cons.Get(fmt.Sprintf("%v", i))
		if err == nil {
			res[node] ++
		}
	 }

	fmt.Println(res)


}
