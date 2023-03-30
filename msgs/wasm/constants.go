package wasm

const (
	TEXT_FILE_PATH     = "./msgs/wasm/files/random.wasm"
	WASM_FILE_PATH     = "./msgs/wasm/files/test_sc.wasm"
	MIGRATE_TO_CODE_ID = 6330
	NEW_ADMIN_ADDRESS  = "osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p"
)

const (
	InitMsg           = `{"count": 0}`
	ExecuteMsgInvalid = `{"decrement": {}}`
	ExecuteMsg        = `{"increment": {}}`
	MigrateMsg        = `{}`
)
