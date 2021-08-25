import service from '@/utils/request'

const mockData = {
    "code": 0,
    "data": {
        "baseinfo": {
            "app_ver": {
                "desc": "app版本",
                "editable": 0,
                "data": "1.3.4"
            },
            "is_new_user": {
                "desc": "是否为新用户",
                "editable": 0,
                "data": 1
            },
            "istester": {
                "desc": "是否为测试用户",
                "editable": 0,
                "data": 1
            },
            "is_fbuser": {
                "desc": "是否绑定faceback",
                "editable": 0,
                "data": ""
            },
            "npc_step": {
                "desc": "新手引导步数",
                "editable": 0,
                "data": 0
            },
            "vip_statue": {
                "desc": "vip状态",
                "editable": 1,
                "data": "SILVER"
            },
            "vip_score": {
                "desc": "vip积分",
                "editable": 1,
                "data": 1710
            },
            "paynum": {
                "desc": "支付数额",
                "editable": 1,
                "data": 6
            },
            "firstlogin": {
                "desc": "第一次登录时间",
                "editable": 0,
                "data": 1616665325
            },
            "lastlogin": {
                "desc": "上次登录时间",
                "editable": 0,
                "data": 1629196088
            },
            "lev": {
                "desc": "用户等级",
                "editable": 1,
                "data": 127
            },
            "levup_percent": {
                "desc": "当前经验值",
                "editable": 1,
                "data": 11.9791666667
            },
            "sum_bet": {
                "desc": "总下注",
                "editable": 0,
                "data": 14009400000
            },
            "sum_win": {
                "desc": "总赢钱",
                "editable": 0,
                "data": 124490106219
            },
            "bigwin_num": {
                "desc": "bigwin次数",
                "editable": 0,
                "data": 0
            },
            "map": {
                "desc": "地图相关",
                "editable": 0,
                "data": 0
            },
            "map_lev": {
                "desc": "地图等级",
                "editable": 0,
                "data": 5
            },
            "cur_map_themeid": {
                "desc": "当前所玩地图主题",
                "editable": 0,
                "data": "map_theme16"
            },
            "coins": {
                "desc": "用户金币",
                "editable": 1,
                "data": 399592234175
            },
            "daily_bonus": {
                "desc": "登录奖励",
                "editable": 1,
                "data": {
                    "collect": 1,
                    "date": "2021-08-17",
                    "days": 6,
                    "version": 2,
                    "week_day": 1
                }
            },
            "dynamic_bet": {
                "desc": "主题动态bet_id",
                "editable": 1,
                "data": 1000
            },
            "jackpot_grand": {
                "desc": "grand相关",
                "editable": 1,
                "data": {
                    "lev": 0
                }
            },
            "win_it_again": {
                "desc": "win it again",
                "editable": 1,
                "data": ""
            }
        },
        "themeinfo": {
            "theme15": {
                "desc": "财神",
                "spin_bet_GRAND": 0,
                "spin_bet_MAJOR": 15,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 30,
                "spin_bet_MINOR": 21
            },
            "theme17": {
                "desc": "金猪",
                "spin_bet_GRAND": 1000,
                "spin_bet_MAJOR": 150375,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 501250,
                "spin_bet_MINOR": 250625
            },
            "theme38": {
                "desc": "海岛度假",
                "fs_mode": "low",
                "spin_bet_GRAND": 1,
                "spin_bet_MAJOR": 10,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 0,
                "spin_bet_MINOR": 20
            },
            "theme41": {
                "desc": "迪斯科",
                "spin_bet_GRAND": 50,
                "spin_bet_MAJOR": 70,
                "spin_bet_MAXI": 50,
                "spin_bet_MINI": 2500,
                "spin_bet_MINOR": 1250
            },
            "theme59": {
                "desc": "小恶魔",
                "respin_actions": {},
                "spin_bet_GRAND": 3175,
                "spin_bet_MAJOR": 6350,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 12700,
                "spin_bet_MINOR": 9525
            },
            "theme62": {
                "desc": "火神",
                "spin_bet_GRAND": 661,
                "spin_bet_MAJOR": 3305,
                "spin_bet_MAXI": 1652,
                "spin_bet_MINI": 6610,
                "spin_bet_MINOR": 4957
            },
            "theme65": {
                "box_location": [
                    [
                        0,
                        0,
                        0,
                        0,
                        1,
                        0,
                        0,
                        0,
                        1
                    ],
                    [
                        0,
                        0,
                        0,
                        1,
                        0,
                        0,
                        0,
                        1,
                        0
                    ],
                    [
                        1,
                        0,
                        0,
                        1,
                        0,
                        0,
                        0,
                        0,
                        0
                    ],
                    [
                        1,
                        0,
                        0,
                        0,
                        1,
                        0,
                        0,
                        0,
                        0
                    ]
                ],
                "desc": "幸运熊猫",
                "fs_mode": "low",
                "has_over": 0,
                "spin_bet_GRAND": 6,
                "spin_bet_MAJOR": 15,
                "spin_bet_MINI": 30,
                "spin_bet_MINOR": 21
            },
            "theme68": {
                "desc": "企鹅",
                "spin_bet_GRAND": 481,
                "spin_bet_MAJOR": 965,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 1930,
                "spin_bet_MINOR": 1446
            },
            "theme69": {
                "desc": "太空",
                "jackpot_spin_bet_5rapidfire": 0,
                "jackpot_spin_bet_6rapidfire": 0,
                "jackpot_spin_bet_7rapidfire": 0,
                "jackpot_spin_bet_8rapidfire": 0,
                "jackpot_spin_bet_9rapidfire": 0
            },
            "theme70": {
                "desc": "非洲狮",
                "spin_bet_GRAND": 0,
                "spin_bet_MAJOR": 0,
                "spin_bet_MINI": 0,
                "spin_bet_MINOR": 0
            },
            "theme72": {
                "box_location": [
                    [
                        0,
                        0,
                        0,
                        0,
                        0,
                        1,
                        "fs",
                        0,
                        0
                    ],
                    [
                        2170000,
                        2170000,
                        "fs",
                        0,
                        0,
                        0,
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        1,
                        0,
                        0,
                        0,
                        0,
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        0,
                        0,
                        1,
                        0,
                        1,
                        0,
                        0,
                        0
                    ]
                ],
                "current_page": 1,
                "desc": "小红帽",
                "has_over": 4,
                "spin_bet_GRAND": 2635,
                "spin_bet_MAJOR": 15110,
                "spin_bet_MAXI": 7528,
                "spin_bet_MINI": 17680,
                "spin_bet_MINOR": 17348
            },
            "theme76": {
                "desc": "法老",
                "spin_bet_GRAND": 1355,
                "spin_bet_MAJOR": 140965,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 0,
                "spin_bet_MINOR": 211272
            },
            "theme77": {
                "desc": "金鱼",
                "spin_bet_GRAND": 7,
                "spin_bet_MAJOR": 3742,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 18420,
                "spin_bet_MINOR": 12894
            },
            "theme78": {
                "desc": "怪博士",
                "spin_bet_GRAND": 6,
                "spin_bet_MAJOR": 15,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 30,
                "spin_bet_MINOR": 21
            },
            "theme80": {
                "desc": "富豪",
                "rs_mode": "low",
                "spin_bet_GRAND": 10,
                "spin_bet_MAJOR": 40,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 50,
                "spin_bet_MINOR": 35
            },
            "theme81": {
                "desc": "美国鹰",
                "fs_mode": "high",
                "spin_bet_GRAND": 103778,
                "spin_bet_MAJOR": 206765,
                "spin_bet_MAXI": 207570,
                "spin_bet_MINI": 413530,
                "spin_bet_MINOR": 310108
            },
            "theme82": {
                "desc": "猛犸",
                "spin_bet_GRAND": 2,
                "spin_bet_MAJOR": 5,
                "spin_bet_MAXI": 5,
                "spin_bet_MINI": 10,
                "spin_bet_MINOR": 7
            },
            "theme84": {
                "desc": "糖果熊",
                "spin_bet_GRAND": 0,
                "spin_bet_MAJOR": 0,
                "spin_bet_MINOR": 0
            },
            "theme85": {
                "desc": "聚宝盆",
                "spin_bet_GRAND": 20002,
                "spin_bet_MAJOR": 400570,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 670060,
                "spin_bet_MINOR": 600750
            },
            "theme88": {
                "desc": "章鱼海盗",
                "spin_bet_GRAND": 0,
                "spin_bet_MAJOR": 20,
                "spin_bet_MAXI": 0,
                "spin_bet_MINI": 40,
                "spin_bet_MINOR": 28
            }
        },
        "activityinfo": {
            "daily_task_info": {
                "chest_bingo": 0,
                "chest_left": 0,
                "end_time": 1629208800,
                "misc_reward": {
                    "1002": [
                        0
                    ],
                    "2003": [
                        0
                    ],
                    "3001": [
                        0
                    ]
                },
                "super_task": {
                    "end_time": 1629802349,
                    "misc_reward": {
                        "5001": [
                            0
                        ],
                        "5002": [
                            0
                        ],
                        "5003": [
                            0,
                            2,
                            1,
                            3
                        ],
                        "5004": [
                            0
                        ],
                        "5005": [
                            0
                        ],
                        "5006": [
                            0,
                            3,
                            1
                        ],
                        "5007": [
                            0
                        ]
                    },
                    "task_group": [
                        5002,
                        5005,
                        5003,
                        5001,
                        5006,
                        5007,
                        5004
                    ],
                    "task_group_content": {
                        "5001": {
                            "coins": 2.7,
                            "desc": "Win {1} Coins in a single spin {0} times",
                            "ref": [
                                15,
                                24.89
                            ],
                            "type": 108,
                            "week_score": 81
                        },
                        "5002": {
                            "coins": 3.64,
                            "desc": "Bet {0} Coins",
                            "ref": [
                                495.83
                            ],
                            "type": 102,
                            "week_score": 51
                        },
                        "5003": {
                            "coins": 9.3,
                            "desc": "Collect {0} Lucky Cards of {1} Stars",
                            "ref": [
                                12,
                                4
                            ],
                            "type": 112,
                            "week_score": 83
                        },
                        "5004": {
                            "coins": 3.5,
                            "desc": "Win {0} Coins",
                            "ref": [
                                446.07
                            ],
                            "type": 104,
                            "week_score": 83
                        },
                        "5005": {
                            "coins": 4.82,
                            "desc": "Level up {0} times",
                            "ref": [
                                12
                            ],
                            "type": 106,
                            "week_score": 60
                        },
                        "5006": {
                            "coins": 3.21,
                            "desc": "Collect {0} Stars from Lucky Cards",
                            "ref": [
                                131
                            ],
                            "type": 110,
                            "week_score": 80
                        },
                        "5007": {
                            "coins": 4,
                            "desc": "Get {0} Big Wins",
                            "ref": [
                                35
                            ],
                            "type": 101,
                            "week_score": 62
                        }
                    },
                    "task_group_statue": {
                        "5001": 3,
                        "5002": 3,
                        "5003": 3,
                        "5004": 0,
                        "5005": 3,
                        "5006": 1,
                        "5007": 3
                    },
                    "task_id": 5006,
                    "task_statue": {
                        "card_stars": 7
                    }
                },
                "task_group": [
                    1002,
                    2003,
                    3001
                ],
                "task_group_content": {
                    "1002": {
                        "coins": 0.83,
                        "desc": "Get {0} Big Wins",
                        "ref": [
                            5
                        ],
                        "type": 101,
                        "week_score": 31
                    },
                    "2003": {
                        "coins": 1.24,
                        "desc": "Bet {0} Coins",
                        "ref": [
                            194.66
                        ],
                        "type": 102,
                        "week_score": 48
                    },
                    "3001": {
                        "coins": 3,
                        "desc": "Win {1} Coins in {0} spins",
                        "ref": [
                            100,
                            165
                        ],
                        "type": 105,
                        "week_score": 88
                    }
                },
                "task_group_statue": {
                    "1002": 1,
                    "2003": 0,
                    "3001": 0
                },
                "task_statue": {
                    "big_win_time": 3
                },
                "update_date": "2021-08-17",
                "update_lev": 116,
                "update_rb": 1500,
                "update_rb2": 200,
                "week_end_time": 1629640800,
                "week_reward": {
                    "free": {
                        "1": [
                            [
                                1,
                                3300000
                            ]
                        ],
                        "3": [
                            [
                                2,
                                0
                            ]
                        ],
                        "4": [
                            [
                                13,
                                2
                            ]
                        ],
                        "7": [
                            [
                                15,
                                4
                            ]
                        ],
                        "9": [
                            [
                                1,
                                10000000
                            ]
                        ],
                        "10": [
                            [
                                11,
                                5
                            ]
                        ]
                    },
                    "pay": {
                        "1": [
                            [
                                6,
                                12
                            ],
                            [
                                1,
                                7000000
                            ]
                        ],
                        "2": [
                            [
                                2,
                                0
                            ]
                        ],
                        "3": [
                            [
                                5,
                                12
                            ],
                            [
                                1,
                                9000000
                            ]
                        ],
                        "4": [
                            [
                                12,
                                3
                            ],
                            [
                                1,
                                11000000
                            ]
                        ],
                        "5": [
                            [
                                17,
                                0
                            ]
                        ],
                        "6": [
                            [
                                3,
                                12
                            ],
                            [
                                1,
                                13000000
                            ]
                        ],
                        "7": [
                            [
                                14,
                                1
                            ],
                            [
                                1,
                                16000000
                            ]
                        ],
                        "8": [
                            [
                                16,
                                0
                            ]
                        ],
                        "9": [
                            [
                                9,
                                5
                            ],
                            [
                                1,
                                20000000
                            ]
                        ],
                        "10": [
                            [
                                10,
                                1
                            ],
                            [
                                1,
                                30000000
                            ]
                        ],
                        "11": [
                            [
                                1,
                                15000000
                            ]
                        ]
                    }
                },
                "week_score": 0,
                "week_statue": {
                    "free": {
                        "1": 0,
                        "3": 0,
                        "4": 0,
                        "7": 0,
                        "9": 0,
                        "10": 0
                    },
                    "pay": {
                        "1": 0,
                        "2": 0,
                        "3": 0,
                        "4": 0,
                        "5": 0,
                        "6": 0,
                        "7": 0,
                        "8": 0,
                        "9": 0,
                        "10": 0,
                        "11": 0
                    }
                }
            },
            "world_tour_info": {
                "coins": 4103372764,
                "completeTimes": 0,
                "keys": 0,
                "keys_buy": 28,
                "keys_free": 2,
                "map_detail": [
                    [
                        0
                    ],
                    [
                        0
                    ],
                    [
                        0
                    ],
                    [
                        0
                    ]
                ],
                "map_energy": 27355818,
                "map_location": 1,
                "paynum": 0,
                "rand": 75,
                "sn": 1,
                "user_lev": 110,
                "xp": 37
            }
        }
    },
    "msg": "获取成功"
}
// 查询用户信息
export const getUserInfo = (data) => {
    return mockData
    // return service({
    //     url: "/rpc/userInfo",
    //     method: 'post',
    //     data: data
    // })
}

