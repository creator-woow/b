package shared

import "strconv"

type ID uint64

func ParseID(s string) (ID, error) {
	id, parseErr := strconv.ParseUint(s, 10, 32)
	return ID(id), parseErr
}
