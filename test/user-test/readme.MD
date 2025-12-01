# CREATE USER
go run user-test/main.go create 123456 John Doe teacher

# UPDATE USER
go run user-test/main.go update 1 Jane Doe admin

# GET USER BY ID
go run user-test/main.go get 1

# LIST USERS
go run user-test/main.go list

# GET USER BY TELEGRAM ID
go run user-test/main.go getbytg 123456

# GET TELEGRAM IDS BY USER IDS
go run user-test/main.go tgids 1 2 3
