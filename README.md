# BEV and PHEV Data Visualization Application

This repository contains a web application designed for visualizing data related to Battery Electric Vehicles (BEVs) and Plug-in Hybrid Electric Vehicles (PHEVs). It uses data from the Washington State Department of Licensing (DOL), focusing on the current registration statistics for these types of vehicles. 

The frontend is built using Solid.js, providing an interactive and responsive user interface, while the backend is powered by a Golang HTTP server, handling data processing and serving.

## Prerequisites

Before you begin, ensure you have the following installed:
- [Node.js](https://nodejs.org/) (with npm or yarn)
- [Go](https://golang.org/dl/) (version 1.21 or later)
- An active OpenAI API key


## Getting Started

Follow these steps to get your development environment set up:

### 1. Clone the Repository

Clone this repository to your local machine using:

```bash
git clone https://github.com/fcmdias/CSVAnalysis.git
cd CSVAnalysis
```

### 2. Environment Variables
```bash
export AIToken='your_openai_api_key'
```

### 3. Start data backend service 
```bash
cd services/backend/cmd/data
go run main.go
```

### 4. Start AI backend service 
```bash
cd services/backend/cmd/ai
go run main.go
```

### 5. Start frontend service 
```bash
cd services/frontend
npm install 
npm run dev
```

### 6. Accessing the Application

After starting both the frontend and backend servers, you can access the application at http://localhost:3000.

