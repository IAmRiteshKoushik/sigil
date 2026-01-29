# Certificate Dispatch System

A Go-based system for processing CSV data and sending personalized 
certificates via email using RabbitMQ for queue management and 
concurrent processing.

## Overview

This system reads student data from CSV files, validates eligibility, 
creates event-specific RabbitMQ queues, and processes certificates through 
a concurrent workflow. Each event gets dedicated `cert_` and `dispatch_` 
queues for certificate processing and email delivery.

## Setup

1. **Start RabbitMQ:**
   ```bash
   docker-compose up -d
   ```

2. **Build the application:**
   ```bash
   go build -o sigil .
   ```

3. **Create queues for events:**
   ```bash
   ./sigil create events.txt
   ```

## Usage

The CLI tool currently supports queue creation:

```bash
./sigil create [events-file]
```

This reads events from the specified file and creates:
- `cert_{event}` queues for certificate processing
- `dispatch_{event}` queues for email dispatch

Configure RabbitMQ connection in `config.toml`.

## Current Status

Phase 1 implementation complete with:
- ✅ RabbitMQ container setup
- ✅ CLI queue management
- ✅ Event file processing
- ✅ Dynamic queue creation per event

Phases 2-4 (CSV processing, certificate generation, email dispatch) in development.
