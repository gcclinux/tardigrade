
## "tardigrade" is small and simple no-SQL database app & mod for GO.
*updated:  Sun 19 Feb 2023 15:15:12 GMT*<br>
*release:  0.1.0*

<br>

## Getting Started with App Code
>git clone [https://github.com/gcclinux/tardigrade.git](https://github.com/gcclinux/tardigrade.git)

## Getting started with module in GO
>go get [https://github.com/gcclinux/tardigrade-mod](https://github.com/gcclinux/tardigrade-mod)

## Getting started with ready compiled binary
>Download [https://github.com/gcclinux/tardigrade/tree/main/bin](https://github.com/gcclinux/tardigrade/tree/main/bin) 

<BR>

Current structure and available functions() through the MOD

```
 -createdb		"CREATE a new database"
 -copydb 		"CREATE a backup / copy of main database"
 -deletedb 		"DELETE main database"
 -deletef 		"DELETE a single row / field from database"
 -insert 		"INSERT new row / field into the database"
 -change 		"CHANGE existing row / field in the dtabase"
 -selectf 		"SELECT FIRST row / field from database"
 -selectfx 		"SELECT FIRST X of rows / fields from database"
 -selecti 		"SELECT specific row / field from database"
 -selectl 		"SELECT LAST row / field from database"
 -selectlx 		"SELECT LAST X of rows / fields from database"
 -total 		"SHOW number of entries"
```


# HOW-TO-USE STANDALONE BINARY

<BR>

**CreateDB - This function will create a database file if it does not exist and return true | false**
>function: CreateDB()
```
Example:
	bin/tardigrade-linux-x86_64 -createdb

```

**DeleteDB - WARNING - this function delete the database file return true | false**
>function: DeleteDB()
```
Example:
	bin/tardigrade-linux-x86_64 -deletedb

```
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
>function: CreatedDBCopy()

```
Example:
	bin/tardigrade-linux-x86_64 -copydb

```

**AddField() function take in ((key)string, (Value) string) and add to database.**

>function: AddField()

```
Example: 
	bin/tardigrade-linux-x86_64 -insert "key free text" "value free text string"
```

**CountSize() function will return number of rows in the database**

>function: CountSize()

````
Example:
	bin/tardigrade-linux-x86_64 -total 
````

**FirstField func returns the first entry of gojsondb.db in all formats \[ raw | json | id | key | value ] specify format required**

>function: FirstField()

```
Example:
	bin/tardigrade-linux-x86_64 -selectf 
```

**LastField() func takes an id and returns the last entry in multi-format \[ raw | json | id | key | value ]**

>function: LastField()

```
Example:
	bin/tardigrade-linux-x86_64 -selectl
```

**SelectByID func take an id and format and returns an entry string for a specific id in all formats \[ raw | json | id | key | value ]**
>function: SelectByID()

```
Example:
	bin/tardigrade-linux-x86_64 -selecti "id" "format"
```

**FirstXFields returns last X number of entries from db in byte[] format**
>function: FirstXFields()

```
Example:
	bin/tardigrade-linux-x86_64 -selectfx "id"
```

**LastXFields returns last X number of entries from db in values byte[] format**
>function: LastXFields()

```
Example:
	bin/tardigrade-linux-x86_64 -selectlx "id"
```

**RemoveField function takes an unique field id as an input and remove the matching field entry**
>function: RemoveField()

```
Example:
	bin/tardigrade-linux-x86_64 -deletef "id"
```

**ModifyField function takes ID, Key, Value and update row = ID with new information provided**
> ModifyField(2, "Updated key", "Updated data set with new inforation")

```
Example: 
	bin/tardigrade-linux-x86_64 -change "id" "new key free text" "new value free text string"

```

RELEASE NOTE:

```
** release 0.0.1 - Initial version
** release 0.0.2 - Updated README.md and corrected some issues.
** release 0.0.3 - Modified to use structure method
** release 0.0.5 - External command input
** release 0.1.0 - Build the first binary
```

OUTSTANDING:
```
** Write and share some test functions
```
