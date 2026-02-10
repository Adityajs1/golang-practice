package main
import "fmt"
type post struct{
	views int
}

func (p *post) inc(){
	p.views += 1
}


func mutexx(){
	postCount := post{views: 0}
	postCount.inc()
	postCount.inc()
	postCount.inc()

	fmt.Println(postCount.views)
}