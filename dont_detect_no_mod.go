package dont_detect_no_mod

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome Back Again! dont_detect_ni_mod!", name)
    return message
}
