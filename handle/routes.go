package handle

import "net/http"

func setupRoutes() {
	http.HandleFunc("/", home)
}
