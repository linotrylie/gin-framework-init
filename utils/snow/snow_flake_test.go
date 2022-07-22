package snow

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

type ZybService struct {
}
type YunService struct {
}

func (y *YunService) Show() {
	fmt.Println("11111")
}

func TestSnowFlake(t *testing.T) {
	map1 := make(map[string]struct{})
	map1["zyb"] = ZybService{}
	map1["yun"] = YunService{}
	var y YunService
	y = map1["yun"]
	y.Show()
	fmt.Println(y)
	i := 1
	for i <= 5 {
		go ShowId()
		i++
	}
}

var mutx sync.Mutex

func ShowId() {
	mutx.Lock()
	sID, err := SFlake.GetID()
	defer mutx.Unlock()
	if err != nil {
		log.Fatalf("snow flake get id err: %v\n", err)
	}
	fmt.Printf("Get sid from snow flake: %v\n", sID)
	return
}
