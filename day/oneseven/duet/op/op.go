package op

type Value struct {
	Integer *int     `  @Integer`
	Variable *string `| @Ident`
}

type Snd struct {
	V Value `"s" "n" "d" @@`
}

type Set struct {
	L string `"s" "e" "t" @Ident`
	V Value  `@@`
}

type Add struct {
	L string `"a" "d" "d" @Ident`
	V Value  `@@`
}

type Mul struct {
	L string `"m" "u" "l" @Ident`
	V Value  `@@`
}

type Mod struct {
	L string `"m" "o" "d" @Ident`
	V Value  `@@`
}

type Rcv struct {
	V Value `"r" "c" "v" @@`
}

type Jgz struct {
	C Value `"j" "g" "z" @@`
	O Value `@@`
}

type Op struct {
	Snd *Snd `  @@`
	Set *Set `| @@`
	Add *Add `| @@`
	Mul *Mul `| @@`
	Mod *Mod `| @@`
	Rcv *Rcv `| @@`
	Jgz *Jgz `| @@`
}

func (o Op) Op() interface{} {
	switch {
	case o.Snd != nil:
		return *o.Snd
	case o.Set != nil:
		return *o.Set
	case o.Add != nil:
		return *o.Add
	case o.Mul != nil:
		return *o.Mul
	case o.Mod != nil:
		return *o.Mod
	case o.Rcv != nil:
		return *o.Rcv
	case o.Jgz != nil:
		return *o.Jgz
	default:
		return nil
	}
}