package foobarbaz

import "github.com/google/wire"

var SuperSet = wire.NewSet( ProviderBar, ProvideBaz,ProvideFoo)
