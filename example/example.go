package main

import (
	"fmt"

	"github.com/coreos/etcd/client"
	_ "github.com/pengux/go-logz-api-client"
)

func main() {
	c := client.New("{LOGZ_API_TOKEN}")

	alerts := c.ListAlerts()
	fmt.Println(alerts)
}
