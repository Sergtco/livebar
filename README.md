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
![image](https://github.com/Sergtco/livebar/assets/75071620/eddec4a1-ae72-4f45-9b5a-ec77bab8a3d3)
![image](https://github.com/Sergtco/livebar/assets/75071620/a9ede021-bb1e-4f75-a75e-5fd39b5c4204)
