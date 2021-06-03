package structproto

type NamedValues map[string]interface{}

func (values NamedValues) Iterate() <-chan KeyValuePair {
	c := make(chan KeyValuePair, 1)
	go func() {
		for k, v := range values {
			c <- KeyValuePair{k, v}
		}
		close(c)
	}()
	return c
}
