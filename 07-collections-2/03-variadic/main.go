package main

func DebugLog(args ...string) []string {
	log := []string{"[DEBUG]"}
	log = append(log, args...)
	return log
}

func InfoLog(args ...string) []string {
	log := []string{"[INFO]"}
	log = append(log, args...)
	return log
}

func ErrorLog(args ...string) []string {
	return append([]string{"[ERROR]"}, args...)
}
