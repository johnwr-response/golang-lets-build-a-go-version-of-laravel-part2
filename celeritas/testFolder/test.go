package testFolder

import "net/http"

func TestHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("It works"))
}
