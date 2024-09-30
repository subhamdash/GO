package main
import (
	"fmt" 
	"time"


)

/*
• If your method modifies the receiver, you must use a pointer receiver.
• If your method needs to handle nil instances then it must use a pointer receiver.
• If your method doesn’t modify the receiver, you can use a value receiver.
*/


type persons struct {
	FirstName string
	LastName string
	Age int
   }
   func (p persons) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
   }




   type Counter struct {
	total int
	lastUpdated time.Time
   }
   func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
   }
   func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
   }





   func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
   }
   func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
   }


func main08(){

	type Score int
	type Converter func(string)Score
	type TeamScores map[string]Score

	p := persons {
		FirstName: "Fred",
		LastName:"Fredson",
		Age: 52,
	   }
	   output := p.String()
	   fmt.Println(output)
	   var c Counter
doUpdateWrong(c)
 fmt.Println("in main:", c.String())
 doUpdateRight(&c)
 fmt.Println("in main:", c.String())

}
/*
Fred Fredson, age 52
in doUpdateWrong: total: 1, last updated: 2024-05-13 13:05:58.640411 +0530 IST m=+0.000185876
in main: total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
in doUpdateRight: total: 1, last updated: 2024-05-13 13:05:58.640623 +0530 IST m=+0.000397668
in main: total: 1, last updated: 2024-05-13 13:05:58.640623 +0530 IST m=+0.000397668
*/