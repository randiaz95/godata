package main

import ("fmt"; "os"; "bufio"; "regexp"; "strings")

func main() {
    
    var filename string = "test_input.csv"
    
    var output []string = extract_email(filename)
    
    fmt.Println(output)

}


func extract_email(filename string) []string {
    // Get filename and get all emails

    var output []string = []string{}
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }

    reg, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if err != nil {
        fmt.Println("ERROR: ", err)
        panic(err)
    }

    file_scanner := bufio.NewScanner(file)
    var word string
    for file_scanner.Scan() {
        for _, word = range strings.Fields(file_scanner.Text()) {
            if reg.MatchString(word) {
                output = append(output, word)
            }
        }
    }

    return output

}
