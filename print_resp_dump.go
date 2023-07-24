package gonep

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func printRespDump(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
