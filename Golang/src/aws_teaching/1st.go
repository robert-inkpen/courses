package main
import (
	"fmt"
	"testing"
)
// func (c *Client) ReceiveMessage(option1 string, option2 int) ([]Message, error)
type MessageReciever interface {
	ReceiveMessage(string, int) ([]Message, error)
}
type Message string
func getMessagesFromQueue(n int, receiver MessageReceiver) []Message {
	read, err := reciever.ReceiveMessage("queue1", n)
	return read
}
func main() {
	//cli := aws.CLient{}
	// output, err := cli.ReceiveMessage(isfiafno)
	//messages := getMessagesFromQueue(5, cli)
	fmt.Println("Hi")
}
type DummyQueue []Message
func (dq DummyQueue) ReceiveMessage(option1 string, option2 int) ([]Message, error) {
	return dq, nil
}
func TestGetMessagesFromQueue(t *testing.T) {
	cases := []struct{
		queue DummyQueue
		want []Message
	}{
		{
			queue: DummyQueue{"hi", "there"},
			want: []Message{"hi", "there"},
		},
		{
			queue: DummyQueue{"there"},
			want: []Message{"there"},
		},
	}
	for _, case := range cases {
		got, gotErr := getMessagesFromQueue("test", case.queue)
		if case.want != case.got {
			t.Errorf("Wanted: %v\n Got %v\n", case.want, got)
	}
}