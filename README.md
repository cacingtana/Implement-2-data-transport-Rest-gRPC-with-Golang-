# golang-test :star:
 2 transport rest http and grpc

## Data initialization
 using gorm as library to write log each search call with method below.
 
 Design
 - Explicitly separate User-Side, Business Logic, and Server-Side
 - Dependencies are going from User-Side and Server-Side to the Business Logic

## How To Consume The rest API http

	//Get all data movies
	METHOD GET: "/movies"
	
	//using queryParams
	example http://localhost:8000/movies?s=Batman&page=2
	return data movies with title Batman and rowPage 2
	
	//Get movies by ID
	METHOD GET: "/movies/id"
	
	//using queryParams
	example http://localhost:8000/movies/id?i=tt0106364
	return data movies with title ID tt0106364
	
