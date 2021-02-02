package public_private

// // AlertCounter is an exported type that
// // contains an integer counter for alerts.
// //(1) THIS IS PUBLIC VARIABLE
// type AlertCounter int

// //(2) THIS IS PRIVATE VARIABLE
// type alertCounter int

//(3) THIS IS PRIVATE VARIABLE
// BUT IT CAN ALSO ACCESS IN ANOTHER PACKAGE VIA PUBLIC METHOD IN THIS PACKAGE
type alertCounter int

// NewAlertCounter creates and returns objects of
// the unexported type alertCounter.
func NewAlertCounter(value int) alertCounter {
	return alertCounter(value)
}
