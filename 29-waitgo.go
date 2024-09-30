package main
import(
	"fmt"
	"time"
	"sync"
)

func main29(){

	ch:=make(chan string,3)


	wg:= &sync.WaitGroup{}
	wg.Add(3)
	
	go fetch_user_details(ch,wg)
	go fetch_user_preferance(ch,wg)
	go fetch_user_recomendation(ch,wg)
	
	wg.Wait()

	close(ch)
	for a:= range ch{
		fmt.Println(a)
	}



}

func fetch_user_details(ch chan string,wg *sync.WaitGroup){
	time.Sleep(time.Millisecond*80)
	ch<-"I am subham dash";
	wg.Done()

}


func fetch_user_preferance(ch chan string,wg *sync.WaitGroup){
	time.Sleep(time.Millisecond*120)
	ch<-"I Like Ice cream";
	wg.Done()

}

func fetch_user_recomendation(ch chan string,wg *sync.WaitGroup){
	time.Sleep(time.Millisecond*50)
	ch<-"Trife, Moose , Pina Colda Lush, Tiramisu";
	wg.Done()

}

/*
A sync.WaitGroup doesn’t need to be initialized, just declared, as its zero value is use‐
ful. There are three methods on sync.WaitGroup: Add, which increments the counter
of goroutines to wait for; Done, which decrements the counter and is called by a
goroutine when it is finished; and Wait, which pauses its goroutine until the counter
hits zero. Add is usually called once, with the number of goroutines that will be
launched. Done is called within the goroutine. To ensure that it is called, even if the
goroutine panics, we use a defer.

We need to pass the pointer to pass the wait group so that we can get the value from it
*/