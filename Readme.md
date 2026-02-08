
### Install Database
- Run the docker compose file: `docker compose up`

### Database Connection
*Reference* - https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- Install Dependencies
  - GORM - `go get -u gorm.io/gorm`
  - Driver for postgres - `go get -u gorm.io/driver/postgres`

### Fix Go sum Issue
- Run `go mod tidy` before `go run main.go`



