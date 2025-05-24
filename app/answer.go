package main

import "encoding/binary"

type AnswerSection struct {
	Records []ResourceRecord
}

func NewAnswerSection() *AnswerSection {
	return &AnswerSection{
		Records: []ResourceRecord{*NewResourceRecord()},
	}
}

func (answerSection *AnswerSection) ToBytes() []byte {
	var buf []byte

	for _, record := range answerSection.Records {
		buf = append(buf, record.ToBytes()...)
	}

	return buf
}

type ResourceRecord struct {
	Name   []string
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   []byte
}

func NewResourceRecord() *ResourceRecord {
	return &ResourceRecord{
		Name:   []string{"codecrafters", "io"},
		Type:   1,
		Class:  1,
		TTL:    60,
		Length: 4,
		Data:   []byte("8.8.8.8"),
	}
}

func (record *ResourceRecord) ToBytes() []byte {
	var buf []byte

	for _, name := range record.Name {
		buf = append(buf, byte(len(name)))
		buf = append(buf, []byte(name)...)
	}

	tmp := make([]byte, 12)
	binary.BigEndian.PutUint16(tmp[0:2], record.Type)
	binary.BigEndian.PutUint16(tmp[2:4], record.Class)
	binary.BigEndian.PutUint32(tmp[4:8], record.TTL)
	binary.BigEndian.PutUint32(tmp[8:12], uint32(record.Length))

	buf = append(buf, tmp...)
	buf = append(buf, record.Data...)

	return buf
}
