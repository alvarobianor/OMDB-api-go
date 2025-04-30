# 🎬 OMDB API Go

A simple and efficient Go-based API wrapper for the OMDB (Open Movie Database) API. This project provides a clean interface to search for movies and retrieve movie information.

## 🚀 Features

- 🔍 Movie search functionality
- 🔐 Secure API key handling
- 🌐 CORS enabled
- ⚡ Fast and efficient
- 🛡️ Error handling and logging

## 🛠️ Prerequisites

- Go 1.24 or higher
- OMDB API key (get it from [OMDB API](http://www.omdbapi.com/apikey.aspx))

## ⚙️ Installation

1. Clone the repository:

```bash
git clone https://github.com/alvarobianor/OMDB-api-go.git
cd OMDB-api-go
```

2. Create a `.env` file in the root directory:

```bash
API_KEY=your_omdb_api_key_here
```

3. Install dependencies:

```bash
go mod download
```

## 🚀 Running the Application

Start the server:

```bash
go run main.go
```

The server will start on port 8081.

## 📝 API Usage

### Search Movies

```http
GET http://localhost:8081/?movie=inception
```

Headers required:

```
X-CUSTOM-HEADER-API-KEY: your_api_key
```

## 🛠️ Built With

- [Go](https://golang.org/) - The programming language
- [Chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## 👨‍💻 Author

- **Alvaro Bianor** - [alvarobianor](https://github.com/alvarobianor)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [OMDB API](http://www.omdbapi.com/) for providing the movie database API
