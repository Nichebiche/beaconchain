package handlers

import (
	"net/http"

	types "github.com/gobitfly/beaconchain/pkg/api/types"

	"github.com/gorilla/mux"
)

// All handler function names must include the HTTP method and the path they handle
// Internal handlers may only be authenticated by an OAuth token

// --------------------------------------
// Authenication

func (h HandlerService) InternalPostOauthAuthorize(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalPostOauthToken(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalPostApiKeys(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

// --------------------------------------
// Ad Configurations

func (h HandlerService) InternalPostAdConfigurations(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) InternalGetAdConfigurations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalPutAdConfiguration(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalDeleteAdConfiguration(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

// --------------------------------------
// Dashboards

func (h HandlerService) InternalGetUserDashboards(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

// --------------------------------------
// Account Dashboards

func (h HandlerService) InternalPostAccountDashboards(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) InternalGetAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalDeleteAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) InternalPostAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) InternalDeleteAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) InternalPostAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) InternalGetAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalDeleteAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) InternalPutAccountDashboardAccount(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalPostAccountDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) InternalPutAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalDeleteAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) InternalGetAccountDashboardTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalPutAccountDashboardTransactionsSettings(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

// --------------------------------------
// Validator Dashboards

func (h HandlerService) InternalPostValidatorDashboards(w http.ResponseWriter, r *http.Request) {
	var err error
	userId, err := getUser(r)
	if err != nil {
		returnUnauthorized(w, err)
		return
	}
	req := struct {
		Name    string `json:"name"`
		Network string `json:"network"`
	}{}
	if internalErr := checkBody(&err, &req, r.Body); internalErr != nil {
		returnInternalServerError(w, internalErr)
		return
	}
	name := checkNameNotEmpty(&err, req.Name)
	network := checkNetwork(&err, req.Network)
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.CreateValidatorDashboard(userId, name, network)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}
	returnCreated(w, response)
}

func (h HandlerService) InternalGetValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardOverview(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h HandlerService) InternalDeleteValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	err = h.dai.RemoveValidatorDashboardOverview(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	returnNoContent(w)
}

func (h HandlerService) InternalPostValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	var err error
	req := struct {
		Name string `json:"name"`
	}{}
	if internalErr := checkBody(&err, &req, r.Body); internalErr != nil {
		returnInternalServerError(w, internalErr)
		return
	}
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	name := checkNameNotEmpty(&err, req.Name)
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.CreateValidatorDashboardGroup(dashboardId, name)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	// TODO check group limit reached

	returnCreated(w, response)
}

func (h HandlerService) InternalDeleteValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	groupId := checkGroupId(&err, vars["group_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	err = h.dai.RemoveValidatorDashboardGroup(dashboardId, groupId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}

	returnNoContent(w)
}

func (h HandlerService) InternalPostValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	var err error
	req := struct {
		Validators []string `json:"validators"`
		GroupId    string   `json:"group_id,omitempty"`
	}{}
	if internalErr := checkBody(&err, &req, r.Body); internalErr != nil {
		returnInternalServerError(w, internalErr)
		return
	}
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	groupId := checkGroupId(&err, req.GroupId)
	validators := checkValidatorArray(&err, req.Validators)
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.AddValidatorDashboardValidators(dashboardId, groupId, validators)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	// TODO check validator limit reached
	returnCreated(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) InternalDeleteValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	var err error
	// TODO check body for validators, ignore query param if body is present
	vars := mux.Vars(r)
	q := r.URL.Query()
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	validators := checkValidatorList(&err, q.Get("validators"))
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	err = h.dai.RemoveValidatorDashboardValidators(dashboardId, validators)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}

	returnNoContent(w)
}

func (h HandlerService) InternalPostValidatorDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	var err error
	req := struct {
		Name          string `json:"name"`
		ShareSettings struct {
			GroupNames bool `json:"group_names"`
		} `json:"share_settings"`
	}{}
	if internalErr := checkBody(&err, &req, r.Body); internalErr != nil {
		returnInternalServerError(w, internalErr)
		return
	}
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	name := checkNameNotEmpty(&err, req.Name)
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.CreateValidatorDashboardPublicId(dashboardId, name, req.ShareSettings.GroupNames)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	//TODO check public id limit reached
	returnCreated(w, response)
}

func (h HandlerService) InternalPutValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	var err error
	req := struct {
		Name          string `json:"name"`
		ShareSettings struct {
			GroupNames bool `json:"group_names"`
		} `json:"share_settings"`
	}{}
	if internalErr := checkBody(&err, &req, r.Body); internalErr != nil {
		returnInternalServerError(w, internalErr)
		return
	}
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	publicDashboardId := checkPublicDashboardId(&err, vars["public_dashboard_id"])
	name := checkNameNotEmpty(&err, req.Name)
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.UpdateValidatorDashboardPublicId(dashboardId, publicDashboardId, name, req.ShareSettings.GroupNames)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h HandlerService) InternalDeleteValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	publicDashboardId := checkPublicDashboardId(&err, vars["public_dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	err = h.dai.RemoveValidatorDashboardPublicId(dashboardId, publicDashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}

	returnNoContent(w)
}

func (h HandlerService) InternalGetValidatorDashboardSlotViz(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}

	data, err := h.dai.GetValidatorDashboardSlotViz(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSlotVizResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardSummary(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	sortingParams := checkSortingParams[types.VDBSummaryTableColumn](&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardSummary(dashboardId, pagingParams.cursor, sortingParams, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSummaryResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardGroupSummary(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	groupId := checkGroupId(&err, vars["group_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardGroupSummary(dashboardId, groupId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupSummaryResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardSummaryChart(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardSummaryChart(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSummaryChartResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardRewards(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	sortingParams := checkSortingParams[types.VDBRewardsTableColumn](&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardRewards(dashboardId, pagingParams.cursor, sortingParams, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardRewardsResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardGroupRewards(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	groupId := checkGroupId(&err, vars["group_id"])
	epoch := checkUint(&err, vars["epoch"], "epoch")
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardGroupRewards(dashboardId, groupId, epoch)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupRewardsResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardRewardsChart(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardRewardsChart(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardRewardsChartResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardDuties(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	epoch := checkUint(&err, vars["epoch"], "epoch")
	pagingParams := checkPagingParams(&err, r)
	sortingParams := checkSortingParams[types.VDBDutiesTableColumn](&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardDuties(dashboardId, epoch, pagingParams.cursor, sortingParams, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardDutiesResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardBlocks(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	sortingParams := checkSortingParams[types.VDBBlocksTableColumn](&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardBlocks(dashboardId, pagingParams.cursor, sortingParams, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardBlocksResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardHeatmap(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardHeatmap(dashboardId)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardHeatmapResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardGroupHeatmap(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	groupId := checkGroupId(&err, vars["group_id"])
	epoch := checkUint(&err, vars["epoch"], "epoch")
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardGroupHeatmap(dashboardId, groupId, epoch)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupHeatmapResponse{
		Data: data,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardExecutionLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardElDeposits(dashboardId, pagingParams.cursor, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardExecutionLayerDepositsResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardConsensusLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardClDeposits(dashboardId, pagingParams.cursor, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardConsensusLayerDepositsResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}

func (h HandlerService) InternalGetValidatorDashboardWithdrawals(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	dashboardId := checkDashboardId(&err, vars["dashboard_id"])
	pagingParams := checkPagingParams(&err, r)
	sortingParams := checkSortingParams[types.VDBWithdrawalsTableColumn](&err, r)
	if err != nil {
		returnBadRequest(w, err)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardWithdrawals(dashboardId, pagingParams.cursor, sortingParams, pagingParams.search, pagingParams.limit)
	if err != nil {
		returnInternalServerError(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardWithdrawalsResponse{
		Data:   data,
		Paging: paging,
	}
	returnOk(w, response)
}