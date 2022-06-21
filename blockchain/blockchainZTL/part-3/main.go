package main

func Test3() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}

func main() {

	Test3()

}
