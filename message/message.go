package message

type Kind int

const (
	KindUser Kind = iota
	KindServer
)

type Message struct {
	SenderName string
	Content    string
	Kind       Kind
}
