package main

import (
        "net/http"
        "fmt"

        "github.com/go-chi/chi"
        "github.com/go-chi/chi/middleware"
)

func main() {
        r := chi.NewRouter()
        r.Use(middleware.RequestID)
        r.Use(middleware.Logger)
        r.Use(middleware.RealIP)

        //fmt.Println(&middleware.RealIP)

        r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println(*middleware.RealIP)
                w.Write([]byte("middleware.RealIP"))
        })

        http.ListenAndServe(":8888", r)
}
