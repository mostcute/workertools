module github.com/filecoin-project/yz-watch-manager/workertools

go 1.14
require (
	github.com/gorilla/websocket v1.4.2
	github.com/filecoin-project/yz-watch-manager/msg v0.0.0
	github.com/filecoin-project/yz-watch-manager/tools v0.0.0
)

replace github.com/filecoin-project/yz-watch-manager/tools => ../tools

replace github.com/filecoin-project/yz-watch-manager/msg => ../msg
