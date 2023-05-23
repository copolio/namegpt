# ERD
```mermaid
erDiagram
    User {
      id int PK
      name string
    }
    Query {
      id int PK
      keyword string
    }
    Domain {
      id int PK
      name string
      checkDate datetime
    }

    User ||--o{ Query: search
    Query ||--o{ Domain: return
```