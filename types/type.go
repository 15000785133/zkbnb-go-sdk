package types

const (
	CodeOK = 100
)

type Result struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type Status struct {
	Status    uint32 `json:"status"`
	NetworkId uint32 `json:"network_id"`
}

type AccountAsset struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
	Price   string `json:"price"`
}

type Account struct {
	Status          uint32          `json:"status"`
	Index           int64           `json:"index"`
	L1Address       string          `json:"l1_address"`
	Name            string          `json:"name"`
	Pk              string          `json:"pk"`
	Nonce           int64           `json:"nonce"`
	Assets          []*AccountAsset `json:"assets"`
	TotalAssetValue string          `json:"total_asset_value"`
}

type SimpleAccount struct {
	Index     int64  `json:"index"`
	L1Address string `json:"l1_address"`
	Pk        string `json:"pk"`
}

type Accounts struct {
	Total    uint32           `json:"total"`
	Accounts []*SimpleAccount `json:"accounts"`
}

type Asset struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Decimals   uint32 `json:"decimals"`
	Symbol     string `json:"symbol"`
	Address    string `json:"address"`
	Price      string `json:"price"`
	IsGasAsset uint32 `json:"is_gas_asset"`
	Icon       string `json:"icon"`
}

type Assets struct {
	Total  uint32   `json:"total"`
	Assets []*Asset `json:"assets"`
}

type ProtocolRate struct {
	ProtocolRate string `json:"protocol_rate"`
}

type Block struct {
	Commitment                      string `json:"commitment"`
	Height                          int64  `json:"height"`
	StateRoot                       string `json:"state_root"`
	PriorityOperations              int64  `json:"priority_operations"`
	PendingOnChainOperationsHash    string `json:"pending_on_chain_operations_hash"`
	PendingOnChainOperationsPubData string `json:"pending_on_chain_operations_pub_data"`
	CommittedTxHash                 string `json:"committed_tx_hash"`
	CommittedAt                     int64  `json:"committed_at"`
	VerifiedTxHash                  string `json:"verified_tx_hash"`
	VerifiedAt                      int64  `json:"verified_at"`
	Txs                             []*Tx  `json:"txs"`
	Status                          int64  `json:"status"`
	Size                            uint16 `json:"size"`
}

type Blocks struct {
	Total  uint32   `json:"total"`
	Blocks []*Block `json:"blocks"`
}

type CurrentHeight struct {
	Height int64 `json:"height"`
}

type ContractAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type SignBody struct {
	SignBody string `json:"sign_body"`
}

type Layer2BasicInfo struct {
	BlockCommitted            int64             `json:"block_committed"`
	BlockVerified             int64             `json:"block_verified"`
	TotalTransactionCount     int64             `json:"total_transaction_count"`
	YesterdayTransactionCount int64             `json:"yesterday_transaction_count"`
	TodayTransactionCount     int64             `json:"today_transaction_count"`
	YesterdayActiveUserCount  int64             `json:"yesterday_active_user_count"`
	TodayActiveUserCount      int64             `json:"today_active_user_count"`
	ContractAddresses         []ContractAddress `json:"contract_addresses"`
}

type GasFee struct {
	GasFee string `json:"gas_fee"`
}

type GasAccount struct {
	Status    int64  `json:"status"`
	Index     int64  `json:"index"`
	L1Address string `json:"l1_address"`
}

type GasFeeAssets struct {
	Assets []Asset `json:"assets"`
}

type Search struct {
	DataType int32 `json:"data_type"`
}

type NftIndex struct {
	Index             int64  `json:"index"`
	IpfsId            string `json:"ipfs_id"`
	IpnsId            string `json:"ipns_id"`
	Metadata          string `json:"metadata"`
	MutableAttributes string `json:"mutable_attributes"`
}

type MaxCollectionId struct {
	CollectionId uint64 `json:"collection_id"`
}

type Mutable struct {
	IpnsId string `json:"ipns_id"`
}

type Tx struct {
	Hash             string `json:"hash"`
	Type             int64  `json:"type,range=[1:64]"`
	Amount           string `json:"amount"`
	Info             string `json:"info"`
	Status           int64  `json:"status"`
	Index            int64  `json:"index"`
	GasFeeAssetId    int64  `json:"gas_fee_asset_id"`
	GasFee           string `json:"gas_fee"`
	NftIndex         int64  `json:"nft_index"`
	CollectionId     int64  `json:"collection_id"`
	AssetId          int64  `json:"asset_id"`
	AssetName        string `json:"asset_name"`
	NativeAddress    string `json:"native_address"`
	ExtraInfo        string `json:"extra_info"`
	Memo             string `json:"memo"`
	AccountIndex     int64  `json:"account_index"`
	L1Address        string `json:"l1_address"`
	Nonce            int64  `json:"nonce"`
	ExpiredAt        int64  `json:"expire_at"`
	BlockHeight      int64  `json:"block_height"`
	CreatedAt        int64  `json:"created_at"`
	StateRoot        string `json:"state_root"`
	FromAccountIndex int64  `json:"from_account_index"`
	FromL1Address    string `json:"from_l1_address"`
	ToAccountIndex   int64  `json:"to_account_index"`
	ToL1Address      string `json:"to_l1_address"`
}

type Txs struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type TxHash struct {
	TxHash string `json:"tx_hash"`
}

type NextNonce struct {
	Nonce uint64 `json:"nonce"`
}

type EnrichedTx struct {
	Tx
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
	AssetAId    int64 `json:"asset_a_id"`
	AssetBId    int64 `json:"asset_b_id"`
}

type MaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
}

type Nft struct {
	Index               int64  `json:"index"`
	CreatorAccountIndex int64  `json:"creator_account_index"`
	CreatorL1Address    string `json:"creator_l1_address"`
	OwnerAccountIndex   int64  `json:"owner_account_index"`
	OwnerL1Address      string `json:"owner_l1_address"`
	ContentHash         string `json:"content_hash"`
	NftContentType      int64  `json:"nft_content_type"`
	L1Address           string `json:"l1_address"`
	L1TokenId           string `json:"l1_token_id"`
	RoyaltyRate         int64  `json:"royalty_rate"`
	CollectionId        int64  `json:"collection_id"`
	IpfsId              string `json:"ipfs_id"`
	IpnsId              string `json:"ipns_id"`
	Metadata            string `json:"metadata"`
	MutableAttributes   string `json:"mutable_attributes"`
}

type NftEntity struct {
	Nft *Nft `json:"nft"`
}

type Nfts struct {
	Total int64  `json:"total"`
	Nfts  []*Nft `json:"nfts"`
}

type Rollback struct {
	FromBlockHeight int64  `json:"from_block_height"`
	FromTxHash      string `json:"from_tx_hash"`
	ID              uint   `json:"id"`
	CreatedAt       int64  `json:"created_at"`
}

type ReqGetRollbacks struct {
	FromBlockHeight int64  `form:"from_block_height"`
	Offset          uint16 `form:"offset,range=[0:100000]"`
	Limit           uint16 `form:"limit,range=[1:100]"`
}

type Rollbacks struct {
	Total     uint32      `json:"total"`
	Rollbacks []*Rollback `json:"rollbacks"`
}

type ReqGetNft struct {
	NftIndex int64 `form:"nft_index"`
}
