// cmd/chainedgraphs/main.go
package main

import (
"flag"
"log"
"os"

"chainedgraphs/internal/chainedgraphs"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chainedgraphs.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
