// простой HTTP/JSON сервер , принимет и откликается JSON
// curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'
// curl -X GET localhost:8080 -d '{"offset": 0}'
package main

import (
	"log"

	"github.com/KVexcavator/chasm-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
