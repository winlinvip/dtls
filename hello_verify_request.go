package dtls

import (
	"bytes"
	"fmt"
)

type handshakeHelloVerifyRequest struct {
	ServerVersion protocolVersion
	Cookie        []byte
}

func (hvr handshakeHelloVerifyRequest) Bytes() []byte {
	b := make([]byte, 0, len(hvr.Cookie)+3)
	b = append(b, hvr.ServerVersion.Bytes()...)
	b = append(b, byte(len(hvr.Cookie)))
	b = append(b, hvr.Cookie...)
	return b
}

func (hvr handshakeHelloVerifyRequest) String() string {
	return fmt.Sprintf("HelloVerifyRequest{ ServerVersion: %s, Cookie: %x }", hvr.ServerVersion, hvr.Cookie)
}

func readHandshakeHelloVerifyRequest(byts []byte) (hvr handshakeHelloVerifyRequest, err error) {
	buffer := bytes.NewBuffer(byts)
	if buffer.Len() < 3 {
		return hvr, InvalidHandshakeError
	}
	if hvr.ServerVersion, err = readProtocolVersion(buffer); err != nil {
		return
	}
	cookieLength, err := buffer.ReadByte()
	if err != nil {
		return
	}
	if buffer.Len() < int(cookieLength) {
		return hvr, InvalidHandshakeError
	}
	hvr.Cookie = buffer.Next(int(cookieLength))
	return
}
