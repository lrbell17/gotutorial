package main

/*
 1. Beware of when goroutines start
    --> give time for them to start before program ends

 2. Beware of goroutines launch as closures
    --> closures referencing the *same variable*

 3. Beware of sync.WaitGroup
    --> don't copy sync.WaitGroup object
*/

func main() {

}
