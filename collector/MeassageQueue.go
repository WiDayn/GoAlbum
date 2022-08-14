package collector

type DefaultRequestMessage struct {
	SourceType string
	Id         string
}

type MessageQueue struct {
	data []DefaultRequestMessage
}

var DefaultMessageQueue MessageQueue

func (q *MessageQueue) Add(k DefaultRequestMessage) {
	q.data = append(q.data, k)
}

func (q *MessageQueue) Pop() DefaultRequestMessage {
	if len(q.data) == 0 {
		return DefaultRequestMessage{}
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *MessageQueue) Top() DefaultRequestMessage {
	if len(q.data) == 0 {
		return DefaultRequestMessage{}
	}
	return q.data[0]
}

func (q *MessageQueue) Size() int {
	return len(q.data)
}

func (q *MessageQueue) Contains(data DefaultRequestMessage) bool {
	for _, message := range q.data {
		if message.SourceType == data.SourceType && message.Id == data.Id {
			return true
		}
	}
	return false
}
