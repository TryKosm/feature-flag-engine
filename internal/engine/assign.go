package engine

import "hash/fnv"

func Bucket(userID string, rolloutPercent int) bool {
	h := fnv.New32a()
	_, _ = h.Write([]byte(userID))
	return int(h.Sum32()%100) < rolloutPercent
}
