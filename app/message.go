package main

type DNSMessage struct {
	Header          DNSHeader
	QuestionSection QuestionSection
}

func NewDNSMessage() *DNSMessage {
	return &DNSMessage{
		Header: *NewDNSHeader(),
	}
}

func (message *DNSMessage) ToBytes() []byte {
	var buf []byte

	buf = append(buf, message.Header.ToBytes()...)
	buf = append(buf, message.QuestionSection.ToBytes()...)

	return buf
}
