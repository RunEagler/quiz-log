---
description: generate Go models from schema
---

Execute the following steps in the backend directory:

**Note:** Only generate new files that don't already exist.

**PostgreSQL Connection Information:**

- Host: `localhost`
- Port: `5432`
- Database: `quizlog`
- User: `postgres`
- Password: `postgres`

**Generate Directory:**

`backend/models/[table_name].go`

**Generated File Specifications:**

- One file per table
- File name matches the table name
- Struct name should be singular (remove trailing 's' from table name)
  - Example: `quizzes` table → `Quiz` struct
  - Example: `questions` table → `Question` struct
  - Example: `attempts` table → `Attempt` struct
  - Example: `tags` table → `Tag` struct
- Generate struct with embedded `bun.BaseModel`
- Add `bun` tags for field mapping
- Include appropriate Go types for PostgreSQL types
- Add `time.Time` fields for timestamps
- Use pointers for nullable fields
- generate `getter` methods for each field

