package protocol

type Protocol int

const (
	UNIX Protocol = iota
	TCP
)

func (ct Protocol) String() string {
	switch ct {
	case UNIX:
		return "unix"
	case TCP:
		return "tcp"
	default:
		return "Unknown"
	}
}