* [ ] 

```markdown
# Go Pipeline Example

This project demonstrates a simple pipeline implementation in Go that processes a stream of random integers.
 It includes a generator, squaring, and summing stages, all running concurrently with goroutines.

```


```markdown
## How to Run the Program

1.**Clone the Repository**
```bash
   git clone https://github.com/yourusername/go-pipeline.git
   cd go-pipeline
```



## Code Explanation

* **Generator** : Produces random integers and sends them to a channel.
* **Square** : Receives integers, squares them, and sends the results to another channel.
* **Sum** : Sums up squared integers and prints the final result.

### Error Handling and Shutdown

* The program includes proper error handling through bounded channels and ensures graceful shutdown of goroutines using `sync.WaitGroup`.
