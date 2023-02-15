
## Small and simple noSQL JSON database for small apps and easy use.
<br>

## Getting Started
>go get wagemaker.no-ip.co.uk:3000/ricardo/gojsondb

<BR>

# HOW-TO-USE

<BR>

**CreateDB - This function will create a database file if it does not exist and return true | false**
>function: CreateDB()
```
Return:
true | false
```

**DeleteDB - WARNING - this function delete the database file return true | false**
>function: DeleteDB()
```
Return:
true | false
```
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
>function: CreatedDBCopy()

```
Return:
PATH: /Users/ricardowagemaker/gojsontmp.db
true | false
```

***EmptyDB function - WARNING - this will destroy the database and all data stored in it!***

>function: EmptyDB() 

```
Return:
true | false
```

***AddField() function take in ((key)string, (Value) string) and add to database.***

>function: AddField("New string Entry", "string of data representing a the value")

```
Return:
true | false
```

***CountSize() function will return number of rows in the gojsondb.db***

>function: CountSize()

````
Result:
44
````

***LastField() func returns the last entry in multi-format [ raw | json | id | key | value ]***

>function: LastField("raw")

```
Result:
{"id":44,"key":"New Entry","data":"string of data representing a the value"}
```

>function: LastField("id")

```
Result:
44
```

>function: LastField("key")

```
Result:
New Entry
```

>function: LastField("value")

```
Result:
string of data representing a the value
```

>function: LastField("json")

```
Result:
{
        "id": 44,
        "key": "New Entry",
        "data": "string of data representing a the value"
}
```
***FirstField func returns the first entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required***

>function: FirstField("raw")

```
Result:
{"id":1,"key":"one","data":"string data test"}
```

***SelectByID func returns an entry string for a specific id in all formats [ raw | json | id | key | value ]***
>function: SelectByID(10, "raw")

```
Result:
{"id":10,"key":"Roman","data":"string of data representing a the value of X"}
```

>function: SelectByID(10, "value")

```
Result:
string of data representing a the value of X
```

***UniqueID function returns an int for the last used UniqueID***
>function: UniqueID()

```
Result:
54
```

**LastXFields returns last X number of entries from db in byte[] format**
>function: LastXFields(2)

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
>function: FirstXFields(2)

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

**RemoveField function takes an unique field id as an input and remove the matching field entry**
>function: RemoveField(2)

```
Example:

// Delete the field Id 2 (this can not be undone)
fmt.Println("Return: ", RemoveField(2))
fmt.Println()

// Check results by printing the first 2 rows in database
var received = FirstXFields(2)
bytes := received
var data []MyStruct
json.Unmarshal(bytes, &data)

for l := range data {
        fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
        fmt.Println()
}

Result:

Return:  true

id: 1, key: New Entry, data: string of data representing a the value
id: 3, key: New Entry, data: string of data representing a the value

```

# STILL TO DO

```
NEW
* ModifyField - Takes an id, Key & Value (all 3 fields) and update with information provided for key & value

CHECK
** Write some test functions
** Review functions code on empty database reaction
** Review functions code on number of rows less than available in database
```



