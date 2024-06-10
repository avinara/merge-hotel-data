# Merge Hotel Data 

This project merges hotel data from multiple suppliers, cleans it, and delivers it via an API endpoint.

### Getting Started

Follow these instructions to set up and run the project on your local machine.

### Prerequisites

- Make sure you have [Docker Desktop](https://www.docker.com/products/docker-desktop) installed and running.

### Installation

1. **Clone the repository**:

   ```git clone https://github.com/avinara/merge-hotel-data.git
      cd merge-hotel-data```

2. **Make sure Docker Desktop is running**:

 - Docker Desktop should be running to ensure the Docker daemon is switched on.

3. **Navigate to the project directory**:

            ```cd merge-hotel-data```

4. **Build the Docker image**:

            ```docker-compose build```

5.**Run the application**:

            ```docker-compose up```

### Usage
Once the application is running, you can access the API at http://localhost:8080.

### API Endpoints
            ```localhost:8080/hotels?searchStr=destination_id&searchValue=1122``` 

            ```localhost:8080/hotels?searchStr=id&searchValue=f8c9,iJhz```

### Query Params

    `searchStr =  "id" , searchValue =  "f8c9,iJhz" `

    `searchStr =  "destination_id", searchValue = "1122"`