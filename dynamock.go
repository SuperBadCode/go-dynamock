package dynamock

type Dynamock struct {
	ordered  bool
	requests []requestExpectation
	errors []error
}

type requestExpectation interface{}

type ignoreField interface{}

func (d *Dynamock) UnOrderedExpectations(ordered bool) {
	d.ordered = ordered
}

func (d *Dynamock) ExpectPutItemRequest(table string) {

}

func (d *Dynamock) ExpectGetItemRequest(table string) {

}

func (d *Dynamock) ExpectBatchWriteRequest(table string) {

}

func (d *Dynamock) ExpectBatchGetRequest(table string) {

}

func (d *Dynamock) ExpectCreateBackupRequest(table string) {

}

func (d *Dynamock) ExpectCreateGlobalTableRequest(table string) {

}

func (d *Dynamock) ExpectCreateTableRequest(table string) {

}

func (d *Dynamock) ExpectDeleteBackupRequest(table string) {

}

func (d *Dynamock) ExpectDeleteItemRequest(table string) {

}

func (d *Dynamock) ExpectDeleteTableRequest(table string) {

}

func (d *Dynamock) ExpectDescribeBackupRequest(table string) {

}

func (d *Dynamock) ExpectDescribeContinuousBackupsRequest(table string) {

}

func (d *Dynamock) ExpectDescribeEndpointsRequest(table string) {

}

func (d *Dynamock) ExpectDescribeGlobalTableRequest(table string) {

}

func (d *Dynamock) ExpectDescribeGlobalTableSettingsRequest(table string) {

}

func (d *Dynamock) ExpectDescribeLimitsRequest(table string) {

}

func (d *Dynamock) ExpectDescribeTableRequest(table string) {

}

func (d *Dynamock) ExpectDescribeTimeToLiveRequest(table string) {

}

func (d *Dynamock) ExpectListBackupsRequest(table string) {

}

func (d *Dynamock) ExpectListGlobalTablesRequest(table string) {

}

func (d *Dynamock) ExpectListTablesRequest(table string) {

}

func (d *Dynamock) ExpectListTagsOfResourceRequest(table string) {

}

func (d *Dynamock) ExpectQueryRequest(table string) {

}

func (d *Dynamock) ExpectRestoreTableFromBackupRequest(table string) {

}

func (d *Dynamock) ExpectRestoreTableToPointInTimeRequest(table string) {

}

func (d *Dynamock) ExpectScanRequest(table string) {

}

func (d *Dynamock) ExpectTagResourceRequest(table string) {

}

func (d *Dynamock) ExpectTransactGetItemsRequest(table string) {

}

func (d *Dynamock) ExpectTransactWriteItemsRequest(table string) {

}

func (d *Dynamock) ExpectUntagResourceRequest(table string) {

}

func (d *Dynamock) ExpectUpdateContinuousBackupsRequest(table string) {

}

func (d *Dynamock) ExpectUpdateGlobalTableRequest(table string) {

}

func (d *Dynamock) ExpectUpdateGlobalTableSettingsRequest(table string) {

}

func (d *Dynamock) ExpectUpdateItemRequest(table string) {

}

func (d *Dynamock) ExpectUpdateTableRequest(table string) {

}

func (d *Dynamock) ExpectUpdateTimeToLiveRequest(table string) {

}

func (d *Dynamock) ExpectWaitUntilTableExists(table string) {

}

func (d *Dynamock) ExpectWaitUntilTableExistsWithContext(table string) {

}

func (d *Dynamock) ExpectWaitUntilTableNotExists(table string) {

}

func (d *Dynamock) ExpectWaitUntilTableNotExistsWithContext(table string) {

}
