package models

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
	var (
		uid uuid.UUID
		oid [16]byte
	)

	uid = uuid.UUID(d)
	oid = OrderedUUIDv1.ToOrderedUuid(uid)

	return oid, nil
}

func (d UUID) Guid() (uid uuid.UUID) {
	uid = uuid.UUID(d)
	return
}