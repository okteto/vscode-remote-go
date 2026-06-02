# Remote Development over Kubernetes with Visual Studio Code and Okteto

This example shows how to leverage [Okteto](https://github.com/okteto/okteto) to launch a remote development environment in Kubernetes, and use Visual Studio Code to develop directly in it.

This demo targets Go 1.25 or newer.

## Run the demo

Install and configure the latest Okteto CLI, then start the development container:

```sh
okteto up hello-world
```

From the Okteto shell, run the app:

```sh
go run main.go
```

The app is available at `localhost:8080`.

## Debug from VS Code

Inside the Okteto shell, start Delve:

```sh
dlv debug --headless --listen=:2345 --log --api-version=2
```

Then run the `Connect to okteto` configuration in VS Code.
