<p align="center">
  <img src="./docs/pwncat.png" width="50%" alt="pwncat">
</p>

<h1 align="center">pwncat // Web Fuzzing Tool</h1>
<div align="center">

<a href="https://github.com/wvverez/pwncat/stargazers"><img src="https://img.shields.io/github/stars/wvverez/pwncat?style=for-the-badge&logo=starship&color=C9CBFF&logoColor=C9CBFF&labelColor=302D41" alt="stars"></a>&nbsp;&nbsp;
<a href="https://github.com/wvverez/pwncat/forks"><img src="https://img.shields.io/github/forks/wvverez/pwncat?style=for-the-badge&logo=git&logoColor=f9e2af&label=Forks&labelColor=302D41&color=f9e2af" alt="forks"></a>&nbsp;&nbsp;
<a href="https://github.com/wvverez/pwncat/issues"><img src="https://img.shields.io/github/issues/wvverez/pwncat?style=for-the-badge&logo=github&logoColor=eba0ac&label=Issues&labelColor=302D41&color=eba0ac" alt="issues"></a>&nbsp;&nbsp;
<a href="https://github.com/wvverez/pwncat/commits/main"><img src="https://img.shields.io/github/last-commit/wvverez/pwncat?style=for-the-badge&logo=github&logoColor=white&label=Last%20Commit&labelColor=302D41&color=A6E3A1" alt="last commit"></a>&nbsp;&nbsp;
<a href="https://github.com/wvverez/pwncat/blob/main/LICENSE"><img src="https://img.shields.io/github/license/wvverez/pwncat?style=for-the-badge&logo=open-source-initiative&color=CBA6F7&logoColor=CBA6F7&labelColor=302D41" alt="license"></a>&nbsp;&nbsp;

</div>


<div align="center">

Made with ♥ by wvverez 

</div>

---

**Pwncat** is a modern, high-performance web fuzzing tool written in [Go](https://go.dev/), designed for professional [penetration testing](https://owasp.org/www-community/Penetration_Testing) and security research. It enables security professionals to systematically discover hidden directories, files, subdomains, and parameters on web applications through automated [HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP) requests.

For installation instructions, usage examples, and available options, see the [Documentation](https://github.com/wvverez/pwncat#readme). If you encounter a bug or would like to request a new feature, please open an [Issue](https://github.com/wvverez/pwncat/issues). Contributions are welcome through [Pull Requests](https://github.com/wvverez/pwncat/pulls).

## 🏗️ pwncat tree architecture

```sh
pwncat/
├── cmd/
│   └── pwncat/
│       └── main.go
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── flags.go
│   ├── engine/
│   │   ├── engine.go
│   │   └── worker.go
│   ├── filter/
│   │   ├── excluder.go
│   │   └── matcher.go
│   ├── http/
│   │   ├── client.go
│   │   ├── request.go
│   │   └── response.go
│   ├── input/
│   │   └── provider.go
│   └── output/
│       └── display.go
├── pkg/
│   └── utils/
│       └── color.go
├── wordlists/
│   └── common.txt
├── go.mod
```

## 🎯 Use Cases

**pwncat** can be used as part of an authorized security testing workflow for tasks such as:

- Discovering hidden web directories and files.
- Enumerating accessible application endpoints.
- Identifying potentially exposed resources.
- Performing subdomain discovery.
- Fuzzing URL paths and parameters.
- Mapping the attack surface of web applications.
- Supporting reconnaissance and enumeration phases during penetration tests.
- Conducting security research in controlled and authorized environments.

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). For the full license text, see the [LICENSE](https://github.com/wvverez/pwncat/blob/main/LICENSE) file.


## 👉🏼 Way to Install

```sh
git clone https://github.com/wvverez/pwncat.git && cd pwncat
go build -o pwncat cmd/pwncat/main.go
./pwncat <SNIP> <SNIP>
```

## 📋 Parameter Reference

| Parameter | Shorthand | Value | Description | Example |
|-----------|-----------|-------|-------------|---------|
| `--url` | `-u` | URL | Target URL with `PWN` as placeholder | `-u "http://example.com/PWN"` |
| `--wordlist` | `-w` | Path or spec | Wordlist file path or range specification | `-w wordlists/common.txt` or `-w "range:1-100"` |
| `--method` | `-X` | GET/POST/PUT/etc | HTTP method to use | `-X POST` |
| `--threads` | `-t` | Number | Number of concurrent workers | `-t 50` |
| `--rate` | `-r` | Number | Requests per second (0 = unlimited) | `-r 100` |
| `--timeout` | `-to` | Duration | Request timeout duration | `-to 10s` |
| `--retry` | `-re` | Number | Retry attempts on error | `-re 3` |
| `--delay` | `-dl` | Duration | Delay between requests | `-dl 100ms` |
| `--status` | `-s` | HTTP codes (comma-separated) | Status codes to match | `-s 200,301,302` |
| `--exclude` | `-e` | HTTP codes (comma-separated) | Status codes to exclude | `-e 404,500` |
| `--size-match` | `-sm` | Min-Max | Response size range to match | `-sm 100-2000` |
| `--exclude-size` | `-ex` | Min-Max | Response size range to exclude | `-ex 0-100` |
| `--regex` | `-rg` | Regex pattern | Regex pattern to match in response | `-rg "admin\|login"` |
| `--regex-exclude` | `-rx` | Regex pattern | Regex pattern to exclude | `-rx "error\|notfound"` |
| `--output` | `-o` | File path | Output file for results | `-o results.json` |
| `--format` | `-f` | json/csv/html | Output format | `-f json` |
| `--no-color` | `-nc` | Flag | Disable colored output | `-nc` |
| `--verbose` | `-v` | Flag | Enable verbose mode | `-v` |
| `--debug-log` | `-log` | File path | Debug log file | `-log debug.log` |
| `--config` | `-cfg` | File path | JSON/YAML configuration file | `-cfg config.json` |
| `--replay` | `-rp` | URL | Replay matched requests to this URL | `-rp http://localhost:8080` |
| `--cert` | - | File path | TLS certificate file | `--cert cert.pem` |
| `--key` | - | File path | TLS private key file | `--key key.pem` |
| `--insecure` | `-k` | Flag | Skip TLS verification | `-k` |
| `--version` | - | Flag | Show version information | `--version` |

## Basic directory fuzzing example

<p align="center">
  <img src="./docs/pwncat.gif" alt="pwncat demo" width="800">
</p>

## Disclaimer

> [!NOTE]
> This project is provided for **ethical, educational, and authorized security testing only**.

Do not use it against systems you do not own or have explicit permission to test.

## 📋 To-Do List

- [ ] Exit automatically when the `wordlist is fully processed`.

- [x] Improve and redesign the `ASCII art`.

- [ ] Add `line-related` parameters and configuration options.

- [ ] Fix unwanted line breaks in the output.

- [x] Add response size in `bytes` information to the output.

- [x] Add an option to `disable status code colors`.


## 🤝 Contributions

A brief **thank you** to everyone who contributed their little piece to this project :)


<a href="https://github.com/CuriosidadesDeHackers">
  <img src="https://github.com/CuriosidadesDeHackers.png?size=50" width="50" alt="CuriosidadesDeHackers" style="border-radius: 50%;">
</a>
