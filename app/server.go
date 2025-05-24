package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type DNSServer struct {
	port                  int
	conn                  *net.UDPConn
	questionCount         uint16
	answerRecordCount     uint16
	authorityRecordCount  uint16
	additionalRecordCount uint16
}

func NewServer(port int) *DNSServer {
	return &DNSServer{
		port: port,
	}
}

func (s *DNSServer) ListenAndServe() {
	conn, err := s.listen()
	if err != nil {
		panic("Could not start server")
	}
	defer conn.Close()

	s.conn = conn
	s.serve()
}

func (s *DNSServer) listen() (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", s.port))
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return nil, err
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return nil, err
	}

	return udpConn, nil
}

func (s *DNSServer) serve() {
	buf := make([]byte, 512)
	for {
		size, source, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		go s.handleMessage(receivedData, source)
	}
}

func (s *DNSServer) handleMessage(data string, source *net.UDPAddr) {
	header := s.parseHeader([]byte(data[:12]))
	response := NewDNSMessage(*header)

	_, err := s.conn.WriteToUDP(response.ToBytes(), source)
	if err != nil {
		fmt.Println("Failed to send response:", err)
	}
}

func (s *DNSServer) parseHeader(headerBytes []byte) *DNSHeader {
	flags := binary.BigEndian.Uint16(headerBytes[2:4])
	opcode := uint16((flags & 0x7800)) >> 11

	responseCode := 0
	if opcode != 0 {
		responseCode = 4
	}

	opts := []Option{
		WithId(binary.BigEndian.Uint16(headerBytes[:2])),
		WithOpcode(opcode),
		WithResponseCode(uint16(responseCode)),
		WithAnswerRecordCount(s.answerRecordCount + 1),
		WithQuestionCount(s.questionCount + 1),
	}

	if (flags & 0x0100) != 0 {
		opts = append(opts, WithRecursionDesired())
	}

	return NewDNSHeader(opts...)
}
