git The Swagger UI will provide an interactive interface where you can test your API endpoints directly from the browser. Remember to regenerate docs with `swag init` whenever you update your API comments.

Here's how to implement Swagger in a Fiber Go application:

## 1. Install Required Dependencies

```bash
go get github.com/gofiber/fiber/v2
go get github.com/swaggo/swag/cmd/swag
go get github.com/gofiber/swagger
go get github.com/swaggo/files
```

## 2. Add Swagger Comments to Your Main File

## 3. Generate Swagger Documentation

Run this command in your project root to generate the Swagger docs:

```bash
swag init
```

This creates a `docs` folder with generated files including `docs.go`, `swagger.json`, and `swagger.yaml`.

## 4. Key Swagger Annotations Explained

- **General Info**: `@title`, `@version`, `@description` - Basic API information
- **Contact & License**: `@contact.name`, `@license.name` - API metadata
- **Host & BasePath**: Define your API's base URL structure
- **Route Documentation**: `@Summary`, `@Description`, `@Tags` - Describe endpoints
- **Parameters**: `@Param` - Define query params, path params, and request bodies
- **Responses**: `@Success`, `@Failure` - Define response structures
- **Router**: `@Router` - Map the endpoint path and HTTP method

## 5. Access Your Swagger UI

After running your application, visit:
```
http://localhost:3000/swagger/index.html
```

