package cockpit

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *AppHandler) HelloGetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`hello`))
	if err != nil {
		log.Fatal(err)
	}
}

func (h *AppHandler) HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	x, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal(e)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(fmt.Sprintf("Failed closing the body stream: %s", err.Error()))
		}
	}(r.Body)

	_, err := w.Write(x)
	if err != nil {
		log.Println(err)
	}
}
