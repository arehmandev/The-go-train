package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

func AddGigasecond(mytime time.Time) time.Time {
	return mytime.Add(1E9 * time.Second)
}
