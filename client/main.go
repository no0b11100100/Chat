package main

// https://dev-gang.ru/article/golang-prostoy-server-tcp-i-tcp-klient/

func main() {
	client := NewClient()
	client.Run()
	defer client.Destroy()
}
