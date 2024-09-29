IMAGE_NAME=application-name

fmt:
	go fmt ./...

mock:
	go generate -v ./...

update-module:
	go mod tidy

execute_GET: 
	curl --request GET --url http://localhost:8000/order --header 'User-Agent: insomnia/10.0.0'

execute_POST: 
	curl --request POST --url http://localhost:8000/order --header 'Content-Type: application/json' --header 'User-Agent: insomnia/10.0.0' \
  	--data '{"id":"order 1", "price": 10.5, "tax": 0.9 }'

	curl --request POST --url http://localhost:8000/order --header 'Content-Type: application/json' --header 'User-Agent: insomnia/10.0.0' \
  	--data '{"id":"order 2", "price": 1.5, "tax": 2.0 }'

	curl --request POST --url http://localhost:8000/order --header 'Content-Type: application/json' --header 'User-Agent: insomnia/10.0.0' \
  	--data '{"id":"order 3", "price": 100.5, "tax": 20.0 }'

	curl --request POST --url http://localhost:8000/order --header 'Content-Type: application/json' --header 'User-Agent: insomnia/10.0.0' \
  	--data '{"id":"order 4", "price": 2.5, "tax": 3.0 }'

	curl --request POST --url http://localhost:8000/order --header 'Content-Type: application/json' --header 'User-Agent: insomnia/10.0.0' \
  	--data '{"id":"order 5", "price": 10.9, "tax": 4.0 }'