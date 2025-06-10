# Gateway CLI

Gateway CLI is a lightweight API Gateway tool written in Go. It routes HTTP requests to backend services based on a YAML-defined configuration. It supports rate limiting, access control, and middleware injection, and is fully configurable via `config.yaml` and `services_config.yaml`.

> 🚧 **Note:** This project is under active development and may contain bugs. Contributions and suggestions are welcome!

---

## 🧩 Features

- 🔀 Request routing to multiple backend services
- 📄 Configuration via YAML files
- 🚦 Built-in rate limiting
- 🧩 Middleware support (`/gateway/middleware`)
- 🛠 CLI utility to manage and start the gateway

---

## 📁 Project Structure

- `config.yaml`: Gateway configuration (env, port, service config path)
- `services_config.yaml`: List of backend services, endpoints, routes, API keys, and rules

---

## 📂 Example Configuration

### `config.yaml`
```yaml
env: env
http_server:
  address: localhost:4001
services_config_path: config/services_config.yaml
```

### `services_config.yaml`
```yaml
services:
  - name: user-service
    baseUrl: http://localhost:5001
    service_endpoint: /service1
    routes:
      - method: "GET"
        endpoint: /users
      - method: "GET"
        endpoint: /users/{id}
      - method: "POST"
        endpoint: /users
    rate_limits: 1
    api_key: xyc
    allowlist:
      - "192.168.1.5"
    enabled: true
```

---

## 🚀 Usage

## Clone and Build
1. Clone the repository:
    ```
     git clone <repository-url> 
     cd <repository-dir> 
    ```
2. Build the binary:
    ```
     go build -o gateway-cli
    ```

## Setup an Alias
To make the tool globally accessible:

1. Open `~/.bashrc (or ~/.zshrc if using Zsh)`:
    ```
    nano ~/.bashrc
    ```
2. Add the following line:
    ```
    alias gateway-cli="$HOME/DIR WHERE THE BUILD IS PLACED/gateway-cli" 
    ```
3. Refresh your shell configuration:
    ```
     . ~/.bashrc
    ```


### Start the CLI
```bash
gatewat-cli help
```

### CLI Commands
| Command         | Description                                                  |
|----------------|--------------------------------------------------------------|
| `config`        | Load services config from specified path                     |
| `help`          | Display help info                                            |
| `service`       | Add/Remove/List services to/from the YAML config             |
| `start`         | Start the API gateway server                                 |
| `status`        | Check gateway server status                                  |
| `stop`          | Stop the gateway server                                      |

### Start Gateway Server
```bash
gateway-cli start -d
```
- `-d`, `--daemon`: Run server in background  
- `-h`, `--help`: Help for start

---

## 📌 Routing Example

If a client sends a request to:

```
<gateway_base_url>/service1/users
```

The gateway will proxy it to:

```
http://localhost:5001/users
```

---

## ⚙️ Middleware

You can modify or extend the gateway functionality by adding middleware to `/gateway/middleware`. Examples include:

- Custom logging
- Authentication
- Caching
- Retry logic

---

## 🛡 Known Limitations

- Project is in early development.
- YAML validation is minimal.
- Rate limiter is basic (single token per service, adjustable per your needs).

---

## 🔮 Planned Features

Here are some upcoming improvements and considerations for the project:

1. ~~**Service-Specific Middleware**  
   Separate service proxies to allow applying different middleware stacks for each individual service~~

2. **Reviewing `api_key` Field**  
   The `api_key` field in `services_config.yaml` is currently unused. Will evaluate its purpose; it may be removed or repurposed in future versions.

3. **Redis-Based Rate Limiter**  
   Replace or enhance the current in-memory rate limiter with a distributed Redis-based implementation for scalability and resilience.
4. **IP whitelisting per service**
---


## 📬 Contributions

Got feedback or improvements?  
Feel free to submit a pull request or open an issue.

---