# Users Service
## MVC architectural pattern
The app is divided into three components, exposing a REST API to manage users resources.
### 1. Controllers
They work as endpoints entry points. They're responsible for validating incoming data and calling services to handle business logic.
### 2. Services
They handle business logic.
### 3. Domain
Here we model our business entities, in this case: users.
## Data persistence pattern: DAO (Data access object)
This patterns consists in delegating interacting with the database to each business Entity. In this case, only the User entity knows how to interact with the database.
## Database
For this project i chose to use a SQL Database, particularly Mysql as it integrates well with Go.


