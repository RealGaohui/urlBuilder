# urlBuilder


```
package main

import (
	"fmt"
	"github.com/RealGaohui/urlBuilder"
)

func main() {
	uriBuilder := urlBuilder.URLBuilder()
	uriBuilder.SetBase("http://localhost:8080")
	uriBuilder.SetPath("/test")
	uriBuilder.SetParameter("a", "b", "c", "d", "e", "f")
	fmt.Println(uriBuilder.ToString())
}

```
```bigquery
http://localhost:8080/test?a=b&c=d&e=f

Process finished with the exit code 0
```

