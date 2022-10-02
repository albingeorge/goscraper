package custom

type Custom interface {
	Run(input interface{}) (error, interface{})
}
