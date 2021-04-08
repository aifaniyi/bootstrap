# bootstrap

A golang web application bootstraping framework.

## Installation
```bash
go get github.com/aifaniyi/bootstrap
```

## Usage

The bootstrap command creates a new web application template and generates models, database repos (postgres) and controllers to allow CRUD api calls.

```bash
# linux, macos
$ bootstrap new -i schema.spec.json -p newproject -o $GOPATH/src/dir
cd $GOPATH/src/dir/newproject
go mod vendor
```

It accepts a schema.spec.json file of the form
```json
{
    "lang": "golang",
    "entities": [
        {
            "name": "Address",
            "description": "address model",
            "relations": [],
            "properties": [{
                "name": "number",
                "type": {
                    "name": "integer"
                },
                "width": 10,
                "nullable": true,
                "dto": true,
                "description": "house number on street",
                "unique": false
            }, {
                "name": "line1",
                "type": {
                    "name": "string"
                },
                "width": 255,
                "nullable": true,
                "dto": true,
                "description": "street name, postalcode or other info",
                "unique": false
            }]
        }
    ]
}
```

## .spec.json Documentation
TODO