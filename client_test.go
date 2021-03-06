package gochatwork

import (
        "testing"
        "reflect"
)

const ApiKey = ``

func expect(t *testing.T, a interface{}, b interface{}) {
        if a != b {
                t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
        }
}

func refute(t *testing.T, a interface{}, b interface{}) {
        if a == b {
                t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
        }
}

func TestNewClient(t *testing.T) {
        c := NewClient(ApiKey)
        refute(t, c, nil)
}
