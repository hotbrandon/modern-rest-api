# modern-rest-api

Modern_REST_API_Development_in_Go

## packages

go get -u github.com/golang-jwt/jwt/v5


## Issues during development

I forgot to exclude sqlite.db in .gitignore, and here's how to ask git to stop tracking it and remove it from Github.

```shell
git rm --cached sqlite.db && git commit -m "Remove sqlite.db from version control" && git push
```