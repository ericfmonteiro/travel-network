# travel-network

Social network about travel tips.

This was an academic assignment, therefore it was not focused on clean code (specially in frontend)
I did not find a lot of examples of social networks in golang, so feel free to use it to have some examples of Golang code and suggest changes 

## How to run:

Run database (Must have Docker): 
`docker run --name postgres-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres# travel-network`

Start server:
`go run main.go`

Open frontend app with live server.
