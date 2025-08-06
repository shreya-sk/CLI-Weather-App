# CLI-Weather-App
Learning project #2! Trying to extend my golang capabilities to experiment with API retrievals :)
# CLI-Weather-App (Docker Optimized)

A command-line weather application built in Go that fetches real-time weather data from OpenWeatherMap API. This branch demonstrates **Docker multi-stage builds** for production-ready containerization.

## 🐳 Docker Versions Available

- **Basic Docker**: See this branch for simple containerization
- **Optimized Docker**: Check `docker-optimization` branch for multi-stage builds (98% smaller!)

## 🐳 Docker Optimization Results

| Version | Image Size | Build Type |
|---------|------------|------------|
| Basic Docker | 1.37GB | Single-stage |
| **Optimized** | **28.3MB** | **Multi-stage** |

**98% size reduction** achieved through multi-stage builds!

## 🚀 Quick Start with Docker

### Prerequisites
- Docker installed
- OpenWeatherMap API key ([get one free here](https://openweathermap.org/api))

### Build the Image
```bash
# Clone and navigate to project
git clone https://github.com/shreya-sk/CLI-Weather-App.git
cd CLI-Weather-App
git checkout docker-optimization

# Build the optimized image
docker build -t weather-app:optimized .
```

### Run the Application
```bash
# Interactive mode (recommended)
docker run -it \
  -e OPENWEATHER_API_KEY=your_api_key_here \
  -e OPENWEATHER_API_URL=https://api.openweathermap.org/data/2.5/weather \
  weather-app:optimized

# Or with city parameter
docker run -it \
  -e OPENWEATHER_API_KEY=your_api_key_here \
  -e OPENWEATHER_API_URL=https://api.openweathermap.org/data/2.5/weather \
  weather-app:optimized \
  -city "New York"
```

## 📋 Command Reference

### Docker Commands Explained

#### Build Commands
```bash
# Build the optimized image
docker build -t weather-app:optimized .
# -t: Tags the image with a name
# .: Uses current directory as build context

# View all local images
docker images
# Shows all Docker images on your machine with sizes

# View specific image
docker images | grep weather-app
# Filters to show only weather-app images
```

#### Run Commands
```bash
# Interactive mode with environment variables
docker run -it \
  -e OPENWEATHER_API_KEY=abc123 \
  -e OPENWEATHER_API_URL=https://api.openweathermap.org/data/2.5/weather \
  weather-app:optimized

# -it: Interactive terminal (allows user input)
# -e: Sets environment variables inside container
# weather-app:optimized: Image name and tag to run
```

#### Debug Commands
```bash
# Access container shell for debugging
docker run -it weather-app:optimized sh
# Opens shell inside container to inspect files

# View container logs
docker logs container_id
# Shows output from a running container

# List running containers
docker ps
# Shows currently running containers
```

## 🏗️ Multi-Stage Build Explained

This Dockerfile uses **two stages** to minimize final image size:

### Stage 1: Builder (The Factory)
- Uses `golang:1.23` image (~800MB)
- Downloads dependencies
- Compiles Go source code to binary
- Creates statically-linked binary for Linux

### Stage 2: Runtime (The Car)
- Uses `alpine:latest` (~5MB base)
- Installs HTTPS certificates for API calls
- Copies **only the binary** from Stage 1
- Runs the application

### Key Optimizations
```dockerfile
# Cross-compilation for Alpine Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-app main.go

# Copy only the binary (not source code or build tools)
COPY --from=builder /some/weather-app .

# Install certificates for HTTPS API calls
RUN apk --no-cache add ca-certificates
```

## 🌦️ Application Features

- **Interactive mode**: Type city names and get real-time weather
- **Multiple city handling**: Automatically resolves ambiguous city names
- **Colored output**: Beautiful terminal display with ASCII weather icons
- **Global coverage**: Works with cities worldwide
- **Detailed weather info**: Temperature, humidity, wind speed, conditions

## 🔧 Environment Variables

| Variable | Required | Description | Example |
|----------|----------|-------------|---------|
| `OPENWEATHER_API_KEY` | Yes | Your OpenWeatherMap API key | `abc123def456` |
| `OPENWEATHER_API_URL` | Yes | OpenWeatherMap API endpoint | `https://api.openweathermap.org/data/2.5/weather` |
| `OPENWEATHER_UNITS` | No | Temperature units (default: metric) | `metric`, `imperial` |

## 🔍 Troubleshooting

### Common Issues

**"no such file or directory"**
- Ensure cross-compilation flags are set: `CGO_ENABLED=0 GOOS=linux`
- Verify binary exists: `docker run -it weather-app:optimized sh` then `ls -la`

**"API base URL not found"**
- Set both required environment variables when running container
- Check environment variable names match your config.go file

**HTTPS certificate errors**
- Ensure `ca-certificates` is installed in Alpine stage
- Verify internet connectivity from container

## 📚 Learning Outcomes

This project demonstrates:
- **Docker containerization** of Go applications
- **Multi-stage builds** for production optimization
- **Cross-compilation** for different architectures
- **Environment variable management** in containers
- **Debugging containerized applications**
- **Image size optimization** (98% reduction achieved)

## 🔗 Branch Comparison

- **`main`**: Basic Dockerfile (single-stage, 1.37GB)
- **`docker-optimization`**: Multi-stage build (28.3MB) ← You are here
- Compare branches to see the optimization difference!

## 🛠️ Built With

- **Go 1.23.2** - Programming language
- **Docker** - Containerization
- **Alpine Linux** - Minimal base image
- **OpenWeatherMap API** - Weather data source

---

*This is a learning project exploring Docker optimization techniques and Go development practices.*