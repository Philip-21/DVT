package transform

func ErrTime() string {
	errt :=
		`
	Ensure you have the right time stamp format [GMT, WAT, CET, EET, CST]
	XXXX-XX-XXTXX:XX:XX+00:00 GMT,
	XXXX-XX-XXTXX:XX:XX+01:00 WAT,
	XXXX-XX-XXTXX:XX:XX+01:00 CET,
	XXXX-XX-XXTXX:XX:XX+02:00 EET,
	XXXX-XX-XXTXX:XX:XX-06:00 CST,
	`
	return errt
}

var Errtime = ErrTime()
