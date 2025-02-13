package main

import (
	"third_site/app"
	"flag"
)

func main() {
	port := flag.Int("p", 8000, "WEB端口")
	flag.Parse()
	app.Start(*port)
}
