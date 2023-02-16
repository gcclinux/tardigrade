
## "tardigrade" is small and simple noSQL database for small apps and easy use.
*updated:  Wed 15 Feb 21:52:17 GMT 2023*<br>
*release:  0.0.3*

<br>

## Getting Started
>go get [https://github.com/gcclinux/tardigrade](https://github.com/gcclinux/tardigrade)

<BR>

Current structure and available functions()

```
type Tardigrade struct{}

func (*Tardigrade).AddField(key string, data string) bool
func (*Tardigrade).CountSize() int
func (*Tardigrade).CreateDB() (msg string, status bool)
func (*Tardigrade).CreatedDBCopy() (string, bool)
func (*Tardigrade).DeleteDB() (msg string, status bool)
func (*Tardigrade).EmptyDB() (msg string, status bool)
func (*Tardigrade).FirstField(f string) string
func (*Tardigrade).FirstXFields(count int) []byte
func (*Tardigrade).GetUpdated() (updated string)
func (*Tardigrade).GetVersion() (release string)
func (*Tardigrade).LastField(f string) string
func (*Tardigrade).LastXFields(count int) []byte
func (*Tardigrade).ModifyField(id int, k string, v string) bool
func (*Tardigrade).RemoveField(id int) (string, bool)
func (*Tardigrade).SelectByID(id int, f string) string
func (*Tardigrade).UniqueID() int
```

# HOW-TO-USE

<BR>

**CreateDB - This function will create a database file if it does not exist and return true | false**
>function: CreateDB()
```
Example 1: (ignore return)
	tar := Tardigrade{}
	tar.CreateDB()

Example 2 (capture return):
	tar := Tardigrade{}
	msg, status := tar.CreateDB()
	fmt.Println(msg, status)

Return:
	Created: <full_path>/tardigrade.db true
	Exist: <full_path>/tardigrade.db false

```

**DeleteDB - WARNING - this function delete the database file return true | false**
>function: DeleteDB()
```
Example 1: (ignore return)
	tar := Tardigrade{}
	tar.DeleteDB()

Example 2 (capture return):
	tar := Tardigrade{}
	msg, status := tar.DeleteDB()
	fmt.Println(msg, status)

Return:
	Removed: <full_path>/tardigrade.db true
	Unavailable: <full_path>/tardigrade.db false

```
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
>function: CreatedDBCopy()

```
Example 1: (ignore return)
	tar := Tardigrade{}
	tar.CreatedDBCopy()

Example 2 (capture return):
	tar := Tardigrade{}
	msg, status := tar.CreatedDBCopy()
	fmt.Println(msg, status)

Return:
	Copy: <full_path>/tardigradecopy.db true
	Failed: database tardigrade.db missing! false
	Failed: buffer error failed to create database! false
	Failed: permission error failed to create database! false

```

**EmptyDB function - WARNING - this will destroy the database and all data stored in it!**

>function: EmptyDB() 

```
Example 1: (ignore return)
	tar := Tardigrade{}
	tar.EmptyDB()

Example 2 (capture return):
	tar := Tardigrade{}
	msg, status := tar.EmptyDB()
	fmt.Println(msg, status)

Return:
	Empty: database now clean! true
	Failed: no permission to re-create! false
	Missing: could not find database false! false

```

**AddField() function take in ((key)string, (Value) string) and add to database.**

>function: AddField()

```
Example 1: (ignore return)
	tar := Tardigrade{}
	tar.AddField("New string Entry", "string of data representing a the value")

Example 2 (capture return):
	tar := Tardigrade{}
	status := tar.AddField("New string Entry", "string of data representing a the value")
	fmt.Println(status)

Return:
	true | false

```

**CountSize() function will return number of rows in the gojsondb.db**

>function: CountSize()

````
Example (capture return):
	tar := Tardigrade{}
	fmt.Println(tar.CountSize())

Result:
	44
````

**FirstField func returns the first entry of gojsondb.db in all formats \[ raw | json | id | key | value ] specify format required**

>function: FirstField()

```
Example 1: (true | failed)
	tar := Tardigrade{}
	fmt.Println(tar.FirstField("raw"))

Result: 
	{"id":1,"key":"one","data":"string data test"}
	Failed: database tardigrade.db is empty!
	Failed: database tardigrade.db missing!

Example 2: (true)
	tar := Tardigrade{}
	fmt.Println(tar.FirstField("json"))

Result:
{
        "id": 1,
        "key": "New string Entry",
        "data": "string of data representing a the value"
}
```

**LastField() func returns the last entry in multi-format \[ raw | json | id | key | value ]**

>function: LastField()

```
Example 1: (true | failed)
	tar := Tardigrade{}
	fmt.Println(tar.FirstField("raw"))

Result: 
	{"id":44,"key":"New Entry","data":"string of data representing a the value"}
	Failed: database tardigrade.db is empty!
	Failed: database tardigrade.db missing!

Example 2: (true)
	tar := Tardigrade{}
	fmt.Println(tar.LastField("value"))

Result:
	string of data representing a the value

Example 3: (true)
	tar := Tardigrade{}
	fmt.Println(tar.LastField("key"))

Result:
	New Entry

Example: 4 (true)
	tar := Tardigrade{}
	fmt.Println(tar.LastField("json"))

Result:
{
        "id": 44,
        "key": "New Entry",
        "data": "string of data representing a the value"
}
```

**SelectByID func returns an entry string for a specific id in all formats \[ raw | json | id | key | value ]**
>function: SelectByID()

```
Example 1: (true)
	tar := Tardigrade{}
	fmt.Println(tar.SelectByID(10, "raw"))

Result:
	{"id":10,"key":"Roman","data":"string of data representing a the value of X"}

Example 2: (false)
	tar := Tardigrade{}
	fmt.Println(tar.SelectByID(100, "raw"))

Result:
	Record 100 is empty!

Example 3: (true)
	tar := Tardigrade{}
	fmt.Println(tar.SelectByID(25, "json"))

Result:
{
        "id": 25,
        "key": "New string Entry 23",
        "data": "string of data representing a the value"
}
```

**UniqueID function returns an int for the last used UniqueID**
>function: UniqueID()

```
Example: (always true)
	tar := Tardigrade{}
	fmt.Println(tar.UniqueID())

Result:
	52
```

**LastXFields returns last X number of entries from db in values byte[] format**
>function: LastXFields()

```
Example 1: (true)
	tar := Tardigrade{}
	var received = tar.LastXFields(2)

	bytes := received
	var data []MyStruct
	json.Unmarshal(bytes, &data)

	for l := range data {
		fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
		fmt.Println()
	}

Result:
	id: 51, key: New string Entry 49, data: string of data representing a the value
	id: 52, key: New string Entry 50, data: string of data representing a the value
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
msg, status := RemoveField(2)
fmt.Println(msg, status)
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

Return:  false

Record 2 is empty! false
id: 1, key: New string Entry, data: string of data representing a the value
id: 3, key: New string Entry, data: string of data representing a the value


```

**ModifyField - Takes an id, Key & Value (all 3 fields) and update with information provided for key & value**
> ModifyField(2, "Updated key", "Updated data set with new inforation")

```
Example:
// Check current information in ROW 2 BEFORE CHANGE
fmt.Println(SelectByID(2, "raw"))

// MODIFY ROW 2 with NEW information provided in key & value
var change = ModifyField(2, "Updated key", "Updated data set with new inforation")
fmt.Println("Changed: ", change)

// Check current information in ROW 2 AFTER CHANGE
fmt.Println(SelectByID(2, "raw"))

Result:

{"id":2,"key":"Before Entry","data":"string of data representing a the value"}
Changed:  true
{"id":2,"key":"Updated key","data":"Updated data set with new inforation"}
```

RELEASE NOTE:

```
** release 0.0.1 - Initial version
** release 0.0.2 - Cleanup comments and unused files
```

OUTSTANDING
```
** Write some test functions
```



