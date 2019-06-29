# go-logz-api-client
Go client for [Logz Public API](https://github.com/logzio/public-api)

## Usage
```
package main

import (
	"fmt"

	client "github.com/pengux/go-logz-api-client"
)

func main() {
	c := client.New("{LOGZ_API_TOKEN}")

	alerts := c.ListAlerts()
	fmt.Println(alerts)
}
```



## Install dependencies
```
make deps
```

## Tests
```
make test
```
