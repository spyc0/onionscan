package main

func ScanHTTP(onion string) bool {
	conn, err := Connect(onion + ":80")
	if err != nil {
		return false
	}
	defer conn.Close()
	/* should add code to check if
	 * HTTP functionality works
	 * normally as in it can send,
	 * and receive requests.
	*/
	return true
}
