package api

import "strconv"

func ConvertToID(id string) (uint, error) {
	uid, err := strconv.ParseUint(id, 0, 32)

	return uint(uid), err
}
