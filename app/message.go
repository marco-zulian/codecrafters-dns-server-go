package main

type DNSMessage struct {
	Header          DNSHeader
	QuestionSection QuestionSection
	AnswerSection   AnswerSection
}

func NewDNSMessage() *DNSMessage {
	return &DNSMessage{
		Header:          *NewDNSHeader(),
		QuestionSection: *NewQuestionSection(),
		AnswerSection:   *NewAnswerSection(),
	}
}

func (message *DNSMessage) ToBytes() []byte {
	var buf []byte

	message.Header.QuestionCount += uint16(len(message.QuestionSection.Questions))
	message.Header.AnswerRecordCount += uint16(len(message.AnswerSection.Records))
	buf = append(buf, message.Header.ToBytes()...)
	buf = append(buf, message.QuestionSection.ToBytes()...)
	buf = append(buf, message.AnswerSection.ToBytes()...)

	return buf
}
