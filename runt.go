package main

import (
    "fmt"
    "encoding/json"
    "tropo"
)

func main() {
    message := tropo.Message(
        tropo.Say(
            "Hello world",
            tropo.Voice("Roger"),
            tropo.As("Jonathan"),
            tropo.AllowSignals("lol"),
            tropo.Required(true)),
        "+1555555555",
        tropo.Name("Schwartz"))

    on := tropo.On(
        "failure",
        tropo.Say("That sure was bad"))

    m, _ := json.MarshalIndent(message, "", "  ")
    o, _ := json.MarshalIndent(on, "", "  ")
    fmt.Println(string(m))
    fmt.Println(string(o))

}
