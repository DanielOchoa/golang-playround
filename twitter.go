package main

import (
    "encoding/json"
    "fmt"
    "reflect"
    "time"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

// using `var val map[string]interface{}` works but just uses string as value,
// if we want time, we change the map vlaue type to time.Time but this current
// str time format is not the default for go so we need to convert it to parsable
// time.
// instead if interface{...} type, we adopt time.Time interface.

// json.Unmarshal will now use this call below, because this is similar to the default
// implementation of Unmarshal. We pass a pointer reference so the caller knows
// how the argument changed, and not just copied a value.

// In other words, when json.Unmarshall tries to typecast? the value into Timestamp portion
// of our `val` map, it will call the `UnmarshalJSON` method for the `Timestamp` interface.
// our Interface inherits from time.Time it seems? That's how it knows to call
// `UnmarshalJSON` ?
type Timestamp time.Time
func (t *Timestamp) UnmarshalJSON(b []byte) error {
    // slice of string is removing trailing and prefix double quotes.
    v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
    if err != nil {
        return err
    }
    *t = Timestamp(v)
    return nil
}

func main() {
    // our target will be of type map[string]interface{}, which is a
    // pretty generic type that will give us a hashtable whose keys
    // are strings, and whose values are of type interface{}

    var val map[string]Timestamp

    if err := json.Unmarshal([]byte(input), &val); err != nil {
        panic(err)
    }

    for k, v := range val {
        fmt.Println(k, reflect.TypeOf(v))
    }

    fmt.Println(time.Time(val["created_at"]))
}
