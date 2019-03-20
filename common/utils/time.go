package utils

func HoursToUnixSeconds(hours int) int64 {
	return int64(hours) * 60 * 60
}
