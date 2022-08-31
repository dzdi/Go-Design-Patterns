package template14

import "testing"

func TestSms_Send(t *testing.T) {
	tel := NewTelecomSms()
	tel.send("test", 12345555)
}
