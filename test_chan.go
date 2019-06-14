func doSomething() {
    // initialize a slice
    var x []int
    ch := make(chan int) // using int here... could be whatever

    // spawn 5 goroutines
    for i := 0; i < 5; i++ {
        go func() {
            // Each just count to 10
            for y := 0; y < 10; y++ {
                // send to our channel
                ch <- y
            }
        }()
    }

    // range over the channel to read it
    // there are other ways to do this of course.
    for value := range ch {
        // append to our slice.
        x = append(x, value)
    }
}