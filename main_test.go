package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestHandleCommandArabic(t *testing.T) {
	assert.Equal(t, "20", handleCommand("10 + 10"))
	assert.Equal(t, "10", handleCommand("9 + 1"))
	assert.Equal(t, "2", handleCommand("10 / 5"))
	assert.Equal(t, "5", handleCommand("9 - 4"))
	assert.Equal(t, "2", handleCommand("9 / 4"))
	assert.Equal(t, "0", handleCommand("9 / 10"))
	assert.Equal(t, "36", handleCommand("9 * 4"))
	assert.Equal(t, "-9", handleCommand("1-10"))

}

func TestHandleCommandRoman(t *testing.T) {
	assert.Equal(t, "V", handleCommand("X / II"))
	assert.Equal(t, "VIII", handleCommand("X - II"))
	assert.Equal(t, "XX", handleCommand("X * II"))
	assert.Equal(t, "XIV", handleCommand("X + IV"))
	assert.Equal(t, "I", handleCommand("X / X"))
	assert.Equal(t, "I", handleCommand("III / II"))
	assert.Equal(t, "C", handleCommand("X * X"))
}

func TestHandleCommandPanicArabic(t *testing.T) {
	assert.Panics(t, func() { handleCommand("11+10") })
	assert.Panics(t, func() { handleCommand("1+1+2") })
	assert.Panics(t, func() { handleCommand("10 - 11") })
	assert.Panics(t, func() { handleCommand("-10+10") })
	assert.Panics(t, func() { handleCommand("0+10") })
	assert.Panics(t, func() { handleCommand("1&2") })
	assert.Panics(t, func() { handleCommand("1 @ 2") })
	assert.Panics(t, func() { handleCommand("1 + 2.2") })
	assert.Panics(t, func() { handleCommand("1 + 2 - 1") })
}

func TestHandleCommandPanicRoman(t *testing.T) {
	assert.Panics(t, func() { handleCommand("V-VI") })
	assert.Panics(t, func() { handleCommand("V/VI") })
	assert.Panics(t, func() { handleCommand("V a I") })
	assert.Panics(t, func() { handleCommand("I - I") })
	assert.Panics(t, func() { handleCommand("XX - I") })
	assert.Panics(t, func() { handleCommand("IIII - I") })
}
