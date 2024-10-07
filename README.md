# Transactions Listener Service
This service listens to the transactions from the source chain to destination chain which specified in the `.env` file and prints the transactions params to the console.

# Running the service
To run the service, you need to have the following installed:
- Go 1.22 or later

Then, you can run the service by running the following command:
```bash
git clone github.com/neor-it/transactions-listener
cd transactions-listener
go install
```

To run the service, you need to create a `.env` file in the directory of the service (e.g. `cmd/tx_listener`) from the `.env.example` file and set the environment variables in the `.env` file.

Then, you can run the service by running the following command:
```bash
cd cmd/tx_listener
go run tx_listener_main.go
```
