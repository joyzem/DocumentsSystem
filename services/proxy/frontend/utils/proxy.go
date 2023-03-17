package utils

import (
	"fmt"

	"github.com/joyzem/documents/services/proxy/dto"
	"github.com/levigross/grequests"
)

func GetProxyById(id int) (*dto.ProxyByIdResponse, error) {
	proxyUrl := fmt.Sprintf("%s/proxy/%d", GetProxiesAddress(), id)
	proxyResp, err := grequests.Get(proxyUrl, &grequests.RequestOptions{
		JSON: dto.ProxyByIdRequest{
			Id: id,
		}})

	var proxy dto.ProxyByIdResponse
	proxyResp.JSON(&proxy)
	return &proxy, err
}
