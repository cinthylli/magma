/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"fmt"
	"net/http"

	"magma/orc8r/cloud/go/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	eventdh "magma/orc8r/cloud/go/services/eventd/obsidian/handlers"
	models2 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/lib/go/service/config"

	"github.com/labstack/echo"
	"github.com/olivere/elastic/v7"
)

const (
	Networks                           = "networks"
	ListNetworksPath                   = obsidian.V1Root + Networks
	RegisterNetworkPath                = obsidian.V1Root + Networks
	ManageNetworkPath                  = obsidian.V1Root + Networks + obsidian.UrlSep + ":network_id"
	ManageNetworkNamePath              = ManageNetworkPath + obsidian.UrlSep + "name"
	ManageNetworkTypePath              = ManageNetworkPath + obsidian.UrlSep + "type"
	ManageNetworkDescriptionPath       = ManageNetworkPath + obsidian.UrlSep + "description"
	ManageNetworkFeaturesPath          = ManageNetworkPath + obsidian.UrlSep + "features"
	ManageNetworkDNSPath               = ManageNetworkPath + obsidian.UrlSep + "dns"
	ManageNetworkDNSRecordsPath        = ManageNetworkDNSPath + obsidian.UrlSep + "records"
	ManageNetworkDNSRecordByDomainPath = ManageNetworkDNSRecordsPath + obsidian.UrlSep + ":domain"

	Gateways                     = "gateways"
	ListGatewaysPath             = ManageNetworkPath + obsidian.UrlSep + Gateways
	ManageGatewayPath            = ListGatewaysPath + obsidian.UrlSep + ":gateway_id"
	ManageGatewayNamePath        = ManageGatewayPath + obsidian.UrlSep + "name"
	ManageGatewayDescriptionPath = ManageGatewayPath + obsidian.UrlSep + "description"
	ManageGatewayConfigPath      = ManageGatewayPath + obsidian.UrlSep + "magmad"
	ManageGatewayDevicePath      = ManageGatewayPath + obsidian.UrlSep + "device"
	ManageGatewayStatePath       = ManageGatewayPath + obsidian.UrlSep + "status"
	ManageGatewayTierPath        = ManageGatewayPath + obsidian.UrlSep + "tier"

	Channels               = "channels"
	ListChannelsPath       = obsidian.V1Root + Channels
	ManageChannelPath      = obsidian.V1Root + Channels + obsidian.UrlSep + ":channel_id"
	Tiers                  = "tiers"
	ListTiersPath          = ManageNetworkPath + obsidian.UrlSep + Tiers
	ManageTiersPath        = ListTiersPath + obsidian.UrlSep + ":tier_id"
	ManageTierNamePath     = ManageTiersPath + obsidian.UrlSep + "name"
	ManageTierVersionPath  = ManageTiersPath + obsidian.UrlSep + "version"
	ManageTierImagesPath   = ManageTiersPath + obsidian.UrlSep + "images"
	ManageTierImagePath    = ManageTierImagesPath + obsidian.UrlSep + ":image_name"
	ManageTierGatewaysPath = ManageTiersPath + obsidian.UrlSep + "gateways"
	ManageTierGatewayPath  = ManageTierGatewaysPath + obsidian.UrlSep + ":gateway_id"

	LogSearchQueryPath = ManageNetworkPath + obsidian.UrlSep + "logs" + obsidian.UrlSep + "search"
	LogCountQueryPath  = ManageNetworkPath + obsidian.UrlSep + "logs" + obsidian.UrlSep + "count"
)

// GetObsidianHandlers returns all plugin-level obsidian handlers for orc8r
func GetObsidianHandlers() []obsidian.Handler {
	ret := []obsidian.Handler{
		// Magma V1 Network
		{Path: ListNetworksPath, Methods: obsidian.GET, HandlerFunc: listNetworks},
		{Path: RegisterNetworkPath, Methods: obsidian.POST, HandlerFunc: registerNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.GET, HandlerFunc: getNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.PUT, HandlerFunc: updateNetwork},
		{Path: ManageNetworkPath, Methods: obsidian.DELETE, HandlerFunc: deleteNetwork},

		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.POST, HandlerFunc: CreateDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.GET, HandlerFunc: ReadDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.PUT, HandlerFunc: UpdateDNSRecord},
		{Path: ManageNetworkDNSRecordByDomainPath, Methods: obsidian.DELETE, HandlerFunc: DeleteDNSRecord},

		// Magma V1 Gateways
		{Path: ListGatewaysPath, Methods: obsidian.GET, HandlerFunc: ListGatewaysHandler},
		{Path: ListGatewaysPath, Methods: obsidian.POST, HandlerFunc: CreateGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.GET, HandlerFunc: GetGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.PUT, HandlerFunc: UpdateGatewayHandler},
		{Path: ManageGatewayPath, Methods: obsidian.DELETE, HandlerFunc: DeleteGatewayHandler},
		{Path: ManageGatewayStatePath, Methods: obsidian.GET, HandlerFunc: GetStateHandler},

		// Upgrades
		{Path: ListChannelsPath, Methods: obsidian.GET, HandlerFunc: listChannelsHandler},
		{Path: ListChannelsPath, Methods: obsidian.POST, HandlerFunc: createChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.GET, HandlerFunc: readChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.PUT, HandlerFunc: updateChannelHandler},
		{Path: ManageChannelPath, Methods: obsidian.DELETE, HandlerFunc: deleteChannelHandler},
		{Path: ListTiersPath, Methods: obsidian.GET, HandlerFunc: listTiersHandler},
		{Path: ListTiersPath, Methods: obsidian.POST, HandlerFunc: createTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.GET, HandlerFunc: readTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.PUT, HandlerFunc: updateTierHandler},
		{Path: ManageTiersPath, Methods: obsidian.DELETE, HandlerFunc: deleteTierHandler},
		{Path: ManageTierImagesPath, Methods: obsidian.POST, HandlerFunc: createTierImage},
		{Path: ManageTierImagePath, Methods: obsidian.DELETE, HandlerFunc: deleteImage},
		{Path: ManageTierGatewaysPath, Methods: obsidian.POST, HandlerFunc: createTierGateway},
		{Path: ManageTierGatewayPath, Methods: obsidian.DELETE, HandlerFunc: deleteTierGateway},

		// Magmad commands
		{Path: RebootGatewayV1, Methods: obsidian.POST, HandlerFunc: rebootGateway},
		{Path: RestartServicesV1, Methods: obsidian.POST, HandlerFunc: restartServices},
		{Path: GatewayPingV1, Methods: obsidian.POST, HandlerFunc: gatewayPing},
		{Path: GatewayGenericCommandV1, Methods: obsidian.POST, HandlerFunc: gatewayGenericCommand},
		{Path: TailGatewayLogsV1, Methods: obsidian.POST, HandlerFunc: tailGatewayLogs},
	}
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkNamePath, new(models.NetworkName), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkTypePath, new(models.NetworkType), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDescriptionPath, new(models.NetworkDescription), "")...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkFeaturesPath, &models2.NetworkFeatures{}, orc8r.NetworkFeaturesConfig)...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDNSPath, &models2.NetworkDNSConfig{}, orc8r.DnsdNetworkType)...)
	ret = append(ret, GetPartialNetworkHandlers(ManageNetworkDNSRecordsPath, new(models2.NetworkDNSRecords), "")...)

	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayNamePath, new(models.GatewayName))...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayDescriptionPath, new(models.GatewayDescription))...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayConfigPath, &models2.MagmadGatewayConfigs{})...)
	ret = append(ret, GetPartialGatewayHandlers(ManageGatewayTierPath, new(models2.TierID))...)
	ret = append(ret, GetGatewayDeviceHandlers(ManageGatewayDevicePath)...)

	ret = append(ret, GetPartialEntityHandlers(ManageTierNamePath, "tier_id", new(models2.TierName))...)
	ret = append(ret, GetPartialEntityHandlers(ManageTierVersionPath, "tier_id", new(models2.TierVersion))...)
	ret = append(ret, GetPartialEntityHandlers(ManageTierImagesPath, "tier_id", new(models2.TierImages))...)
	ret = append(ret, GetPartialEntityHandlers(ManageTierGatewaysPath, "tier_id", new(models2.TierGateways))...)

	// Elastic
	elasticConfig, err := config.GetServiceConfig(orc8r.ModuleName, "elastic")
	if err != nil {
		ret = append(ret, obsidian.Handler{Path: LogSearchQueryPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
		ret = append(ret, obsidian.Handler{Path: LogCountQueryPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
		ret = append(ret, obsidian.Handler{Path: eventdh.EventsPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
	} else {
		elasticHost := elasticConfig.MustGetString("elasticHost")
		elasticPort := elasticConfig.MustGetInt("elasticPort")

		client, err := elastic.NewSimpleClient(elastic.SetURL(fmt.Sprintf("http://%s:%d", elasticHost, elasticPort)))
		if err != nil {
			ret = append(ret, obsidian.Handler{Path: LogSearchQueryPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
			ret = append(ret, obsidian.Handler{Path: LogCountQueryPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsRootPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsCountPath, Methods: obsidian.GET, HandlerFunc: getInitErrorHandler(err)})
		} else {
			ret = append(ret, obsidian.Handler{Path: LogSearchQueryPath, Methods: obsidian.GET, HandlerFunc: GetQueryLogHandler(client)})
			ret = append(ret, obsidian.Handler{Path: LogCountQueryPath, Methods: obsidian.GET, HandlerFunc: GetCountLogHandler(client)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsRootPath, Methods: obsidian.GET, HandlerFunc: eventdh.GetMultiStreamEventsHandler(client)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsCountPath, Methods: obsidian.GET, HandlerFunc: eventdh.GetEventCountHandler(client)})
			ret = append(ret, obsidian.Handler{Path: eventdh.EventsPath, Methods: obsidian.GET, HandlerFunc: eventdh.GetEventsHandler(client)})
		}
	}

	ret = append(ret, obsidian.Handler{
		Path:    "/",
		Methods: obsidian.GET,
		HandlerFunc: func(c echo.Context) error {
			return c.JSON(
				http.StatusOK,
				"hello",
			)
		},
	})
	return ret
}

func getInitErrorHandler(err error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return obsidian.HttpError(fmt.Errorf("initialization Error: %v", err), 500)
	}
}
