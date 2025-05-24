package main

func main() {
	server := NewServer(2053)
	server.ListenAndServe()
}
