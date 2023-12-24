# Dynamic bar for your needs
# This project is in development mode and works for unix only!

# Example
```go
import (
	"fmt"
	"time"
	"github.com/Sergtco/livebar"
)

func main() {
    bar := livebar.NewBar(25, livebar.PacMan)
    fmt.Println("Processing")
    for i := 0; i <= 100; i++ {
        bar.Update(i)
        time.Sleep(time.Millisecond * 100)
    }
    fmt.Println("Processed")
}
```
![output](https://github.com/Sergtco/livebar/assets/75071620/582cf4bb-5f4c-4fda-b217-2b4b9a33bb01)
