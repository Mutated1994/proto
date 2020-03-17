// Copyright (c) 2018 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package proto

import "context"

// Handler is a type of function that accepts a Visitee.
type Handler func(ctx context.Context, v Visitee)

// Walk recursively pays a visit to all Visitees of a Proto and calls each handler with it.
func Walk(ctx context.Context, proto *Proto, handlers ...Handler) {
	walk(ctx, proto, handlers...)
}

func walk(ctx context.Context, container elementContainer, handlers ...Handler) {
	for _, eachElement := range container.elements() {
		for _, eachFilter := range handlers {
			eachFilter(ctx, eachElement)
		}
		if next, ok := eachElement.(elementContainer); ok {
			walk(ctx, next, handlers...)
		}
	}
}

// WithImport returns a Handler that will call the apply function when the Visitee is an Import.
func WithImport(apply func(context.Context, *Import)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Import); ok {
			apply(ctx, s)
		}
	}
}

// WithMessage returns a Handler that will call the apply function when the Visitee is a Message.
func WithMessage(apply func(context.Context, *Message)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Message); ok {
			apply(ctx, s)
		}
	}
}

// WithOption returns a Handler that will call the apply function when the Visitee is a Option.
func WithOption(apply func(context.Context, *Option)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Option); ok {
			apply(ctx, s)
		}
	}
}

// WithEnum returns a Handler that will call the apply function when the Visitee is a Enum.
func WithEnum(apply func(context.Context, *Enum)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Enum); ok {
			apply(ctx, s)
		}
	}
}

// WithOneof returns a Handler that will call the apply function when the Visitee is a Oneof.
func WithOneof(apply func(context.Context, *Oneof)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Oneof); ok {
			apply(ctx, s)
		}
	}
}

// WithService returns a Handler that will call the apply function when the Visitee is a Service.
func WithService(apply func(context.Context, *Service)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Service); ok {
			apply(ctx, s)
		}
	}
}

// WithRPC returns a Handler that will call the apply function when the Visitee is a RPC.
func WithRPC(apply func(context.Context, *RPC)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*RPC); ok {
			apply(ctx, s)
		}
	}
}

// WithPackage returns a Handler that will call the apply function when the Visitee is a Package.
func WithPackage(apply func(context.Context, *Package)) Handler {
	return func(ctx context.Context, v Visitee) {
		if s, ok := v.(*Package); ok {
			apply(ctx, s)
		}
	}
}
