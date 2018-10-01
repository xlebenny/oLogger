# Optional Logger
Use to optional log some struct field

## Installation
go get github.com/gitlab/xlebenny/optionalLogger

## Usage
````go
type structB struct {
    field1 string `oLog:"2"`
}
type structA struct {
    fieldA string  `oLog:"2"`
    fieldB structB `oLog:"2"`
    fieldC string  ``
}
obj := structA{
    fieldA: "valueA",
    fieldB: structB{field1: "valueB"},
    fieldC: "valueC", // because haven't `oLog` tag, this haven't log
}
fmt.Println(oLogger.Log(4, "  ", obj))

// {
//   "structA": {
//     "fieldA": "valueA",
//     "fieldB": {
//       "field1": "valueB"
//     }
//   }
// }
````

## Methods
### optionalLogger.Log
Parameter Name | Description
--- | ---
logLevel | Log when `LogLevel >= struct oLog`
indentString | Direct pass to json.MarshalIndent / json.Marshal
obj | Your object

## TODO
Haven't test Pointer