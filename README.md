# Golang Concurrency

This repo is based on the Rob Pike video on Golang's Concurrency patterns called:
[Google I/O 2012 - Go Concurrency Patterns](https://youtu.be/f6kdp27TYZs).

## Sections

- [Non Concurrent](https://github.com/aldrickdev/go-concurrency/blob/1.non-concurrent/main.go)  
This branch just has a simple project that doesn't take advantage of go routines or channels.

- [Go Routine](https://github.com/aldrickdev/go-concurrency/blob/2.run-in-background/main.go)  
Here we can see how to use go routines to run functions in the background.

- [Go Routines while main continues](https://github.com/aldrickdev/go-concurrency/blob/3.run-in-backgroun-continue-main/main.go)  
Here we show how you can run a go routine in the background while main continues to execute.

- [Using Channels to communicate](https://github.com/aldrickdev/go-concurrency/blob/4.concurrency-and-channels/main.go)  
Here you learn to use channels to communicate between concurrently executing code.

- [The Generator Pattern](https://github.com/aldrickdev/go-concurrency/blob/5.Pattern-Generator/main.go)  
Here you learn how to use the Generator Pattern to simplify how the client receives data from another func by putting the go rountine in the function generating the data instead of in the client (the client being main in this case). In Go, a Generator is a function that returns a channel. 

- [Using a Generator to create service-like functions](https://github.com/aldrickdev/go-concurrency/blob/5.1.Using-generator-for-services/main.go)  
In this section we make use of the generator by lanuching it twice in the background.  

- [Using the Fan In pattern](https://github.com/aldrickdev/go-concurrency/blob/5.2.Sequencing-for-concurrent-services/main.go)  
In this section we see how we can use the FanIn Pattern to receive all of the results from the services from 1 channel. Here we also use a wait channel in order to syncronize the data we receive from the services.

- [Using Select](https://github.com/aldrickdev/go-concurrency/blob/6.Select/main.go)  
Here we learn how to use Select instead of using multiple go routines in the fanIn function making the function easier to read. Select is like a switch statement but it performs the first communication case available, this allows you to perform different operations depending on the first communication available.

- [Select with Timeout](https://github.com/aldrickdev/go-concurrency/blob/6.1.Select-with-timeout/main.go)  
In the section we make use of the time.After function that creates a channel and returns a value after the time has elapsed. This allows us to use Select as a way to timeout communications that are taking longer than expected. In this case we are using a time.After channel to detect when a single communication channel is taking too long, meaning during each execution of the loop, we restart the timer. However, in [This section](https://github.com/aldrickdev/go-concurrency/blob/6.2.Select-with-total-timeout/main.go) we use the time.after function to timeout after the total time elapsed, meaning we do not restart the timer for each iteration of the loop.

- [Select with Quit Channel](https://github.com/aldrickdev/go-concurrency/blob/6.3.Select-quit-channel/main.go)  
Here we learn how to stop a service using a quit channel allowing you to stop a service 

- [Select with Quit channel and Cleanup](https://github.com/aldrickdev/go-concurrency/blob/6.4.Select-quit-channel-with-cleanup/main.go)  
In this section we add to the quit channel by allowing the service to run some code before it stops in case it needs to perform some cleanup code before exiting.

- [Daisy Chaining](https://github.com/aldrickdev/go-concurrency/blob/7.Daisy-chain/main.go)  
Here we show a simple example of how you can daisy chain together channels and run thousands of go routines to pass data from one end of a chain to the other in seconds. This showcases how lightweight the go rountines are.

- []()


