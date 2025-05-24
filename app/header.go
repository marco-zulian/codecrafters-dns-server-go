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

type Option func(*DNSHeader)

func NewDNSHeader(opts ...Option) *DNSHeader {
	header := &DNSHeader{
		QRIndicator: true,
	}

	for _, opt := range opts {
		opt(header)
	}

	return header
}

func WithId(id uint16) Option {
	return func(d *DNSHeader) {
		d.Id = id
	}
}

func WithOpcode(code uint16) Option {
	return func(d *DNSHeader) {
		d.Opcode = code
	}
}

func AsAuthoritativeAnswer() Option {
	return func(d *DNSHeader) {
		d.AuthoritativeAnswer = true
	}
}

func Truncated() Option {
	return func(d *DNSHeader) {
		d.Truncation = true
	}
}

func WithRecursionDesired() Option {
	return func(d *DNSHeader) {
		d.RecursionDesired = true
	}
}

func WithRecursionAvailable() Option {
	return func(d *DNSHeader) {
		d.RecursionAvailable = true
	}
}

func WithResponseCode(code uint16) Option {
	return func(d *DNSHeader) {
		d.ResponseCode = code
	}
}

func WithQuestionCount(count uint16) Option {
	return func(d *DNSHeader) {
		d.QuestionCount = count
	}
}

func WithAnswerRecordCount(count uint16) Option {
	return func(d *DNSHeader) {
		d.AnswerRecordCount = count
	}
}

func WithAuthorityRecordCount(count uint16) Option {
	return func(d *DNSHeader) {
		d.AuthorityRecordCount = count
	}
}

func WithAdditionalRecordCount(count uint16) Option {
	return func(d *DNSHeader) {
		d.AdditionalRecordCount = count
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
