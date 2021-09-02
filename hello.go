package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Sep  1 15:23:40 voucher-service php[30600]: [PHP Worker Info]: [date: 2021-09-01 15-23-40] [tags: [masterId: RID-voucher_service-proxy-612f7edcbce217.99085409][lastId: RID-voucher_service-api-612f7edcbf77b8.24993753][currentId: RID-worker-612f7edcc34411.20432524]] [worker-1] [memory usage: 2.53 MB] Command handler COMPLETE. Type: Voucher\\Create\\Single [Exec time: 0.019s]")
	_, err := fmt.Fprintf(w, "pong\n")

	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	http.HandleFunc("/ping", pong)

	http.ListenAndServe(":8090", nil)
}