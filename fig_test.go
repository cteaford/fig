package fig

import (
	"testing"
    "sync"
    "os"
)

func TestConfigure(t *testing.T) {
	e := Entry{ Name: "TEST", Source: "env", DefaultValue: "test" }
	c := Config{ Entries: []Entry{e}, RefreshInterval: 0 }
	container, _ := Configure(c)
    resp := container.Get("TEST")
	if resp != "test" {
		t.Errorf("HALP")
	}
}

func TestConfigureMultiRoutine(t *testing.T) {
	e1 := Entry{ Name: "TEST1", Source: "env", DefaultValue: "test1" }
	e2 := Entry{ Name: "TEST2", Source: "env", DefaultValue: "test2" }
	e3 := Entry{ Name: "TEST3", Source: "env", DefaultValue: "test3" }
	c := Config{ Entries: []Entry{e1, e2, e3}, RefreshInterval: 0 }
	container, _ := Configure(c)
    pass := true
    var wg sync.WaitGroup
    go func () {
        wg.Add(1)
        resp := container.Get("TEST1")
        if resp != "test1" {
            pass = false
        }
        wg.Done()
    }()
    go func () {
        wg.Add(1)
        resp := container.Get("TEST2")
        if resp != "test2" {
            pass = false
        }
        wg.Done()
    }()
    go func () {
        wg.Add(1)
        resp := container.Get("TEST4")
        if resp != "test3" {
            pass = false
        }
        wg.Done()
    }()
    wg.Wait()
    if !pass {
        t.Errorf("HALP")
    }
}

func TestConfigureThreadRipperEdition(t *testing.T) {
	e := Entry{ Name: "TEST", Source: "env", DefaultValue: "test" }
	c := Config{ Entries: []Entry{e}, RefreshInterval: 0 }
	container, _ := Configure(c)

    var wg sync.WaitGroup
    pass := true
    for _ = range 1000 {
        go func () {
            wg.Add(1)
            resp := container.Get("TEST")
            if resp != "test" {
                pass = false
            }
            wg.Done()
        }()
    }

    wg.Wait()

    if !pass {
        t.Errorf("HALP")
    }
}

func TestConfigureDefaultValue(t *testing.T) {
    const expectedValue = "test"
	e := Entry{ Name: "TEST", Source: "env", DefaultValue: "test" }
	c := Config{ Entries: []Entry{e}, RefreshInterval: 0 }
	container, _ := Configure(c)

    resp := container.Get("TEST")
    if resp != expectedValue {
        t.Errorf("HALP")
    }
}

func TestConfigureEnvSource(t *testing.T) {
    const expectedValue = "testing"
	e := Entry{ Name: "TEST", Source: "env", DefaultValue: "test" }
	c := Config{ Entries: []Entry{e}, RefreshInterval: 0 }
    os.Setenv("TEST", "testing")
	container, _ := Configure(c)

    resp := container.Get("TEST")
    if resp != expectedValue {
        t.Errorf("HALP")
    }
}
