[![Build Status](https://travis-ci.org/knspriggs/messenger.svg?branch=master)](https://travis-ci.org/knspriggs/messenger)
##Messenger
This is a first pass of a library that returns a golang channel that contains a system call's output (line-by-line)
####Motivation
I wanted a system call implementation that would let me make the call then look at the output later.
####Use
Example use:
```go
cmd_obj := New("date", []string{}, "", 10)
out_chan, err := cmd_obj.Run()
if err != nil {
  log.Println(err)
}
<- out_chan //do what you want with the output
```
