package parallel

import (
	"reflect"
	"sync"
)

// Group is a goroutine group
type Group struct {
	errs chan error
	wg   sync.WaitGroup
}

var (
	// Reflect Type of error, used to check function return parameters
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

// Go run a goroutine
// fun MUST be a function, if the last return paramter is an error, I'll be push to errs chan
// args are fun parameters, optional
func (g *Group) Go(fun interface{}, args ...interface{}) {

	v := reflect.ValueOf(fun)
	t := v.Type()

	// Check for function type
	if t.Kind() != reflect.Func {
		panic("bad function")
	}

	// Create WaitGroup & chan if needed
	if g.errs == nil {
		g.errs = make(chan error)
	}

	// Build input parameters
	var in []reflect.Value
	for _, a := range args {
		in = append(in, reflect.ValueOf(a))
	}

	// Register goroutine in WaitGroup
	g.wg.Add(1)

	// TODO: support panic ?
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		// Call it !
		out := v.Call(in)

		// Check if last return parameter is an errorType
		if n := t.NumOut(); n > 0 && t.Out(n-1) == errorType {
			if errv := out[n-1]; !errv.IsNil() {
				g.errs <- errv.Interface().(error)
			}
		}
	}(&g.wg)
}

// Wait block until all goroutine has finished
// return the last error if any
// it may skip error since it only return the last one, workaround ?
func (g *Group) Wait() error {

	// Local errors
	errs := make(chan error)
	go func() {
		// Read errs channel
		for e := range g.errs {
			errs <- e
		}
		close(errs)
	}()

	g.wg.Wait()
	close(g.errs)

	err := <-errs
	return err
}
