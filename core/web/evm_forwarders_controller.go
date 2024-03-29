package web

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/DeAI-Artist/MintAI/core/chains/evm/forwarders"
	"github.com/DeAI-Artist/MintAI/core/chains/evm/logpoller"
	ubig "github.com/DeAI-Artist/MintAI/core/chains/evm/utils/big"
	"github.com/DeAI-Artist/MintAI/core/logger/audit"
	"github.com/DeAI-Artist/MintAI/core/services/chainlink"
	"github.com/DeAI-Artist/MintAI/core/services/pg"
	"github.com/DeAI-Artist/MintAI/core/utils/stringutils"
	"github.com/DeAI-Artist/MintAI/core/web/presenters"

	"github.com/gin-gonic/gin"
)

// EVMForwardersController manages EVM forwarders.
type EVMForwardersController struct {
	App chainlink.Application
}

// Index lists EVM forwarders.
func (cc *EVMForwardersController) Index(c *gin.Context, size, page, offset int) {
	orm := forwarders.NewORM(cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database())
	fwds, count, err := orm.FindForwarders(0, size)

	if err != nil {
		jsonAPIError(c, http.StatusBadRequest, err)
		return
	}

	var resources []presenters.EVMForwarderResource
	for _, fwd := range fwds {
		resources = append(resources, presenters.NewEVMForwarderResource(fwd))
	}

	paginatedResponse(c, "forwarder", size, page, resources, count, err)
}

// TrackEVMForwarderRequest is a JSONAPI request for creating an EVM forwarder.
type TrackEVMForwarderRequest struct {
	EVMChainID *ubig.Big      `json:"evmChainId"`
	Address    common.Address `json:"address"`
}

// Track adds a new EVM forwarder.
func (cc *EVMForwardersController) Track(c *gin.Context) {
	request := &TrackEVMForwarderRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, err)
		return
	}
	orm := forwarders.NewORM(cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database())
	fwd, err := orm.CreateForwarder(request.Address, *request.EVMChainID)

	if err != nil {
		jsonAPIError(c, http.StatusBadRequest, err)
		return
	}

	cc.App.GetAuditLogger().Audit(audit.ForwarderCreated, map[string]interface{}{
		"forwarderID":         fwd.ID,
		"forwarderAddress":    fwd.Address,
		"forwarderEVMChainID": fwd.EVMChainID,
	})
	jsonAPIResponseWithStatus(c, presenters.NewEVMForwarderResource(fwd), "forwarder", http.StatusCreated)
}

// Delete removes an EVM Forwarder.
func (cc *EVMForwardersController) Delete(c *gin.Context) {
	id, err := stringutils.ToInt64(c.Param("fwdID"))
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, err)
		return
	}

	filterCleanup := func(tx pg.Queryer, evmChainID int64, addr common.Address) error {
		chain, err2 := cc.App.GetRelayers().LegacyEVMChains().Get(big.NewInt(evmChainID).String())
		if err2 != nil {
			// If the chain id doesn't even exist, or logpoller is disabled, then there isn't any filter to clean up.  Returning an error
			// here could be dangerous as it would make it impossible to delete a forwarder with an invalid chain id
			return nil
		}

		if chain.LogPoller() == logpoller.LogPollerDisabled {
			// handle same as non-existent chain id
			return nil
		}
		return chain.LogPoller().UnregisterFilter(forwarders.FilterName(addr), pg.WithQueryer(tx))
	}

	orm := forwarders.NewORM(cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database())
	err = orm.DeleteForwarder(id, filterCleanup)

	if err != nil {
		jsonAPIError(c, http.StatusInternalServerError, err)
		return
	}

	cc.App.GetAuditLogger().Audit(audit.ForwarderDeleted, map[string]interface{}{"id": id})
	jsonAPIResponseWithStatus(c, nil, "forwarder", http.StatusNoContent)
}
