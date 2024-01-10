# .envの内容をmake変数として読み込む
include .env

PRISMA_SCHEMA = blog-schema/prisma/schema.prisma

# .envのDATABASE_URLはdockerコンテナ内のやり取りで使うのでhostnameがdbになっている
# こっちはコンテナ外なのでlocalhost
DATABASE_URL = postgresql://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)

.PHONY: gen

gen:
	goa gen github.com/matumoto1234/blog-backend/app/design
	chmod -R 755 gen/

swagger-ui:
	docker compose up swagger-ui

up:
	docker compose down \
	&& docker compose up db -d \
	&& DATABASE_URL=$(DATABASE_URL) pnpm dlx prisma migrate deploy --schema=$(PRISMA_SCHEMA) \
	&& docker compose up server
