

/* 
 * racer: a program to use retry to mitigate a data-race.
 *	I could use a lock or a channel, but both slow down
 *	the producer and consumer. 
 *	Instead, I use read-copy-update to keep the data 
 *	stable, and the same retry-until-stable read as was
 *	used in an error-tolerant UofT unix filesystem.
 *
 *	The latter is just reread a few times, until we get two
 *	identical structs.  This works best if the struct is read-
 *	mostly, but it races a lot if the struct is constantly
 *	written to. It can still work, but by doing a lot of tries.
 *
 *	The number of tries is a parameter, usually three, but it
 *	can be decreased or increased for pretty-satic data or 
 *	constantly-changing data. The higher the tries, the more
 *	time that can be spent in the routine
 */

func racer(func grab, int tries) error {
	 grab()
	 for i := 0; i < limit; i++ {
		if grab() {
			return null // non-error
		}
	}
	return err
}

/*
 * requires a grab and a compare fucntion with local storage.
 * to get data out, requires a third fucntion that return the data
 * in the form of a json string.
 */
type interestingType struct {}
var rcuVariable interestingType // empty

func grab() string {
	var grabResult interestingType

	grabResult = someVar
	if grabResult == rcuVariable {
		s, err := jsonify()
		if err != nil {
			panic()
		}
		return s
	}
	rcuVariable = grabResult
	return false
}

func jsonify() (string, error) {
	x, err := json.Marshal(rcuVariable)
    	if err != nil {
    		panic()
        	//return err, ""
	}
	return string(x),  nil
}
