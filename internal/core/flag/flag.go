package flag

import "flag"

var (
	isParsed = false
	flagProduction = flag.Bool("production", false, "production mode")
	flagWatch = flag.Bool("watch", false, "watch mode")
	flagDebug = flag.Bool("debug", false, "debug mode")
)

func parse() {
	if !isParsed {
		isParsed = true
		flag.Parse()
	}
}

func IsProduction() bool {
	parse()
	return *flagProduction
}

func IsDevelopment() bool {
	return !IsProduction()
}

func IsWatch() bool {
	parse()
	return *flagWatch
}

func IsDebug() bool {
	parse()
	return *flagDebug
}
