# go-server

A lightweight HTTP server built with Go that serves static files and handles form submissions.

## Project Structure

```
go-server/
├── main.go          # Server setup and route registration
├── handlers.go      # Route handler functions
├── go.mod           # Go module definition
└── static/
    ├── index.html   # Landing page
    └── form.html    # Form page
```

## Routes

| Method | Path         | Description                        |
|--------|--------------|------------------------------------|
| GET    | `/`          | Serves static landing page         |
| GET    | `/form.html` | Serves the form page               |
| POST   | `/form`      | Handles form submission            |
| GET    | `/hello`     | Returns a hello page               |

## Request Flow

```
Browser
   │
   ├── GET /              ──► FileServer ──► static/index.html
   │
   ├── GET /form.html     ──► FileServer ──► static/form.html
   │
   ├── POST /form         ──► formHandler ──► result page
   │                                             └── body: name, profession
   │
   └── GET /hello         ──► helloHandler ──► hello page
```

## Getting Started

```bash
# Initialize module (first time only)
go mod init go-server

# Run the server
go run .

# Build binary
go build -o server .
```

Visit `http://localhost:8080` in your browser.
