package git

import (
	_ "github.com/libgit2/git2go/v28/git2"
	_ "github.com/libgit2/git2go/v28/git2/sys"
)

/*
#include <git2.h>
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR} -lgit2-darwin-amd64 -framework CoreFoundation -framework Security -lz -liconv
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR} -lgit2-darwin-arm64 -framework CoreFoundation -framework Security -lz -liconv
#cgo linux,amd64 LDFLAGS: -L${SRCDIR} -lgit2-linux-amd64 -lssl -lcrypto -lz -lpthread -ldl

#if LIBGIT2_VER_MAJOR != 0 || LIBGIT2_VER_MINOR != 28
# error "Invalid libgit2 version; this git2go supports libgit2 v0.28"
#endif

*/
import "C"
