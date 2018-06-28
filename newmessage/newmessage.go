package newmessage

// NewMessage -
type NewMessage struct {
	Name    string
	Message string
}

// NewMess -
func (m NewMessage) NewMess() string {
	return m.Message

}
