package dynamock

type dynamock struct {
	ordered  bool
	requests []requestExpectation
}

type requestExpectation interface {
	complete bool
	err error
}

func (d *dynamock) UnOrderedExpectations(ordered bool) {
	d.ordered = ordered
}

func (d *dynamock) ExpectPutItemRequest(table string) {

}

func (d *dynamock) ExpectGetItemRequest(table string) {

}

func (d *dynamock) ExpectBatchWriteRequest(table string) {

}

func (d *dynamock) ExpectBatchGetRequest(table string) {

}

func (d *dynamock) ExpectCreateBackupRequest(table string) {

}

func (d *dynamock) ExpectCreateGlobalTableRequest(table string) {

}

func (d *dynamock) ExpectCreateTableRequest(table string) {

}

func (d *dynamock) ExpectDeleteBackupRequest(table string) {

}

func (d *dynamock) ExpectDeleteItemRequest(table string) {

}

func (d *dynamock) ExpectDeleteTableRequest(table string) {

}

func (d *dynamock) ExpectDescribeBackupRequest(table string) {

}

func (d *dynamock) ExpectDescribeContinuousBackupsRequest(table string) {

}

func (d *dynamock) ExpectDescribeEndpointsRequest(table string) {

}

func (d *dynamock) ExpectDescribeGlobalTableRequest(table string) {

}

func (d *dynamock) ExpectDescribeGlobalTableSettingsRequest(table string) {

}

func (d *dynamock) ExpectDescribeLimitsRequest(table string) {

}

func (d *dynamock) ExpectDescribeTableRequest(table string) {

}

func (d *dynamock) ExpectDescribeTimeToLiveRequest(table string) {

}

func (d *dynamock) ExpectListBackupsRequest(table string) {

}

func (d *dynamock) ExpectListGlobalTablesRequest(table string) {

}

func (d *dynamock) ExpectListTablesRequest(table string) {

}

func (d *dynamock) ExpectListTagsOfResourceRequest(table string) {

}

func (d *dynamock) ExpectQueryRequest(table string) {

}

func (d *dynamock) ExpectRestoreTableFromBackupRequest(table string) {

}

func (d *dynamock) ExpectRestoreTableToPointInTimeRequest(table string) {

}

func (d *dynamock) ExpectScanRequest(table string) {

}

func (d *dynamock) ExpectTagResourceRequest(table string) {

}

func (d *dynamock) ExpectTransactGetItemsRequest(table string) {

}

func (d *dynamock) ExpectTransactWriteItemsRequest(table string) {

}

func (d *dynamock) ExpectUntagResourceRequest(table string) {

}

func (d *dynamock) ExpectUpdateContinuousBackupsRequest(table string) {

}

func (d *dynamock) ExpectUpdateGlobalTableRequest(table string) {

}

func (d *dynamock) ExpectUpdateGlobalTableSettingsRequest(table string) {

}

func (d *dynamock) ExpectUpdateItemRequest(table string) {

}

func (d *dynamock) ExpectUpdateTableRequest(table string) {

}

func (d *dynamock) ExpectUpdateTimeToLiveRequest(table string) {

}

func (d *dynamock) ExpectWaitUntilTableExists(table string) {

}

func (d *dynamock) ExpectWaitUntilTableExistsWithContext(table string) {

}

func (d *dynamock) ExpectWaitUntilTableNotExists(table string) {

}

func (d *dynamock) ExpectWaitUntilTableNotExistsWithContext(table string) {

}