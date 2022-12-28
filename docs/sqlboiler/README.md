# `sqlboiler`

Database-First ORM

https://blog.logrocket.com/introduction-sqlboiler-go-framework-orms/

## Workflow

1. Create a `sqlboiler.toml` in the root of your project.
2. Paste this configuration: https://github.com/volatiletech/sqlboiler#full-example and update the necessary fields.
3. Define a SQL Schema

   ```sql
   drop table if exists students;
   drop table if exists departments;
   drop table if exists staffs;
   drop table if exists classes;

   create table students (
      id serial not null primary key,
      firstname varchar not null,
      lastname varchar not null,
      email varchar not null,
      admission_number varchar not null,
      year int not null,
      cgpa float not null
   );

   create table classes (
      id serial not null primary key,
      title varchar not null,
      code varchar not null,
      unit int not null,
      semester int not null,
      location varchar not null
   );

   create table departments (
      id serial not null primary key,
      name varchar not null,
      code varchar not null,
      telephone varchar not null,

      foreign key (user_id) references users (id)
   );

   create table staffs (
      id serial not null primary key,
      firstname varchar not null,
      lastname varchar not null,
      email varchar not null,
      telephone varchar not null,
      salary bigint not null,
   );

   create table classes_students (
      class_id int not null,
      student_id int not null,

      primary key (class_id, student_id),
      foreign key (student_id) references students (id),
      foreign key (class_id) references classes (id)
   );

   create table classes_instructors (
      class_id int not null,
      staff_id int not null,

      primary key (class_id, staff_id),
      foreign key (staff_id) references staffs (id),
      foreign key (class_id) references classes (id)
   );

   insert into users (name) values ('Franklin');
   insert into users (name) values ('Theressa');
   ```

4. Load schema into DB:

   ```bash
   psql --username <user> --password <password> < schema.sql
   ```

5. Generate your SQL models:

   ```bash
   sqlboiler psql -c sqlboiler.toml --wipe --no-tests
   ```

## Usage

### Query Mods

Use **query mods** to augment your queries dynamically:

```go
//...
// fetch classes including students with cgpa >= 2.6
classes, err := models.Classes(qm.Load(models.ClassRels.Student, qm.Where("cgpa >= ?", 2.6))).All(ctx, db)
if err != nil {} // handle err
```

```go
//...
// fetch department including students
department, err := models.Departments(qm.Where("id=?", 1), qm.Load(models.DepartmentRels.Student)).One(ctx, db)
if err != nil {} // handle err
```

### CRUD Operations

```go
//...
// create a department
var department models.Department
department.Name = "Computer Science"
department.Code = "CSC"
department.Telephone = "+1483006541"
err := department.Insert(ctx, db, boil.Infer())
if err != nil {} // handle err
```

...

## Transactions

```go
//...
// start a transaction
tx, err := db.BeginTx(ctx, nil)
if err != nil {} // handle err

// create a department
var department models.Department
department.Name = "Computer Science"
department.Code = "CSC"
department.Telephone = "+1483006541"
err = department.Insert(ctx, tx, boil.Infer())
if err != nil {
  // rollback transaction
  tx.Rollback()
}

// create a class
var class models.Class
class.Title = "Database Systems"
class.Code = "CSC 215"
class.Unit = 3
class.Semester = "FIRST"
err = class.Insert(ctx, tx, boil.Infer())
if err != nil {
  // rollback transaction
  tx.Rollback()
}

// add class to department
class, err := models.Classes(qm.Where("code=?", "CSC 215")).One(ctx, tx)
department, err := models.Departments(qm.Where("code=?", "CSC")).One(ctx, tx)
err = department.AddClasses(ctx, tx, class)
if err != nil {
  // rollback transaction
  tx.Rollback()
}

// commit transaction
tx.Commit()
```
