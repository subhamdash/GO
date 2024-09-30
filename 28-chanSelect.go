
package main
import (
    "fmt"
	//"time"

)

func main28() {
    ch := make(chan int) 
	ch2 := make(chan int)

	go func(){
		ch <- 10; 
		a:= <-ch2
		fmt.Println(a)

	}()
	var a int
	select
	{
	case ch2<-10:
	case a=<-ch:
	}
	fmt.Println(a)


}

/*


The
select algorithm is simple: it picks randomly from any of its cases that can go for‐
ward; order is unimportant. This is very different from a switch statement, which
always chooses the first case that resolves to true. It also cleanly resolves the starva‐
tion problem, as no case is favored over another and all are checked at the same time.
Another advantage of select choosing at random is that it prevents one of the most
common causes of deadlocks: acquiring locks in an inconsistent order. If you have
two goroutines that both access the same two channels, they must be accessed in the
same order in both goroutines, or they will deadlock. This means that neither one can
proceed because they are waiting on each other.

*/