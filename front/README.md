# WATEC Teacher Frontend

React + Vite frontend for the Teacher feature.

## Run locally

```bash
cd front
npm install
npm run dev
```

Default API base URL:

```text
http://localhost:8080/api/v1
```

You can change it in the UI or before running:

```bash
VITE_API_BASE_URL=http://localhost:8080/api/v1 npm run dev
```

## Expected Gateway endpoints

The UI calls:

```text
POST   /api/v1/teachers
GET    /api/v1/teachers/:name
PUT    /api/v1/teachers/:name
DELETE /api/v1/teachers/:name
```

Request body for create/update:

```json
{
  "name": "Aidos",
  "free": [
    { "roomNumber": 204, "from": 9, "to": 12 }
  ]
}
```
