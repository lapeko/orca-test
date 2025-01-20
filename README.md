# Attack Surface Service

## Description
The **Attack Surface Service** provides customers with the ability to query the attack surface of their cloud environment. Specifically, it identifies which virtual machines (VMs) can potentially attack a specified VM based on defined firewall rules.

## Features
1. **Attack Surface Query**: Identify which VMs can access a specified VM.
2. **Service Statistics**: Provides the number of VMs, total requests, and average request processing time.

## Endpoints

### `/api/v1/attack`
- **Description**: Returns a list of VM IDs that can attack the specified VM.
- **Method**: `GET`
- **Query Parameters**:
    - `vm_id` (required): The ID of the virtual machine to analyze.
- **Response**:
    - **200 OK**: JSON array of attacking VM IDs.
    - **400 Bad Request**: Missing `vm_id` query parameter.
    - **404 Not Found**: The specified VM is not found in the environment.
    - **500 Internal Server Error**: Failed to process the request.
- **Example**:
  ```bash
  curl "http://localhost/api/v1/attack?vm_id=vm-xxxxxx"
  ```
  **Response**:
  ```json
  ["vm-xxxxxxxxxx"]
  ```

### `/api/v1/stats`
- **Description**: Returns service statistics since the application startup.
- **Method**: `GET`
- **Response**:
    - **200 OK**: JSON object with statistics.
    - **500 Internal Server Error**: Failed to process the request.
- **Example**:
  ```bash
  curl "http://localhost/api/v1/stats"
  ```
  **Response**:
  ```json
  {
    "vm_count": 2,
    "request_count": 3,
    "average_request_time": 0.013522666666666665
  }
  ```

## How to Run

### Prerequisites
- **Go**: Ensure Go is installed (version 1.23.5+ recommended).
- **JSON Input File**: A valid JSON file describing the cloud environment.

### Usage
1. Unzip the archive and change CWD to the unpacked project:
   ```bash
   unzip attack-surface-service-v1.0.zip
   cd attack-surface-service-v1.0
   ```

2. Run the application:
   ```bash
   go run ./cmd/api/... /path/to/input.json
   ```

3. The service will start on port `80`. You can use Postman, cURL, or a browser to query the endpoints.

### Example
```bash
sudo go run ./cmd/api/... ./data/input-0.json
```

## Input File Format
The input JSON file must have the following structure:
```json
{
  "vms": [
    {
      "vm_id": "vm-xxxxxxx",
      "name": "example-vm",
      "tags": ["tag1", "tag2"]
    }
  ],
  "fw_rules": [
    {
      "fw_id": "fw-xxxxx",
      "source_tag": "tag1",
      "dest_tag": "tag2"
    }
  ]
}
```

## Logging
- Logs are printed to the console with timestamps, request details, and error messages.
- Example log:
  ```
    INFO    [2025-01-20 17:01:22.554] Server is working on port :80
    INFO    [2025-01-20 17:01:24.656] GET /api/v1/attack?vm_id=vm-xxxxxx
    INFO    [2025-01-20 17:01:24.656] GET /api/v1/attack?vm_id=vm-xxxxxx 200
    INFO    [2025-01-20 17:01:27.373] GET /api/v1/stats
    INFO    [2025-01-20 17:01:27.373] GET /api/v1/stats 200
  ```

## Future Improvements
- Add unit and integration tests.
- Implement more detailed error handling.
- Optimize performance for large datasets.

## Author
[Vitali Lapeka](mailto:lapeko88@gmail.com)

