# Password Generator

## Overview
A simple command-line tool written in Go to generate secure passwords based on user-defined criteria such as length and inclusion of uppercase letters, numbers, and special characters.

## Features
- Generate passwords of customizable length.
- Include or exclude:
  - Uppercase letters
  - Numbers
  - Special characters
- Secure randomization using `crypto/rand`.

## Prerequisites
- Go installed on your machine (version 1.17 or later).

## Usage

### Clone the Repository
```bash
git clone [https://github.com/gaurikarkhile001/Password-Generator-in-Golang.git](https://github.com/gaurikarkhile001/Password-Generator-in-Golang.git)
cd password-generator
```

### Run the Program
```bash
go run main.go
```

### Available Flags
| Flag         | Description                          | Default |
|--------------|--------------------------------------|---------|
| `-length`    | Length of the password               | `12`    |
| `-upper`     | Include uppercase letters           | `true`  |
| `-numbers`   | Include numbers                     | `true`  |
| `-specials`  | Include special characters          | `true`  |

### Example Output
```bash
Generated Password: A8@dZ2!xKbP$g9
```

## License
This project is licensed under the MIT License. Feel free to use and modify it as needed.

---

**Happy Coding!**

