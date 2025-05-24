package main

import (
	"encoding/binary"
)

type QuestionSection struct {
	Questions []DNSQuestion
}

func NewQuestionSection() *QuestionSection {
	return &QuestionSection{
		Questions: []DNSQuestion{*NewDNSQuestion()},
	}
}

func (questionSection *QuestionSection) ToBytes() []byte {
	var buf []byte

	for _, question := range questionSection.Questions {
		buf = append(buf, question.ToBytes()...)
	}

	return buf
}

type DNSQuestion struct {
	Name         []string
	QuestionType uint16
	Class        uint16
}

func NewDNSQuestion() *DNSQuestion {
	return &DNSQuestion{
		Name:         []string{"codecrafters", "io"},
		QuestionType: 1,
		Class:        1,
	}
}

func (question *DNSQuestion) ToBytes() []byte {
	var buf []byte

	for _, label := range question.Name {
		buf = append(buf, byte(len(label)))
		buf = append(buf, []byte(label)...)
	}
	buf = append(buf, 0)

	tmp := make([]byte, 4)
	binary.BigEndian.PutUint16(tmp[0:2], question.QuestionType)
	binary.BigEndian.PutUint16(tmp[2:4], question.Class)
	buf = append(buf, tmp...)

	return buf
}
