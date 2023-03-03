// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	nft "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/handler/nft"
	root "github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/handler/root"
	"github.com/ZecreyGaming/zecrey_treasure_hunt/game/api/server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: root.GetStatusHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/asset/mintNft",
				Handler: nft.MintNftHandler(serverCtx),
			},
		},
	)
}
