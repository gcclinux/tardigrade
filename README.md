
## Small and simple noSQL JSON database for small apps and easy use.
<br>

## Getting Started
>go get wagemaker.no-ip.co.uk:3000/ricardo/gojsondb

<BR>

# HOW-TO-USE

<BR>

***AddField() function take in ((key)string, (Value) string) and add to database.***

>AddField("New string Entry", "string of data representing a the value")

```
Return:
true | false
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
string of data representing a the value
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

***SelectByID func returns an entry string for a specific id in all formats [ raw | json | id | key | value ]***
>SelectByID(10, "raw")

```
Result:
{"id":10,"key":"Roman","data":"string of data representing a the value of X"}
```

>SelectByID(10, "value")

```
Result:
string of data representing a the value of X
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

**CreateDB - This function will create a database file if it does not exist and return true | false**
>CreateDB()
```
Return:
true | false
```

**DeleteDB - WARNING - this function delete the database file return true | false**
>DeleteDB()
```
Return:
true | false
```
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
> CreatedDBCopy()

```
Return:
PATH: /Users/ricardowagemaker/gojsontmp.db
true | false
```

***EmptyDB function - WARNING - this will destroy the database and all data stored in it!***

> EmptyDB() 

```
Return:
true | false
```

# STILL TO DO

```
NEW
* RemoveField - WARNING: takes an id as an input and remove add matching the unique id
* ModifyField - Takes ID, Key, Value (all 3 fields) and update with information provided in k & V

CHECK
** Write some test functions
** Review functions code on empty database reaction
** Review functions code on number of rows less than available in database
```



