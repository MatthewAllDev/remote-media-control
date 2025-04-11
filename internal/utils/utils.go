package utils

import (
	"errors"
	"log"
	"net"
	"net/http"
)

func ChainMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func CspMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:;")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func GetServerIp() (net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, interf := range interfaces {
		addrs, err := interf.Addrs()
		if err != nil {
			return nil, err
		}
		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				if ip.IP.DefaultMask() != nil && !ip.IP.IsLoopback() && ip.IP.IsPrivate() {
					return ip.IP, nil
				}
			}
		}
	}
	return nil, errors.New("no private IP address found")
}
