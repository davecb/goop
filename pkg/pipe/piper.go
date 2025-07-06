/* 
 * piper: a program to use channels to mitigate a data-race.
 * 	After updating a struct, and IFF an atomic boolean is
 *	set, writes the struct to a pipe.
 *	That pipe goes to a goroutine that implements read-
 *	copy update.
 *	A select statement there mediates the race, by allowing
 *	either a read or write to the copy of the data, but
 *	not both.
 */


/* 
 * Example sequence:
 *	the handler program allows http requests on a port
 *	the path identifies what is to be read.
 *	when it's handled, it first turns on the variable,
 *	then issues a read on an "output" channel in the
 *	goroutine, which sleeps until something is written 
 *	to the "input" channel.  When that happens, the read 
 *	suceeds and the handler turns off the variable.
 *	The goroutine needs to get the output request, and
 *	give it a non-nil struct that arrived at the input
 *	channel... somehow. Probably by being handed a request
 *	on a "start now" channel that sets the variable to nil,
 *	and the output channel won't send a <nothing>.
 */

 // http handler

 // in-line data copier with variable

 // goroutine with three channels

/*
 * Over in cmd, a command-line program looks at the variable to read,
 * selects the right path, does an http request and pretty-prints the
 * json it gets back.  Then loops.
 */

