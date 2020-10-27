package pleasantries

// Welcomer welcomes people.
type Welcomer interface {
	// Welcome makes a welcome message for somebody.
	Welcome(WelcomeRequest) WelcomeResponse
}

// WelcomeRequest is the request object for Welcomer.Welcome.
type WelcomeRequest struct {
	// To is the address of the person to send the message to.
	// META(example): "your@email.com"
	// META(featured): true
	To string
	// Name is the name of the person to welcome.
	// META(example): "John Smith"
	Name string
	// The number of times to send the message.
	// META(example): 3
	Times int
	// NewCustomer indicates whether this is a new customer
	// or not.
	// META(example): true
	NewCustomer bool
}

// WelcomeResponse is the response object for Welcomer.Welcome.
type WelcomeResponse struct {
	// Message is the welcome message.
	// example: "Welcome John Smith."
	Message string
}
