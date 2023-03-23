# Table of Contents
- [Jaeger](#about-jaeger)
- [Utils](#about-utils-function)

# About Jaeger
Jaeger SDK is the code for start tracing the service's operation in distributed system

# Getting Start

## Configuration

```go
type JaegerConfig struct {
    Host        string `mapstructure:"host"`
    Environment string `mapstructure:"env"`
    ServiceName string `mapstructure:"service-name"`
}
```
| name              | description                                             | example                |
|-------------------|---------------------------------------------------------|------------------------|
| Host              | The host of the Jaeger in format `http://hostname:port` | http://localhost:14268 |
| Environment       | Environment of current service                          | local                  |
| ServiceName       | The name of service                                     | gateway                |


## Usage

### Initialize
Initialize by calling `gosdk.SetupTracer`

```go
if err := gosdk.SetupTracer(*JaegerConfig, tracerName){
    // handle error	
}
```

#### Parameters

| name         | description          | example |
|--------------|----------------------|---------|
| JaegerConfig | Jaeger Configuration |         |
| tracerName   | name of the tracer   | gateway |

### StartTracer
Start to trace

```go
newCtx, span := gosdk.StartTracer(tracerName, spanName, ctx, opt...)
if span != nil {
    defer span.End()
}
```

#### Parameters
| name       | description                                  | example               |
|------------|----------------------------------------------|-----------------------|
| tracerName | name of tracer                               | verify-ticket-handler |
| spanName   | name of span                                 | verify-ticket         |
| ctx        | context to pass to another service with span |                       |
| opt        | span start option (optional)                 |                       |

#### Returns
| name    | description                 | example |
|---------|-----------------------------|---------|
| newCtx  | context from tracer         |         |
| span    | interface of span in tracer |         |

# About Utils Function

Utils function is the function that can be use in a various scenario

# Getting Start

## Functions

### BoolAdr
the function use for convert `bool` to `*bool`

```go
boolPtr := gosdk.BoolAdr(boolean)
```

#### Parameters
| name    | description       | example |
|---------|-------------------|---------|
| boolean | the boolean value | true    |

#### Return

| name    | description         | example |
|---------|---------------------|---------|
| boolPtr | the boolean pointer |         |

### StringAdr
the function use for convert `string` to `*string`

```go
stringPtr := gosdk.StringAdr(string)
```

#### Parameters
| name   | description      | example       |
|--------|------------------|---------------|
| string | the string value | "hello world" |

#### Return

| name      | description        | example |
|-----------|--------------------|---------|
| stringPtr | the string pointer |         |

### IntAdr
the function use for convert `int` to `*int`

```go
intPtr := gosdk.IntAdr(int)
```

#### Parameters
| name | description   | example |
|------|---------------|---------|
| int  | the int value | 999     |

#### Return

| name   | description     | example |
|--------|-----------------|---------|
| intPtr | the int pointer |         |

### UUIDAdr
the function use for convert `uuid.UUID` to `*uuid.UUID`

```go
uuidPtr := gosdk.UUIDAdr(uuid)
```

#### Parameters
| name | description    | example |
|------|----------------|---------|
| uuid | the uuid value |         |

#### Return

| name    | description      | example |
|---------|------------------|---------|
| uuidPtr | the uuid pointer |         |

### GetCurrentTimePtr
the function use to get the current time pointer

```go
currentTimePtr := gosdk.GetCurrentTimePtr()
```

#### Return

| name           | description                          | example |
|----------------|--------------------------------------|---------|
| currentTimePtr | current time pointer in `*time.Time` |         |


### GetCurrentYear2Digit

the function use to get the 2 last digit of current year

```go
year2Digit := gosdk.GetCurrentYear2Digit()
```

#### Return

| name        | description                          | example |
|-------------|--------------------------------------|---------|
| year2Digit  | current year in 2 last digit `int`   | 66      |


### CalYearFromID
the function use to get the current year of student from the student id

```go
year, err := gosdk.CalYearFromID(studentID)
if err != nil{
	// handle error
}
```

#### Parameters
| name      | description    | example    |
|-----------|----------------|------------|
| studentID | the student id | 633xxxxx21 |

#### Return

| name  | description              | example |
|-------|--------------------------|---------|
| year  | the year from student id | 3       |

### IsExisted
Check is the variable existed in map

```go
if ok := gosdk.IsExisted(map, key); !ok {
	// handle error
}
```


#### Parameters
| name | description                           | example   |
|------|---------------------------------------|-----------|
| map  | the map structure                     |           |
| key  | the key of map that you want to check | "hello"   |

#### Return

| name | description        | example |
|------|--------------------|---------|
| ok   | boolean is existed | true    |

### MergeStringSlice
merge slices of string into one slice

```go
resultSlice := gosdk.MergeStringSlice(slice1, slice2, ...)
```

#### Parameters
| name  | description        | example   |
|-------|--------------------|-----------|
| slice | the `string` slice |           |

#### Return

| name         | description                         | example |
|--------------|-------------------------------------|---------|
| resultSlice  | the slice of string that was merged | true    |

### TrimInList
Trim the string if existed in list

```go
result := gosdk.TrimInList(word, sep, trimList)
```

#### Parameters
| name     | description                                               | example                       |
|----------|-----------------------------------------------------------|-------------------------------|
| word     | the input word                                            | /v1/path                      |
| sep      | the separate work                                         | /                             |
| trimList | the list of word that want to trim in `map[string]struct` | map[string]struct{}{"v1": {}} |

#### Return

| name    | description                    | example |
|---------|--------------------------------|---------|
| result  | result string that was trimmed | /path   |
