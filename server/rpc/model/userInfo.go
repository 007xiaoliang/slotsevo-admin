package model

type RpcUserInfo struct {
	BaseInfo     BaseInfo               `json:"baseinfo"`
	ThemeInfo    map[string]interface{} `json:"themeinfo"`
	ActivityInfo map[string]interface{} `json:"activityinfo"`
}
type Column struct {
	Desc     string      `json:"desc"`
	Editable int         `json:"editable"`
	Data     interface{} `json:"data"`
}
type BaseInfo struct {
	AppVer        Column `json:"app_ver"`
	IsNewUser     Column `json:"is_new_user"`
	IsTester      Column `json:"istester"`
	IsFbUser      Column `json:"is_fbuser"`
	NpcStep       Column `json:"npc_step"`
	VipStatue     Column `json:"vip_statue"`
	VipScore      Column `json:"vip_score"`
	PayNum        Column `json:"paynum"`
	FirstLogin    Column `json:"firstlogin"`
	LastLogin     Column `json:"lastlogin"`
	Lev           Column `json:"lev"`
	LevupPercent  Column `json:"levup_percent"`
	SumBet        Column `json:"sum_bet"`
	SumWin        Column `json:"sum_win"`
	BigwinNum     Column `json:"bigwin_num"`
	Map           Column `json:"map"`
	MapLev        Column `json:"map_lev"`
	CurMapThemeid Column `json:"cur_map_themeid"`
	Coins         Column `json:"coins"`
	DailyBonus    Column `json:"daily_bonus"`
	DynamicBet    Column `json:"dynamic_bet"`
	JackpotGrand  Column `json:"jackpot_grand"`
	WinItAgain    Column `json:"win_it_again"`
}
//
//type ThemeInfo struct {
//}
//
//type ActivityInfo struct {
//}
