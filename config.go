package main

// Env environment variable structure
type Env struct {
	DBHost     string `env:"DBHOST" envDefault:"127.0.0.1:3306"`
	DBUser     string `env:"DBUSER" envDefault:"test"`
	DBPass     string `env:"DBPASS" envDefault:"test"`
	DBName     string `env:"DBNAME" envDefault:"test"`
	DBSocket   string `env:"DBSOCKET" envDefault:"/var/mysql/mysql.sock"`
	DBConnType int64  `env:"DBCONNTYPE" envDefault:"1"`
	Port       string `env:"PORT" envDefault:"8080"`
	AppSocket  string `env:"APPSOCKET" envDefault:"/tmp/app.sock"`
	ListenIP   string `env:"LISTENIP" envDefault:"127.0.0.1"`
	ListenType int64  `env:"LISTENTYPE" envDefault:"1"`
	Production bool   `env:"PRODUCTION" envDefault:"false"`
	SMTPUser   string `env:"SMTPUSER" envDefault:""`
	SMTPPass   string `env:"SMTPPASS" envDefault:""`
	SMTPServer string `env:"SMTPSERVER" envDefault:"127.0.0.1:25"`
	UploadDir  string `env:"UPLOADDIR" envDefault:""`
	Secret     string `env:"SECRET" envDefault:""`
}

const siteTitle = "My Website Name"

// Location to use for dates
const loc = "Australia/Sydney"
