package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting hello-world server...")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, `
  ____  _  _______ _____ _____ ___
 / __ \| |/ /_   _| ____|_   _/ _ \
| |  | | ' /  | | |  _|   | || | | |
| |__| | . \  | | | |___  | || |_| |
 \____/|_|\_\ |_| |_____| |_| \___/

        +------------------+
        | VS Code Extension |
        +------------------+

Hello from Okteto!

This Go service is running in an Okteto development container.
Edit locally with the VS Code extension, sync instantly, and debug
through the forwarded Delve port.
`)
}
