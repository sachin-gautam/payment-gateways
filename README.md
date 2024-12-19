# Payment Gateway Integration Assessment

This task implements a payment gateway system with two main endpoints.

## Endpoints

- `/deposit`: For processing deposit transactions.
- `/withdrawal`: For processing withdrawal transactions.
- `/swagger/`: For API documentation

In addition to these endpoints the api handles callback responses from gateways asynchronously to update the transaction status.

### Database

Postgres database has been implemented for the assessment which runs in the postgres container in docker.

### Features

1. **Endpoints Implementation:**
    - `/deposit` and `/withdrawal` endpoints to process transactions.
    - Each endpoint accept a JSON payload with details such as `user_id`, `gateway_id`, and `amount`.
    
2. **Callback Handling:**
    - Handle the callback from third-party gateways to update the transaction status asynchronously using go-routines.
    - The callback includes information like transaction status and is used to update the corresponding transaction in the database.
    
3. **Transaction Status:**
    - Each transactions includes a status field ("pending", "success", "failed") which should be updated when the callback is received.

4. **Fault Tolerance:**
    - Fault tolerance with retry mechanisms for database connection.

7. **Security:**
    - Masking the transaction amount using base64 encoding.
    
### How to Get Started

1. **Clone the Repository:**
    Clone the repository to your local machine:

    ```bash
    git clone [<repository_url>]
    cd <project_directory>
    ```

2. **Setup Docker:**
    Docker is configured to run PostgreSQL, Kafka, and Redis. Use the following command to start all the services:

    ```bash
    docker compose up --build
    ```

    This will start:
    - PostgreSQL on port `5433`
    - Application on port `8080`

3. **Database Migration:**
    The migration file `db/init.sql` is already provided. Once the Docker services are up and running, the database will be initialized automatically, and the tables will be created.






---


