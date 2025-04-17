# Go + Tailwind CSS Hot Reload Demo

This project demonstrates how to set up hot reloading for both Go templates and Tailwind CSS in a web application.

## Prerequisites

Before you begin, make sure you have the following installed:

- Go (1.16 or later)
- Node.js and npm
- Air (Go hot reload tool): `go install github.com/cosmtrek/air@latest`

## Project Structure

```
├── .air.toml            # Air configuration for Go hot reloading
├── go.mod               # Go module file
├── main.go              # Main Go application
├── package.json         # Node.js dependencies and scripts
├── tailwind.config.js   # Tailwind CSS configuration
├── static/              # Static assets
│   └── css/
│       ├── input.css    # Source CSS with Tailwind directives
│       └── styles.css   # Generated CSS (will be created)
└── templates/           # HTML templates
    └── index.html       # Main template file
```

## Setup

1. Install Node.js dependencies:

```bash
npm install
```

2. Build the initial Tailwind CSS file:

```bash
npm run tailwind:build
```

## Running with Hot Reload

To start the application with hot reloading for both Go and Tailwind CSS:

```bash
npm run dev
```

This command will:
- Watch for changes in your Go files and templates and automatically rebuild and restart the server
- Watch for changes in your CSS and HTML files and automatically rebuild the Tailwind CSS

Your application will be available at http://localhost:8080

## How It Works

### Go Hot Reloading with Air

Air watches for changes in your Go files and templates. When a change is detected, it automatically rebuilds and restarts your application.

### Tailwind CSS Hot Reloading

The Tailwind CLI watches your CSS input file and template files for changes. When a change is detected, it automatically rebuilds the CSS output file.

### Concurrently

The `concurrently` package is used to run both the Air and Tailwind CLI watch commands simultaneously.

## Making Changes

1. Edit the Go files or templates to see the Go hot reloading in action
2. Edit the CSS or add Tailwind classes to the HTML to see the Tailwind CSS hot reloading in action

## Production Build

For production, you should build an optimized version of your Tailwind CSS:

```bash
npm run tailwind:build
```

Then build and run your Go application as usual:

```bash
go build -o app .
./app
```