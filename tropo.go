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
    CHOICES = "choices"
    ATTEMPTS = "attempts"
    HEADERS = "headers"
    RECORDING = "recording"
    BARGEIN = "bargein"
    INTERDIGITTIMEOUT = "interDigitTimeout"
    MINCONFIDENCE = "minConfidence"
    RECOGNIZER = "recognizer"
)


type JSFields map[string]interface{ }

func JSSet(available Constrained, fields JSFields,
        field string, value interface{ }) {
    _, valid := available[field]
    if valid {
        fields[field] = value
    }
}

type Constrained map[string]bool
func Constrain(fields... string) Constrained {
    constraints := Constrained{ }
    for _, field := range fields {
        constraints[field] = true
    }
    return constraints
}

// Ask command

var ASK_FIELDS Constrained = Constrain(
    CHOICES, ALLOWSIGNALS, ATTEMPTS, BARGEIN, INTERDIGITTIMEOUT,
    MINCONFIDENCE, NAME, RECOGNIZER, REQUIRED, SAY, TIMEOUT, VOICE)

type _Ask struct {
    Fields JSFields `json:"ask"`
}

func (ask _Ask) Set(field string, value interface{ }) Setter {
    JSSet(ASK_FIELDS, ask.Fields, field, value)
    return Setter(ask)
}

func Ask(args... Arg) _Ask {
    ask := _Ask{ JSFields{ } }
    AddArgs(ask, args...)
    return ask
}

/*
func Ask(choices Choices, args... Arg) _Ask {
    ask := _Ask{ JSFields{ } }
    AddArgs(ask, choices, args...)
    return ask
}
*/

// Call command

var CALL_FIELDS Constrained = Constrain(
    TO, ALLOWSIGNALS, ANSWERONMEDIA, CHANNEL, FROM, HEADERS,
    NAME, NETWORK, RECORDING, REQUIRED, TIMEOUT)

type _Call struct {
    Fields JSFields `json:"call"`
}

func (call _Call) Set(field string, value interface{ }) Setter {
    JSSet(CALL_FIELDS, call.Fields, field, value)
    return Setter(call)
}

func Call(to string, args... Arg) _Call {
    call := _Call{ JSFields{ } }
    AddArgs(call, append(args, To(to))...)
    return call
}

// Say command
var SAY_FIELDS Constrained = Constrain(
    VALUE, ALLOWSIGNALS, AS, NAME, REQUIRED, VOICE)

type _Say struct {
    Fields JSFields `json:"say"`
}

func (say _Say) Set(field string, value interface{ }) Setter {
    JSSet(SAY_FIELDS, say.Fields, field, value)
    return Setter(say)
}

func (say _Say) AddArg(setter Setter) Setter {
    return setter.Set(SAY, say)
}

func Say(value string, args... Arg) _Say {
    say := _Say{ JSFields{ } }
    AddArgs(say, append(args, Value(value))...)
    return say
}


// Message command

var MESSAGE_FIELDS Constrained = Constrain(
    SAY, TO, ANSWERONMEDIA, CHANNEL, FROM, NAME, NETWORK,
    REQUIRED, TIMEOUT, VOICE)

type _Message struct {
    Fields JSFields `json:"message"`
}

func (message _Message) Set(field string, value interface{ }) Setter {
    JSSet(MESSAGE_FIELDS, message.Fields, field, value)
    return Setter(message)
}

func Message(say _Say, to string, args... Arg) _Message {
    message := _Message{ JSFields{ } }
    AddArgs(message, append(args, say, To(to))...)
    return message
}

// On command

var ON_FIELDS Constrained = Constrain(
    EVENT, NEXT, SAY)

type _On struct {
    Fields JSFields `json:"on"`
}

func (on _On) Set(field string, value interface{ }) Setter {
    JSSet(ON_FIELDS, on.Fields, field, value)
    return Setter(on)
}

func On(event string, args... Arg) _On {
    on := _On{ JSFields{ } }
    AddArgs(on, append(args, Event(event))...)
    return on
}

// Arguments

type Setter interface {
    Set(string, interface{ }) Setter
}

type Arg interface {
    AddArg(Setter) Setter
}

func AddArgs(setter Setter, args... Arg) {
    for _, arg := range args {
        arg.AddArg(setter)
    }
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

type Attempts int
func (i Attempts) AddArg(setter Setter) Setter {
    return setter.Set(ATTEMPTS, int(i))
}

type Bargein bool
func (b Bargein) AddArg(setter Setter) Setter {
    return setter.Set(BARGEIN, bool(b))
}

type InterdigitTimeout int
func (i InterdigitTimeout) AddArg(setter Setter) Setter {
    return setter.Set(INTERDIGITTIMEOUT, int(i))
}

type MinConfidence int
func (i MinConfidence) AddArg(setter Setter) Setter {
    return setter.Set(MINCONFIDENCE, int(i))
}

type Recognizer string
func (s Recognizer) AddArg(setter Setter) Setter {
    return setter.Set(RECOGNIZER, string(s))
}

// DEAL WITH HEADERS / CHOICES / RECORDING

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
        "failure",
        Say("That sure was bad"))

    fmt.Println(heh)
    j, _ := json.Marshal(heh)
    fmt.Println(string(j))

}
