package enums

type Status int64

const (
	Open Status = 0
	TransitionPending = 1
	Processed = 2
	TransitionWithoutProcessed = 3
	Closed = 4
)

func (status Status) Name() string {
	switch status {
	case Open:
		return "OPEN"
	case Processed:
		return "PROCESSED"
	case TransitionPending:
		return "TRANSITION PENDING"
	case TransitionWithoutProcessed:
		return "TRANSITION WITHOUT PROCESSED"
	case Closed:
		return "CLOSED"
	}
	return "UNKNOWN"
}

func (status Status) Value() int {
 return int(status)
}
