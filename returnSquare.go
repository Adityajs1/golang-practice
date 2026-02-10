package main
import ("fmt")


func returnSquare(){
	number := make(chan int)
	squares := make(chan int)

	go func(){
		for i:=0; i<=5; i++{
			number <- i;
		}
		close(number)
	}();

	go func(){
     for n := range number{
		squares <- n*n;
	 }
	 close(squares)
	}()

	for s := range squares{
		fmt.Println(s)
	}
}