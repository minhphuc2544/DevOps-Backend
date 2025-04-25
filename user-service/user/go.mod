module github.com/minhphuc2544/DevOps-Backend/user-service/user

go 1.24.1

require github.com/julienschmidt/httprouter v1.3.0

require (
	github.com/go-sql-driver/mysql v1.9.2
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.37.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.24.0 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
