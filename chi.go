package main

import (
        "net/http"
        "net"
        "github.com/go-chi/chi"
        "github.com/go-chi/chi/middleware"
)


const (
        XForwardedFor = "X-Forwarded-For"
        XRealIP       = "X-Real-IP"
)

func RemoteIp(req *http.Request) string {
        remoteAddr := req.RemoteAddr
        if ip := req.Header.Get(XRealIP); ip != "" {
                remoteAddr = ip
        } else if ip = req.Header.Get(XForwardedFor); ip != "" {
                remoteAddr = ip
        } else {
                remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
        }

        if remoteAddr == "::1" {
                remoteAddr = "127.0.0.1"
        }

        return remoteAddr
}


func main() {
        r := chi.NewRouter()
        r.Use(middleware.RequestID)
        r.Use(middleware.Logger)
        r.Use(middleware.RealIP)
        r.Use(middleware.GetHead)


        r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte(RemoteIp(r)))
        })

        http.ListenAndServe(":8888", r)
}
