package constants

// SecretConstants are constants which will be used for generating random string
var SecretConstants = []string{
	"01234567",                   // include Octals (weakest)
	"89abcdef",                   // include Hexadecimals (weaker)
	"qwrtyuiopsghjklzxvnm",       // include rest lower case alphabets (weak)
	"QWERTYUIOPASDFGHJKLZXCVBNM", // ^^ include all upper case alphabets (still weak)
	"~!@#$%^&*()",                // include some symbols (seems strong)
}

// Version string to store the command version
var Version = "v0.8.0-alpha"

// BuildDate string to store the built date
var BuildDate = "2019-05-18 13:11:28 IST"
