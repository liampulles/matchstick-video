package domain

// Runnable encapsulates logic that can just be
// run - it requires no further input or setup.
type Runnable func() error
