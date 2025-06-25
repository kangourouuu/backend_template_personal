# Pull:
git checkout main
git checkout "your branch"
git merge main

# Push:
commit and push to your branch 
git checkout main
git merge "your branch"

go install github.com/swaggo/swag/cmd/swag@latest
swag init create Docs
go get github.com/swaggo/swag@v1.16.4 upgrate swagger for fix bug LeftDelim, RightDelim
