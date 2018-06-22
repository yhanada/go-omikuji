package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yhanada/go-omikuji/omikuji"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	now := time.Now()

	result := omikuji.Draw(now.In(jst))
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(result); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
