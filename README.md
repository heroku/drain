# Go library for writing Heroku drains

See [heroku/draincat](https://github.com/heroku/draincat) for example use case.

## Example

```Go
// main.go
package main

import (
	"github.com/heroku/drain"
)

func main() {
	drn := drain.NewDrain()
	http.HandleFunc("/logs", drn.LogsHandler)
	go http.ListenAndServe(":8080", nil)
	for line := range drn.Logs() {
		if lerr := line.Err(); lerr != nil {
			fmt.Printf("Logplex error: %+v\n", lerr)
		} else {
			fmt.Printf("Log line: %s\n", line.Data)
		}
	}
}
```
