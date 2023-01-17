# Clean Architecture API 

### Models
- also known as Domains / Entities
- defined by use of a struct, similar to classes in OOP
  ```
  type Person struct {
    FirstName string `json:"firstname" binding:"required"`
    LastName  string `json:"lastname" binding:"required"`
    Age       int8   `json:"age" binding:"gte=1,lte=130"`
    Email     string `json:"email" binding:"required,email"`
  }

  type Video struct {
    Title       string `json:"title" binding:"min=2,max=50" validate:"is-cool"`
    Description string `json:"description" binding:"max=20"`
    URL         string `json:"url" binding:"required,url"`
    Author      Person `json:"author" binding:"required"`
  }
  ```
  
### Repository layer
  - responsible for any database related activities eg fetching data, posting data
  
### Service layer
  - holds the core business logic
  - Business logic may include actions that might require to make a database call or make an external API request
  - we need to inject our repository into the services.

### Controllers
- static handlers for the incoming HTTP / HTTPS requests
- delegates the data from the request to the service layer for implementation of business logic
- returns the response for the requests

### templates 
- displays what the user will view --> the frontend

