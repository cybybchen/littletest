package io

import (
	"io"
	"encoding/binary"
	"errors"
	"hash/adler32"
)

var RPC_MAGIC = [4]byte{'p', 'y', 'x', 'i'}

func DecodePacket(r io.Reader) ([]byte, error) {
	var totalsize uint32
	err := binary.Read(r, binary.BigEndian, &totalsize)
	if err != nil {
		return nil, errors.Annotate(err, "read total size")
	}

	// at least len(magic) + len(checksum)
	if totalsize < 8 {
		return nil, errors.Errorf("bad packet. header:%d", totalsize)
	}

	sum := adler32.New()
	rr := io.TeeReader(r, sum)

	var magic [4]byte
	err = binary.Read(rr, binary.BigEndian, &magic)
	if err != nil {
		return nil, errors.Annotate(err, "read magic")
	}
	if magic != RPC_MAGIC {
		return nil, errors.Errorf("bad rpc magic:%v", magic)
	}

	payload := make([]byte, totalsize-8)
	_, err = io.ReadFull(rr, payload)
	if err != nil {
		return nil, errors.Annotate(err, "read payload")
	}

	var checksum uint32
	err = binary.Read(r, binary.BigEndian, &checksum)
	if err != nil {
		return nil, errors.Annotate(err, "read checksum")
	}

	if checksum != sum.Sum32() {
		return nil, errors.Errorf("checkSum error, %d(calc) %d(remote)", sum.Sum32(), checksum)
	}
	return payload, nil
}
