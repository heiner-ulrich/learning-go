
# install go
brew install go

# go env
reload:
  source ~/.zshrc


# install simpleforce
go install github.com/simpleforce/simpleforce@latest


# create mod file
go mod init sf_query_simpleforce.go


