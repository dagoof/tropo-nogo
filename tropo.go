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
    TERMINATOR = "terminator"
    MODE = "mode"
    BEEP = "beep"
    FORMAT = "format"
    MAXSILENCE = "maxSilence"
    MAXTIME = "maxTime"
    METHOD = "method"
    TRANSCRIPTION = "transcription"
    URL = "url"
    PASSWORD = "password"
    USERNAME = "username"
    ID = "id"
    MUTE = "mute"
    PLAYTONES = "playTones"
    ACTIONS = "actions"
    CALLID = "callId"
    COMPLETE = "complete"
    ERROR = "error"
    SEQUENCE = "sequence"
    SESSIONDURATION = "sessionDuration"
    SESSIONID = "sessionId"
    STATE = "state"

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

// Choices structure

var CHOICES_FIELDS Constrained = Constrain(
    VALUE, TERMINATOR, MODE)

type _Choices struct {
    Fields JSFields `json:"choices"`
}

func (choices _Choices) Set(field string, value interface{ }) Setter {
    JSSet(CHOICES_FIELDS, choices.Fields, field, value)
    return Setter(choices)
}

func (choices _Choices) AddArg(setter Setter) Setter {
    return setter.Set(CHOICES, choices)
}

func Choices(value string, args... Arg) _Choices {
    choices := _Choices{ JSFields{ } }
    AddArgs(choices, append(args, Value(value))...)
    return choices
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

func Ask(choices _Choices, name string, args... Arg) _Ask {
    ask := _Ask{ JSFields{ } }
    AddArgs(ask, append(args, choices, Name(name))...)
    return ask
}

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

// Conference command

var CONFERENCE_FIELDS Constrained = Constrain(
    ID, ALLOWSIGNALS, INTERDIGITTIMEOUT, MUTE, NAME, PLAYTONES,
    REQUIRED, TERMINATOR)

type _Conference struct {
    Fields JSFields `json:"message"`
}

func (conference _Conference) Set(field string,  value interface{ }) Setter {
    JSSet(CONFERENCE_FIELDS, conference.Fields, field, value)
    return Setter(conference)
}

func Conference(id string, args... Arg) _Conference {
    conference := _Conference{ JSFields{ } }
    AddArgs(call, append(args, Id(id))...)
    return conference
}

// Hangup command

var HANGUP_FIELDS Constrained = Constrain(HEADERS)

type _Hangup struct {
    Fields JSFields `json:"hangup"`
}

func (hangup _Hangup) Set(field string, value interface{ }) Setter {
    JSSet(HANGUP_FIELDS, hangup.Fields, field, value)
    return Setter(hangup)
}

func Hangup(args... Arg) _Hangup {
    hangup := _Hangup{ JSFields{ } }
    AddArgs(hangup, args...)
    return hangup
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

// Record command

var RECORD_FIELDS Constrained = Constrain(
    ATTEMPTS, ALLOWSIGNALS, BARGEIN, BEEP, CHOICES, SAY,
    FORMAT, MAXSILENCE, MAXTIME, METHOD, MINCONFIDENCE,
    NAME, REQUIRED, TRANSCRIPTION, URL, PASSWORD, USERNAME,
    TIMEOUT, INTERDIGITTIMEOUT, VOICE)

type _Record struct {
    Fields JSFields `json:"record"`
}

func (record _Record) Set(field string, value interface{ }) Setter {
    JSSet(RECORD_FIELDS, record.Fields, field, value)
    return Setter(record)
}

func (record _Record) AddArg(setter Setter) Setter {
    return setter.Set(RECORDING, record)
}

func Record(name string, url string, args... Arg) _Record {
    record := _Record{ JSFields{ } }
    AddArgs(record, append(args, Name(name), Url(url))...)
    return record
}

var Recording func(string, string, ...Arg) _Record = Record

// Redirect command

var REDIRECT_FIELDS Constrained = Constrain(
    TO, NAME, REQUIRED)

type _Redirect struct {
    Fields JSFields `json:"redirect"`
}

func (redirect _Redirect) Set(field string, value interface{ }) Setter {
    JSSet(REDIRECT_FIELDS, redirect.Fields, field, value)
    return Setter(redirect)
}

func Redirect(to string, args... Arg) _Redirect {
    redirect := _Redirect{ JSFields{ } }
    AddArgs(redirect, append(args, To(to))...)
    return redirect
}

// Reject command

var REJECT_FIELDS Constrained = Constrain()

type _Reject struct {
    Fields JSFields `json:"reject"`
}

func (reject _Reject) Set(field string, value interface{ }) Setter {
    JSSet(REJECT_FIELDS, reject.Fields, field, value)
    return Setter(reject)
}

func Reject(args... Arg) _Reject {
    reject := _Reject{ JSFields{ } }
    AddArgs(reject, args...)
    return reject
}

// Result command

var RESULT_FIELDS Constrained = Constrain(
    ACTIONS, CALLID, COMPLETE, ERROR, SEQUENCE,
    SESSIONDURATION, SESSIONID, STATE)


type _Result struct {
    Fields JSFields `json:"result"`
}

func (result _Result) Set(field string, value interface{ }) Setter {
    JSSet(RESULT_FIELDS, result.Fields, field, value)
    return Setter(result)
}

func Result(args... Arg) _Result {
    result := _Result{ JSFields{ } }
    AddArgs(result, args...)
    return result
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

type Terminator string
func (s Terminator) AddArg(setter Setter) Setter {
    return setter.Set(TERMINATOR, string(s))
}

type Mode string
func (s Mode) AddArg(setter Setter) Setter {
    return setter.Set(MODE, string(s))
}

type Beep bool
func (b Beep) AddArg(setter Setter) Setter {
    return setter.Set(BEEP, bool(b))
}

type Format string
func (s Format) AddArg(setter Setter) Setter {
    return setter.Set(FORMAT, string(s))
}

type MaxSilence float32
func (f MaxSilence) AddArg(setter Setter) Setter {
    return setter.Set(MAXSILENCE, float32(f))
}

type MaxTime float32
func (f MaxTime) AddArg(setter Setter) Setter {
    return setter.Set(MAXTIME, float32(f))
}

type Method string
func (s Method) AddArg(setter Setter) Setter {
    return setter.Set(METHOD, string(s))
}

//type Transcription 

type Url string
func (s Url) AddArg(setter Setter) Setter {
    return setter.Set(URL, string(s))
}

type Password string
func (s Password) AddArg(setter Setter) Setter {
    return setter.Set(PASSWORD, string(s))
}

type Username string
func (s Username) AddArg(setter Setter) Setter {
    return setter.Set(USERNAME, string(s))
}

type Id string
func (s string) AddArg(setter Setter) Setter {
    return setter.Set(STRING, string(s))
}

type Mute bool
func (b Mute) AddArg(setter Setter) Setter {
    return setter.Set(MUTE, bool(b))
}

type PlayTones bool
func (b PlayTones) AddArg(setter Setter) Setter {
    return setter.Set(PLAYTONES, bool(b))
}

//type Actions

type CallId string
func (s CallId) AddArg(setter Setter) Setter {
    return setter.Set(CALLiD, string(s))
}

type Complete bool
func (b type) AddArg(setter Setter) Setter {
    return setter.Set(tYPE, bool(b))
}

type Error string
func (e Error) AddArg(setter Setter) Setter {
    return setter.Set(ERROR, string(e))
}

type Sequence int
func (i Sequence) AddArg(setter Setter) Setter {
    return setter.Set(SEQUENCE, int(i))
}

type SessionDuration int
func (i SessionDuration) AddArg(setter Setter) Setter {
    return setter.Set(SESSIONdURATION, int(i))
}

type SessionId string
func (s SessionId) AddArg(setter Setter) Setter {
    return setter.Set(SESSIONiD, string(s))
}

type State string
func (s State) AddArg(setter Setter) Setter {
    return setter.Set(STATE, string(s))
}



// DEAL WITH HEADERS / TRANSCRIPTION / ACTIONS

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
