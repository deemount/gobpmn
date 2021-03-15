package utils

import "github.com/deemount/gobpmn/repository"

func N(n int64) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.N = n
		return o
	}
}

func Counter(c int64) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.Counter = c
		return o
	}
}

func DefHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.DefHash = h
		return o
	}
}

func ProcHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.ProcHash = h
		return o
	}
}

func FlowHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.FlowHash = h
		return o
	}
}

func EventHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.EventHash = h
		return o
	}
}

func PartHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.ParticipantHash = h
		return o
	}
}

func CollaboHash(h string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.CollaboHash = h
		return o
	}
}

func CVersionTag(c string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.CVersionTag = c
		return o
	}
}

func Name(n string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.Name = n
		return o
	}
}

func FormKey(f string) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.FormKey = f
		return o
	}
}

func HasStartEvent(b bool) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.HasStartEvent = b
		return o
	}
}

func HasFirstState(b bool) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.HasFirstState = b
		return o
	}
}

func HasCollab(b bool) repository.BPMNFOption {
	return func(o repository.Options) repository.Options {
		o.HasCollab = b
		return o
	}
}
