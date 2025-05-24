package main

import "encoding/binary"

type DNSHeader struct {
	Id                    uint16
	QRIndicator           bool
	Opcode                uint16
	AuthoritativeAnswer   bool
	Truncation            bool
	RecursionDesired      bool
	RecursionAvailable    bool
	DNSSecQueries         uint16
	ResponseCode          uint16
	QuestionCount         uint16
	AnswerRecordCount     uint16
	AuthorityRecordCount  uint16
	AdditionalRecordCount uint16
}

func NewDNSHeader() *DNSHeader {
	return &DNSHeader{
		Id:                    1234,
		QRIndicator:           true,
		Opcode:                0,
		AuthoritativeAnswer:   false,
		Truncation:            false,
		RecursionDesired:      false,
		RecursionAvailable:    false,
		DNSSecQueries:         0,
		ResponseCode:          0,
		QuestionCount:         0,
		AnswerRecordCount:     0,
		AuthorityRecordCount:  0,
		AdditionalRecordCount: 0,
	}
}

func (header *DNSHeader) ToBytes() []byte {
	buf := make([]byte, 12)

	binary.BigEndian.PutUint16(buf[0:2], header.Id)

	var flags uint16
	if header.QRIndicator {
		flags |= 1 << 15
	}

	flags |= (header.Opcode & 0xF) << 11
	if header.AuthoritativeAnswer {
		flags |= 1 << 10
	}
	if header.Truncation {
		flags |= 1 << 9
	}
	if header.RecursionDesired {
		flags |= 1 << 8
	}
	if header.RecursionAvailable {
		flags |= 1 << 7
	}
	flags |= header.DNSSecQueries & 0x7
	flags |= header.ResponseCode & 0xF

	binary.BigEndian.PutUint16(buf[2:4], flags)
	binary.BigEndian.PutUint16(buf[4:6], header.QuestionCount)
	binary.BigEndian.PutUint16(buf[6:8], header.AnswerRecordCount)
	binary.BigEndian.PutUint16(buf[8:10], header.AuthorityRecordCount)
	binary.BigEndian.PutUint16(buf[10:12], header.AdditionalRecordCount)

	return buf
}
