package storage

import (
	"log"

	"gitee.com/sailinst/pallet_dag/palletdb"
	"github.com/palletone/go-palletone/dag/config"
)

var (
	PACKET_PREFIX                    = []byte("p") // packet_prefix  + mci + hash
	UNIT_PREFIX                      = []byte("u") // unit_prefix + mci + hash
	ALL_UNITS_PREFIX                 = []byte("au")
	UNITAUTHORS_PREFIX               = []byte("ua")
	SKIPLIST_PREFIX                  = []byte("sl")
	HASH_TREE_BALLS_PREFIX           = []byte("ht")
	PARENTHOODS_PREFIX               = []byte("pt")
	UNIT_WITNESS_PREFIX              = []byte("uw")
	WITNESS_LIST_HASHES_PREFIX       = []byte("wl")
	DEFINITIONS_PREFIX               = []byte("de")
	ADDRESS_PREFIX                   = []byte("ad")
	ADDRESS_DEFINITION_CHANGE_PREFIX = []byte("ac")
	AUTHENTIFIERS_PREFIX             = []byte("au")
	MESSAGES_PREFIX                  = []byte("me")
	POLL_PREFIX                      = []byte("po")
	VOTE_PREFIX                      = []byte("vo")
	ATTESTATION_PREFIX               = []byte("at")
	ASSET_PREFIX                     = []byte("as")
	ASSET_ATTESTORS                  = []byte("ae")
	DATA_FEED_PREFIX                 = []byte("df")
	SPEED_PROOFS                     = []byte("sp")
	EAENED_HEADERS_COMMISSION        = "earned_headers_commossion"
	ALL_UNITS                        = "array_units"
)

func init() {
	var err error
	if Dbconn == nil {
		if config.DConfig.DbPath == "" {
			config.DConfig.DbPath = DBPath
		}
		Dbconn, err = palletdb.NewLDBDatabase(config.DConfig.DbPath, 0, 0)
		if err != nil {
			log.Println("new dbconn error:", err)
		}
		log.Println("db_n:", Dbconn.Path())
	}
}
func ReNewDbConn() *palletdb.LDBDatabase {
	log.Println("renew dbconn start...")
	if config.DConfig.DbPath == "" {
		config.DConfig.DbPath = DBPath
	}
	if dbconn, err := palletdb.NewLDBDatabase(config.DConfig.DbPath, 0, 0); err != nil {
		log.Println("renew dbconn error:", err)
		return nil
	} else {
		return dbconn
	}
}