package main

import ("io/ioutil"; "fmt")

func main() {
	
	fmt.Println(read_file("test.txt"))

}


func read_file(filename string) string {
	// Input filename and return string
	
	stream, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	return string(stream)
}