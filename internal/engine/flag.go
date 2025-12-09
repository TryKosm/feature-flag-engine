package engine

type Flag struct {
	Key            string
	RolloutPercent int
}

func Enabled(flag Flag, userID string) bool {
	return Bucket(userID, flag.RolloutPercent)
}
