package main

import "fmt"

//==================================================

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}

//===================================================

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

//================================================

func main() {
	ch := make(chan int)
	ch <- 5	//program will panic here
}

//=================================================

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)	//create unidirectional channel
	go sendData(sendch)
	fmt.Println(<-sendch)
}

//===================================================

//use bidirectional channel as unidirectional

func sendData(sendch chan<- int) {	//converted to send only channel
	sendch <- 10
}

func main() {
	chnl := make(chan int)	//// bidirectional channel
}
	go sendData(chnl)
	fmt.Println(<-chnl)

// channel is send only inside the sendData Goroutine but it's bidirectional in the main Goroutine.

//====================================================

//Closing channels

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl) //close the channel
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch //ok will be false if we are reading from closed channel
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

//==================================================

//using for range loop

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
