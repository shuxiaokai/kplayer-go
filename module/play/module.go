package play

import (
    "encoding/json"
    "github.com/bytelang/kplayer/module"
    "github.com/bytelang/kplayer/module/play/provider"
    kptypes "github.com/bytelang/kplayer/types"
    "github.com/bytelang/kplayer/types/config"
    "github.com/spf13/cobra"
)


type AppModule struct {
    *provider.Provider
}

var _ module.AppModule = &AppModule{}

func NewAppModule() AppModule {
    return AppModule{provider.NewProvider()}
}

func (m AppModule) GetModuleName() string {
    return provider.ModuleName
}

func (m AppModule) GetCommand() *cobra.Command {
    return provider.GetCommand()
}

func (m AppModule) InitConfig(ctx kptypes.ClientContext, data json.RawMessage) {
    var cfg config.Play
    if err := json.Unmarshal(data, &cfg); err != nil {
        panic(err)
    }

    m.InitModuleConfig(ctx, cfg)
}
