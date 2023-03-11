// Package cookies is used to manage cookies, when we have determined cookies are stale
// we need to get updated ones by notifying the team in Discord or by SMS as it is urgent.
// This will prevent the crawler from getting stuck in a loop and will not be able to continue
// with the rest of the invites in the queue.
package cookies

// TODO: implement

// Service represents a service that can be used to store and retrieve cookies
type Service interface {
}
