# Go CRUD API

---

## Model
---
```
type Candidate struct {
	gorm.Model

	Name         string `gorm:"type:varchar(100)"`
	Phone_number string `gorm:"type:varchar(20)"`
	Status       string `gorm:"type:varchar(10)"`
}
```

## Table
---
* Table Name : `candidates`

* Fields
```
id | created_at | updated_at | deleted_at | name | phone_number | status
---+------------+------------+------------+------+--------------+-------
   |            |            |            |      |              | 
```


## End Points
---
### POST
* `/candidates`  => Create Candidate.

### GET
* `/candidates/:id`  => Returns Candidate with specified id.

* `/candidates`  => Returns List of all Candidates.

### DELETE
* `/candidates/:id` => Delete Candidate with specified id.

### PUT
* `/candidates/:id` => Update Candidate with specified id.


## TODO
---
* Validations on CRUD
* Appropriate Resonse for various cases (Object not found with id, Request format, Object Not created etc).

