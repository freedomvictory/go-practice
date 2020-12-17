# GORM 

It's a fantastic ORM library for Golang aims to be developer friendly 



## A simple demo 

```go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create(same like insert)
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
}

```

## Declaring Models 

Models are normal structs with basic Go types, pointer/alias of them or custom types implementing `Scanner` and `Valuer` interfaces

For Example : 

```go
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivedAt    sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

### Conventions 

`Gorm` prefer convention over configuration, by default , `Gorm` uses `ID` as primary key.  


### `gorm.Model`

```go
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

`gorm.Model` defined by `GORM`, You can embed it into your struct to include those fields 

### **Advanced** 

#### Fieled-Level Permission 


1. write
2. read 
... 


```
//TODO 
```


#### Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking 

`GORM` use `CreatedAt` , `UpdatedAt` to track creating/updating time by convention, and GORM will set the current time when creating/updating if the fields are defined 

To use fields with a different name , you can configure those fields with tag `autoCreateTime`, `autoUpdateTime`

If you prefer to save `UNIX` (mill/nano) seconds instead of time, you can simply change the field's data type from `time.Time` to `int`

```go 
type User struct {
  CreatedAt time.Time // Set to current time if it is zero on creating
  UpdatedAt int       // Set to current unix seconds on updaing or if it is zero on creating
  Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
  Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
}
```

#### Filed Tags 

```
//TODO 
```