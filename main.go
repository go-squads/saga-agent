package main

func main() {
	saga := App{}
	saga.Initialize()
	saga.Run(":9200")
}
