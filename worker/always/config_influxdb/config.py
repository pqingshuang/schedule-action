from dataclasses import dataclass, field
import os

@dataclass
class Config:
    algo_1_param_1: int = field(default=0)
    algo_1_param_2: int = field(default=0)

    algo_2_param_1: int = field(default=0)
    algo_2_param_2: int = field(default=0)


def load_config(*opts) -> Config:
    c = Config()

    for opt in opts:
        opt(c)

    return c

def load_algo_X_wrapper_opt(c: Config, version: int = 0) -> None:
    version = version or 0
    if version == 0: # default
        load_config_0(c)
    elif version == 1:
        load_config_1(c)
    elif version == 2:
        load_config_2(c)

def load_config_0(c: Config) -> None:
    print("loading default config")
    c.algo_1_param_1 = 11
    c.algo_1_param_2 = 12
    c.algo_2_param_1 = 21
    c.algo_2_param_2 = 22

def load_config_1(c: Config) -> None:
    print("loading config 1")
    c.algo_1_param_1 = 111
    c.algo_1_param_2 = 121
    c.algo_2_param_1 = 211
    c.algo_2_param_2 = 221

def load_config_2(c: Config) -> None:
    print("loading config 2")
    c.algo_1_param_1 = 112
    c.algo_1_param_2 = 122
    c.algo_2_param_1 = 212
    c.algo_2_param_2 = 222
