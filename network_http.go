package main

func ScanHTTP(onion string) bool {
	conn, err := Connect(onion + ":80")
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
