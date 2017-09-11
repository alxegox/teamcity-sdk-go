TeamCity API bindings
=====================

This is a simple wrapper around the TeamCity API.

[![GoDoc](https://godoc.org/github.com/Cardfree/teamcity-sdk-go?status.png)](https://godoc.org/github.com/Cardfree/teamcity-sdk-go)

Sample usage:

```
package main

import "github.com/Cardfree/teamcity-go-sdk/teamcity"

func main() {
	client := teamcity.New("myinstance.example.com", "username", "password")

	b, err := client.QueueBuild("Project_build_task", "master", nil)
	if err != nil {
		fmt.Printf("You're outta luck: %s\n", err)
		return
	}

	fmt.Printf("Build: %#v\n", b)
}
```
