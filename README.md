
## Small and simple noSQL JSON database for small apps and easy use.
<br>

## Getting Started
>go get wagemaker.no-ip.co.uk:3000/ricardo/gojsondb

<br>
# HOW-TO-USE

***AddField() function take in ((key)string, (Value) string) and add to database.***

>AddField("New string Entry", "string of data representing a the value")

```
Result:
true
```

***CountSize() function will return number of rows in the gojsondb.db***

>CountSize()

````
Result:
44
````

***LastField() func returns the last entry in multi-format [ raw | json | id | key | value ]***

>LastField("raw")

```
Result:
{"id":44,"key":"New Entry","data":"string of data representing a the value"}
```

>LastField("id")

```
Result:
44
```

>LastField("key")

```
Result:
New Entry
```

>LastField("value")

```
Result:
string of data is the value
```

>LastField("json")

```
Result:
{
        "id": 44,
        "key": "New Entry",
        "data": "string of data representing a the value"
}
```
***FirstField func returns the first entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required***

>FirstField("raw")

```
Result:
{"id":1,"key":"one","data":"string data test"}
```

***EmptyDB function - WARNING - this will destroy all data stored in gojsondb.db!***

> EmptyDB() 

```
Result:
true
```

***SelectByID func returns an entry string for a specific id in all formats [ raw | json | id | key | value ]***
>SelectByID(10, "raw")

```
Result:
{"id":10,"key":"Roman","data":"string of data representing a the value of X"}
```

>SelectByID(10, "value")

```
Result:
string data test X
```

***UniqueID function returns an int for the last used UniqueID***
> UniqueID()

```
Result:
54
```

**LastXFields returns last X number of entries from db in byte[] format**
> LastXFields(2)

```
Example:

var received = LastXFields(2)

bytes := received
var data []MyStruct
json.Unmarshal(bytes, &data)

for l := range data {
        fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
        fmt.Println()
}

Result:

id: 43, key: New Entry, data: string of data representing a the value
id: 44, key: New Entry, data: string of data representing a the value
```

**FirstXFields returns last X number of entries from db in byte[] format**
> FirstXFields(2)

```
Example:

var received = FirstXFields(2)

bytes := received
var data []MyStruct
json.Unmarshal(bytes, &data)

for l := range data {
        fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
        fmt.Println()
}

Result:

id: 1, key: New Entry, data: string of data representing a the value
id: 2, key: New Entry, data: string of data representing a the value
```