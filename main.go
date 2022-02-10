package main

import (
	"github.com/rudderlabs/analytics-go"
)

func main() {
	// Instantiates a client to use send messages to the Rudder API.
	/*client := analytics.New("24u1IBzLs4qz7e6Jfb7VYkvL82Z", "http://localhost:8080")

	// Enqueues a track event that will be sent asynchronously.
	client.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "test-snippet",
	})

	// Flushes any queued messages and closes the client.
	client.Close()

	fmt.Println("end")*/

	clientCloud := analytics.New("24rFakUrc9PSv0skP14kOtcmkIM", "http://localhost:8081")
	// Enqueues a track event that will be sent asynchronously.
	clientCloud.Enqueue(analytics.Track{
		UserId: "test-user-cloud",
		Event:  "test-snippet-cloud",
	})

	// Flushes any queued messages and closes the client.
	clientCloud.Close()
}
