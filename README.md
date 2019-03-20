# Example

```
import (
	"fmt"
	"math/rand"

	"github.com/dr4ds/worker-pool"
)

const WORKERS = 50

func processorMultiplyByTwo(job interface{}) interface{} {
	return job.(int) * 2
}

func main() {
	todo := make([]int, 10000)
	for i := 0; i < len(todo); i++ {
		todo[i] = rand.Intn(100)
	}

	fmt.Println(todo)

	arr := make([]interface{}, len(todo))
	for i, v := range todo {
		arr[i] = v
	}

	r := worker_pool.WorkerPool(arr, WORKERS, processorMultiplyByTwo)

	res := make([]int, 0)
	for v := range r {
		res = append(res, v.(int))
	}

	fmt.Println(res)
}
```
