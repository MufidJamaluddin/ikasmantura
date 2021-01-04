package utils

import (
	"database/sql/driver"
	"fmt"
	OrderedUUIDv1 "github.com/MufidJamaluddin/uuidv1_orderer"
	uuid "github.com/satori/go.uuid"
)

type UUID uuid.UUID

func (d *UUID) Scan(value interface{}) (err error) {
	var (
		uid uuid.UUID
		oid []byte
		nid [16]byte
	)

	oid = value.([]byte)

	if len(oid) == 0 {
		*d = UUID(uuid.Nil)
	}

	if len(oid) != 16 {
		err = fmt.Errorf("uuid: UUID must be exactly 16 bytes long, got %d bytes", len(oid))
		return
	}

	copy(nid[:], oid)

	uid = OrderedUUIDv1.FromOrderedUuid(nid)
	*d = UUID(uid)

	return
}

func (d UUID) Value() (driver.Value, error) {
	var oid [16]byte
	oid = d.OrderedValue()
	return oid, nil
}

func (d UUID) OrderedValue() uuid.UUID {
	var (
		uid uuid.UUID
		oid [16]byte
	)

	uid = uuid.UUID(d)

	if uid == uuid.Nil {
		return uuid.Nil
	}

	oid = OrderedUUIDv1.ToOrderedUuid(uid)
	return oid
}

func (d UUID) Guid() (uid uuid.UUID) {
	uid = uuid.UUID(d)
	return
}