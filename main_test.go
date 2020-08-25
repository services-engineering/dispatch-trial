package main

import (
    "testing"
)

func Test_World(t *testing.T) {
    actual := World()
    if actual != "World" {
        t.Fail()
    }
}
