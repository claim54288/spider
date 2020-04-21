package rpcdemo

import "github.com/pkg/errors"

//Service.Method 来调用

type DemoService struct{}

type Args struct {
	A, B int
}

//这边是rpc框架要求，参数是两个，其中第二个一定要是指针,第一个无所谓，返回值是一个error
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero ")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
