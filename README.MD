# MOVING
**MOVING** is a restful api server written in **Golang** for [html5-moving-vue](https://github.com/alydnhrealgang/html5-moving-vue).

## Overview

- The REST API server is described by an OpenAPI 2.0 specification file located at `server/swagger.yaml`, which was generated using [go-swagger](https://github.com/go-swagger/go-swagger).

- The **MOVING** data are stored in [dgraph-io/badger](https://github.com/dgraph-io/badger) located at `data/db`.
> Please note that there's a **jounery log** located at `data/db/jounery.log` which records all operations in sequence.

- All **ARTICLE** and **BOX** image files located at `data/assets` and theirs metadata are aslo stored in [dgraph-io/badger](https://github.com/dgraph-io/badger)

- There's a command tool generated using [spf13/cobra](https://github.com/spf13/cobra) which supports import/export data from/to an Excel file by selected item's tags, fields and properties. It's easily view/edit all **Moving Items**  in Excel application.

- The `certs` folder contains SSL certificate files that enable the server to run under the HTTPS scheme.

> Because the client-side is running in SSL mode, the server must also be run in SSL mode.

## Build & Run
```bash
> cd moving
> go mod tidy // run only once
> cd server/cmd
> go build -o moving-server main.go
> ./moving-server --tls-certificate=ssl/192.168.31.49.crt --tls-key=ssl/192.168.31.49.key --tls-host=192.168.31.49 --tls-port=8443 --scheme=https
```