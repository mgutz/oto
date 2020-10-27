package pleasantries

import (
	"github.com/pacedotdev/oto/testdata/services"
)

// GreeterService is a polite API.
// You will love it.
// META(strapline): "A lovely greeter service"
type GreeterService interface {
	// Greet creates a Greeting for one or more people.
	// META(featured): true
	Greet(GreetRequest) GreetResponse
	// GetGreetings gets a range of saved Greetings.
	// META(featured): false
	GetGreetings(GetGreetingsRequest) GetGreetingsResponse
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Names are the names of the people to greet.
	// example: ["Mat", "David"]
	Names []string
}

// GreetResponse is the response object containing a
// person's greeting.
type GreetResponse struct {
	// Greeting is the greeted person's Greeting.
	Greeting Greeting
}

// GetGreetingsRequest is the request object for GreeterService.GetGreetings.
// META(featured): true
type GetGreetingsRequest struct {
	// Page describes which page of data to get.
	Page services.Page `tagtest:"value,option1,option2"`
}

// GetGreetingsResponse is the response object for GreeterService.GetGreetings.
// META(featured): false
type GetGreetingsResponse struct {
	Greetings []Greeting
}

// Greeting contains the pleasentry.
type Greeting struct {
	// Text is the message.
	// example: "Hello there"
	Text string
}
