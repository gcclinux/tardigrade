## "tardigrade" is small and simple no-SQL database for small apps and easy use.
*updated:  Fri 03 Mar 12:15:49 GMT 2023*<br>
*release:  1.0.0*

<br>

## Getting Started with App Code
>git clone [https://github.com/gcclinux/tardigrade.git](https://github.com/gcclinux/tardigrade.git)

## Getting started with module in GO
>go get [github.com/gcclinux/tardigrade-mod](https://github.com/gcclinux/tardigrade-mod)

## Getting started with ready compiled binary / standalone executable
> Download: [https://github.com/gcclinux/tardigrade/tree/main/bin](https://github.com/gcclinux/tardigrade/tree/main/bin)

<BR>

Current structure and available functions()

```
--upgrade               "Check for newer version and upgrade the compiled application!"
--createdb              "CREATE new database"
--copydb                "CREATE backup (copy) of the database"
--deletedb              "DELETE database"
--deletef               "DELETE <id> specific row from database"
--search                "SEARCH <Word(s)> <format> match all words and return results"
--selectf               "SELECT <format> TOP row from database "
--selectl               "SELECT <format> LAST row from database"
--selectfx              "SELECT <number> <format> TOP rows from database"
--selectlx              "SELECT <number> <format> LAST rows from database"
--selecti               "SELECT <id> <format> return specific row from database"
--insert                "INSERT <field one> <field two> for new entry"
--change                "CHANGE <id> <field one> <field two> on existing row "
--total                 "SHOW number of entries in database"
--version               "SHOW local (App) & (Mod) build version & date

```


# HOW-TO-USE

<BR>

**CreateDB - This function will create a database file if it does not exist and return true | false**
>function: CreateDB()
```
Example:
	bin/tardigrade-linux-x86_64 --createdb

```

**DeleteDB - WARNING - this function delete the database file return true | false**
>function: DeleteDB()
```
Example:
	bin/tardigrade-linux-x86_64 --deletedb

```
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
>function: CreatedDBCopy()

```
Example:
	bin/tardigrade-linux-x86_64 --copydb

```

**AddField() function take in ((key)string, (Value) string) and add to database.**

>function: AddField()

```
Example: 
	bin/tardigrade-linux-x86_64 --insert "key free text" "value free text string"
```

**CountSize() function will return number of rows in the database**

>function: CountSize()

````
Example:
	bin/tardigrade-linux-x86_64 --total 
````

**FirstField func returns the first entry of gojsondb.db in all formats \[ raw | json | id | key | value ] specify format required**

>function: FirstField()

```
Example:
	bin/tardigrade-linux-x86_64 --selectf 
```

**LastField() func takes an id and returns the last entry in multi-format \[ raw | json | id | key | value ]**

>function: LastField()

```
Example:
	bin/tardigrade-linux-x86_64 --selectl
```

**SelectByID func take an id and format and returns an entry string for a specific id in all formats \[ raw | json | id | key | value ]**
>function: SelectByID()

```
Example:
	bin/tardigrade-linux-x86_64 --selecti "id" "format"
```

**FirstXFields returns last X number of entries from db in byte[] format**
>function: FirstXFields()

```
Example:
	bin/tardigrade-linux-x86_64 --selectfx "id" "format"
```

**LastXFields returns last X number of entries from db in values byte[] format**
>function: LastXFields()

```
Example:
	bin/tardigrade-linux-x86_64 --selectlx "id" "format"
```

**RemoveField function takes an unique field id as an input and remove the matching field entry**
>function: RemoveField()

```
Example:
	bin/tardigrade-linux-x86_64 --deletef "id"
```

**ModifyField function takes ID, Key, Value and update row = ID with new information provided**
> ModifyField(2, "Updated key", "Updated data set with new inforation")

```
Example: 
	bin/tardigrade-linux-x86_64 --change "id" "new key free text" "new value free text string"

```

**SelectSearch function takes (comma, separated pattern(s)) and format if true returning byte[]**
> SelectSearch("pattern1,pattern2", "format")

```
Example: 
	bin/tardigrade-linux-x86_64 --search "pattern1,pattern2" "json"

```
**RunUpgrade function will check release note for current version and then upgrade if required.**
> RunUpgrade()

```
Example: 
	bin/tardigrade-linux-x86_64 --upgrade

Result:
	Upgraded tardigrade-linux-x86_64 to latest version (1.0.0) ....
```


RELEASE NOTE:

```
** release 0.0.1 - Initial version
** release 0.0.2 - Updated README.md and corrected some issues.
** release 0.0.3 - Modified to use structure method
** release 0.0.5 - External command input
** release 0.1.0 - Build the first binary
** release 0.1.1 - Fixed some issues
** release 0.1.2 - Added version for format to selectfx & selectlx
** release 0.1.3 - Added search function -search (pattern(s)), format.
** release 0.1.4 - Added upgrade option for supported systems
** release 0.1.5 - Bug fix storing string with encoder.SetEscapeHTML(false)
** release 1.0.0 - Initial major release CLI interface
```

OUTSTANDING:
```
** Create a WEB UI integrations
```
