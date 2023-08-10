package data

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func ToSQLNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func FromSQLNullString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func ToSQLNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true,
	}
}

func FromSQLNullBool(b sql.NullBool) *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

func ToSQLNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}

func FromSQLNullTime(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func ToSQLNullInt16(i *int16) sql.NullInt16 {
	if i == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{
		Int16: *i,
		Valid: true,
	}
}

func FromSQLNullInt16(i sql.NullInt16) *int16 {
	if !i.Valid {
		return nil
	}
	return &i.Int16
}

func ToSQLNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *i,
		Valid: true,
	}
}

func FromSQLNullInt32(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}
	return &i.Int32
}

func ToSQLNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

func FromSQLNullInt64(i sql.NullInt64) *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

func ToSQLNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

func FromSQLNullFloat64(f sql.NullFloat64) *float64 {
	if !f.Valid {
		return nil
	}
	return &f.Float64
}

func ToSQLNullByte(b *byte) sql.NullByte {
	if b == nil {
		return sql.NullByte{}
	}
	return sql.NullByte{
		Byte:  *b,
		Valid: true,
	}
}

func FromSQLNullByte(b sql.NullByte) *byte {
	if !b.Valid {
		return nil
	}
	return &b.Byte
}

func ToNullDecimal(d *decimal.Decimal) decimal.NullDecimal {
	if d == nil {
		return decimal.NullDecimal{}
	}
	return decimal.NullDecimal{
		Decimal: *d,
		Valid:   true,
	}

}

func FromNullDecimal(d decimal.NullDecimal) *decimal.Decimal {
	if !d.Valid {
		return nil
	}
	return &d.Decimal
}

func ToNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u == nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{
		UUID:  *u,
		Valid: true,
	}

}

func FromNullUUID(u uuid.NullUUID) *uuid.UUID {
	if !u.Valid {
		return nil
	}
	return &u.UUID
}
