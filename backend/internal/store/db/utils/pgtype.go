package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func UUIDFrom(pgUUID pgtype.UUID) uuid.UUID {
	return uuid.UUID(pgUUID.Bytes)
}

func PGUUIDFromUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}
