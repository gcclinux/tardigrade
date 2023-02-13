
## Small and simple noSQL JSON database for small apps and easy use.
<br>

## Getting Started
>go get wagemaker.no-ip.co.uk:3000/ricardo/gojsondb

<br>
# HOW-TO-USE

***AddField() function take in ((key)string, (Value) string) and add to database.***

>AddField("New string Entry", "new value string entry")

```true```

***CountSize() function will return number of rows in the gojsondb.db***

>CountSize()

````44````

***LastField() func returns the last entry in multi-format [ raw | json | id | key | value ]***

>LastField("raw")

```{"id":44,"key":"New Entry","data":"string of data is the value"}```

>LastField("id")

```44```

>LastField("key")

```New Entry```

>LastField("value")

```string of data is the value```

>LastField("json")

```
Last Line JSON  :  {
        "id": 44,
        "key": "New Entry",
        "data": "string of data representing a the value"
}
```
***FirstField func returns the first entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required***

>FirstField("raw")

```{"id":1,"key":"one","data":"string data test"}```

***EmptyDB function - WARNING - this will destroy all data stored in gojsondb.db!***

> EmptyDB() 

```true```

***SelectByID func returns an entry string for a specific id in all formats [ raw | json | id | key | value ]***
>SelectByID(10, "raw")

```{"id":10,"key":"Roman","data":"string data test X"}```

>SelectByID(10, "value")

```string data test X```

***UniqueID function returns an int for the last used UniqueID***
> UniqueID()

```54```
