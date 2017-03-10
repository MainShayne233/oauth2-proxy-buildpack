// +build heroku

package main

// Heroku Logplex already adds proper timestamps:
// so we don't need to emit date or time.
const logDateAndTime = false
