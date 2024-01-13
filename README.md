<div align="center">
 <h1> Simple Real-Time Truck Location Tracker </h1>
</div>

<div align="center">

  [![Go](https://img.shields.io/badge/Go-v1.21-blue.svg)](https://golang.org/)
  ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/user/real-time-truck-tracker/main)
  ![GitHub Repo stars](https://img.shields.io/github/stars/user/real-time-truck-tracker)
  ![License](https://img.shields.io/badge/license-MIT-green)
  [![CI](https://github.com/user/real-time-truck-tracker/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/user/real-time-truck-tracker/actions/workflows/ci.yaml)
  ![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/user/real-time-truck-tracker/main)
  ![Github Repository Size](https://img.shields.io/github/repo-size/user/real-time-truck-tracker)
  ![GitHub contributors](https://img.shields.io/github/contributors/user/real-time-truck-tracker)
  ![Github Open Issues](https://img.shields.io/github/issues/user/real-time-truck-tracker)
</div>

<div align="center">
  <img src="images/truck-logo.png" alt="Project Logo" width="200">
</div>

## Introduction

This project is a distributed microservices application for real-time tracking of truck locations. It captures truck locations and calculates the distance traveled to determine road usage charges.

## Table of Contents

- [System Overview](#system-overview)
- [Installation and Running](#installation-and-running)
- [Features](#features)
- [License](#license)

## System Overview

- **OBU Service:** Connects via WebSocket and transmits location data.
- **Receiver Service:** Receives data and forwards it to Kafka.
- **Calculator Service:** Processes location data from Kafka, calculates distances.
- **Aggregator Service:** Receives calculated distances and determines charges.

## Installation and Running

1. **Go 1.21+**: Ensure Go 1.21 or later is installed. [Download Go](https://golang.org/dl/)

2. **Kafka**: Requires a Kafka broker instance. [Download Kafka](https://kafka.apache.org/downloads)

3. **Make**: Unix-like systems have it pre-installed. For Windows, use [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm).

4. **Protocol Buffers**: Install Protocol Buffers for gRPC. [Protobuf Installation](https://grpc.io/docs/protoc-installation/)

5. **Running Services**: Use the Makefile to build and run services:

   ```bash
   # Build and run individual services
   $ make obu
   $ make receiver
   $ make calculator
   $ make aggregator
   ```

## Features

- 🔱 Clean architecture
- ⚙️ Makefile 
- 🗃️ Kafka
- 🔄 gRPC/HTTP
- 🚦 Graceful shutdown in some repos


## License
This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize the "TODO" sections with specific details about how to use and configure each microservice in your application.