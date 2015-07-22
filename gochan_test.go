package messenger

import (
	"log"
	"testing"
)

func TestRunDate(t *testing.T) {
	cmd_obj := New("date", []string{}, "", 10)
	out_chan, err := cmd_obj.Run()
	if err != nil {
		log.Println(err)
	}
	<-out_chan
}

func TestMultiLineResponse(t *testing.T) {
	cmd_obj := New("echo", []string{"this\nis\na\nmultiline\ntest"}, "", 10)
	out_chan, err := cmd_obj.Run()
	if err != nil {
		log.Println(err)
	}
	expecting := []string{"this", "is", "a", "multiline", "test"}
	for i := 0; i < 5; i++ {
		var m string
		m = <-out_chan
		if m != expecting[i] {
			t.Error("Expected ", expecting[i], " got ", m)
		}
	}
}
