# 🚕 Mini Ride Booking API (Go)

A simple backend service for creating and managing taxi rides — built in Go using Clean Architecture.

---

### ✅ Features Implemented

- Create a new ride → `POST /rides`
- Get ride by ID → `GET /rides/{ride_id}`
- Get all rides → `GET /rides`
- Assign a driver to a ride → `PUT /rides/{ride_id}/driver`
- Automatic status update: `pending` → `accepted`
- In-memory storage with auto-incrementing ride IDs

> 📝 **Note:** Currently, there is no driver and rider creation flow.  
> The  driver_id and rider_id are manually provided as a string by the client (e.g., in Postman), by design for simplicity in this version.

---
### ▶️ How to Run

1. Clone the repo
2. Navigate to the project directory
3. Run the server:
   ```bash
   go run ./cmd/main.go
   ```

---

### 📬 How to Use with Postman

You can interact with the API using the following endpoints:

#### ➕ Create a new ride  
**POST** `/rides`  
Body (JSON):
```json
{
  "passenger_id": "1",
  "origin": "Tel Aviv",
  "destination": "Jerusalem"
}
```

#### 📋 Get all rides  
**GET** `/rides`

#### 🔍 Get ride by ID  
**GET** `/rides/{ride_id}`  
Example: `/rides/1`

#### 👨‍✈️ Assign a driver to a ride  
**PUT** `/rides/{ride_id}/driver`  
Body (JSON):
```json
{
  "driver_id": "1"
}
```

---

