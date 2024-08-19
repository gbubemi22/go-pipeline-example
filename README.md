* [ ] 

```markdown
# Go Pipeline Example

This project demonstrates a simple pipeline implementation in 
Go that processes a stream of random integers.

 It includes a generator, squaring, and summing stages, all running concurrently with goroutines.

```


```markdown
## How to Run the Program

1.**Clone the Repository**
```bash
for http
   https://github.com/gbubemi22/go-pipeline-example.git
for ssh
git@github.com:gbubemi22/go-pipeline-example.git
   cd go-pipeline
```



## Code Explanation

* **Generator** : Produces random integers and sends them to a channel.
* **Square** : Receives integers, squares them, and sends the results to another channel.
* **Sum** : Sums up squared integers and prints the final result.

### Error Handling and Shutdown

* The program includes proper error handling through bounded channels and ensures graceful shutdown of goroutines using `sync.WaitGroup`.
