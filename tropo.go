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
    EVENT = "event"
    NEXT = "next"
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

var ON_FIELDS Constrained = Constrain(
    EVENT, NEXT, SAY)

type JSFields map[string]interface{ }

// Say command
type _Say struct {
    Fields JSFields `json:"say"`
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

// Message command
type _Message struct {
    Fields JSFields `json:"message"`
}

func (message _Message) Set(field string, value interface{ }) Setter {
    _, valid := MESSAGE_FIELDS[field]
    if valid {
        message.Fields[field] = value
    }
    return Setter(message)
}

// On object
type _On struct {
    Fields JSFields `json:"on"`
}

func (on _On) Set(field string, value interface{ }) Setter {
    _, valid := ON_FIELDS[field]
    if valid {
        on.Fields[field] = value
    }
    return Setter(on)
}


// Arguments

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

type Required bool
func (b Required) AddArg(setter Setter) Setter {
    return setter.Set(REQUIRED, bool(b))
}

type Voice string
func (s Voice) AddArg(setter Setter) Setter {
    return setter.Set(VOICE, string(s))
}

type To string
func (s To) AddArg(setter Setter) Setter {
    return setter.Set(TO, string(s))
}

type AnswerOnMedia bool
func (b AnswerOnMedia) AddArg(setter Setter) Setter {
    return setter.Set(ANSWERONMEDIA, bool(b))
}

type Channel string
func (s Channel) AddArg(setter Setter) Setter {
    return setter.Set(CHANNEL, string(s))
}

type From string
func (s From) AddArg(setter Setter) Setter {
    return setter.Set(FROM, string(s))
}

type Network string
func (s Network) AddArg(setter Setter) Setter {
    return setter.Set(NETWORK, string(s))
}

type Timeout float32
func (f Timeout) AddArg(setter Setter) Setter {
    return setter.Set(TIMEOUT, float32(f))
}

type Event string
func (s Event) AddArg(setter Setter) Setter {
    return setter.Set(EVENT, string(s))
}

type Next string
func (s Next) AddArg(setter Setter) Setter {
    return setter.Set(NEXT, string(s))
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

func On(event, next string, args... Arg) _On {
    on := _On{ JSFields{ } }
    Event(event).AddArg(on)
    Next(next).AddArg(on)
    for _, arg := range args {
        arg.AddArg(on)
    }
    return on
}

func main() {
    /*
    heh := Message(
        Say(
            "Hello friend i eat crap",
            Voice("roger"),
            As("Jonathan"),
            AllowSignals("lol"),
            Required(true)),
        "+1555555555",
        Name("cockroach"))
    */

    heh := On(
        "failure", "next/more",
        Say("That sure was bad"))

    fmt.Println(heh)
    j, _ := json.Marshal(heh)
    fmt.Println(string(j))

}
