package main

type DNSMessage struct {
	Header          DNSHeader
	QuestionSection QuestionSection
}

func NewDNSMessage() *DNSMessage {
	return &DNSMessage{
		Header:          *NewDNSHeader(),
		QuestionSection: *NewQuestionSection(),
	}
}

func (message *DNSMessage) ToBytes() []byte {
	var buf []byte

	message.Header.QuestionCount += uint16(len(message.QuestionSection.Questions))
	buf = append(buf, message.Header.ToBytes()...)
	buf = append(buf, message.QuestionSection.ToBytes()...)

	return buf
}
