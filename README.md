```
 ██████╗██╗████████╗██╗   ██╗██████╗ ██╗   ██╗██╗     ███████╗███████╗
██╔════╝██║╚══██╔══╝╚██╗ ██╔╝██╔══██╗██║   ██║██║     ██╔════╝██╔════╝
██║     ██║   ██║    ╚████╔╝ ██████╔╝██║   ██║██║     ███████╗█████╗  
██║     ██║   ██║     ╚██╔╝  ██╔═══╝ ██║   ██║██║     ╚════██║██╔══╝  
╚██████╗██║   ██║      ██║   ██║     ╚██████╔╝███████╗███████║███████╗
 ╚═════╝╚═╝   ╚═╝      ╚═╝   ╚═╝      ╚═════╝ ╚══════╝╚══════╝╚══════╝
```
# CityPulse

A real-time urban mobility heatmap engine built using Go, Python, and Rust.  
CityPulse ingests live real-time road traffic data from TomTom’s Traffic Flow API, transforms segment speeds & congestion levels into a spatial density grid, and visualizes the movement and breathing pattern of a city street network.

This project is structured as a multi-language learning journey:

- Go — Core engine for ingestion and streaming  
- Python — Visualization and analysis  
- Rust — Optional performance modules  

---

## Table of Contents

1. Overview  
2. System Architecture  
3. Phase 1 — Go (Core Engine)  
4. Phase 2 — Python (Visualization Layer)  
5. Phase 3 — Rust (Performance Modules)  
6. Optional Phase 4 — Data & Prediction  
7. Roadmap  
8. Contributing  
9. License  

---

# Overview

CityPulse visualizes real-time transit activity as a continuously evolving heatmap.  
By aggregating GTFS-Realtime vehicle positions onto a spatial grid, the system produces a dynamic portrait of a city's movement patterns.

### Language Roles

- **Go**: concurrency, polling, data streaming  
- **Python**: visualization, GUIs, analysis  
- **Rust**: optimized modules for performance-critical operations  

---

# System Architecture

                    +------------------------------+
                    |        Frontend Layer        |
                    |   (Python GUI / WASM UI)     |
                    +--------------+---------------+
                                   ^
                                   | Live Heatmap Frames
                                   |
    +------------------------------+------------------------------+
    |                         Go Backend                          |
    |-------------------------------------------------------------|
    | - GTFS-Realtime polling (vehicle positions)                 |
    | - Spatial density grid computation                          |
    | - REST API + WebSocket streaming                            |
    +------------------------------+------------------------------+
                                   |
                                   | Optional historical logging
                                   v
                   +-------------------------------+
                   |          Database Layer       |
                   |      (SQLite/PostgreSQL)      |
                   +-------------------------------+
                                   |
                                   | Offline analytics / ML
                                   v
               +--------------------------------------+
               |              Python Lab              |
               |     (Jupyter Notebooks, Analysis)    |
               +--------------------------------------+

             Optional: Rust modules for performance

---

# Phase 1 — Go (Core Engine)

## Purpose
Implement the live backend that powers CityPulse by fetching, processing, and streaming transit data.

## Components

### 1. Traffic Flow Fetcher (TomTom API)
- Polls TomTom Traffic Flow API at fixed intervals 
- Fetches live congestion + road segment geometry (polyline coordinates)  
- Extracts `currentSpeed, freeFlowSpeed, congestion level, list of coordinates`  
- Converts each road segment into one or more heatmap cells
- Maintains current in-memory traffic snapshot for the entire city

### 2. Spatial Density Grid
- Define bounding box for the city  
- Create grid (e.g., 100×100 or 256×256)  
- Map each `(lat, lon)` to grid coordinates  
- Increment cell values  
- Optional smoothing (Gaussian blur/kernel filtering)  

### 3. HTTP + WebSocket Server
Endpoints:
GET /health      -> Service status
GET /vehicles    -> Raw vehicle list (JSON)
GET /heatmap     -> Current density grid
WS  /stream      -> Continuous heatmap frames

### 4. Minimal Verification Output
- ASCII heatmap  
- JSON-only output  
- Optional PNG export  

## Skills Learned
- Go concurrency (goroutines + channels)  
- HTTP servers with `net/http`  
- Protobuf parsing  
- Spatial reasoning & grid systems  

---

# Phase 2 — Python (Visualization Layer)

## Purpose
Build a clean, real-time heatmap viewer without writing manual JavaScript.

## Visualization Approaches

### A. Python Desktop GUI (PyQt / PySide / Kivy)
- Load static OSM map or tile set  
- Overlay heatmap matrix  
- Refresh via polling or WebSocket  

### B. Streamlit or Dash
- Pure Python  
- Libraries generate HTML/JS  
- Real-time updates via polling  
- Heatmaps rendered with Plotly/Matplotlib  

## Optional Features
- Opacity slider  
- Resolution switching  
- Vehicle marker overlay  
- Replay mode  
- Multi-city support  

## Skills Learned
- GUI frameworks  
- Image blending  
- Numerical processing with `numpy`  
- Network communication  

---

# Phase 3 — Rust (Performance Modules)

## Purpose
Enhance computational speed or rendering performance.

## Possible Rust Extensions

### 1. Density Kernel Accelerator
- Compute large heatmaps quickly  
- Integrate via FFI or separate microservice  

### 2. WASM Renderer
- Rust → WebAssembly  
- Browser-based high-performance visualizer  
- Shaders, blending, or particle effects  

### 3. Advanced Spatial Indexes
- Quadtrees, R-trees, geohashes  
- Useful for dense cities  

## Skills Learned
- Ownership & borrowing  
- `cargo` workflows  
- WASM/FFI pipelines  
- High-performance math  

---

# Optional Phase 4 — Data & Prediction

## Purpose
Unlock analytics and forecasting by storing historical data.

## Possible Features
- Replay previous days/hours  
- Average heatmaps per time of day  
- Anomaly detection  
- Forecasting (ARIMA, Prophet, ML models)  

---

# Roadmap

### Core Milestones

[ ] Phase 1 — Go Backend
[ ] API fetcher
[ ] Density grid engine
[ ] API + WebSocket server
[ ] Minimal output

[ ] Phase 2 — Python Visualizer
[ ] GUI or Streamlit interface
[ ] Heatmap overlay rendering
[ ] Interface refinements

[ ] Phase 3 — Rust Enhancements
[ ] High-speed density kernel
[ ] Optional WASM rendering layer

### Optional

[ ] Historical database
[ ] Replay mode
[ ] Time-series analysis and ML

---

# Contributing

Contributions are welcome.  
Helpful areas include:

- Performance improvements  
- Additional visualization modes  
- Documentation improvements  
- Multi-city support  

---

# License

CityPulse is licensed under the MIT License.
