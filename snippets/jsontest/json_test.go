package jsontest

import (
	"encoding/json"
	"testing"
	"time"
)

type op string

const (
	opSubscribe   = op("subscribe")
	opUnsubscribe = op("unsubscribe")
)

type command struct {
	Op   op            `json:"op"`
	Args []interface{} `json:"args"`
}

type foo struct {
	A   string
	B   bool
	T   time.Time
	Bar bar
}

type bar struct {
	X float64
}

func TestInterfaces(t *testing.T) {
	apple := &foo{
		A: "apple",
	}
	alpha := &foo{
		A: "alpha",
		B: true, T: time.Now(),
	}
	cmd := command{
		Op: opSubscribe,
		Args: []interface{}{
			apple,
			alpha,
		},
	}
	bs, err := json.Marshal(cmd)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cmd: %s\n", bs)

	var cmdFromJSON command
	err = json.Unmarshal(bs, &cmdFromJSON)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cmd: %+v\n", cmdFromJSON)
}

func TestBasic(t *testing.T) {
	cmd := command{
		Op: opSubscribe,
		Args: []interface{}{
			1,
			2.0,
			true,
			'A',
			"apple",
		},
	}
	bs, err := json.Marshal(cmd)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cmd: %s\n", bs)

	var cmdFromJSON command
	err = json.Unmarshal(bs, &cmdFromJSON)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cmd: %+v\n", cmdFromJSON)
}
