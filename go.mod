module github.com/mvvershinin/http-shortener

go 1.22.2

require github.com/mvvershinin/http-shortener/internal/app/handler v0.0.0-00010101000000-000000000000
require github.com/mvvershinin/http-shortener/internal/app/strencoder v0.0.0-00010101000000-000000000000 // indirect

replace (
	github.com/mvvershinin/http-shortener/internal/app/handler => ./internal/app/handler
	github.com/mvvershinin/http-shortener/internal/app/strencoder => ./app/strencoder
)