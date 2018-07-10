package newmessage

// NewMessage -
type NewMessage struct {
	name    string
	message string
}

func (m *NewMessage) UpdateName(name string) {
	m.name = name
}

func (m *NewMessage) UpdateMessage(message string) {
	m.message = message
}

// NewMess -
func (m NewMessage) DisplayMsg() string {
	return m.message
}

// DisplayName
func (m NewMessage) DisplayName() string {
	return m.name
}
