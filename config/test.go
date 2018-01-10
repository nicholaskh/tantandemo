package config

import (
    conf "github.com/nicholaskh/jsconf"
)

type ConfigTest struct {
    ListenAddr  string
}

func (this *ConfigTest) LoadConfig(cf *conf.Conf) {
    this.ListenAddr = cf.String("listen_addr", ":8080")
}
