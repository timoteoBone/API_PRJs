package main

import (
	//"fmt"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	calculos "github.com/timoteoBone/API_PRJs/calculatorPrj/calculo"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {

		urlElem := strings.Split(rq.URL.Path, "/")

		if urlElem[1] == "calcular" {

			if _, ok := calculos.Operations[urlElem[2]]; ok {

				//nums := convert(urlElem[2:])
				fmt.Fprintln(w, urlElem[3])

			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found"))
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//http:://localhost:8000/calcular/sumar/4 4 4

}

func convert(nums []string) []int {
	var res []int
	for k, i := range nums {
		if num, err := strconv.Atoi(i); err == nil {
			res[k] = num
		}

	}
	return res
}

