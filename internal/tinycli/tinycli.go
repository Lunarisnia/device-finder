package tinycli

type Program func(ctx Context) error

type app struct {
	ctx     Context
	program Program
}

func New() *app {
	return &app{
		ctx: newContext(),
		program: func(ctx Context) error {
			return nil
		},
	}
}

func (a *app) Argument(arg string) string {
	return a.ctx.Argument(arg)
}

func (a *app) Arguments() map[string]string {
	return a.ctx.Arguments()
}

func (a *app) SetProgram(program Program) {
	a.program = program
}

func (a *app) Run() error {
	return a.program(a.ctx)
}
