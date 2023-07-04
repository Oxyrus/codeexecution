# Code Execution

Creates an API with an endpoint that allows for remote Go code execution.

## Usage

```bash
curl -X POST -d 'package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}' http://localhost:8000/execute
```
