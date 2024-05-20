package transport_utils

import (
	"encoding/binary"
	"errors"
	"io"
	"minerrpc/rpc_common/entities"
	"minerrpc/rpc_core/serializer"
	"net"
)

type ObjectReader struct {
	conn net.Conn
}

func NewObjectReader(conn net.Conn) *ObjectReader {
	return &ObjectReader{conn}
}

func (or *ObjectReader) ReadObject() (*entities.RPCdata, error) {
	magNumByte := make([]byte, 4)
	serCodeByte := make([]byte, 4)
	dataLenByte := make([]byte, 4)

	_, err := io.ReadFull(or.conn, magNumByte)
	if err != nil {
		return nil, err
	}
	magNum := binary.BigEndian.Uint32(magNumByte)
	if magNum != MagicNumber {
		return nil, errors.New("unknown protocol")
	}
	_, err = io.ReadFull(or.conn, serCodeByte)
	if err != nil {
		return nil, err
	}
	serCode := binary.BigEndian.Uint32(serCodeByte)
	ser := serializer.GetByCode(int(serCode))
	if ser == nil {
		return nil, errors.New("unknown serializer")
	}
	_, err = io.ReadFull(or.conn, dataLenByte)
	if err != nil {
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(dataLenByte)
	dataByte := make([]byte, dataLen)
	_, err = io.ReadFull(or.conn, dataByte)
	if err != nil {
		return nil, err
	}
	var data entities.RPCdata
	err = ser.Deserialize(dataByte, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
