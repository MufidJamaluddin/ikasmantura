package utils

import (
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
)

func ToBase64UUID(uid UUID) (nid string)  {
	nid = base64.StdEncoding.EncodeToString(uid.Guid().Bytes())
	return
}

func FromBase64UUID(nid string) (uid UUID, err error)  {
	var (
		oid []byte
		id uuid.UUID
	)

	if oid, err = base64.StdEncoding.DecodeString(nid); err != nil {
		uid = UUID(uuid.Nil)
		return
	}

	if id, err = uuid.FromBytes(oid); err != nil {
		uid = UUID(uuid.Nil)
		return
	}

	uid = UUID(id)
	return
}
