# Go TODO app

Borrowing https://github.com/TKSpectro/go-todo-api

### Getting started

1. Create a `.env` file in the root of the project by copying the `.env.example` file and filling in the correct values

```bash
cp .env.example .env
```

2. Run the migrations (not necessary for in-memory SQlite)

```bash
make migrate-up
```

3. Generate the [TEMPL](https://templ.guide/) files

```bash
 go get -tool github.com/a-h/templ/cmd/templ@latest
 go tool templ generate
```

-> See [this](https://github.com/TKSpectro/go-todo-api/commit/d1f6669f91de0297d28bc0321b616a922e640957)

4. Run the server

```bash
make run
# or (if you have Air installed)
air
```

## Knowledge base

### TEMPL

[TEMPL](https://templ.guide/integrations/)

### JSON parsing with go

Because the BodyParser will default parse null string fields as empty strings, we need a better solution to get actual null values
With the omitempty tag, we also won't get the desired result, because it will omit the field if it's null (and then use the default value - empty string)

The solution is to use the packages `"gopkg.in/guregu/null.v4/zero"` and `"gopkg.in/guregu/null.v4/null"`
And then use the types `zero.String` and `null.String` instead of `string`

Because these will use a struct under the hood, we also want to overwrite the swagger documentation for these fields, so that it will show the correct type in the docs with `swaggertype:"string"` tag

### Migrations

#### SQlite

In-memory, done at runtime

#### MySQL

We can generate the schema migration code with [Atlas](https://atlasgo.io/). GORM does this directly from the models defined in this package. See [`loader/atlasGorm.go`](loader/atlasGorm.go) and [`atlas.hcl`](atlas.hcl) for details.

Atlas installation:

```bash
curl -sSf https://atlasgo.sh | sh
```

##### Generate a new migration file based on the current models

```bash
make migrate-gen name=<migration-name>
```
##### Generate a new empty migration file

```bash
make migrate-new name=<migration-name>
```

##### Apply all migrations up to the latest version

```bash
make migrate-up
```

##### Reverse all migrations down to the given version (version is the timestamp of the migration file)

```bash
make migrate-down version=<version>
```

### Prerequisites

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/) (optional) - for running the database
- [Make](https://www.gnu.org/software/make/) (optional) - for running the Makefile commands (shortcuts for other commands)
- [Air](https://github.com/cosmtrek/air/) (optional) - for hot reloading while developing
- [Ginkgo](https://onsi.github.io/ginkgo/) (optional) - for running the tests
