<!-- # Pull:
git checkout main
git checkout "your branch"
git merge main -->

<!-- # Push:
commit and push to your branch 
git checkout main
git merge "your branch" -->

// install CLI for using swagger on terminal
go install github.com/swaggo/swag/cmd/swag@latest

// install dependency for using docs
go get github.com/swaggo/swag@latest

// Initialize docs
swag init create Docs