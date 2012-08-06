package main

import (
    "fmt"
    "encoding/json"
)

const (
    VALUE = "value"
    ALLOWSIGNALS = "allowSignals"
    AS = "as"
    NAME = "name"
    REQUIRED = "required"
    VOICE = "voice"
    SAY = "say"
    TO = "to"
    ANSWERONMEDIA = "answerOnMedia"
    CHANNEL = "channel"
    FROM = "from"
    NETWORK = "network"
    TIMEOUT = "timeout"
)


type Constrained map[string]bool
func Constrain(fields... string) Constrained {
    constraints := Constrained{ }
    for _, field := range fields {
        constraints[field] = true
    }
    return constraints
}

var SAY_FIELDS Constrained = Constrain(
    VALUE, ALLOWSIGNALS, AS, NAME, REQUIRED, VOICE)

var MESSAGE_FIELDS Constrained = Constrain(
    SAY, TO, ANSWERONMEDIA, CHANNEL, FROM, NAME, NETWORK,
    REQUIRED, TIMEOUT, VOICE)

type JSFields map[string]interface{ }
type _Say struct {
    Fields JSFields `json:"say"`
}

type _Message struct {
    Fields JSFields `json:"message"`
}

func (say _Say) Set(field string, value interface{ }) Setter {
    _, valid := SAY_FIELDS[field]
    if valid {
        say.Fields[field] = value
    }
    return Setter(say)
}

func (say _Say) AddArg(setter Setter) Setter {
    return setter.Set(SAY, say)
}

func (message _Message) Set(field string, value interface{ }) Setter {
    _, valid := MESSAGE_FIELDS[field]
    if valid {
        message.Fields[field] = value
    }
    return Setter(message)
}

type Setter interface {
    Set(string, interface{ }) Setter
}

type Arg interface {
    AddArg(Setter) Setter
}

type Value string
func (s Value) AddArg(setter Setter) Setter {
    return setter.Set(VALUE, string(s))
}

type AllowSignals string
func (s AllowSignals) AddArg(setter Setter) Setter {
    return setter.Set(ALLOWSIGNALS, string(s))
}

type As string
func (s As) AddArg(setter Setter) Setter {
    return setter.Set(AS, string(s))
}

type Name string
func (s Name) AddArg(setter Setter) Setter {
    return setter.Set(NAME, string(s))
}

type Voice string
func (s Voice) AddArg(setter Setter) Setter {
    return setter.Set(VOICE, string(s))
}

type Required bool
func (b Required) AddArg(setter Setter) Setter {
    return setter.Set(REQUIRED, bool(b))
}

type To string
func (s To) AddArg(setter Setter) Setter {
    return setter.Set(TO, string(s))
}


func Say(value string, args... Arg) _Say {
    say := _Say{ JSFields{ } }
    Value(value).AddArg(say)
    for _, arg := range args {
        arg.AddArg(say)
    }
    return say
}

func Message(say _Say, to string, args... Arg) _Message {
    message := _Message{ JSFields{ } }
    say.AddArg(message)
    To(to).AddArg(message)
    for _, arg := range args {
        arg.AddArg(message)
    }
    return message
}

func main() {
    heh := Message(
        Say(
            "Hello friend i eat crap",
            Voice("roger"),
            As("Jonathan"),
            AllowSignals("lol"),
            Required(true)),
        "+1555555555",
        Name("cockroach"))
    fmt.Println(heh)
    j, _ := json.Marshal(heh)
    fmt.Println(string(j))

}
