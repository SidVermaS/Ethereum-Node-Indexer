# Ethereum-Node-Indexer
This application is an indexer for Ethereum's Consensus Layer (Beacon Chain). We can fetch the network's participation rate, an individual validator's participation rate and fetch epochs, validators, validators' status, slots, committees, blocks and states.

## Clone the project
```
git clone http://github.com/SidVermaS/Ethereum-Node-Indexer
```

## Project Setup
In order to run and test the application, there are some software and tools that are needed to be installed. (You can skip if you already have the mentioned software).

1. Download and install Docker desktop according to your operating system from the following link: 
https://www.docker.com/products/docker-desktop/
2. Download and install Visual Studio. It is needed to test the APIs. Download it from the following link:
https://code.visualstudio.com/download
3. Open the **Ethereum-Node-Indexer** repository in vsCode.
4. In order to run the Dockerized application you need to execute the following command
```cmd
docker-compose up
```

## Explore the APIs
Once the docker image is completely deployed then you can start exploring the APIs.
1. Select the extensions icon on the left side bar of vsCode, search "*thunder client*" and install it.

![thunderclient1](./meta/screenshots/thunderclient1.png)

2. A thunder client icon will be visible on the left side bar. You can select it.

![thunderclient2](./meta/screenshots/thunderclient2.png)

3. Select the options icon on top right side of the thunder client side bar. Select **Extension Settings** from the list of options.

![thunderclient3](./meta/screenshots/thunderclient3.png)

4. Check the checkbox for *Thunder*-client: **Save To Workspace**.

![thunderclient4](./meta/screenshots/thunderclient4.png)

5. Ensure that the selected current environment is EthereumNodeIndexer environment. 

![thunderclient5](./meta/screenshots/thunderclient5.png)

6. You can open the collections and start testing the APIs.

![thunderclient6](./meta/screenshots/thunderclient6.png)

## Testing
1. Open Docker desktop application and select the **Containers** option on the left sidebar.

![gotest1](./meta/screenshots/gotest1.png)

2. Select the command line icon of the container with the name: **eni_server**.

![gotest2](./meta/screenshots/gotest2.png)

3. Execute the following command:
```cmd
go test tests/main.go -v
```

![gotest3](./meta/screenshots/gotest3.png)

## Technical Justification
### i. Database
1. We are using *PostgreSQL* database because it's an object relational database. The data that we are saving needs a fixed structured format.
3. It performs better while executing complex and resource-intensive queries.
4. It also supports JSONB data types which provides index support which allows to efficiently query the JSON documents.

### ii. Cache
1. We are using *Redis* cache, it's an in-memory database. The data retrieval latency is very low because it stores the data in RAM which allows extremely fast read and write operations
2. In our application, we are using it for fetching the network's participation rate because it was taking a lot of time to calculate it separately for every request. 
3. We first check the cache for data before going to the backend storage, once we get the data then we save it in our backend.
4. We are updating the cache whenever the data in our database changes, so that we can return the latest data.