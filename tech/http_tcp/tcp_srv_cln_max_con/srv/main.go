package main

import "github.com/radisvaliullin/test-golang/tech/http_tcp/tcp_srv_cln_max_con/srvcln"

func main() {

	srv := srvcln.Srv{
		Addr: "0.0.0.0:7373",
	}

	srv.Run()
}
