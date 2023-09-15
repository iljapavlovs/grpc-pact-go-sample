

### Step 1: Install the Protobuf Plugin

Once you have installed the Pact CLI tools, you can install the plugin by its registered name (protobuf), and it will be automatically discovered and installed into the relevant location on your machine:

```bash
pact-plugin-cli -y install protobuf
```




## Consumer 


1. Start Pact Broker
```
docker-compose up
```

2. Run Tests
```bash
go test -count=1 github.com/iljapavlovs/grpc-pact-go-sample/pact-sample -run 'TestGrpcInteraction'
```

3. Publish Pact contracts to Pact Broker
* download cli

```bash
curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | PACT_CLI_VERSION=v2.0.5 bash
```

* Publish the contract to Pact broker
```
pact-broker \
publish /Users/iljapavlovs/Desktop/PoC/Go/gRPC/grpc-go-sample/pacts/grpcconsumer-grpcprovider.json \
--consumer-app-version 4ac729c \
--branch main \
--broker-base-url http://localhost:9292 
```


4. [Can I Delpoy?](https://docs.pact.io/pact_broker/can_i_deploy) - get the latest result of the contract test verification
```
docker run --rm --network host \
  	-e PACT_BROKER_BASE_URL=http://localhost:9292 \
  	pactfoundation/pact-cli:latest \
  	broker can-i-deploy \
  	--pacticipant grpcconsumer \
  	--latest
```


```
docker run --rm --network host \
  	-e PACT_BROKER_BASE_URL=http://localhost:9292 \
  	pactfoundation/pact-cli:latest \
  	broker can-i-deploy \
  	--pacticipant grpcconsumer --version 4ac729c\
  	--pacticipant grpcprovider --latest "1.0.1"
```




5. After deployment
```bash
./pact-broker record-deployment \
--pacticipant grpcconsumer \
--version 1 \
--environment production \
--broker-base-url http://localhost:9292
```



## Provider

### On Provider side
1. Run Test
* Pact will automatically download the pact (contract in json format)
* When test is executed, Pact will send results to Pact Broker

 ```bash
go test -count=1 github.com/iljapavlovs/grpc-pact-go-sample/pact-sample -run "TestGrpcProvider"
```

